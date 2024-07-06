package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

func CreateLabelSuccess(label entity.Tray) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"trays": label,
		},
		"success": true,
	}
}

func CreateLabelError(label entity.Tray) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"trays": label,
		},
		"success": false,
	}
}
