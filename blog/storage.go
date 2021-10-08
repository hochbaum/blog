package blog

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Storage defines functions for reading from and writing to the database.
type Storage interface {
	// Post returns a single Post from the database using its ID.
	Post(id uint) (*Post, error)

	// AllPosts returns an array containing every single post available.
	AllPosts() ([]*Post, error)

	Migrate() error

	// SavePosts saves the specified posts to the database.
	SavePosts(post ...*Post) error
}

// dummyStorage is an implementation of Storage which only contains one single hard-coded post
// for testing reasons.
type dummyStorage struct{}

// NewDummyStorage returns an instance of dummyStorage.
func NewDummyStorage() Storage {
	return &dummyStorage{}
}

var dummyPosts = []*Post{{
	Title:   "This is a dummy post!",
	Content: "**Testing**",
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
func (storage *dummyStorage) Migrate() error {
	return nil
}

// ...
func (*dummyStorage) SavePosts(posts ...*Post) error {
	return nil
}

// sqliteStorage is an implementation of Storage using SQLite to store data.
type sqliteStorage struct {
	GORM *gorm.DB
}

// NewSQLiteStorage constructs a sqliteStorage and returns the pointer to it. The database file
// will be loaded at the specified path.
func NewSQLiteStorage(path string) (Storage, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &sqliteStorage{GORM: db}, nil
}

// ...
func (storage *sqliteStorage) Post(id uint) (*Post, error) {
	var post Post
	return &post, storage.GORM.Find(&post, id).Error
}

// ...
func (storage *sqliteStorage) AllPosts() ([]*Post, error) {
	var posts []*Post
	return posts, storage.GORM.Order("created_at DESC").Find(&posts).Error
}

// ...
func (storage *sqliteStorage) Migrate() error {
	return storage.GORM.AutoMigrate(&Post{})
}

// ...
func (storage *sqliteStorage) SavePosts(posts ...*Post) error {
	return storage.GORM.Save(posts).Error
}