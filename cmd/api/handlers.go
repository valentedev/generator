package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type MailboxIn struct {
	Seller   string `json:"seller"`
	Customer string `json:"customer"`
	Products string `json:"product"`
}

type MailboxOut struct {
	Seller   []People   `json:"seller,omitempty"`
	Customer []Customer `json:"customer,omitempty"`
	Product  []Product  `json:"product,omitempty"`
}

func (app *application) generateHandler(w http.ResponseWriter, r *http.Request) {
	var mbIn MailboxIn
	var mbOut MailboxOut

	json.NewDecoder(r.Body).Decode(&mbIn)

	seller := mbIn.Seller
	strToNum, _ := strconv.Atoi(seller)
	sellerList := genFullNamesList(strToNum)
	mbOut.Seller = append(mbOut.Seller, sellerList...)

	product := mbIn.Products
	strToNum, _ = strconv.Atoi(product)
	ProdList := genProductList(strToNum)
	mbOut.Product = ProdList

	customer := mbIn.Customer
	strToNum, _ = strconv.Atoi(customer)
	CustList := genCustomerList(strToNum)
	mbOut.Customer = CustList

	jsonList, _ := json.Marshal(mbOut)

	w.Write(jsonList)
}
