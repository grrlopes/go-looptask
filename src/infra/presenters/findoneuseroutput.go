package presenters

import "github.com/grrlopes/go-looptask/src/domain/entity"

func FetchOneUserSuccess(user entity.Users) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"id":     user.ID.Hex(),
			"author": user.Author,
			"email":  user.Email,
		},
		"success": true,
	}
}

func FetchOneUserError(user entity.Users) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"author": user.Author,
			"email":  user.Email,
		},
		"success": false,
	}
}
