package controller

import "refernet/internal/dto"

type SourceGenController struct {
	URL    string
	Method string
	Input  dto.Input
	Output dto.Output
}
