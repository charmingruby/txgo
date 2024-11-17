package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type AuditActor struct {
	id        string
	email     string
	auditID   string
	createdAt time.Time
	updatedAt time.Time
}

type NewAuditActorInput struct {
	Email   string
	AuditID string
}

func NewAuditActor(in NewAuditActorInput) (*AuditActor, error) {
	a := AuditActor{
		id:        core.NewID(),
		email:     in.Email,
		auditID:   in.AuditID,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if err := a.validate(); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *AuditActor) validate() error {
	if a.email == "" {
		return core_err.NewModelErr("email is required")
	}

	if a.auditID == "" {
		return core_err.NewModelErr("auditID is required")
	}

	return nil
}

func (a *AuditActor) ID() string {
	return a.id
}

func (a *AuditActor) Email() string {
	return a.email
}

func (a *AuditActor) AuditID() string {
	return a.auditID
}

func (a *AuditActor) CreatedAt() time.Time {
	return a.createdAt
}

func (a *AuditActor) UpdatedAt() time.Time {
	return a.updatedAt
}
