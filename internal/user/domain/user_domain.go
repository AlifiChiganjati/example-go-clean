package domain

type (
	User struct {
		Id        string `json:"id"`
		UserName  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
