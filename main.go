package main
 
import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Please provide a name parameter", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hiiii %s ::: Nice to Meet You !!", name)
}

func main() {
	http.HandleFunc("/greet", greetHandler)

	fmt.Println("Server starting ..........")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
