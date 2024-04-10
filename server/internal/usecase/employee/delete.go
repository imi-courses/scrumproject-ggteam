package employee

import (
	"context"

	"github.com/google/uuid"
)

func (uc *UseCase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	return uc.repo.Delete(ctx, id)
}
