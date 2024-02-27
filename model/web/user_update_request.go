package web

type UserUpdateRequest struct {
	Id    int32  `validate:"required" json:"id"`
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required" json:"email"`
}
