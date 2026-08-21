package maternalvisits

import (
	"context"
	"fmt"
	"time"
)

type VisitPolicy struct {
	Timeout         time.Duration
	UseServiceClock bool
}

func (p VisitPolicy) Scope(ctx context.Context) (context.Context, context.CancelFunc, error) {
	if ctx == nil {
		return nil, nil, fmt.Errorf("visit: nil context")
	}
	parent := ctx
	if p.UseServiceClock {
		parent = context.Background()
	}
	callCtx, cancel := context.WithTimeout(parent, p.Timeout)
	return callCtx, cancel, nil
}
