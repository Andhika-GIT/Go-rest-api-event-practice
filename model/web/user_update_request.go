package web

type UserUpdateRequest struct {
	Id    string `validate:"required" json:"id"`
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required" json:"email"`
}
