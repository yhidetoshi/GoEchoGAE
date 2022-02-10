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

/* Response
{
  "time": "2022-02-10T14:17:13.354564+09:00",
  "gold": 7536,
  "platinum": 4318
}
*/

const (
	targetUrl   = "https://gold.mmc.co.jp/"
	goldNum     = 1
	platinumNum = 13
	timezone    = "Asia/Tokyo"
	offset      = 9 * 60 * 60
)

type Metal struct {
	Date          time.Time `json:"time"`
	GoldPrice     int       `json:"gold"`
	PlatinumPrice int       `json:"platinum"`
}

func FetchMetal(c echo.Context) error {
	var goldPrice, platinumPrice string
	var count int

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
	// GOLD
	doc.Find("div.p-table-scroll--sticky > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		// Gold
		if count == goldNum {
			goldPrice = s.Children().Find("span.c-table__text--xl").First().Text()
			fmt.Println(goldPrice)
		}
		// Platinum
		if count == platinumNum {
			platinumPrice = s.Children().Find("span.c-table__text--xl").First().Text()
			fmt.Println(platinumPrice)
		}
		count++
	})

	// Format
	strGoldPrice := strings.Replace(goldPrice, ",", "", -1)
	strPlatinumPrice := strings.Replace(platinumPrice, ",", "", -1)

	// Convert string to int
	intGoldPrice, err := strconv.Atoi(strGoldPrice)
	if err != nil {
		fmt.Println(err)
	}
	intPlatinumPrice, err := strconv.Atoi(strPlatinumPrice)
	if err != nil {
		fmt.Println(err)
	}

	// Set value to json
	resMetal := Metal{
		Date:          nowTime,
		GoldPrice:     intGoldPrice,
		PlatinumPrice: intPlatinumPrice,
	}

	bytesDurationBytesJSON, err := json.Marshal(resMetal)
	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, string(bytesDurationBytesJSON))
}
