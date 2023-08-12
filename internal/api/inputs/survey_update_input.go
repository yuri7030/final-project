package inputs

type SurveyUpdatingInput struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
}