package castmember

type CastMemberGateway interface {
	Create(castMember *CastMember) (*CastMember, error)
	Update(castMember *CastMember) (*CastMember, error)
	DeleteByID(id string) error
	FindByID(id string) (*CastMember, error)
	FindAll() ([]CastMember, error)
}
