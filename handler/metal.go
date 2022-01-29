package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
)

/*
{
	"time": "2019-09-22T09:51:57.440183+09:00",
	"Gold": {
		"retailTax": 5674,
		"purchaseTax": 5588
	},
	"Platinum": {
		"retailTax": 5674,
		"purchaseTax": 5588
	}
}
*/

var (
	targetUrl = "https://gold.tanaka.co.jp/commodity/souba/index.php"
)

const (
	timezone = "Asia/Tokyo"
	offset   = 9 * 60 * 60
)

type Metal struct {
	Date     time.Time `json:"time"`
	GoldInfo GoldInfo  `json:"goldInfo"`
	Platinum Platinum  `json:"platinum"`
}

type GoldInfo struct {
	RetailTax   int `json:"retailTax"`
	PurchaseTax int `json:"purchaseTax"`
}

type Platinum struct {
	RetailTax   int `json:"retailTax"`
	PurchaseTax int `json:"purchaseTax"`
}

//func FetchMetal() echo.HandlerFunc {
//	return func(c echo.Context) error {

func FetchMetal(c echo.Context) error {
	var goldRetailTax, goldPurchaseTax, platinumRetailTax, platinumPurchaseTax string
	jst := time.FixedZone(timezone, offset)
	nowTime := time.Now().In(jst)

	res, err := http.Get(targetUrl)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Fetch gold and platinum price
	doc.Find("#metal_price_sp").Each(func(_ int, s *goquery.Selection) {
		// Gold
		goldRetailTax = s.Children().Find("td.retail_tax").First().Text()
		goldPurchaseTax = s.Children().Find("td.purchase_tax").First().Text()
		// Platinum
		platinumRetailTax = s.Children().Find("td.retail_tax").Eq(1).Text()
		platinumPurchaseTax = s.Children().Find("td.purchase_tax").Eq(1).Text()
	})

	// Format
	strGoldRetailTax := strings.Replace(goldRetailTax[0:5], ",", "", -1)
	strGoldPurchaseTax := strings.Replace(goldPurchaseTax[0:5], ",", "", -1)
	strPlatinumRetailTax := strings.Replace(platinumRetailTax[0:5], ",", "", -1)
	strPlatinumPurchaseTax := strings.Replace(platinumPurchaseTax[0:5], ",", "", -1)

	// Convert string to int
	intGoldRetailTax, _ := strconv.Atoi(strGoldRetailTax)
	intGoldPurchaseTax, _ := strconv.Atoi(strGoldPurchaseTax)
	intPlatinumRetailTax, _ := strconv.Atoi(strPlatinumRetailTax)
	intPlatinumPurchaseTax, _ := strconv.Atoi(strPlatinumPurchaseTax)

	// Set value to json
	resMetal := Metal{
		Date: nowTime,
		GoldInfo: GoldInfo{
			RetailTax:   intGoldRetailTax,
			PurchaseTax: intGoldPurchaseTax,
		},
		Platinum: Platinum{
			RetailTax:   intPlatinumRetailTax,
			PurchaseTax: intPlatinumPurchaseTax,
		},
	}

	bytesDurationBytesJSON, err := json.Marshal(resMetal)
	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, string(bytesDurationBytesJSON))
	//	}
}
