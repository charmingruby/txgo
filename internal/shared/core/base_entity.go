package core

import "time"

func NewBaseEntity() BaseEntity {
	return BaseEntity{
		id:        NewID(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
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
