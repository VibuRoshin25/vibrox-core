package payload

// SignInPayload authentication payload
type SignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GenerateTokenPayload token generation payload
type GenerateTokenPayload struct {
	Email string `json:"email"`
	Id    uint   `json:"userId"`
}

// ValidateTokenPayload token validation payload
type ValidateTokenPayload struct {
	Token string `json:"token"`
}
