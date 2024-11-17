package repository

import "github.com/charmingruby/txgo/internal/audit/core/model"

type AuditRepository interface {
	StoreAudit(audit *model.Audit) error
	StoreActor(auditActor *model.AuditActor) error
}
