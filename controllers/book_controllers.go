package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"gofr.dev/pkg/gofr"
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
	PriceStr := formData.Get("Price")
	QuantityAvailableStr := formData.Get("QuantityAvailable")
	
	// Convert string Price to integer
	Price, err := strconv.Atoi(PriceStr)
	if err != nil {
		return nil, err // Handle the error if conversion fails
	}

	// Convert string QuantityAvailable to integer
	QuantityAvailable, err := strconv.Atoi(QuantityAvailableStr)
	if err != nil {
		return nil, err // Handle the error if conversion fails
	}

	fmt.Printf("Title : %s, \n Author : %s,\n Price : %d,\n QuantityAvailable: %d\n", Title, Author, Price, QuantityAvailable)


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
        return nil, errors.New("title cannot be empty")
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

	result, err := ctx.DB().ExecContext(ctx, "UPDATE book SET QuantityAvailable = ? WHERE Title = ?", QuantityAvailable, Title)


	if err != nil {
		fmt.Println("Error updating book:", err)
		return nil, err
	}
	
	// Check rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("no book found with Title: %s", Title)
	}

	fmt.Println("Book updated successfully")
	return "Successfully Updated Quantity Of Book", nil
}


// delete a book 
func DeleteBook(ctx *gofr.Context)(interface{}, error){
	req := ctx.Request()
    queryValues := req.URL.Query()
    bookID := queryValues.Get("id")

    // Perform the delete operation in the database
    result, err := ctx.DB().ExecContext(ctx,"DELETE FROM book WHERE BookID = ?", bookID)
    if err != nil {
        fmt.Println("Error deleting book:", err)
        return result, err
    }

	rowsAffected, err := result.RowsAffected()
    if err != nil {
        fmt.Println("Error getting rows affected:", err)
        return nil, err
    }

	if rowsAffected == 0 {
        return nil, fmt.Errorf("Book with ID %s is not present", bookID)
    }
	
    return "Successfully deleted Book", nil
}


