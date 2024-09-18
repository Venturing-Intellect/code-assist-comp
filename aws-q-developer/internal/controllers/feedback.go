package controllers

import (
	"ai-compare-app/internal/models"
	"ai-compare-app/internal/services"
	"encoding/json"
	"net/http"
)

type FeedbackController struct {
	service services.FeedbackService
}

func NewFeedbackController(service services.FeedbackService) *FeedbackController {
	return &FeedbackController{service: service}
}

func (c *FeedbackController) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback
	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateFeedback(&feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := map[string]string{
		"status":  "success",
		"message": "Feedback message received",
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Set the status code
	w.WriteHeader(http.StatusOK)

	// Encode and send the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
