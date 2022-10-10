package resolvers

import (
	"context"

	"github.com/shaheen-ct/go-hexagonal-architecture/api/graph/models"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/graph/generated"
)

// CreateMeetUp is the resolver for the createMeetUp field.
func (r *mutationResolver) CreateMeetUp(ctx context.Context, input models.MeetupInput) (*models.Meetup, error) {
	return meetups[0], nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	return r.Domain.User.Create(ctx, input)
}
func (r *mutationResolver) GetUser(ctx context.Context, input string) (*models.User, error) {
	return r.Domain.User.Get(ctx, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (*string, error) {
	return nil, r.Domain.User.Delete(ctx, *id)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	return r.Domain.User.Update(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
