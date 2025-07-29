package category

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	timeutils "github.com/williamsbgomes/admin-catalogo-video-go/pkg/time-utils"
)

type CategoryError struct {
	msg string
}

func (c CategoryError) Error() string {
	return c.msg
}

type Category struct {
	ID          string
	Name        string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func NewCategory(name, description string, isActive bool) (*Category, error) {
	now := *timeutils.TimeNow()
	var deletedAt *time.Time
	if !isActive {
		deletedAt = &now
	}
	category := &Category{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Active:      isActive,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   deletedAt,
	}
	err := category.IsValid()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) Activate() {
	c.DeletedAt = nil
	c.Active = true
	c.UpdatedAt = *timeutils.TimeNow()
}

func (c *Category) Deactivate() {
	if c.DeletedAt == nil {
		c.DeletedAt = timeutils.TimeNow()
	}
	c.Active = false
	c.UpdatedAt = *timeutils.TimeNow()
}

func (c *Category) Update(name, description string, isActive bool) error {
	if isActive {
		c.Activate()
	} else {
		c.Deactivate()
	}
	c.Name = name
	c.Description = description
	c.UpdatedAt = *timeutils.TimeNow()
	return c.IsValid()
}

func (c *Category) IsValid() error {
	const (
		minNameLength = 3
		maxNameLength = 255
	)

	if c.ID == "" {
		return CategoryError{"'id' should not be empty"}
	}

	if c.Name == "" {
		return CategoryError{"'name' should not be empty"}
	}
	if len(c.Name) < minNameLength || len(c.Name) > maxNameLength {
		return CategoryError{fmt.Sprintf(
			"'name' must be between %d and %d characters",
			minNameLength, maxNameLength,
		)}
	}
	return nil
}
