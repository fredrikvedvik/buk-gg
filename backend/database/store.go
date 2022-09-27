package database

import (
	"cloud.google.com/go/firestore"
	"context"
	status "github.com/ansel1/merry/v2/grpcstatus"
	"google.golang.org/grpc/codes"
)

type Store struct {
	client *firestore.Client
}

type Collection[T any] struct {
	client     *firestore.Client
	collection string
}

func New(client *firestore.Client) *Store {
	return &Store{
		client: client,
	}
}

func get[T any](ctx context.Context, client *firestore.Client, collection string, key string) (*T, error) {
	doc, err := client.Collection(collection).Doc(key).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}

	var r T
	err = doc.DataTo(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

func set[T any](ctx context.Context, client *firestore.Client, collection string, key string, value T) error {
	ref := client.Collection(collection).Doc(key)
	_, err := ref.Set(ctx, value)
	return err
}

func (c *Collection[T]) Get(ctx context.Context, key string) (*T, error) {
	return get[T](ctx, c.client, c.collection, key)
}

func (c *Collection[T]) Set(ctx context.Context, key string, value T) error {
	return set[T](ctx, c.client, c.collection, key, value)
}
