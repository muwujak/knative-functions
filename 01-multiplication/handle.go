package function

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"encoding/json"
)

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	
	// set response to json
	w.Header().Set("Content-Type", "application/json");

	// parse request body consists of json
	var reqData RequestData;
	err := json.NewDecoder(r.Body).Decode($reqData);
	if err != nil {
		http.Error(w, "Invalid request body");
		return;
	}
	

	dump, err	 := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Received request")
	fmt.Printf("%q\n", dump)
	fmt.Fprintf(w, "%q", dump)
}
