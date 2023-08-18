package answer_type_enums

const (
	TextAnswer int = 1
	RadioAnswer int = 2
	CheckboxAnswer int = 3
)

var AnswerTypes = map[int]string{
	TextAnswer: "Text",
	RadioAnswer: "Radio",
	CheckboxAnswer: "Checkbox",
}
