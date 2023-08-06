package common

var validations = map[string]string{
	"_default":     "The %s is not valid",
	"email":        "The %s field must be a valid email",
	"required":     "The %s field is required",
	"required_if":  "The {field} field is required",
	"alpha":        "The {field} field may only contain alphabetic characters",
	"alpha_num":    "The {field} field may only contain alpha-numeric characters",
	"alpha_dash":   "The {field} field may contain alpha-numeric characters as well as dashes and underscores",
	"alpha_spaces": "The {field} field may only contain alphabetic characters as well as spaces",
	"between":      "The {field} field must be between 0:{min} and 1:{max}",
	"confirmed":    "The {field} field confirmation does not match",
	"digits":       "The {field} field must be numeric and exactly contain 0:{length} digits",
	"integer":      "The {field} field must be an integer",
	"length":       "The {field} field must be 0:{length} long",
	"max_value":    "The {field} field must be 0:{max} or less",
	"max":          "The {field} field may not be greater than 0:{length} characters",
	"min_value":    "The {field} field must be 0:{min} or more",
	"min":          "The {field} field must be at least 0:{length} characters",
	"numeric":      "The {field} field may only contain numeric characters",
	"one_of":       "The {field} field is not a valid value",
	"regex":        "The {field} field format is invalid",
}
