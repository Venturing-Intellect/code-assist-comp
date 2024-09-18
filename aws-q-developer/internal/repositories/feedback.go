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
	query := `INSERT INTO feedbacks (email, message) VALUES ($1, $2)`
	_, err := r.db.Exec(query, feedback.Email, feedback.Message)
	if err != nil {
		return err
	}
	return nil
}
