package validator

import (
	"github.com/akiradomi/workspace/go-practice/back/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TaskValidatorInterface interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() TaskValidatorInterface {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("タイトルの入力は必須です"),
			validation.RuneLength(1, 100).Error("100文字以内で入力してください"),
		),
	)
}
