package models

// User is models
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" firestore:"name" binding:"required,max=100"`
	Email    string `json:"email"  firestore:"email,unique" binding:"required,max=100"`
	Password string `json:"password"  firestore:"password" binding:"min=6"`
}

//LoginUser models
type LoginUser struct {
	Email    string `json:"email"  firestore:"email,unique" binding:"required,max=100"`
	Password string `json:"password"  firestore:"password" binding:"min=6"`
}
