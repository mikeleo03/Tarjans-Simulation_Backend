package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mikeleo03/Tarjans-Algorithm_Backend/algorithm"
)

// Response format
type Response struct {
	Status     bool  				`json:"status,omitempty"`
	Message    string 				`json:"msg,omitempty"`
	Value      [][]string			`json:"value,omitempty"`
	Time       int64				`json:"time,omitempty"`
}

// ProcessSCC, doing Tarjans algorithm to find SCC using algorithms
func ProcessSCC(w http.ResponseWriter, r *http.Request) {
	// Allow all origin to handle cors issue
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty data input
	var graph algorithm.Graph

	// decode the json request to fakul
	if (r.Method == "POST") {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			res := Response {
				Status	: false,
				Message	: "Error reading request body",
			}
			// send the response
			json.NewEncoder(w).Encode(res)
			return
		}

		err = json.Unmarshal(body, &graph)
		if err != nil {
			res := Response {
				Status	: false,
				Message	: "Error parsing JSON, please check your file",
			}
			// send the response
			json.NewEncoder(w).Encode(res)
			return
		}

		// Gather the time process and the result of algo
		start := time.Now() // Get the current time
		scc := algorithm.TarjanSCC(&graph)
		elapsed := time.Since(start) // Calculate the elapsed time

		// Check if there is no mata kuliah after applying the algorithm
		if len(scc) == 0 {
			res := Response {
				Status	: false,
				Message	: "No result.",
			}
			// send the response
			json.NewEncoder(w).Encode(res)
			return
		}
		// Return the result to the frontend
		res := Response {
			Status	: true,
			Value	: scc,
			Time	: elapsed.Nanoseconds(),
			Message	: "",
		}
		// send the response
		json.NewEncoder(w).Encode(res)
	}
}

// ProcessBridges, doing modified Tarjans algorithm to find Bridges using algorithms
func ProcessBridges(w http.ResponseWriter, r *http.Request) {
	// Allow all origin to handle cors issue
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty data input
	var graph algorithm.Graph

	// decode the json request to fakul
	if (r.Method == "POST") {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			res := Response {
				Status	: false,
				Message	: "Error reading request body",
			}
			// send the response
			json.NewEncoder(w).Encode(res)
			return
		}

		err = json.Unmarshal(body, &graph)
		if err != nil {
			res := Response {
				Status	: false,
				Message	: "Error parsing JSON, please check your file",
			}
			// send the response
			json.NewEncoder(w).Encode(res)
			return
		}

		// Gather the time process and the result of algo
		start := time.Now() // Get the current time
		scc := algorithm.TarjanBridges(&graph)
		elapsed := time.Since(start) // Calculate the elapsed time

		// Check if there is no mata kuliah after applying the algorithm
		if len(scc) == 0 {
			res := Response {
				Status	: false,
				Message	: "No result.",
			}
			// send the response
			json.NewEncoder(w).Encode(res)
			return
		}
		// Return the result to the frontend
		res := Response {
			Status	: true,
			Value	: scc,
			Time	: elapsed.Nanoseconds(),
			Message	: "",
		}
		// send the response
		json.NewEncoder(w).Encode(res)
	}
}