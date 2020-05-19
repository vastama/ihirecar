package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//var picturePath = "https://www.ihirecar.com/images/cars/"

type Category struct {
	name, code                                                  string
	code1, seats                                                int
	car                                                         string
	price, a, s, d, f, g, h, j, w, e, k, l, z, x, c, v, b, n, m float32
}



type resultsPage struct {
	Title string
}

var Cat = map[string]Category{
	"B":  {"B", "MBMR", 10, 4, "Suzuki Alto Man.", 10, 22, 10, 22, 63, 147, 9, 21, 610, 790, 10, 5, 12, 0.45, 10, 100, 5, 60, 475},
	"Q":  {"Q", "EBAR", 15, 4, "Fiat 500(2 Doors)", 10, 22, 10, 22, 63, 147, 9, 21, 610, 790, 10, 5, 12, 0.45, 10, 100, 5, 60, 475},
	"C":  {"C", "EDAR", 30, 4, "Kia Picanto", 11, 23, 11, 23, 70, 154, 10, 22, 620, 800, 10, 5, 12, 0.45, 10, 100, 5, 60, 475},
	"D":  {"D", "CBAR", 20, 4, "Hyundai i20", 12, 24, 12, 24, 77, 161, 11, 23, 630, 810, 10, 5, 12, 0.45, 10, 100, 5, 60, 475},
	"E":  {"E", "CCAR", 50, 5, "Ford Fiesta", 13, 25, 13, 25, 84, 168, 12, 24, 645, 825, 10, 5, 12, 0.45, 10, 100, 5, 60, 475},
	"F":  {"F", "IDAR", 60, 5, "Hyundai Accent i25", 16, 28, 16, 28, 98, 182, 14, 26, 750, 930, 11, 5, 12, 0.5, 10, 100, 5, 60, 475},
	"I":  {"I", "SCAR", 90, 5, "VW Golf", 18, 30, 18, 30, 112, 196, 16, 28, 790, 970, 11, 5, 12, 0.5, 10, 100, 5, 60, 475},
	"IW": {"IW", "CWAR", 91, 5, "Seat Leon SW", 20, 32, 20, 32, 126, 210, 18, 30, 820, 1000, 11, 5, 12, 0.5, 10, 100, 5, 60, 475},
	"H":  {"H", "SDAR", 80, 5, "VW Jetta", 24, 36, 24, 36, 154, 238, 22, 34, 960, 1140, 11, 5, 12, 0.5, 10, 150, 5, 70, 715},
	"M":  {"M", "PDAR", 130, 5, "Mazda 6", 31, 46, 31, 46, 196, 301, 28, 43, 1050, 1230, 14, 5, 15, 0.5, 10, 150, 7, 70, 715},
	"MH": {"MH", "PCAR", 135, 5, "Hyundai Sonata", 42, 57, 42, 57, 266, 371, 38, 53, 1450, 1630, 14, 5, 15, 0.75, 10, 150, 7, 70, 715},
	"K":  {"K", "UDAR", 110, 5, "BMW 318i", 72, 87, 72, 87, 455, 560, 65, 80, 1590, 1770, 14, 5, 15, 0.75, 10, 150, 7, 70, 715},
	"R":  {"R", "FCAR", 180, 5, "VW Passat", 58, 73, 58, 73, 364, 469, 52, 67, 1640, 1820, 14, 5, 15, 0.75, 15, 210, 10, 90, 960},
	"P":  {"P", "LDAR", 160, 5, "Nissan Maxima", 79, 109, 79, 109, 497, 707, 71, 101, 2170, 2410, 35, 10, 40, 0.75, 15, 210, 10, 90, 960},
	"W":  {"W", "LCBR", 230, 5, "Audi A6", 122, 152, 122, 152, 770, 980, 110, 140, 2690, 2930, 35, 10, 40, 0.75, 15, 210, 10, 90, 960},
	"FX": {"FX", "DBAR", 65, 5, "Renault Megane Coupe(2 Doors)", 17, 29, 17, 29, 105, 189, 15, 27, 760, 940, 11, 5, 12, 0.5, 10, 100, 5, 60, 475},
	"T":  {"T", "CPMR", 200, 5, "VW Caddy Man.", 39, 69, 39, 69, 245, 455, 35, 65, 1050, 1290, 15, 10, 20, 0.75, 10, 100, 5, 60, 475},
	"J":  {"J", "SFBR", 100, 5, "Hyundai Tucson Aut.", 42, 72, 42, 72, 266, 476, 38, 68, 1240, 1480, 15, 10, 20, 0.75, 10, 150, 7, 70, 715},
	"O":  {"O", "LFBR", 150, 5, "Ford Edge Aut.", 74, 104, 74, 104, 469, 679, 67, 97, 2250, 2550, 35, 10, 40, 0.75, 15, 210, 10, 90, 960},
	"U":  {"U", "IVAR", 210, 7, "Opel Zafira Aut.", 39, 69, 39, 69, 245, 455, 35, 65, 1220, 1520, 15, 10, 20, 0.75, 15, 210, 10, 90, 715},
	"V":  {"V", "SVAR", 220, 7, "Mitsubishi Outlander Aut.", 58, 88, 58, 88, 364, 574, 52, 82, 1550, 1850, 15, 10, 20, 0.75, 15, 210, 10, 90, 960},
	"V8": {"V8", "FVAR", 225, 8, "Kia Carnival Aut", 83, 113, 83, 113, 525, 735, 75, 105, 1870, 2170, 35, 10, 40, 0.75, 15, 210, 10, 90, 960},
	"Y":  {"Y", "FVMR", 250, 9, "Renault Traffic Man.", 55, 85, 55, 85, 343, 553, 49, 79, 1600, 1900, 35, 10, 40, 0.75, 15, 210, 10, 90, 960},
	"Z":  {"Z", "LVAR", 260, 9, "VW Transporter Aut.", 95, 125, 95, 125, 602, 812, 86, 116, 2120, 2420, 35, 10, 40, 0.75, 15, 210, 10, 90, 960},
}

func handler(w http.ResponseWriter, r *http.Request) {
	for i := range Cat {
		fmt.Fprintf(w, "\n Category: %3s - %40s or similar, starting from: %3.02f USD/Day", Cat[i].name, Cat[i].car, Cat[i].price)
	}

}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
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
	PriceListUrl = priceListUrl{
		pickupDate:          r.FormValue("pickupDate"),
		dropoffDate:         r.FormValue("dropoffDate"),
		driverAge:           r.FormValue("driverAge"),
		dropOffLocationCode: "300",
		pickupLocationCode:  "300",
		userType:            "Regular",
		priceListId:         "1029",
	}

	tmpl.Execute(w, struct{ Success bool }{true})
	log.Print("formHandler: Price list url:",PriceListUrl)
	JsonRaw := ReadPriceList(PriceListUrl)
	log.Print("formHandler: JsonRaw:",JsonRaw)
}

var AlbarRegPriceList = CarCategory{}

func main() {
	log.Print("Main: Price list url:",PriceListUrl)
	JsonRaw := ReadPriceList(PriceListUrl)
	err := json.Unmarshal([]byte(JsonRaw), &AlbarRegPriceList)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Starting web server on port 50003")
	http.HandleFunc("/form/", formHandler)
	http.HandleFunc("/results/", ResultsHandler)
	log.Fatal(http.ListenAndServe(":50003", nil))

}
