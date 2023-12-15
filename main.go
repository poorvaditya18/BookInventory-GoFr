package main

import (
	"BOOK-INVENTORY/controllers"
	"fmt"

	"gofr.dev/pkg/gofr"
	// "BOOK-INVENTORY/middlewares"
)

// create Book table

func main(){
	
	// goFr instance 
	app := gofr.New()

	// createBook - > http://localhost:3000/books
	// /books-> a book is getting created 
	// (H.W)  How will you add multiple books ? 
	// app.Server.UseMiddleware(middlewares.CreateBookMiddleware())
	app.POST("/book",controllers.CreateBook)

	// getBook 
	app.GET("/book",controllers.GetBook)

	// getAll
	app.GET("/books",controllers.GetAllBooks)

	// updateBook -
	// partial update 
	app.PATCH("/book",controllers.UpdateBookQuantity)


	// deleteBook 
	app.DELETE("/book",controllers.DeleteBook)
    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
	fmt.Println("Starting server ..... ")
    app.Start()
	fmt.Println("Server started.")
}