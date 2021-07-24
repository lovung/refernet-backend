package main

import (
	"refernet/internal/ent/proto/entpb"
	"testing"
)

func TestUserProto(t *testing.T) {
	user := entpb.User{
		Username: "rotemtam",
		Email:    "rotemtam@example.com",
	}
	if user.GetUsername() != "rotemtam" {
		t.Fatal("expected user name to be rotemtam")
	}
	if user.GetEmail() != "rotemtam@example.com" {
		t.Fatal("expected email address to be rotemtam@example.com")
	}
}
