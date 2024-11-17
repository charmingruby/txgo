package mysql

import (
	"fmt"

	"github.com/charmingruby/txgo/internal/audit/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/internal/shared/database/mysql"
)

const (
	AUDITS_REPOSITORY       = "audits"
	AUDIT_ACTORS_REPOSITORY = "audit_actors"
)

func NewAuditRepository(db mysql.Database) *AuditRepository {
	return &AuditRepository{db: db}
}

type AuditRepository struct {
	db mysql.Database
}

func (r *AuditRepository) StoreAudit(audit *model.Audit) error {
	query := fmt.Sprintf("INSERT INTO %s (id, module, context, message, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", AUDITS_REPOSITORY)

	_, err := r.db.Exec(query, audit.ID(), audit.Module(), audit.Context(), audit.Message(), audit.CreatedAt(), audit.UpdatedAt())
	if err != nil {
		return core_err.NewPersistenceErr(err, "audit store", "mysql")
	}

	return nil
}

func (r *AuditRepository) StoreActor(auditActor *model.AuditActor) error {
	query := fmt.Sprintf("INSERT INTO %s (id, email, audit_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", AUDIT_ACTORS_REPOSITORY)

	_, err := r.db.Exec(query, auditActor.ID(), auditActor.Email(), auditActor.AuditID(), auditActor.CreatedAt(), auditActor.UpdatedAt())
	if err != nil {
		return core_err.NewPersistenceErr(err, "audit actor store", "mysql")
	}

	return nil
}
