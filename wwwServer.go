package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func ResultsHandler(w http.ResponseWriter, r *http.Request) {

	//log.Printf("%v",AlbarRegPriceList[0])
	//log.Print(html.EscapeString(AlbarRegPriceList[0].CarCategoryName))
	for i := range AlbarRegPriceList {
		AlbarRegPriceList[i].CarCategoryName = strings.TrimPrefix(AlbarRegPriceList[i].CarCategoryName, "<strong>")
		AlbarRegPriceList[i].CarCategoryName = strings.TrimSuffix(AlbarRegPriceList[i].CarCategoryName, "</strong> or Similar")
		AlbarRegPriceList[i].CarCategoryName = strings.Replace(AlbarRegPriceList[i].CarCategoryName, "</strong> or Similar", " ", 1)
		//log.Print(AlbarRegPriceList[i].CarCategoryName)
	}

	page := AlbarRegPriceList

	t, err := template.ParseFiles("templates/results.html")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	t.Execute(w, page)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	PriceListUrl := priceListUrl{
		pickupDate:          r.FormValue("pickupDate"),
		dropoffDate:         r.FormValue("dropoffDate"),
		driverAge:           r.FormValue("driverAge"),
		dropOffLocationCode: "300",
		pickupLocationCode:  "300",
		//loginUserName:       "vastama_709",
		//userType:            "Agency", //Regular
		priceListId: DefaultPriceList.Value, //1029
	}

	tmpl.Execute(w, struct{ Success bool }{true})
	//log.Print("formHandler: Price list url:", PriceListUrl)
	JsonRaw := ReadPriceList(PriceListUrl)

	err = json.Unmarshal([]byte(JsonRaw), &AlbarRegPriceList)
	if err != nil {
		panic(err.Error())
	}
	//log.Print("formHandler: JsonRaw:", JsonRaw)

}

type M map[string]interface{}

var AlbarRegPriceList = CarCategory{}
var PriceListUrl = priceListUrl{}
var DefaultPriceList = AlbarPricelistList{0, "", "1", "", "1035"}
var PROMO = AlbarPricelistList{29, "PROMO", "1", "Pricelist_1040", "1040"}
var RELIGIOUS = AlbarPricelistList{30, "RELIGIOUS", "1", "Pricelist_1025", "1025"}
var ISRAEL = AlbarPricelistList{16, "ISRAEL", "1", "Pricelist_1027", "1027"}

func main() {
	//AlbarLogin()
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))

	fmt.Println("Starting web server on port 50002")
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/results/", ResultsHandler)
	log.Fatal(http.ListenAndServe(":50002", nil))

}
