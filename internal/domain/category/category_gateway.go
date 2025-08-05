package category

import "github.com/williamsbgomes/admin-catalogo-video-go/internal/domain/pagination"

type CategoryGateway interface {
	Create(category *Category) (*Category, error)
	Update(category *Category) (*Category, error)
	DeleteByID(id string) error
	FindByID(id string) (*Category, error)
	FindAll(query pagination.SearchQuery) (*pagination.Pagination[Category], error)
}
