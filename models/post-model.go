package models

import "time"

// Post is ...
type Post struct {
	ID        string        `json:"id" firestore:"ID"`
	Title     string        `json:"title" firestore:"title" binding:"required"`
	Text      string        `json:"text" firestore:"text" binding:"required"`
	Date      time.Time     `json:"date" firestore:"date"`
	Price     int64         `json:"price" firestore:"price" binding:"required,numeric"`
	Authors   []interface{} `json:"authors" firestore:"authors" binding:"required"`
	Published bool          `json:"published" firestore:"published"`
}
