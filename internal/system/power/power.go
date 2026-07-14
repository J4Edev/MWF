package power

import "context"

type Manager interface {
	CurrentPlan(context.Context) (string, error)
	SetPlan(context.Context, string) error
}
