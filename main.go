package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//---------------------
//new struct that contains transactions  was defined
//----------------
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
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

var transactions Transactions

//--------------------------------------------
// function to read json db file
//-------------------------------------------
func readDBjson(fileName string) error {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully opened: %s \n", fileName)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &transactions)

	return nil
}

//--------------------------------------------
// function to update json db file with 0644 permission
//-------------------------------------------
func updateDBjson() {
	file, _ := json.MarshalIndent(transactions, "", " ")
	_ = ioutil.WriteFile("operations.json", file, 0644)
}

//--------------------------------------------
// CRUD functions for API
//-------------------------------------------
func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set json content-type
	json.NewEncoder(w).Encode(transactions)            // encode transcations as json to responsewritter
}

func getTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //params

	for _, transaction := range transactions.Transactions {
		if transaction.ID == params["id"] {
			json.NewEncoder(w).Encode(transaction)
			return
		}
	}
	err := errors.New("error: client defined non-existing ID")
	fmt.Println(err)
	json.NewEncoder(w).Encode("Error: Such ID does not exists, try new ID")
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, transaction := range transactions.Transactions {
		if transaction.ID == params["id"] {
			transactions.Transactions = append(transactions.Transactions[:index], transactions.Transactions[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(transactions)
	updateDBjson()
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	for _, value := range transactions.Transactions {
		if transaction.ID == value.ID {
			err := errors.New("error: client adding the transaction with the same ID")
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error: Such ID exists, try new ID")
			return
		}
	}

	transactions.Transactions = append(transactions.Transactions, transaction)
	json.NewEncoder(w).Encode(transactions)
	updateDBjson()
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//loop over the transactions
	//delete transaction with the ID that client sent
	//add a new transaction
	for index, transaction := range transactions.Transactions {
		if transaction.ID == params["id"] {
			transactions.Transactions = append(transactions.Transactions[:index], transactions.Transactions[index+1:]...)

			var transaction Transaction
			_ = json.NewDecoder(r.Body).Decode(&transaction)
			transaction.ID = params["id"]
			transactions.Transactions = append(transactions.Transactions, transaction)
			json.NewEncoder(w).Encode(transactions)
			updateDBjson()
		}
	}
}

//------------------- Main function --------------------------
func main() {

	//read json Db file
	if err := readDBjson("operations.json"); err != nil {
		log.Fatal(err)
	}

	//Implement a new router and dispatcher for matching incoming requests to their respective handler
	r := mux.NewRouter()

	//Handler functions for CRUD commands of client
	r.HandleFunc("/transactions", getTransactions).Methods("GET")
	r.HandleFunc("/transactions/{id}", getTransaction).Methods("GET")
	r.HandleFunc("/transactions", createTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", deleteTransaction).Methods("DELETE")
	r.HandleFunc("/transactions/{id}", updateTransaction).Methods("PUT")

	//Start server at port 8000 on the localhost
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
