package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/shaheen-ct/go-hexagonal-architecture/api/graph/models"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/graph/generated"
)

var meetups = []*models.Meetup{
	{
		Name:        "meet1",
		Description: "cool",
		ID:          "1",
	},
	{
		Name:        "meet2",
		Description: "cool2",
		ID:          "2",
	},
}

var users = []*models.User{
	{
		Username: "meet1",
		Email:    "cool@gmail.com",
		ID:       "1",
	},
	{
		Username: "meet2",
		Email:    "cool2@gmail.com",
		ID:       "2",
	},
}

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return users[0], nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

type meetupResolver struct{ *Resolver }
