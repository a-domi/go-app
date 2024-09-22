package validator

import (
	"github.com/akiradomi/workspace/go-practice/back/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserValidatorInterface interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() UserValidatorInterface {
	return &userValidator{}
}

func (tv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("メールアドレスの入力は必須です"),
			is.Email.Error("メールアドレスを入力して下さい"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードの入力は必須です"),
			validation.RuneLength(8, 20).Error("パスワードは8文字以上20字以内で入力してください"),
		),
	)
}
