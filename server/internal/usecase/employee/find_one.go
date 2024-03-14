package employee

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) FindOne(c context.Context, data entity.Employee) (*entity.Employee, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if len(data.Email) != 0 {
		return uc.repo.FindOneByEmail(ctx, data.Email)
	} else if uuid.Nil != data.ID {
		return uc.repo.FindOneById(ctx, data.ID)
	} else if len(data.RefreshToken) != 0 {
		return uc.repo.FindOneByRefreshToken(ctx, data.RefreshToken)
	}

	return nil, errors.New("record not found")
}
