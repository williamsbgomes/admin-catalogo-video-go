package castmember

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	timeutils "github.com/williamsbgomes/admin-catalogo-video-go/pkg/time-utils"
)

type CastMemberError struct {
	msg string
}

func (c CastMemberError) Error() string {
	return c.msg
}

type CastMemberType string

const (
	Actor    CastMemberType = "ACTOR"
	Director CastMemberType = "DIRECTOR"
)

type CastMember struct {
	ID        string
	Name      string
	Type      CastMemberType
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCastMember(name string, castMemberType CastMemberType) (*CastMember, error) {
	now := *timeutils.TimeNow()

	castMember := &CastMember{
		ID:        uuid.New().String(),
		Name:      name,
		Type:      castMemberType,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := castMember.IsValid(); err != nil {
		return nil, err
	}
	return castMember, nil
}

func (c *CastMember) Update(name string, castMemberType CastMemberType) error {
	c.Name = name
	c.Type = castMemberType
	c.UpdatedAt = *timeutils.TimeNow()
	return c.IsValid()
}

func (c *CastMember) IsValid() error {
	const (
		minNameLength = 3
		maxNameLength = 255
	)

	if c.ID == "" {
		return CastMemberError{"'id' should not be empty"}
	}

	if c.Name == "" {
		return CastMemberError{"'name' should not be empty"}
	}
	if len(c.Name) < minNameLength || len(c.Name) > maxNameLength {
		return CastMemberError{fmt.Sprintf(
			"'name' must be between %d and %d characters",
			minNameLength, maxNameLength,
		)}
	}

	if c.Type == "" {
		return CastMemberError{"'type' should not be empty"}
	}
	if c.Type != Actor && c.Type != Director {
		return CastMemberError{"'type' must be either 'ACTOR' or 'DIRECTOR'"}
	}
	return nil
}
