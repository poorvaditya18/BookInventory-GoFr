package controllers

import (
	"fmt"
	"io"
	"net/url"
	"errors"
	"gofr.dev/pkg/gofr"
	"strconv"
)

// create book structure
type Book struct {
	BookID               int    `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Price            int    `json:"price"`
	QuantityAvailable int    `json:"quantity_available"`
}


// create book controller 
func CreateBook(ctx *gofr.Context)(interface{}, error){

	// requestBody --> it is in url-encoded 
	requestBody, err := io.ReadAll(ctx.Request().Body)
		
	// check whether any error is there or not 
	if err != nil {
		return nil, err
	}


	// Parse the URL-encoded form data
	formData, err := url.ParseQuery(string(requestBody))

	if err != nil {
		return nil, err
	}

	Title := formData.Get("Title")
	Author := formData.Get("Author")
	Price := formData.Get("Price")
	QuantityAvailable := formData.Get("QuantityAvailable")
	

	fmt.Printf("Title : %s, \n Author : %s,\n Price : %s,\n QuantityAvailable: %s\n", Title, Author, Price, QuantityAvailable)


	// db layer ------
	resp, err := ctx.DB().ExecContext(ctx, "INSERT INTO BOOK (Title,Author,Price,QuantityAvailable) VALUES (?,?,?,?)", Title,Author,Price,QuantityAvailable)


	fmt.Print(resp)

	if err != nil {
		return "Not able to create", err
	}
	return "Successfully Created Book", err
}


// get all books controller 
func GetAllBooks(ctx *gofr.Context)(interface{}, error){

	//books array
	var books []Book

		// Getting the customer from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM book")

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var book Book
			if err := rows.Scan(&book.BookID, &book.Title,&book.Author, &book.Price,&book.QuantityAvailable); err != nil {
				return nil, err
			}

			books = append(books, book)
		}

		// return the customer
		return books, nil
}


// get a book controller by Title
func GetBook(ctx *gofr.Context)(interface{}, error){
	req := ctx.Request()

    // Extract the query parameters from the request URL
    queryValues := req.URL.Query()
	title := queryValues.Get("Title")

	if title == "" {
        return nil, errors.New("Title cannot be empty")
    }

	fmt.Printf("Fetching book from db....")

	// row object 
	row :=ctx.DB().QueryRowContext(ctx, "SELECT * FROM book WHERE Title = ?", title)

	// Prepare variables to store the result of row object 
    var bookID int
    var bookTitle, bookAuthor string
    var bookPrice, bookQuantity int

    // Scan the result into variables
    err := row.Scan(&bookID, &bookTitle, &bookAuthor, &bookPrice, &bookQuantity)
    if err != nil {
        return nil, err // handle error
    }

    // Create a book object or a map with the retrieved values
    book := map[string]interface{}{
        // "BookID":           bookID,
        "Title":            bookTitle,
        "Author":           bookAuthor,
        "Price":            bookPrice,
        "QuantityAvailable": bookQuantity,
    }

    return book, nil
}



// update a book by price 
func UpdateBookQuantity(ctx *gofr.Context)(interface{}, error){
	requestBody, err := io.ReadAll(ctx.Request().Body)
    if err != nil {
        return nil, err
    }

    formData, err := url.ParseQuery(string(requestBody))
    if err != nil {
        return nil, err
    }

    Title := formData.Get("Title")
    QuantityAvailableStr := formData.Get("QuantityAvailable")
    QuantityAvailable, err := strconv.Atoi(QuantityAvailableStr)
    if err != nil {
        return nil, err
    }

    resp, err := ctx.DB().ExecContext(ctx, "UPDATE book SET QuantityAvailable = ? WHERE Title = ?", QuantityAvailable, Title)
    if err != nil {
        return nil, err
    }

    return resp, nil
}


// delete a book 



