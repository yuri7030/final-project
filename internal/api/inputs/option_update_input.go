package inputs

type OptionUpdatingInput struct {
	OptionText string `json:"optionText" binding:"required"`
}
