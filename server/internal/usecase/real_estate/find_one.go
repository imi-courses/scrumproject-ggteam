package realestate

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

func (uc *UseCase) FindOne(c context.Context, data entity.RealEstate) (*entity.RealEstate, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if uuid.Nil != data.ID {
		return uc.repo.FindOneById(ctx, data.ID)
	} else if uuid.Nil != data.ClientID {
		return uc.repo.FindOneByClientId(ctx, data.ClientID)
	}

	return nil, errors.New("record not found")
}
