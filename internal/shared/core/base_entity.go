package core

import "time"

func NewBaseEntity() BaseEntity {
	return BaseEntity{
		id:        NewID(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

type NewBaseEntityFromInput struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBaseEntityFrom(input NewBaseEntityFromInput) BaseEntity {
	return BaseEntity{
		id:        input.ID,
		createdAt: input.CreatedAt,
		updatedAt: input.UpdatedAt,
	}
}

func (e *BaseEntity) ID() string {
	return e.id
}

func (e *BaseEntity) CreatedAt() time.Time {
	return e.createdAt
}

func (e *BaseEntity) UpdatedAt() time.Time {
	return e.updatedAt
}

type BaseEntity struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
}
