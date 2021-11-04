package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readCsvFile(filePath string) map[int]string {
	mapa := make(map[int]string)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, j := range records {
		for _, m := range j {
			mapa[i] = m
		}
	}
	return mapa
}

func genRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func genFullNamesList(n int) []string {
	var list []string
	fileNames := readCsvFile("./data/Names.csv")
	fileLastNames := readCsvFile("./data/LastNames.csv")
	a := 0
	b := len(fileNames)
	x := 0
	y := len(fileLastNames)

	for i := 0; i < n; i++ {
		num1 := genRandNum(a, b)
		num2 := genRandNum(x, y)
		list = append(list, fileNames[num1-1]+" "+fileLastNames[num2-1])
	}
	return list
}

type Product struct {
	Number      string `json:"number"`
	Name        string `json:"name"`
	Phase       string `json:"-"`
	Description string `json:"description"`
}

type ProductCSV struct {
	Name  string
	Phase string
}

func genProductList(n int) []Product {
	var csvlist []ProductCSV
	var list []Product
	prodCode := 500000

	PackageType := map[int]string{
		1: "25kg bag",
		2: "500kg bag",
		3: "1000kg bag",
		4: "Drum PE 200 liters",
		5: "Drum Steel 200 liters",
		6: "Mini tank 1000 liters",
		7: "Metal IBC 1000 liters",
	}

	f, _ := os.Open("./data/Products.csv")
	defer f.Close()
	var arquivo = csv.NewReader(f)
	r, _ := arquivo.ReadAll()

	for _, j := range r {
		p := ProductCSV{Name: j[0], Phase: j[1]}
		csvlist = append(csvlist, p)
	}

	a := 0
	b := len(csvlist)

	for i := 0; i < n; i++ {
		numProd := genRandNum(a, b)
		rp := csvlist[numProd]
		if rp.Phase == "solid" {
			numPackage := genRandNum(1, 4)
			prod := Product{Number: strconv.Itoa(prodCode), Name: rp.Name, Description: PackageType[numPackage]}
			list = append(list, prod)
		} else {
			numPackage := genRandNum(5, 8)
			prod := Product{Number: strconv.Itoa(prodCode), Name: rp.Name, Description: PackageType[numPackage]}
			list = append(list, prod)
		}
		prodCode++
	}

	return list
}

type Customer struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
}

type Places struct {
	State string
	City  string
}

func genCustomerList(n int) []Customer {

	var customer []Customer
	var placesList []Places

	fileAdjectives := readCsvFile("./data/adjectives.csv")
	filePlanets := readCsvFile("./data/planets.csv")

	companyType := map[int]string{
		1: "LLC.",
		2: "Incorporated",
		3: "Corporation",
		4: "Bros.",
	}

	f, _ := os.Open("./data/stateCities.csv")
	defer f.Close()
	var arquivo = csv.NewReader(f)
	r, _ := arquivo.ReadAll()

	for _, j := range r {
		p := Places{State: j[0], City: j[1]}
		placesList = append(placesList, p)
	}

	a := 0
	b := len(fileAdjectives)
	x := 0
	y := len(filePlanets)
	s := 0
	t := len(placesList)

	for i := 0; i < n; i++ {
		num1 := genRandNum(a, b)
		num2 := genRandNum(x, y)
		num3 := genRandNum(1, 5)
		num4 := genRandNum(s, t)
		c := Customer{Name: strings.Title(fileAdjectives[num1]) + " " + filePlanets[num2] + " " + companyType[num3], City: placesList[num4].City, State: placesList[num4].State}
		customer = append(customer, c)
	}
	return customer
}
