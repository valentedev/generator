package main

import (
	"embed"
	"encoding/csv"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//go:embed data
var csvdata embed.FS

func readCsvFileEmbed(filePath string) [][]string {
	f, err := csvdata.Open(filePath)
	if err != nil {
		fmt.Printf("error inside readCsvFileEmbed %s\n", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func genRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

type People struct {
	Name     string
	LastName string
}

func genFullNamesList(n int) []People {
	var people []People

	names := make(map[int]string)
	fileNames := readCsvFileEmbed("data/Names.csv")
	for i, j := range fileNames {
		names[i] = j[0]
	}

	lastNames := make(map[int]string)
	fileLastNames := readCsvFileEmbed("data/LastNames.csv")
	for i, j := range fileLastNames {
		lastNames[i] = j[0]
	}

	b := len(fileNames)
	y := len(fileLastNames)

	for i := 0; i < n; i++ {
		num1 := genRandNum(0, b)
		num2 := genRandNum(0, y)
		nam := People{Name: string(names[num1-1]), LastName: string(lastNames[num2-1])}
		people = append(people, nam)
	}

	return people
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

	fileProducts := readCsvFileEmbed("data/Products.csv")

	for _, j := range fileProducts {
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

	adj := make(map[int]string)
	fileAdjectives := readCsvFileEmbed("data/adjectives.csv")
	for i, j := range fileAdjectives {
		adj[i] = j[0]
	}

	plan := make(map[int]string)
	filePlanets := readCsvFileEmbed("data/planets.csv")
	for i, j := range filePlanets {
		plan[i] = j[0]
	}

	strn := make(map[int]string)
	fileStreets := readCsvFileEmbed("data/streetNames.csv")
	for i, j := range fileStreets {
		strn[i] = j[0]
	}

	fileStateCity := readCsvFileEmbed("data/stateCities.csv")
	var placesList []Places
	for _, j := range fileStateCity {
		scm := Places{State: j[0], City: j[1]}
		placesList = append(placesList, scm)
	}

	companyType := map[int]string{
		1: "LLC.",
		2: "Incorporated",
		3: "Corporation",
		4: "Bros.",
	}

	roadType := map[int]string{
		1:  "Rd.",
		2:  "St.",
		3:  "Ave.",
		4:  "Blvd",
		5:  "Ln.",
		6:  "Dr.",
		7:  "Way",
		8:  "Hwy.",
		9:  "Fwy.",
		10: "Pkwy.",
	}

	adjLen := len(adj)
	planetLen := len(plan)
	placesLen := len(placesList)
	streetsLen := len(strn)

	for i := 0; i < n; i++ {
		num1 := genRandNum(0, adjLen)
		num2 := genRandNum(0, planetLen)
		//CompanyType
		num3 := genRandNum(1, 5)
		num4 := genRandNum(0, placesLen)
		num5 := genRandNum(0, streetsLen)
		//AddressNumber
		num6 := genRandNum(66, 787)
		//RoadType
		num7 := genRandNum(1, 11)

		c := Customer{Name: strings.Title(adj[num1]) + " " + plan[num2] + " " + companyType[num3], Address: strconv.Itoa(num6) + " " + strn[num5] + " " + roadType[num7], City: placesList[num4].City, State: placesList[num4].State}

		customer = append(customer, c)
	}
	return customer
}
