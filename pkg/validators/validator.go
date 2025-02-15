package validators

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ValidationRule defines a function type for validation rules.
type ValidationRule func(string) error

// ValidateWithPrefix applies validation rules to a value and prefixes any error messages.
func ValidateWithPrefix(prefix, value string, rules ...ValidationRule) error {
	for _, rule := range rules {
		if err := rule(value); err != nil {
			return fmt.Errorf("%s: %w", prefix, err)
		}
	}
	return nil
}

// Validate applies multiple validation rules to a value.
func Validate(value string, rules ...ValidationRule) error {
	for _, rule := range rules {
		if err := rule(value); err != nil {
			return err
		}
	}
	return nil
}

// Required checks if the field is non-empty.
func Required(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("field is required")
	}
	return nil
}

// MustContain ensures the value contains the specified substring.
func MustContain(substring string) ValidationRule {
	return func(value string) error {
		if !strings.Contains(value, substring) {
			return fmt.Errorf("must contain '%s'", substring)
		}
		return nil
	}
}

// MinLength ensures the value has at least the specified number of characters.
func MinLength(length int) ValidationRule {
	return func(value string) error {
		if len(value) < length {
			return fmt.Errorf("must be at least %d characters long", length)
		}
		return nil
	}
}

// ValidateStrongPassword ensures a password meets security requirements.
func ValidateStrongPassword(value string) error {
	if len(value) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if match, _ := regexp.MatchString(`[A-Z]`, value); !match {
		return errors.New("password must include at least one uppercase letter")
	}

	if match, _ := regexp.MatchString(`[a-z]`, value); !match {
		return errors.New("password must include at least one lowercase letter")
	}

	if match, _ := regexp.MatchString(`\d`, value); !match {
		return errors.New("password must include at least one digit")
	}

	if match, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, value); !match {
		return errors.New("password must include at least one special character")
	}

	return nil
}
