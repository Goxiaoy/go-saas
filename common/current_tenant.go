package common

import "context"

type CancelFunc func() context.Context
type CurrentTenant interface {
	IsAvailable(ctx context.Context) bool
	Id(ctx context.Context) string
	//change to one tenant and change back when cancel called
	Change(ctx context.Context, id string, name string) (context.Context, CancelFunc)
}
