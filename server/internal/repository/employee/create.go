package employee

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) Create(ctx context.Context, data dto.CreateEmployee) (*entity.Employee, error) {
	employee := &entity.Employee{
		Firstname:  data.Firstname,
		Surname:    data.Surname,
		Middlename: data.Middlename,
		Email:      data.Email,
		Password:   data.Password,
	}
	if err := r.WithContext(ctx).Create(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}
