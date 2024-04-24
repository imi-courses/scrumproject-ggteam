package client

import (
	"context"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
)

func (uc *UseCase) Update(c context.Context, uuid uuid.UUID, data dto.UpdateClient) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	return uc.repo.Update(ctx, uuid, data)
}
