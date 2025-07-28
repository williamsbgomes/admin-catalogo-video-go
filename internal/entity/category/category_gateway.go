package category

type CategoryRepository interface {
	Create(category *Category) (*Category, error)
	Update(category *Category) (*Category, error)
	Delete(id string) error
	FindByID(id string) (*Category, error)
	FindAll() ([]Category, error)
}
