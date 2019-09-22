package handler

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	targetUrl = "https://gold.tanaka.co.jp/commodity/souba/index.php"
)

const (
	timezone    = "Asia/Tokyo"
	offset      = 9 * 60 * 60
)


type GoldInfo struct {
	Date time.Time `json:"time"`
	RetailTax int `json:"retailTax"`
	PurchaseTax int `json:"purchaseTax"`
}

func FetchMetal() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する

		jst := time.FixedZone(timezone, offset)
		nowTime := time.Now().In(jst)

		var goldRetailTax, goldPurchaseTax string

		doc, err := goquery.NewDocument(targetUrl)
		if err != nil {
			fmt.Println(err)
		}

		// Fetch gold and platinum price
		doc.Find("#metal_price_sp").Each(func(_ int, s *goquery.Selection) {
			// Gold
			goldRetailTax = s.Children().Find("td.retail_tax").First().Text()
			goldPurchaseTax = s.Children().Find("td.purchase_tax").First().Text()
			// Platinum
		})




		// Format
		strGoldRetailTax := strings.Replace(goldRetailTax[0:5], ",", "", -1)
		strGoldPurchaseTax := strings.Replace(goldPurchaseTax[0:5], ",", "", -1)

		// Convert string to int
		intGoldRetailTax, _ := strconv.Atoi(strGoldRetailTax)
		intGoldPurchaseTax, _ := strconv.Atoi(strGoldPurchaseTax)



		// create json
		resGold := GoldInfo{
			Date: nowTime,
			RetailTax: intGoldRetailTax,
			PurchaseTax: intGoldPurchaseTax,
		}

		bytesDurationBytesJSON, _ := json.Marshal(resGold)


		return c.String(http.StatusOK, string(bytesDurationBytesJSON))
	}
}