package scope

import (
	"context"
	"github.com/evenyosua18/tracing"
)

func (u *UsecaseScope) DeleteScope(ctx context.Context, filter interface{}) error {
	// tracing
	sp := tracing.StartChild(ctx, filter)
	defer tracing.Close(sp)

	return u.scope.DeleteScope(tracing.Context(sp), filter)
}
