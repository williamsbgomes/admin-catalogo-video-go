package category

type CategoryGateway interface {
	Create(category *Category) (*Category, error)
	Update(category *Category) (*Category, error)
	DeleteByID(id string) error
	FindByID(id string) (*Category, error)
	FindAll() ([]Category, error)
}
