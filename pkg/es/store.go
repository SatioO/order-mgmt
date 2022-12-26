package es

import "context"

type AggregateStore interface {
	Load(ctx context.Context, aggregate Aggregate) error

	Save(ctx context.Context, aggregate Aggregate) error

	Exists(ctx context.Context, streamID string) error
}
