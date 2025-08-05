package pagination

type Pagination[T any] struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	Items       []T   `json:"items"`
}

func (p *Pagination[T]) Map(mapper func(T) any) *Pagination[any] {
	newItems := make([]any, len(p.Items))
	for i, item := range p.Items {
		newItems[i] = mapper(item)
	}

	return &Pagination[any]{
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		Total:       p.Total,
		Items:       newItems,
	}
}
