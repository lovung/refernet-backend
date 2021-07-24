package userusecase

import (
	"context"
	"errors"
	"refernet/internal/dto"
	"refernet/internal/dto/userdto"
	"refernet/internal/ent"
)

// UpdateUserUseCase help to update user basic information
type UpdateUserUseCase struct {
	client *ent.Client
}

// Do executes the use case
func (u *UpdateUserUseCase) Do(ctx context.Context, input dto.Input) (dto.Output, error) {
	inputStruct, ok := input.(userdto.UpdateUserInput)
	if !ok {
		return nil, errors.New("wrong interface")
	}
	updateBuilder := u.client.User.Update()
	if inputStruct.Username != nil {
		updateBuilder.SetUsername(*inputStruct.Username)
	}
	if inputStruct.Email != nil {
		updateBuilder.SetEmail(*inputStruct.Email)
	}
	if inputStruct.Fullname != nil {
		updateBuilder.SetFullname(*inputStruct.Fullname)
	}
	if inputStruct.Phone != nil {
		updateBuilder.SetPhone(*inputStruct.Phone)
	}
	if inputStruct.Bio != nil {
		updateBuilder.SetBio(*inputStruct.Bio)
	}
	if inputStruct.Intro != nil {
		updateBuilder.SetIntro(*inputStruct.Intro)
	}
	if inputStruct.ProfilePictureURL != nil {
		updateBuilder.SetProfilePictureURL(*inputStruct.ProfilePictureURL)
	}

	_, err := updateBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
