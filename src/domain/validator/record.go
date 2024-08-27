package validator

import (
	"errors"
	"fmt"
	"time"

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
		entity.LabelTrayStack |
		entity.TrayStacked
}

func Validate[T _validate](entity *T) (error bool, field FieldValidation) {
	validate := validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterValidation("notZeroTime", notZeroTime)

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
	for _, err := range validatorErrs {
		translatedErr := fmt.Errorf(err.Translate(trans))
		if err.Field() == "CreatedAt" {
			if err.Tag() == "required" || err.Tag() == "notZeroTime" {
				fieldErr = append(fieldErr, errors.New("created_at is a required field and must not be empty."))
				return true, fieldErr
			} else {
				fmt.Printf("Validation error on field '%s': %s\n", err.Field(), err.Error())
			}
		}
		fieldErr = append(fieldErr, translatedErr)
	}
	return true, fieldErr
}

// Custom validator to check if the field is of type time.Time and is not the zero value
func notZeroTime(fl validator.FieldLevel) bool {
	t, ok := fl.Field().Interface().(time.Time)
	return ok && !t.IsZero()
}
