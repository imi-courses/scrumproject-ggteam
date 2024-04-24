package client

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) FindOne(c context.Context, data entity.Client) (*entity.Client, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if len(data.Email) != 0 {
		return uc.repo.FindOneByEmail(ctx, data.Email)
	} else if uuid.Nil != data.ID {
		return uc.repo.FindOneById(ctx, data.ID)
	} else if len(data.Phone) != 0 {
		return uc.repo.FindOneByPhone(ctx, data.Phone)
	}

	return nil, errors.New("record not found")
}
