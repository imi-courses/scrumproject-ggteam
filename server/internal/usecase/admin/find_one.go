package admin

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) FindOne(c context.Context, p dto.FindOneAdmin) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.FindOne(ctx, p)
	if err != nil {
		return nil, err
	}
	return user, nil
}
