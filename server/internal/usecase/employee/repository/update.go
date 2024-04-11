package repository

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) Update(
	ctx context.Context,
	id uuid.UUID,
	data dto.UpdateEmployee,
) error {
	employee := &entity.Employee{ID: id}
	if err := r.WithContext(ctx).Model(employee).Updates(entity.Employee{
		Email:      data.Email,
		Firstname:  data.Firstname,
		Surname:    data.Surname,
		Middlename: data.Middlename,
	}).Error; err != nil {
		return err
	}
	return nil
}
