package forms

type CreateTaskCommand struct {
	ID    string `json:"id"`
	Title string `json:"title" binding:"required"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	Completed bool `json:"completed"`
}

type UpdateTaskCommand struct {
	ID    string `json:"id"`
	Title string `json:"title" binding:"required"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	Completed bool `json:"completed"`
}
