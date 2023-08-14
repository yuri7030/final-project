package inputs

type OptionsAddingInput struct {
	OptionTexts []string `json:"optionTexts" binding:"required"`
}
