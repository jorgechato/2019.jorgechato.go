package subscription

type (
	Subscritor struct {
		EmailAddress string `json:"email_address" binding:"required"`
		Status       string `json:"status_if_new"`
	}
)
