package scope

import (
	"context"
	"github.com/evenyosua18/tracing"
)

func (u *UsecaseScope) UpdateScope(ctx context.Context, in, filter interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, in, filter)
	defer tracing.Close(sp)

	return u.scope.UpdateScope(tracing.Context(sp), in, filter)
}
