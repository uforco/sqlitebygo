package errorhandle

import "github.com/go-playground/validator/v10"

type ValidationErrorType struct {
	status  string
	Message string
}

func ValidationError(errs validator.ValidationErrors) ValidationErrorType {
	var errorMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMsgs = append(errorMsgs, err.Field()+" is required")
		case "email":
			errorMsgs = append(errorMsgs, err.Field()+" must be a valid email")
		case "min":
			errorMsgs = append(errorMsgs, err.Field()+" must be at least "+err.Param()+" characters")
		case "max":
			errorMsgs = append(errorMsgs, err.Field()+" must be at most "+err.Param()+" characters")
		default:
			errorMsgs = append(errorMsgs, "Invalid input for "+err.Field())
		}
	}

	return ValidationErrorType{
		status:  "error",
		Message: errorMsgs[0],
	}
}
