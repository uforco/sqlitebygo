package types

type UserSchema struct {
	Name     string `validate:"required"`
	Age      int    `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
