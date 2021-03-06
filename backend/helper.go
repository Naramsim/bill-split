package main

import (
	"encoding/json"
	"github.com/Baumanar/bill-split/backend/data"
	"log"
	"net/http"
	"strings"
)

// Convenience function to redirect to the error message page
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}



func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}



func PopulateDB(){

	//#####################################################################################
	billSplit, err := data.CreateBillSplit("Flat sharing")
	billSplit.CreateParticipants([]string{"Robin", "John", "Paul", "Mary"})

	expense, err := billSplit.CreateExpense("Beers", 30.0, "Robin")
	if err != nil {
		log.Fatal(err)
	}
	err = expense.AddParticipants([]string{"John", "Robin", "Mary"})
	if err != nil {
		log.Fatal(err)
	}
	expense, err = billSplit.CreateExpense("Pizza", 33.5, "Mary")
	if err != nil {
		log.Fatal(err)
	}
	err = expense.AddParticipants([]string{"Mary", "Robin", "Paul"})
	if err != nil {
		log.Fatal(err)
	}
	expense, err = billSplit.CreateExpense("Party", 33.5, "John")
	if err != nil {
		log.Fatal(err)
	}
	err = expense.AddParticipants([]string{"Mary", "Robin", "Paul", "John"})
	if err != nil {
		log.Fatal(err)
	}
	//#####################################################################################

	billSplit, err = data.CreateBillSplit("Holidays")
	billSplit.CreateParticipants([]string{"Emma", "Steve", "Sophia", "Bill", "Patrick", "Lisa"})


	expense, err = billSplit.CreateExpense("Groceries", 30.65, "Bill")
	if err != nil {
		log.Fatal(err)
	}

	err = expense.AddParticipants([]string{"Lisa", "Bill", "Patrick"})
	if err != nil {
		log.Fatal(err)
	}

	expense, err = billSplit.CreateExpense("Trip", 130.20, "Lisa")
	if err != nil {
		log.Fatal(err)
	}
	err = expense.AddParticipants([]string{"Lisa", "Bill", "Patrick", "Sophia", "Emma"})
	if err != nil {
		log.Fatal(err)
	}
	expense, err = billSplit.CreateExpense("Picnic", 80.77, "Sophia")
	if err != nil {
		log.Fatal(err)
	}
	err = expense.AddParticipants([]string{"Lisa", "Bill", "Sophia", "Emma"})
	if err != nil {
		log.Fatal(err)
	}

	//#####################################################################################

}
