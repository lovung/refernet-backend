package userusecase

import (
	"context"
	"errors"
	"refernet/internal/dto"
	"refernet/internal/dto/userdto"
	"refernet/internal/ent"
)

type CreateUserUseCase struct {
	client *ent.Client
}

func (u *CreateUserUseCase) Do(ctx context.Context, input dto.Input) (dto.Output, error) {
	inputStruct, ok := input.(userdto.CreateUserInput)
	if !ok {
		return nil, errors.New("wrong interface")
	}
	user, err := u.client.User.
		Create().
		SetUsername(inputStruct.Username).
		SetFullname(inputStruct.Fullname).
		SetEmail(inputStruct.Email).
		SetPhone(inputStruct.Phone).
		SetBio(inputStruct.Bio).
		SetIntro(inputStruct.Intro).
		SetProfilePictureURL(inputStruct.ProfilePictureURL).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
