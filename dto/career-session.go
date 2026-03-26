package dto

type CareerSessionCreateRequest struct {
	CareerID string `json:"career_id" binding:"required,uuid"`
}

type CareerSessionResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CareerID  string `json:"career_id"`
	Status    string `json:"status"`
	StartedAt string `json:"started_at"`
}
