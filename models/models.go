package models

type Task struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	IsDone    bool   `json:"done"`
	CreatedAt int64  `json:"createdAt"`
}
