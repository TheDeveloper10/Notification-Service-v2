package dto

import (
	"net/mail"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type NotificationTarget struct {
	Email                *string           `json:"email"`
	PhoneNumber          *string           `json:"phoneNumber"`
	FCMRegistrationToken *string           `json:"fcmRegistrationToken"`
	Placeholders         map[string]string `json:"placeholders"`
}

func (nt *NotificationTarget) Validate() error {
	if nt.Email == nil && nt.PhoneNumber == nil && nt.FCMRegistrationToken == nil {
		return fiber.NewError(fiber.StatusBadRequest, "You must provide an Email, Phone Number or FCM Registration Token")
	}

	if nt.Email != nil {
		if _, err := mail.ParseAddress(*nt.Email); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid email")
		}
	}

	if nt.PhoneNumber != nil {
		pattern := `^\+?[1-9]\d{1,14}$`
		regex := regexp.MustCompile(pattern)

		if !regex.MatchString(*nt.PhoneNumber) {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid phone number")
		}
	}

	// if nt.FCMRegistrationToken != nil {
	// nothing
	// }

	return nil
}
