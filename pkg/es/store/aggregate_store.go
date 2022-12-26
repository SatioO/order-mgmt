package store

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/satioO/order-mgmt/pkg/es"
	"github.com/satioO/order-mgmt/pkg/logger"
)

type aggregateStore struct {
	log logger.Logger
	db  *dynamodb.Client
}

func NewAggregateStore(log logger.Logger, db *dynamodb.Client) *aggregateStore {
	return &aggregateStore{log, db}
}

func (a *aggregateStore) Load(ctx context.Context, aggregate es.Aggregate) error {
	return nil
}

func (a *aggregateStore) Save(ctx context.Context, aggregate es.Aggregate) error {
	return nil
}

func (a *aggregateStore) Exists(ctx context.Context, streamID string) error {
	return nil
}
