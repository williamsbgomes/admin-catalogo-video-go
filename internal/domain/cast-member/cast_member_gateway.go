package castmember

import "github.com/williamsbgomes/admin-catalogo-video-go/internal/domain/pagination"

type CastMemberGateway interface {
	Create(castMember *CastMember) (*CastMember, error)
	Update(castMember *CastMember) (*CastMember, error)
	DeleteByID(id string) error
	FindByID(id string) (*CastMember, error)
	FindAll(query pagination.SearchQuery) (*pagination.Pagination[CastMember], error)
}
