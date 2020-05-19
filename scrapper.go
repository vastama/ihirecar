package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const albarPriceListUrl = "https://rent.albar.co.il/umbraco/surface/Fleet/GetResults"

type CarCategory []struct {
	Id                        string  `json:"Id"`
	CarCategoryId             string  `json:"CarCategoryId"`
	CarCategoryCode           string  `json:"CarCategoryCode"`
	CarCategoryName           string  `json:"CarCategoryName"`
	Currency                  string  `json:"Currency"`
	CurrencySymbol            string  `json:"CurrencySymbol"`
	NumberOfSmallBags         int8    `json:"NumberOfSmallBags"`
	NumberOfLargeBags         int     `json:"NumberOfLargeBags"`
	AirConditioner            bool    `json:"AirConditioner"`
	Airbags                   int     `json:"Airbags"`
	Radio                     bool    `json:"Radio"`
	PowerSteering             bool    `json:"PowerSteering"`
	Gear                      string  `json:"Gear"`
	Doors                     int     `json:"Doors"`
	PriceExtraKm              float32 `json:"PriceExtraKm"`
	KmQuota                   float32 `json:"KmQuota"`
	UnitPrice                 float32 `json:"UnitPrice"`
	Days                      int     `json:"Days"`
	Price                     float32 `json:"Price"`
	Unit                      string  `json:"Unit"`
	ImmidiateConfirmation     bool    `json:"ImmidiateConfirmation"`
	DiscountRate              int     `json:"DiscountRate"`
	PriceAfterDiscount        float32 `json:"PriceAfterDiscount"`
	RateCodeSelected          string  `json:"RateCodeSelected"`
	LocationIDSupplier        string  `json:"LocationIDSupplier"`
	SupplierLogoImageLink     string  `json:"SupplierLogoImageLink"`
	MinimumAge                int     `json:"MinimumAge"`
	RentalDays                int     `json:"RentalDays"`
	MinimumDays               int     `json:"MinimumDays"`
	DrivingExperienceRequired int     `json:"DrivingExperienceRequired"`
	WinterTiresFee            string  `json:"WinterTiresFee"`
	RentalContractFee         string  `json:"RentalContractFee"`
	SupplierCode              string  `json:"SupplierCode"`
	Terms                     string  `json:"Terms"`
	IncludedInVoucher         string  `json:"IncludedInVoucher"`
	Fees                      string  `json:"Fees"`
	CarGuid                   string  `json:"CarGuid"`
	ComparePrice              float32 `json:"ComparePrice"`
}

var age int8

var url3 = []string{"https://rent.albar.co.il/en/rent-in-israel/step/1"}
var url2 = []string{"https://rent.albar.co.il/umbraco/surface/Fleet/GetResults?categoryId=&countryCode=&driverAge=25&dropOffLocationCode=339&dropoffDate=2020-05-24T09:00:00&guid=78284d45-4c0d-4cec-a8fb-ecfd217956f5&pickupDate=2020-05-21T09:00:00&pickupLocationCode=339&priceListId=1035&promoCode=&userType=Regular"}
var url1 = []string{"https://rent.albar.co.il/umbraco/surface/Fleet/GetResults?categoryId=&countryCode=&driverAge=28&dropOffLocationCode=300&dropoffDate=2020-05-28T09:00:00&guid=d4f5ebf9-c72a-4cb3-bd8e-e5a0c7ae005e&pickupDate=2020-05-25T09:00:00&pickupLocationCode=300&priceListId=1029&promoCode=&userType=Regular"}
var urls = []string{
	albarPriceListUrl + "?" +
		"categoryId=" +
		"&countryCode=" +
		"&driverAge=21" +
		"&dropOffLocationCode=300" +
		"&dropoffDate=2020-05-27T09:00:00" +
		"&guid=f11d41a2-b1f3-4ae4-8eb2-40d619c8a0ca" +
		"&pickupDate=2020-05-24T09:00:00" +
		"&pickupLocationCode=300" +
		"&priceListId=1029" +
		"&promoCode=" +
		"&userType=Regular",
}

type priceListUrl struct {
	url                 string `json:"url"`
	categoryId          string `json:"categoryId"`
	countryCode         string `json:"countryCode"`
	driverAge           string `json:"driverAge"`
	dropOffLocationCode string    `json:"dropOffLocationCode"`
	dropoffDate         string `json:"dropoffDate"`
	guid                string `json:"guid"`
	pickupDate          string `json:"pickupDate"`
	pickupLocationCode  string    `json:"pickupLocationCode"`
	priceListId         string    `json:"priceListId"`
	promoCode           string `json:"promoCode"`
	userType            string `json:"userType"`
}

var PriceListUrl priceListUrl

func ReadPriceList(urlGet priceListUrl) string {
	log.Print("ReadPriceList: Price list url:", PriceListUrl)
	urlGet1 := albarPriceListUrl + "?" + "categoryId=&countryCode=" +
		"&driverAge=" +	urlGet.driverAge +
		"&dropOffLocationCode=" + urlGet.dropOffLocationCode +
		"&dropoffDate=" + urlGet.dropoffDate +  "T09:00:00" +
		"&guid=78284d45-4c0d-4cec-a8fb-ecfd217956f5" +
		"&pickupDate=" + urlGet.pickupDate +"T09:00:0" +
		"0&pickupLocationCode=" + urlGet.pickupLocationCode +
		"&priceListId=" + urlGet.priceListId +
		"&promoCode=&userType=" + urlGet.userType
	log.Print("ReadPriceList: Price list urlGet1:", urlGet1)
	//"&dropOffLocationCode=339&dropoffDate=2020-05-24T09:00:00&guid=78284d45-4c0d-4cec-a8fb-ecfd217956f5&pickupDate=2020-05-21T09:00:00&pickupLocationCode=339&priceListId=" + urlGet.priceListId + "&promoCode=&userType=" + urlGet.userType

	resp, err := http.Get(urlGet1)
	if err != nil {
		log.Fatal("Fatal on Get Url:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Err: %v", err)
	}

	str := string(body)
	var z []string = strings.Split(str, "[")
	z = strings.Split(z[1], "]")
	JsonRaw := "[" + z[0] + "]"
	return JsonRaw
}

func main1() {

	//var AlbarRegPriceList = CarCategory{}

	/*var priceListUrlHebRegAlbar = priceListUrl{
		albarPriceListUrl,
		"",
		"",
		"28",
		"300",
		"2020-05-24T09:00:00",
		"f11d41a2-b1f3-4ae4-8eb2-40d619c8a0ca",
		"2020-05-21T09:00:00",
		"300",
		"1029",
		"",
		"Regular",
	}*/

	//fmt.Printf("%v\n\n", priceListUrlHebRegAlbar.priceListId)

	//q := url.QueryEscape(strings.Join(priceListUrlHebRegAlbar, " "))

	//JsonRaw := ReadPriceList(urls)
	//err := json.Unmarshal([]byte(JsonRaw), &AlbarRegPriceList)
	//if err != nil {
	//	panic(err.Error())
	//}
	//for x, _ := range AlbarRegPriceList {
	/*	fmt.Printf("%d category: %s %s %d %f\n %v\n", x,
		AlbarRegPriceList[x].CarCategoryCode,
		AlbarRegPriceList[x].CarCategoryName,
		AlbarRegPriceList[x].Days,
		AlbarRegPriceList[x].Price,
		AlbarRegPriceList[x])*/
}

//}

//2020-05-21T09:00:00
//fmt.Println(time.Now())
