package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/tevino/abool"

	"github.com/go-gorp/gorp"

	"github.com/ovh/cds/engine/api/cache"
	"github.com/ovh/cds/engine/api/group"
	"github.com/ovh/cds/engine/api/observability"
	"github.com/ovh/cds/engine/api/permission"
	"github.com/ovh/cds/engine/service"
	"github.com/ovh/cds/sdk"
	"github.com/ovh/cds/sdk/log"
)

// eventsBrokerSubscribe is the information needed to subscribe
type eventsBrokerSubscribe struct {
	UUID    string
	User    *sdk.User
	isAlive *abool.AtomicBool
	w       http.ResponseWriter
	mutex   sync.Mutex
}

// lastUpdateBroker keeps connected client of the current route,
type eventsBroker struct {
	clients          map[string]*eventsBrokerSubscribe
	messages         chan sdk.Event
	dbFunc           func() *gorp.DbMap
	cache            cache.Store
	router           *Router
	chanAddClient    chan (*eventsBrokerSubscribe)
	chanRemoveClient chan (string)
}

//Init the eventsBroker
func (b *eventsBroker) Init(ctx context.Context) {
	// Start cache Subscription
	sdk.GoRoutine(ctx, "eventsBroker.Init.CacheSubscribe", func(ctx context.Context) {
		b.cacheSubscribe(ctx, b.messages, b.cache)
	})

	sdk.GoRoutine(ctx, "eventsBroker.Init.Start", func(ctx context.Context) {
		b.Start(ctx)
	})
}

func (b *eventsBroker) cacheSubscribe(c context.Context, cacheMsgChan chan<- sdk.Event, store cache.Store) {
	pubSub := store.Subscribe("events_pubsub")
	tick := time.NewTicker(50 * time.Millisecond)
	defer tick.Stop()
	for {
		select {
		case <-c.Done():
			if c.Err() != nil {
				log.Error("events.cacheSubscribe> Exiting: %v", c.Err())
				return
			}
		case <-tick.C:
			msg, err := store.GetMessageFromSubscription(c, pubSub)
			if err != nil {
				log.Warning("events.cacheSubscribe> Cannot get message %s: %s", msg, err)
				continue
			}
			var e sdk.Event
			if err := json.Unmarshal([]byte(msg), &e); err != nil {
				log.Warning("events.cacheSubscribe> Cannot unmarshal event %s: %s", msg, err)
				continue
			}

			switch e.EventType {
			case "sdk.EventPipelineBuild", "sdk.EventJob":
				continue
			}
			observability.Record(c, b.router.Stats.SSEEvents, 1)
			cacheMsgChan <- e
		}
	}
}

// Start the broker
func (b *eventsBroker) Start(ctx context.Context) {
	b.chanAddClient = make(chan (*eventsBrokerSubscribe))
	b.chanRemoveClient = make(chan (string))

	tickerMetrics := time.NewTicker(10 * time.Second)
	defer tickerMetrics.Stop()

	for {
		select {
		case <-tickerMetrics.C:
			observability.Record(b.router.Background, b.router.Stats.SSEClients, int64(len(b.clients)))

		case <-ctx.Done():
			if b.clients != nil {
				for uuid := range b.clients {
					delete(b.clients, uuid)
				}
				observability.Record(b.router.Background, b.router.Stats.SSEClients, 0)

			}
			if ctx.Err() != nil {
				log.Error("eventsBroker.Start> Exiting: %v", ctx.Err())
				return
			}

		case receivedEvent := <-b.messages:
			for i := range b.clients {
				c := b.clients[i]
				if c == nil {
					delete(b.clients, i)
					continue
				}

				// Send the event to the client sse within a goroutine
				s := "sse-" + b.clients[i].UUID
				sdk.GoRoutine(ctx, s,
					func(ctx context.Context) {
						if c.isAlive.IsSet() {
							log.Debug("send data to %s", c.UUID)
							if err := c.Send(receivedEvent); err != nil {
								log.Error("eventsBroker> unable to send event to %s: %v", c.UUID, err)
								b.chanRemoveClient <- c.UUID
							}
						}
					},
				)
			}

		case client := <-b.chanAddClient:
			b.clients[client.UUID] = client

		case uuid := <-b.chanRemoveClient:
			client, has := b.clients[uuid]
			if !has {
				continue
			}

			client.isAlive.UnSet()
			delete(b.clients, uuid)
		}
	}
}

