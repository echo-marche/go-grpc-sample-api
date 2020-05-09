package validations

import (
	presencePb "github.com/echo-marche/hack-tech-tips-api/proto/pb/presence"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateUserRegistrationRequest(r *presencePb.UserRegistrationRequest) error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(&r.Password, validation.Required),
		// Streetは空を許容せず、5から50までの長さ
		// validation.Field(&r.Street, validation.Required, validation.Length(5, 50)),
		// // Cityは空を許容せず、5から50までの長さ
		// validation.Field(&r.City, validation.Required, validation.Length(5, 50)),
		// // Stateは空を許容せず、大文字2つの文字列
		// validation.Field(&r.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// // Stateは空を許容せず、数字5つの文字列
		// validation.Field(&r.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}
