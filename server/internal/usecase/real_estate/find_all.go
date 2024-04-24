package realestate

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) FindAll(c context.Context, page dto.Page) ([]entity.RealEstate, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	return uc.repo.FindAll(ctx, page)
}
