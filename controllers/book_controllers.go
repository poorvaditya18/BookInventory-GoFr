package controllers
import (
	"fmt"
	"io"
	"net/url"
	"gofr.dev/pkg/gofr"
)

// create book controller 
func CreateBook(ctx *gofr.Context)(interface{}, error){

	// requestBody --> it is in url-encoded 
	requestBody, err := io.ReadAll(ctx.Request().Body)
		
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




// get a book controller 




// update a book 



// delete a book 



