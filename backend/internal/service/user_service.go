package service

import (
	"time"

	"my-coffee-log/internal/model"
	"my-coffee-log/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetCurrentUser(userID uint) (*model.User, error) {
	return s.userRepo.FindByID(userID)
}

type UpdateUserRequest struct {
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}

func (s *UserService) UpdateUser(userID uint, req UpdateUserRequest) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.AvatarURL != "" {
		user.AvatarURL = req.AvatarURL
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// ---- Onboarding ----

type OnboardingRequest struct {
	PreferredLogMode     string `json:"preferred_log_mode" binding:"omitempty,oneof=quick detailed"`
	PreferredCoffeeTypes string `json:"preferred_coffee_types"`
	Skip                  bool   `json:"skip"`
}

func (s *UserService) CompleteOnboarding(userID uint, req OnboardingRequest) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.OnboardingCompleted = true
	if req.PreferredLogMode != "" {
		user.PreferredLogMode = req.PreferredLogMode
	}
	if req.PreferredCoffeeTypes != "" {
		user.PreferredCoffeeTypes = req.PreferredCoffeeTypes
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) MarkFirstRecord(userID uint) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if user.FirstRecordAt == nil {
		now := time.Now()
		user.FirstRecordAt = &now
		return s.userRepo.Update(user)
	}

	return nil
}
