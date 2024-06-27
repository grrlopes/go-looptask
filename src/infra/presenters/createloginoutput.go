package presenters

import "github.com/grrlopes/go-looptask/src/domain/entity"

func CreateLoginSuccess(user entity.Users, result entity.MongoResul) output {
	return output{
		"error": nil,
		"message": map[string]interface{}{
			"id":      result.ID,
			"name":    user.Name,
			"surname": user.Surname,
			"email":   user.Email,
		},
		"success": true,
	}
}

func CreateLoginError(user entity.Users, result entity.MongoResul) output {
	return output{
		"error": result.Error,
		"message": map[string]interface{}{
			"name":    user.Name,
			"surname": user.Surname,
			"email":   user.Email,
		},
		"success": false,
	}
}
