package GraphQL_Service

import (
	"context"
	"fmt"

	"github.com/boltdb/bolt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) People(ctx context.Context, id string) (*People, error) {
	d, err := bolt.Open("data/data.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	return GetPeopleByID(id, d)
}

func (r *queryResolver) Films(ctx context.Context, id string) (*Films, error) {
	d, err := bolt.Open("data/data.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	return GetFilmByID(id, d)
}

func (r *queryResolver) Planets(ctx context.Context, id string) (*Planets, error) {
	d, err := bolt.Open("data/data.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	return GetPlanetByID(id, d)
}

func (r *queryResolver) Starships(ctx context.Context, id string) (*Starships, error) {
	d, err := bolt.Open("data/data.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	return GetStarshipByID(id, d)
}

func (r *queryResolver) Species(ctx context.Context, id string) (*Species, error) {
	d, err := bolt.Open("data/data.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	return GetSpeciesByID(id, d)
}

func (r *queryResolver) Vehicles(ctx context.Context, id string) (*Vehicles, error) {
	d, err := bolt.Open("data/data.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	return GetVehicleByID(id, d)
}
