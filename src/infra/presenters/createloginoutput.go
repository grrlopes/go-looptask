package presenters

import "github.com/grrlopes/go-looptask/src/domain/entity"

func CreateLoginSuccess(user entity.Users, result entity.MongoResul) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"id":    result.ID,
			"email": user.Email,
		},
		"success": true,
	}
}

func CreateLoginError(user entity.Users, result entity.MongoResul) output {
	return output{
		"error": result.Error,
		"message": map[string]interface{}{
			"author": user.Author,
			"email":  user.Email,
		},
		"success": false,
	}
}
