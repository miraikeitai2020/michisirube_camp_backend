package resolver

import (
	"context"

	"github.com/miraikeitai2020/michisirube_camp_backend/bff/pkg/bff"
	"github.com/miraikeitai2020/michisirube_camp_backend/bff/pkg/server/model"
)

var(
	CITY_ID			= 1
	CITY_NAME		= "函館駅"
	CITY_LATITUDE	= 41.773746
	CITY_LONGITUDE	= 140.726399
	CITY_TIME		= 10
)

func (r *mutationResolver) Evaluat(ctx context.Context, emotion int, value int, locationID int) (bool, error) {

	return true, nil
}

func (r *queryResolver) City(ctx context.Context, emotion int, evaluat float64, location float64, time int) (*model.City, error) {
	return &model.City{
		Location: &model.Location{
			ID:	&CITY_ID,
			Name: &CITY_NAME,
			Latitude: &CITY_LATITUDE,
			Longitude: &CITY_LONGITUDE,
		},
		Time: CITY_TIME,
	}, nil
}

// Mutation returns bff.MutationResolver implementation.
func (r *Resolver) Mutation() bff.MutationResolver { return &mutationResolver{r} }

// Query returns bff.QueryResolver implementation.
func (r *Resolver) Query() bff.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }