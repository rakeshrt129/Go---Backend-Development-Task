package models

import "time"

// ---------- REQUEST MODELS ----------

// CreateUserRequest represents data coming from POST /users
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob"  validate:"required,datetime=2006-01-02"`
}

// UpdateUserRequest represents data coming from PUT /users/:id
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob"  validate:"required,datetime=2006-01-02"`
}

// ---------- RESPONSE MODELS ----------

// UserResponse is what API sends back to client
type UserResponse struct {
	ID   int32     `json:"id"`
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
	Age  int       `json:"age"` // calculated dynamically
}
