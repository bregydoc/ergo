package ergo

// ErrorSummary ...
type ErrorSummary struct {
	ID                 string `json:"id"`
	Code               uint64 `json:"code"`
	Explain            string `json:"explain"`
	EnglishUserMessage string `json:"english_user_message"`
	Type               string `json:"type"`
	Image              string `json:"image"`
	ActionLink         string `json:"action_link"`
	ActionMessage      string `json:"action_message"`
}
