package employee

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) FindAll(c context.Context, page dto.Page) ([]entity.Employee, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	employees, err := uc.repo.FindAll(ctx, page)
	if err != nil {
		return nil, err
	}
	return employees, nil
}
