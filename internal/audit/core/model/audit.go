package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type Audit struct {
	id        string
	module    string
	context   string
	message   string
	createdAt time.Time
	updatedAt time.Time
}

type NewAuditInput struct {
	Module  string
	Context string
	Message string
}

func NewAudit(in NewAuditInput) (*Audit, error) {
	p := Audit{
		id:        core.NewID(),
		module:    in.Module,
		context:   in.Context,
		message:   in.Message,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Audit) validate() error {
	if p.module == "" {
		return core_err.NewModelErr("module is required")
	}

	if p.context == "" {
		return core_err.NewModelErr("context is required")
	}

	if p.message == "" {
		return core_err.NewModelErr("message is required")
	}

	return nil
}

func (a *Audit) ID() string {
	return a.id
}

func (a *Audit) Module() string {
	return a.module
}

func (a *Audit) Context() string {
	return a.context
}

func (a *Audit) Message() string {
	return a.message
}

func (a *Audit) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Audit) UpdatedAt() time.Time {
	return a.updatedAt
}
