package user

import (
	"fmt"
	"strconv"

	"github.com/shaheen-ct/go-hexagonal-architecture/api/graph/models"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/entity"
)

type DTO struct {
}

func (t DTO) toUserEntity(user models.UserInput) (*entity.User, error) {
	rs, err := t.toMeetupsEntity(user.Meetups)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Username: user.Username,
		Email:    user.Email,
		Meetups:  rs,
	}, nil
}

func (t DTO) toUserModel(user *entity.User) *models.User {
	return &models.User{
		ID:       fmt.Sprintf("%v", user.ID),
		Username: user.Username,
		Email:    user.Email,
		Meetups:  t.toMeetupsModel(user.Meetups),
	}
}

// generic
func (t DTO) toMeetupsEntity(meetups []*models.MeetupInput) ([]*entity.Meetup, error) {
	rs := make([]*entity.Meetup, len(meetups))
	for idx, meetup := range meetups {
		e, err := t.toMeetupEntity(meetup)
		if err != nil {
			return nil, err
		}
		rs[idx] = e
	}
	return rs, nil
}

func (t DTO) toMeetupEntity(meetup *models.MeetupInput) (*entity.Meetup, error) {
	var id int64
	if meetup.ID != nil {
		id, err := strconv.ParseInt(*meetup.ID, 0, 10)
		fmt.Println("lolo", id, err)
		if err != nil {
			return nil, err
		}
	}
	return &entity.Meetup{
		ID:          id,
		Name:        meetup.Name,
		Description: meetup.Description,
	}, nil
}

func (t DTO) toMeetupsModel(meetups []*entity.Meetup) []*models.Meetup {
	rs := make([]*models.Meetup, len(meetups))
	for idx, meetup := range meetups {
		rs[idx] = t.toMeetupModel(meetup)
	}
	return rs
}

func (t DTO) toMeetupModel(meetup *entity.Meetup) *models.Meetup {
	return &models.Meetup{
		ID:          fmt.Sprint(meetup.ID),
		Name:        meetup.Name,
		Description: meetup.Description,
	}
}
