package service

import (
	"context"
	"time"

	"user-age-api/db/sqlc"
	"user-age-api/internal/models"
)

// UserService handles business logic related to users
type UserService struct {
	repo *sqlc.Queries
}

// NewUserService creates a new UserService
func NewUserService(repo *sqlc.Queries) *UserService {
	return &UserService{
		repo: repo,
	}
}

// -------------------- CREATE USER --------------------
func (s *UserService) CreateUser(ctx context.Context, name string, dob time.Time) (models.UserResponse, error) {

	user, err := s.repo.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return convertDBUserToResponse(user), nil
}

// -------------------- GET USER BY ID --------------------
func (s *UserService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return convertDBUserToResponse(user), nil
}

// -------------------- LIST USERS --------------------
func (s *UserService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.UserResponse
	for _, user := range users {
		response = append(response, convertDBUserToResponse(user))
	}

	return response, nil
}

// -------------------- UPDATE USER --------------------
func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) (models.UserResponse, error) {

	user, err := s.repo.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return convertDBUserToResponse(user), nil
}

// -------------------- DELETE USER --------------------
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}


// -------------------- HELPER FUNCTIONS --------------------

// calculateAge calculates age from date of birth
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Adjust if birthday has not occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

// convertDBUserToResponse converts SQLC user to API response
func convertDBUserToResponse(user sqlc.User) models.UserResponse {
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob,
		Age:  calculateAge(user.Dob),
	}
}
