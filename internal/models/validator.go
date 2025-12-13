package models

import "github.com/go-playground/validator/v10"

// Validate is a single validator instance used everywhere
var Validate = validator.New()
