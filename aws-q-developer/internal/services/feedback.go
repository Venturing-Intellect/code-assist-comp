package services

import (
	"ai-compare-app/internal/models"
	"ai-compare-app/internal/repositories"
	"ai-compare-app/pkg/utils"
	"fmt"
)

type FeedbackService interface {
	CreateFeedback(feedback *models.Feedback) error
}

type FeedbackServiceImpl struct {
	repo repositories.FeedbackRepository
}

func NewFeedbackService(repo repositories.FeedbackRepository) *FeedbackServiceImpl {
	return &FeedbackServiceImpl{repo: repo}
}

func (s *FeedbackServiceImpl) CreateFeedback(feedback *models.Feedback) error {
	if !utils.IsValidEmail(feedback.Email) {
		return fmt.Errorf("invalid email format")
	}

	return s.repo.CreateFeedback(feedback)
}
