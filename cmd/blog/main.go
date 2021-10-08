package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"hochbaum.dev/blog/blog"
	"os"
)

const dbPath = "blog.db"

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	// Check for this before we let GORM open the database file, because it creates it.
	var firstRun bool
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		firstRun = true
	}

	storage, err := blog.NewSQLiteStorage(dbPath)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "error opening blog.db: ")
		panic(err)
	}

	// If blog.db was just created, that means we are running the blog for the first time.
	// Run migration and publish example posts.
	if firstRun {
		fmt.Println("Running migration and creating example posts for you.")

		if err := storage.Migrate(); err != nil {
			_, _ = fmt.Fprint(os.Stderr, "migration failed: ")
			panic(err)
		}

		if err := storage.SavePosts(examplePosts...); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "could not create example post: ")
			panic(err)
		}
	}

	if err := blog.New(router, storage).Publish(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, "could not publish blog: ")
		panic(err)
	}
}

// examplePosts contains a few blog.Post which will be created on the first execution.
var examplePosts = []*blog.Post{
	{
		Title: "Lorem ipsum",
		Content: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt " +
			"ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores" +
			" et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. " +
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut " +
			"labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores " +
			"et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
	},
	{
		Title: "Lorem ipsum 2",
		Content: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt " +
			"ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores" +
			" et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. " +
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut " +
			"labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores " +
			"et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
	},
}
