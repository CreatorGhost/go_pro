package handlers

import (
	"encoding/json"
	"net/http"
	
)


// SortRequest is the expected request body for sorting endpoints.
type SortRequest struct {
	ToSort [][]int `json:"toSort"`
}

// SortResponse is the response body for sorting endpoints.
type SortResponse struct {
	SortedArrays [][]int `json:"sortedArrays"`
	TimeNs       int64   `json:"timeNs"`
}

// ProcessSingleHandler handles the /process-single endpoint.
func ProcessSingleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SortRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assuming processSingle is a function that sorts the arrays and returns the sorted arrays and time taken.
	// You need to implement this function or import it from the appropriate package.
	sortedArrays, timeNs := processSingle(req.ToSort)

	resp := SortResponse{
		SortedArrays: sortedArrays,
		TimeNs:       timeNs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ProcessConcurrentHandler handles the /process-concurrent endpoint.
func ProcessConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SortRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assuming processConcurrent is a function that sorts the arrays concurrently and returns the sorted arrays and time taken.
	// You need to implement this function or import it from the appropriate package.
	sortedArrays, timeNs := processConcurrent(req.ToSort)

	resp := SortResponse{
		SortedArrays: sortedArrays,
		TimeNs:       timeNs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// You can remove the StartServer function if it's not used elsewhere,
// as the main.go is already starting the server.