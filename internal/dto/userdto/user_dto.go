package userdto

import (
	"strings"
)

// CreateUserInput represent input when creating a new user
type CreateUserInput struct {
	Username          string `json:"username,omitempty"`
	Fullname          string `json:"fullname,omitempty"`
	Email             string `json:"email,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Bio               string `json:"bio,omitempty"`
	Intro             string `json:"intro,omitempty"`
	ProfilePictureURL string `json:"profile_picture_url,omitempty"`
}

// Validate checks if the input is valid
func (i CreateUserInput) Validate() error {
	return nil
}

func (i CreateUserInput) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString("username=")
	builder.WriteString(i.Username)
	builder.WriteString(", fullname=")
	builder.WriteString(i.Fullname)
	builder.WriteString(", email=")
	builder.WriteString(i.Email)
	builder.WriteString(", phone=")
	builder.WriteString(i.Phone)
	builder.WriteString(", bio=")
	builder.WriteString(i.Bio)
	builder.WriteString(", intro=")
	builder.WriteString(i.Intro)
	builder.WriteString(", profile_picture_url=")
	builder.WriteString(i.ProfilePictureURL)
	builder.WriteByte(')')
	return builder.String()
}

// UpdateUserInput represent input when creating a new user
type UpdateUserInput struct {
	Username          *string `json:"username,omitempty"`
	Fullname          *string `json:"fullname,omitempty"`
	Email             *string `json:"email,omitempty"`
	Phone             *string `json:"phone,omitempty"`
	Bio               *string `json:"bio,omitempty"`
	Intro             *string `json:"intro,omitempty"`
	ProfilePictureURL *string `json:"profile_picture_url,omitempty"`
}

func (i UpdateUserInput) String() string {
	return ""
}

// Validate checks if the input is valid
func (i UpdateUserInput) Validate() error {
	return nil
}
