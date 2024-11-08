package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type Audit struct {
	id          string
	actors      []string
	context     string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

type NewAuditInput struct {
	actors      []string
	context     string
	description string
}

func NewAudit(in NewAuditInput) (*Audit, error) {
	p := Audit{
		id:          core.NewID(),
		actors:      in.actors,
		context:     in.context,
		description: in.description,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Audit) validate() error {
	if len(p.actors) == 0 {
		return core_err.NewModelErr("actors are required")
	}

	if p.context == "" {
		return core_err.NewModelErr("context is required")
	}

	if p.description == "" {
		return core_err.NewModelErr("description is required")
	}

	return nil
}
