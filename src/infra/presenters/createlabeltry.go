package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

func CreateLabelTraySuccess(label entity.Labeled, id string) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"id":    id,
			"trays": label.Trays,
		},
		"success": true,
	}
}

func CreateLabelTrayError(label entity.Labeled) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"trays": label.Trays,
		},
		"success": false,
	}
}
