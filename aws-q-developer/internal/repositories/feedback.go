package repositories

import (
	"ai-compare-app/internal/models"
	"database/sql"
)

type FeedbackRepository interface {
	CreateFeedback(feedback *models.Feedback) error
}

type PostgresFeedbackRepository struct {
	db *sql.DB
}

func NewPostgresFeedbackRepository(db *sql.DB) *PostgresFeedbackRepository {
	return &PostgresFeedbackRepository{db: db}
}

func (r *PostgresFeedbackRepository) CreateFeedback(feedback *models.Feedback) error {
	// Implement the database insertion logic here
	return nil
}
