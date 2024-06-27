package presenters

import "github.com/grrlopes/go-looptask/src/domain/entity"

func FetchOneUserSuccess(user entity.Users) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"id":      user.ID.Hex(),
			"name":    user.Name,
			"surname": user.Surname,
			"email":   user.Email,
		},
		"success": true,
	}
}

func FetchOneUserError(user entity.Users) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"name":    user.Name,
			"surname": user.Surname,
			"email":   user.Email,
		},
		"success": false,
	}
}
