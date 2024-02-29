package employee

import (
	"context"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (r *Repository) FindOne(ctx context.Context, p dto.FindOneEmployee) (*entity.Employee, error) {
	var employee *entity.Employee
	if err := r.WithContext(ctx).Where(&entity.Employee{ID: p.ID, Email: p.Email}).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}
