package handlers

import (
	"context"
	"pet/internal/userService"
	"pet/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, item := range allUsers {
		user := users.User{
			Id:       &item.ID,
			Email:    &item.Email,
			Password: &item.Password,
		}

		response = append(response, user)
	}

	return response, nil
}

func (h UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	user := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.PostUsers(user)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h UserHandler) GetUsersId(ctx context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	user, err := h.Service.GetUsersId(request.Id)
	if err != nil {
		return nil, err
	}
	response := users.GetUsersId200JSONResponse{
		Id:       &user.ID,
		Email:    &user.Email,
		Password: &user.Password,
	}

	return response, nil
}

func (h UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	user, err := h.Service.GetUsersId(request.Id)
	if err != nil {
		return nil, err
	}

	if userRequest.Email != nil {
		user.Email = *userRequest.Email
	}

	if userRequest.Password != nil {
		user.Password = *userRequest.Password
	}

	if userRequest.Email == nil && userRequest.Password == nil {
		response := users.PatchUsersId400Response{}
		return response, nil
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &user.ID,
		Email:    &user.Email,
		Password: &user.Password,
	}

	return response, nil

}

func (h UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.Service.DeleteUsersId(request.Id)
	if err != nil {
		return nil, err
	}

	response := users.DeleteUsersId204Response{}

	return response, nil
}
