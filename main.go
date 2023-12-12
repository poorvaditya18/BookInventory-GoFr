package main

import (
	"fmt"
	"gofr.dev/pkg/gofr"
	"BOOK-INVENTORY/controllers"
)

// create Book table
type Book struct {
	// ID               int    `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Price            int    `json:"price"`
	QuantityAvailable int    `json:"quantity_available"`
}



func main(){
	
	// goFr instance 
	app := gofr.New()

	// createBook - > http://localhost:3000/books
	// /books-> a book is getting created 
	// (H.W)  How will you add multiple books ? 
	app.POST("/books",controllers.CreateBook)

	// getBook 
	app.GET("/book",controllers.GetBook)

	// getAll
	app.GET("/books",controllers.GetAllBooks)

	// updateBook -
	// partial update 
	app.PATCH("/book/{id}",controllers.UpdateBook)

    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
	fmt.Println("Starting server ..... ")
    app.Start()
	fmt.Println("Server started.")
}