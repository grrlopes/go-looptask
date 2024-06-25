package presenters

func AuthError() output {
	return output{
		"error":   "Unauthorized",
		"message": "Unable to validate credencial.",
		"success": false,
	}
}
