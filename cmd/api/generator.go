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

type MailboxIn struct {
	Seller   string `json:"seller"`
	Customer string `json:"customer"`
	Products string `json:"product"`
}

type MailboxOut struct {
	Seller   []string  `json:"seller,omitempty"`
	Customer []string  `json:"customer,omitempty"`
	Product  []Product `json:"product,omitempty"`
}

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
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

// type CustomerNameCSV struct {
// 	Adjective  string
// 	Planet string
// 	Type string
// }

func genCustomerList(n int) []string {
	// var customer []Customer
	// var csvCustomerNameList []CustomerNameCSV

	// CompanyType := map[int]string{
	// 	1: "LLC.",
	// 	2: "Incorporated",
	// 	3: "Corporation",
	// 	4: "Bros.",
	// }

	// fileAdjectives, _ := os.Open("./data/adjectives.csv")
	// defer fileAdjectives.Close()
	// var fileA = csv.NewReader(fileAdjectives)
	// adjectives, _ := fileA.ReadAll()

	// filePlanets, _ := os.Open("./data/planets.csv")
	// defer filePlanets.Close()
	// var fileB = csv.NewReader(filePlanets)
	// planets, _ := fileB.ReadAll()

	// for _, j := range adjectives {
	// 	name := CustomerNameCSV{Adjective: j[0]}
	// 	csvCustomerNameList = append(csvCustomerNameList, name)
	// }

	// for _, k := range planets {
	// 	name2 := CustomerNameCSV{Planet: k[0]}
	// 	csvCustomerNameList = append(csvCustomerNameList, name2)
	// }

	// a := 0
	// b := len(name1)

	// for i := 0; i < n; i++ {
	// 	numName := genRandNum(a, b)

	// }

	var list []string
	fileAdjectives := readCsvFile("./data/adjectives.csv")
	filePlanets := readCsvFile("./data/planets.csv")

	companyType := map[int]string{
		1: "LLC.",
		2: "Incorporated",
		3: "Corporation",
		4: "Bros.",
	}

	a := 0
	b := len(fileAdjectives)
	x := 0
	y := len(filePlanets)

	for i := 0; i < n; i++ {
		num1 := genRandNum(a, b)
		num2 := genRandNum(x, y)
		num3 := genRandNum(1, 5)
		list = append(list, strings.Title(fileAdjectives[num1])+" "+filePlanets[num2]+" "+companyType[num3])
	}
	return list

}
