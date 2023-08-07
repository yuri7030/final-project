package inputs

type SurveyCreatingInput struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
}