package telegram

type ValidateOptions struct {
	ExpiresIn int // Expiration in seconds
}

type TeleAuthData struct {
	QueryID  string       `json:"query_id"`
	User     UserTeleAuth `json:"user"`
	AuthDate string       `json:"auth_date"`
	Hash     string       `json:"hash"`
}

type UserTeleAuth struct {
	ID              int64  `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	LanguageCode    string `json:"language_code"`
	IsPremium       bool   `json:"is_premium"`
	AllowsWriteToPm bool   `json:"allows_write_to_pm"`
	PhotoURL        string `json:"photo_url"`
}
