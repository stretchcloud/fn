package mqs

import (
	"context"

	"github.com/fnproject/fn/api/models"
)

type Mock struct {
	FakeApp   *models.App
	Apps      []*models.App
	FakeRoute *models.Route
	Routes    []*models.Route
}

func (mock *Mock) Push(context.Context, *models.Call) (*models.Call, error) {
	return nil, nil
}

func (mock *Mock) Reserve(context.Context) (*models.Call, error) {
	return nil, nil
}

func (mock *Mock) Delete(context.Context, *models.Call) error {
	return nil
}
