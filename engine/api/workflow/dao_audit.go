package workflow

import (
	"github.com/go-gorp/gorp"

	"github.com/ovh/cds/sdk"
)

// InsertAudit insert a workflow audit
func InsertAudit(db gorp.SqlExecutor, a sdk.AuditWorkflow) error {
	audit := auditWorkflow(a)
	return db.Insert(&audit)
}

// LoadAudits Load audits for the given workflow
func LoadAudits(db gorp.SqlExecutor, workflowID int64) ([]sdk.AuditWorkflow, error) {
	query := `
		SELECT * FROM workflow_audit WHERE workflow_id = $1
	`
	var audits []auditWorkflow
	if _, err := db.Select(&audits, query, workflowID); err != nil {
		return nil, sdk.WrapError(err, "workflow.loadAudits> Unable to load audits")
	}

	workflowAudits := make([]sdk.AuditWorkflow, len(audits), len(audits))
	for i := range audits {
		workflowAudits[i] = sdk.AuditWorkflow(audits[i])
	}
	return workflowAudits, nil
}

// LoadAudit Load audit for the given workflow
func LoadAudit(db gorp.SqlExecutor, auditID int64) (sdk.AuditWorkflow, error) {
	var audit auditWorkflow
	if err := db.SelectOne(&audit, "SELECT * FROM workflow_audit WHERE id = $1", auditID); err != nil {
		return sdk.AuditWorkflow{}, sdk.WrapError(err, "workflow.LoadAudit> Unable to load audit")
	}

	return sdk.AuditWorkflow(audit), nil
}
