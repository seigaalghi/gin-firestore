package models

import "time"

// Post is ...
type Post struct {
	Title     string        `firestore:"title" binding:"required"`
	Text      string        `firestore:"text" binding:"required"`
	Date      time.Time     `firestore:"date"`
	Price     int64         `firestore:"price" binding:"required,numeric"`
	Authors   []interface{} `firestore:"authors" binding:"required"`
	Published bool          `firestore:"published"`
}
