package mongodb

import "context"

type Repository[M any] interface {
	FindOne(ctx context.Context, id string) (M, error)
	Find(ctx context.Context, ids []string) ([]M, error)
	FindAll(ctx context.Context) ([]M, error)
	Search(ctx context.Context, filters map[string][]any, opts ...SearchOptions) ([]M, error)

	CountDocuments(ctx context.Context, filters map[string][]any) (int64, error)

	Upsert(ctx context.Context, id string, m M) error

	Delete(ctx context.Context, id string) (M, error)
	DeleteMany(ctx context.Context, filters map[string][]any) error
}
