package realestate

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) Create(
	c context.Context,
	data dto.CreateRealEstate,
) (*entity.RealEstate, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.repo.Create(ctx, data)
}
