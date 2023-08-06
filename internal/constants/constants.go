package constants

type AnswerType string

const (
	BasicAuthUsername                   = "user"
	BasicAuthPassword                   = "12345678"
	JwtSecretKey                        = "QhNYaTIWvlvT0lSZxCEOSbCyd9pJMFz2Wtypgriv59U="
	AnswerTypeNumber         AnswerType = "number"
	AnswerTypeString         AnswerType = "string"
	AnswerTypeDropdown       AnswerType = "dropdown"
	AnswerTypeMultipleChoice AnswerType = "multiple_choice"
	PasswordHashLength                  = 12
)
