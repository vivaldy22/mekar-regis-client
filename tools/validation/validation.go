package validation

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidateInputNotEmpty(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("make sure input not empty")
		case 0:
			return errors.New("make sure input not a zero")
		case nil:
			return errors.New("make sure input not a nil")
		}
	}
	return nil
}

func ValidateUUID(data ...interface{}) error {
	for _, value := range data {
		if len(value.(string)) != 36 {
			return fmt.Errorf("value '%v' is not a valid UUID", value)
		}
	}
	return nil
}

func ValidateDate(date string) error {
	re := regexp.MustCompile(`((19|20)\d\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])`)
	if !re.MatchString(date) {
		return errors.New("date format must be yyyy-mm-dd")
	}
	return nil
}

func ValidateKTP(ktp string) error {
	if len(ktp) != 16 {
		return errors.New("ktp length must be 16 digits")
	}
	return nil
}
