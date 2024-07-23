package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

type FindAllOutput struct {
	TotalRows int              `json:"total_rows"`
	Offset    int              `json:"offset"`
	Data      []entity.Labeled `json:"data"`
}

type FieldValidation struct {
	Error   string  `json:"error"`
	Message []error `json:"message"`
}

type _validate interface {
	entity.Labeled |
		entity.Users |
		entity.Tray |
		entity.UserId |
		entity.TrayId |
		entity.LabelTrayStack
}

func Validate[T _validate](entity *T) (error bool, field FieldValidation) {
	validate := validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(entity)
	checked, errs := handleError(err, trans)

	erros := FieldValidation{
		Error:   "Field not valid",
		Message: errs,
	}

	return checked, erros
}

func handleError(err error, trans ut.Translator) (checked bool, fieldErr []error) {
	if err == nil {
		return false, nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		fieldErr = append(fieldErr, translatedErr)
	}
	return true, fieldErr
}
