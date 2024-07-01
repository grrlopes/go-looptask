package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

func FetchOneLabelTraySuccess(label entity.LabelAggSet) output {
	return output{
		"error":   nil,
		"message": label,
		"success": true,
	}
}

func FetchOneLabelTrayError(label entity.Tray, err error) output {
	return output{
		"error": err.Error(),
		"message": nil,
		"success": false,
	}
}
