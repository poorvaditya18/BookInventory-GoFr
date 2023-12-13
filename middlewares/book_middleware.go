// middleware
package middlewares

import (
	// "fmt"
	"net/http"
	// "strconv"
	// "BOOK-INVENTORY/controllers"
)

// create book middleware
func CreateBookMiddleware() func(handler http.Handler) http.Handler {

	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			// your logic here
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}

			// extract values from body 
			Title := r.FormValue("Title")
			Author := r.FormValue("Author")
			PriceStr := r.FormValue("Price")
			QuantityAvailableStr := r.FormValue("QuantityAvailable")

			// check whether all fields are present or not 
			if Title == "" || Author == "" || PriceStr == "" || QuantityAvailableStr == "" {
				// if present then send to next middleware 
				// Convert Price and QuantityAvailable to integers
                // Price, err := strconv.Atoi(PriceStr)
                // if err != nil {
                //     http.Error(w, "Invalid Price format", http.StatusBadRequest)
                //     return
                // }

                // QuantityAvailable, err := strconv.Atoi(QuantityAvailableStr)
                // if err != nil {
                //     http.Error(w, "Invalid QuantityAvailable format", http.StatusBadRequest)
                //     return
                // }
				// fmt.Printf("All Fields are present.")
				// controllers.CreateBook(w, r, Title, Author, Price, QuantityAvailable)
				// sends the request to the next middleware/request-handler
				http.Error(w, "Invalid Fields", http.StatusBadRequest)
				return
			}
			
		})
	}
}