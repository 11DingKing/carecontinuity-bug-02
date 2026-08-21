package maternalvisits

import (
	"context"
	"fmt"
	"time"
)

type Coordinator struct{ policy VisitPolicy }

func NewCoordinator() *Coordinator {
	return &Coordinator{policy: VisitPolicy{Timeout: time.Second, UseServiceClock: true}}
}

func (c *Coordinator) Run(ctx context.Context, operation func(context.Context) error) error {
	callCtx, cancel, err := c.policy.Scope(ctx)
	if err != nil {
		return err
	}
	defer cancel()
	if err := operation(callCtx); err != nil {
		return fmt.Errorf("maternal visit deadline operation: %w", err)
	}
	return nil
}
