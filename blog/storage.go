package blog

import (
	"time"
)

// Storage defines functions for reading from and writing to the database.
type Storage interface {
	// Post returns a single Post from the database using its ID.
	Post(id uint) (*Post, error)

	// AllPosts returns an array containing every single post available.
	AllPosts() ([]*Post, error)

	// SavePost saves the specified post to the database.
	SavePost(post *Post) error
}

// dummyStorage is an implementation of Storage which only contains one single hard-coded post
// for testing reasons.
type dummyStorage struct{}

// NewDummyStorage returns an instance of dummyStorage.
func NewDummyStorage() Storage {
	return &dummyStorage{}
}

var dummyPosts = []*Post{{
	ID:         0,
	Title:      "This is a dummy post!",
	Timestamp:  time.Now(),
	rawContent: "**Testing**",
}}

// ...
func (storage *dummyStorage) Post(id uint) (*Post, error) {
	return dummyPosts[0], nil
}

// ...
func (storage *dummyStorage) AllPosts() ([]*Post, error) {
	return dummyPosts, nil
}

// ...
func (*dummyStorage) SavePost(post *Post) error {
	return nil
}
