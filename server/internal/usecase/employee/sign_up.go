package employee

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) Create(c context.Context, data dto.CreateEmployee) (*entity.Employee, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	employee, err := uc.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return employee, nil
}
