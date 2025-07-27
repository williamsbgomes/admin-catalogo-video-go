package main

import (
	"fmt"

	"github.com/williamsbgomes/admin-catalogo-video-go/internal/entity/category"
)

func main() {
	c, err := category.NewCategory("Filmes", "A categoria mais assistida", true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Category created: %+v\n", c)
}
