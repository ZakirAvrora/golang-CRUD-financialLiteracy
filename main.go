package main

import (
	"encoding/json"
	"fmt"
	"log"

	//"math/rand"
	"net/http"
	//"strconv"
	"errors"

	"github.com/gorilla/mux"
)

type Operation struct {
	ID       string `json:"id"`
	Price    string `json:"price"`
	Type     string `json:"type"` // income, purchase
	Comment  string `json:"comment"`
	Category string `json:"category"`
	Date     *Date  `json:"date"`
}

type Date struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
}

var operations []Operation

func getOperations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operations)
}

func getOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, operation := range operations {
		if operation.ID == params["id"] {
			json.NewEncoder(w).Encode(operation)
			return
		}
	}
}

func deleteOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, operation := range operations {
		if operation.ID == params["id"] {
			operations = append(operations[:index], operations[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(operations)
}

func createOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var operation Operation
	_ = json.NewDecoder(r.Body).Decode(&operation)
	//operation.ID = strconv.Itoa(rand.Intn(1000000000))
	for _, transac := range operations {
		if operation.ID == transac.ID {
			err := errors.New("error: Such ID exists")
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error: Such ID exists, try new ID")
			return
		}
	}

	operations = append(operations, operation)
	json.NewEncoder(w).Encode(operations)
}

func updateOperation(w http.ResponseWriter, r *http.Request) {
	//set json content-type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the operations
	//delete operation with the i.d that you sent
	//add a new operation
	for index, operation := range operations {
		if operation.ID == params["id"] {
			operations = append(operations[:index], operations[index+1:]...)
			var operation Operation
			_ = json.NewDecoder(r.Body).Decode(&operation)
			operation.ID = params["id"]
			operations = append(operations, operation)
			json.NewEncoder(w).Encode(operations)
		}
	}
}

func main() {

	operations = append(operations, Operation{ID: "1", Price: "2500", Type: "purchase", Comment: "Meal was purchased", Category: "Meal", Date: &Date{Year: "2020", Month: "Jan", Day: "15"}})
	operations = append(operations, Operation{ID: "2", Price: "5000", Type: "income", Comment: "Own product was sold", Category: "Business", Date: &Date{Year: "2020", Month: "Jan", Day: "15"}})
	operations = append(operations, Operation{ID: "3", Price: "2000", Type: "income", Comment: "Mothly interest rate from the deposit ", Category: "Deposit", Date: &Date{Year: "2020", Month: "Jan", Day: "15"}})
	operations = append(operations, Operation{ID: "4", Price: "1200", Type: "purchase", Comment: "Price for Taxi", Category: "Transport", Date: &Date{Year: "2020", Month: "Jan", Day: "15"}})

	r := mux.NewRouter()

	r.HandleFunc("/operations", getOperations).Methods("GET")
	r.HandleFunc("/operations/{id}", getOperation).Methods("GET")
	r.HandleFunc("/operations", createOperation).Methods("POST")
	r.HandleFunc("/operations/{id}", deleteOperation).Methods("DELETE")
	r.HandleFunc("/operations/{id}", updateOperation).Methods("PUT")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
