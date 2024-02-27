package controller

import (
	"net/http"
	"strconv"

	"github.com/Andhika-GIT/Go-REST-Event-Management/helper"
	"github.com/Andhika-GIT/Go-REST-Event-Management/model/web"
	"github.com/Andhika-GIT/Go-REST-Event-Management/service"
	"github.com/go-chi/chi/v5"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		service: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {

	createUserRequest := &web.UserCreateRequest{}

	helper.ReadJsonRequestBody(request, createUserRequest)

	userResponse, err := controller.service.Create(request.Context(), *createUserRequest)

	if err != nil {
		webResponse := &web.WebResponseError{
			Code:    400,
			Message: err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return
	}

	webResponse := &web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {

	updateUserRequest := &web.UserUpdateRequest{}

	helper.ReadJsonRequestBody(request, updateUserRequest)

	userId := chi.URLParam(request, "userId")

	id, _ := strconv.Atoi(userId)

	updateUserRequest.Id = int32(id)

	userResponse, err := controller.service.Update(request.Context(), *updateUserRequest)

	if err != nil {
		webResponse := &web.WebResponseError{
			Code:    400,
			Message: err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return
	}

	webResponse := &web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "userId")

	id, _ := strconv.Atoi(userId)

	err := controller.service.Delete(request.Context(), int32(id))

	if err != nil {
		webResponse := &web.WebResponseError{
			Code:    400,
			Message: err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return
	}

	webResponse := &web.WebResponse{
		Code:   200,
		Status: "Success",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "userId")

	id, _ := strconv.Atoi(userId)

	userResponse, err := controller.service.FindById(request.Context(), int32(id))

	if err != nil {
		webResponse := web.WebResponseError{
			Code:    400,
			Message: err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
