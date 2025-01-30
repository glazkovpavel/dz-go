package verify

type EmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyResponse struct {
	VerificationPassed bool `json:"verification_passed"`
}