func (b *eventsBroker) ServeHTTP() service.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

		// Make sure that the writer supports flushing.
		f, ok := w.(http.Flusher)
		if !ok {
			return sdk.WrapError(fmt.Errorf("streaming unsupported"), "")
		}

		uuid := sdk.UUID()
		client := &eventsBrokerSubscribe{
			UUID:    uuid,
			User:    getUser(ctx),
			isAlive: abool.NewBool(true),
			w:       w,
		}

		// Add this client to the map of those that should receive updates
		b.chanAddClient <- client

		// Set the headers related to event streaming.
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")

		if _, err := w.Write([]byte(fmt.Sprintf("data: ACK: %s \n\n", uuid))); err != nil {
			return sdk.WrapError(err, "events.write> Unable to send ACK to client")
		}
		f.Flush()

		tick := time.NewTicker(time.Second)
		defer tick.Stop()

	leave:
		for {
			select {
			case <-ctx.Done():
				log.Info("events.Http: context done")
				b.chanRemoveClient <- client.UUID
				break leave
			case <-r.Context().Done():
				log.Info("events.Http: client disconnected")
				b.chanRemoveClient <- client.UUID
				break leave
			case <-tick.C:
				if _, err := w.Write([]byte("")); err != nil {
					return sdk.WrapError(err, "events.write> Unable to ping client")
				}
				f.Flush()
			}
		}

		return nil
	}
}

func (client *eventsBrokerSubscribe) manageEvent(event sdk.Event) bool {
	var isSharedInfra bool
	for _, g := range client.User.Groups {
		if g.ID == group.SharedInfraGroup.ID {
			isSharedInfra = true
			break
		}
	}

	if strings.HasPrefix(event.EventType, "sdk.EventProject") {
		if client.User.Admin || isSharedInfra || permission.ProjectPermission(event.ProjectKey, client.User) >= permission.PermissionRead {
			return true
		}
		return false
	}
	if strings.HasPrefix(event.EventType, "sdk.EventWorkflow") || strings.HasPrefix(event.EventType, "sdk.EventRunWorkflow") {
		if client.User.Admin || isSharedInfra || permission.WorkflowPermission(event.ProjectKey, event.WorkflowName, client.User) >= permission.PermissionRead {
			return true
		}
		return false
	}
	if strings.HasPrefix(event.EventType, "sdk.EventApplication") {
		if client.User.Admin || isSharedInfra || permission.ApplicationPermission(event.ProjectKey, event.ApplicationName, client.User) >= permission.PermissionRead {
			return true
		}
		return false
	}
	if strings.HasPrefix(event.EventType, "sdk.EventPipeline") {
		if client.User.Admin || isSharedInfra || permission.PipelinePermission(event.ProjectKey, event.PipelineName, client.User) >= permission.PermissionRead {
			return true
		}
		return false
	}
	if strings.HasPrefix(event.EventType, "sdk.EventEnvironment") {
		if client.User.Admin || isSharedInfra || permission.EnvironmentPermission(event.ProjectKey, event.EnvironmentName, client.User) >= permission.PermissionRead {
			return true
		}
		return false
	}
	if strings.HasPrefix(event.EventType, "sdk.EventBroadcast") {
		if client.User.Admin || isSharedInfra || event.ProjectKey == "" || permission.AccessToProject(event.ProjectKey, client.User, permission.PermissionRead) {
			return true
		}
		return false
	}
	return false
}

// Send an event to a client
func (client *eventsBrokerSubscribe) Send(event sdk.Event) (err error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if client == nil || client.w == nil {
		return nil
	}

	// Make sure that the writer supports flushing.
	f, ok := client.w.(http.Flusher)
	if !ok {
		return sdk.WrapError(fmt.Errorf("streaming unsupported"), "")
	}

	if ok := client.manageEvent(event); !ok {
		return nil
	}

	msg, err := json.Marshal(event)
	if err != nil {
		return sdk.WrapError(err, "Unable to marshall event")
	}

	var buffer bytes.Buffer
	buffer.WriteString("data: ")
	buffer.Write(msg)
	buffer.WriteString("\n\n")

	if !client.isAlive.IsSet() {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	if _, err := client.w.Write(buffer.Bytes()); err != nil {
		return sdk.WrapError(err, "unable to write to client")
	}
	f.Flush()

	return nil
}
