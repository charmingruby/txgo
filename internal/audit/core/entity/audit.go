package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewAuditInput struct {
	actors      []string
	context     string
	description string
}

func NewAudit(in NewAuditInput) (*Audit, error) {
	p := Audit{
		actors:      in.actors,
		context:     in.context,
		description: in.description,
		BaseEntity:  core.NewBaseEntity(),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return &p, nil
}

func NewAuditFrom(in Audit) *Audit {
	return &Audit{
		actors:      in.actors,
		context:     in.context,
		description: in.description,
		BaseEntity:  in.BaseEntity,
	}
}

func (p *Audit) validate() error {
	if len(p.actors) == 0 {
		return core_err.NewEntityErr("actors are required")
	}

	if p.context == "" {
		return core_err.NewEntityErr("context is required")
	}

	if p.description == "" {
		return core_err.NewEntityErr("description is required")
	}

	return nil
}

type Audit struct {
	core.BaseEntity

	actors      []string
	context     string
	description string
}
