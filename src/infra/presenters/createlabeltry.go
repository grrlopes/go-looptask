package presenters

import "github.com/grrlopes/go-looptask/src/domain/entity"

func CreateLabelTraySuccess(label entity.Labeled) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"labeled": label.Trays,
		},
		"success": true,
	}
}

func CreateLabelTrayError(label entity.Labeled) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"labeled": label.Trays,
		},
		"success": false,
	}
}
