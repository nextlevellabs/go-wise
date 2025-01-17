package wise

import (
	"github.com/nextlevellabs/go-wise/wise/internal/filters"
	"github.com/nextlevellabs/go-wise/wise/internal/mongodb"
)

func WithBloomFilter(bf filters.Bloom[string]) mongodb.RepositoryOptions {
	return mongodb.WithBloomFilter(bf)
}

func WithMaxPageSize(size int) mongodb.RepositoryOptions {
	return mongodb.WithMaxPageSize(size)
}

func WithDefaultPageSize(size int) mongodb.RepositoryOptions {
	return mongodb.WithDefaultPageSize(size)
}
