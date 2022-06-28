package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"example/server/commons"
	"example/server/repositories"
	"log"
)

func (r *mutationResolver) StoreDataEntry(ctx context.Context, input model.InputDataEntry) (*model.DataEntry, error) {
	rep, err := repositories.GetRepository(commons.RepositoryTypeDB)
	if err != nil {
		log.Fatal(ctx, err, "Unable to get repository")
	}
	rep.Save(ctx, input)

	var entry = contains(input)
	if entry == nil {
		newEntry := &model.DataEntry{
			ID: input.ID,
		}
		entry = newEntry
	}
	entry.Title = input.Title
	entry.Content = input.Content
	entry.Views = input.Views
	entry.Timestamp = input.Timestamp

	return entry, nil
}

func (r *queryResolver) GetDataEntries(ctx context.Context) ([]*model.DataEntry, error) {
	rep, err := repositories.GetRepository(commons.RepositoryTypeDB)
	if err != nil {
		log.Fatal(ctx, err, "Unable to get repository")
	}

	var retValur []*model.DataEntry
	retValur, _ = rep.Get(ctx)
	return retValur, nil
}

func (r *queryResolver) GetDataEntry(ctx context.Context, id string) (*model.DataEntry, error) {
	rep, err := repositories.GetRepository(commons.RepositoryTypeDB)
	if err != nil {
		log.Fatal(ctx, err, "Unable to get repository")
	}

	var dataEntry *model.DataEntry
	dataEntry, _ = rep.GetById(ctx, id)
	return dataEntry, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var entries []*model.DataEntry

func contains(dataEntry model.InputDataEntry) *model.DataEntry {
	for _, a := range entries {
		if a.ID == dataEntry.ID {
			return a
		}
	}
	return nil
}
