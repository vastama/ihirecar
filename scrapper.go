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
	ImageLink                 string  `json:"ImageLink"`
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


type priceListUrl struct {
	url                 string `json:"url"`
	categoryId          string `json:"categoryId"`
	countryCode         string `json:"countryCode"`
	driverAge           string `json:"driverAge"`
	dropOffLocationCode string `json:"dropOffLocationCode"`
	dropoffDate         string `json:"dropoffDate"`
	guid                string `json:"guid"`
	pickupDate          string `json:"pickupDate"`
	pickupLocationCode  string `json:"pickupLocationCode"`
	priceListId         string `json:"priceListId"`
	promoCode           string `json:"promoCode"`
	userType            string `json:"userType"`
}

var PriceListUrl priceListUrl

func ReadPriceList(urlGet priceListUrl) string {
	log.Print("ReadPriceList: Price list url:", PriceListUrl)
	urlGet1 := albarPriceListUrl + "?" + "categoryId=&countryCode=" +
		"&driverAge=" + urlGet.driverAge +
		"&dropOffLocationCode=" + urlGet.dropOffLocationCode +
		"&dropoffDate=" + urlGet.dropoffDate + ":00" +
		"&guid=78284d45-4c0d-4cec-a8fb-ecfd217956f5" +
		"&pickupDate=" + urlGet.pickupDate + ":00" +
		"&pickupLocationCode=" + urlGet.pickupLocationCode +
		"&priceListId=" + urlGet.priceListId +
		"&promoCode=&userType=" + urlGet.userType
	log.Print("ReadPriceList: Price list urlGet1:", urlGet1)

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
