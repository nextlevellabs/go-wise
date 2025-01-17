package wise

import "github.com/nextlevellabs/go-wise/wise/internal/mongodb"

func WithPageMethod(method string) mongodb.SearchOptions {
	return mongodb.WithPageMethod(method)
}

func WithPageSize(pageSize int) mongodb.SearchOptions {
	return mongodb.WithPageSize(pageSize)
}

func WithPage(page int) mongodb.SearchOptions {
	return mongodb.WithPage(page)
}

func WithSort(sort map[string]int) mongodb.SearchOptions {
	return mongodb.WithSort(sort)
}
