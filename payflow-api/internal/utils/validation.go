package utils

import (
	"fmt"
	"strings"
)

func ValidateRequired(value string, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}

	return nil
}