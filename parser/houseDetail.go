package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/TDTzzz/crawlerLianjia/engine"
	"github.com/TDTzzz/crawlerLianjia/model"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func HouseDetail(contents []byte) engine.ParseResult {
	var res engine.ParseResult
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}
	region := ""
	subRegion := ""
	//解析当前的region+sub_region
	dom.Find("div.position>dl:nth-child(2)>dd>div[data-role=ershoufang]>div>a.selected").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			region = selection.Text()
		} else if i == 1 {
			subRegion = selection.Text()
		}
	})

	dom.Find(".info.clear").Each(func(i int, selection *goquery.Selection) {
		name := selection.Find(".title>a").Text()
		//id, _ := selection.Find(".title>a").Attr("data-housecode")
		houseHref, _ := selection.Find(".title>a").Attr("href")
		id := GetNumFromStr(houseHref)

		community := selection.Find(".flood>.positionInfo>a:first-of-type").Text()
		communityHref, _ := selection.Find(".flood>.positionInfo>a:first-of-type").Attr("href")
		communityInt := GetNumFromStr(communityHref)

		houseInfoStr := selection.Find(".address>.houseInfo").Text()
		houseInfo := getHouseInfo(houseInfoStr)
		//price
		totalPriceStr := selection.Find(".priceInfo>.totalPrice>span").Text()
		unitPriceStr := selection.Find(".priceInfo>.unitPrice>span").Text()

		unitPrice := GetFloatFromStr(unitPriceStr)
		totalPrice, err := strconv.ParseFloat(totalPriceStr, 64)
		if err != nil {
			log.Println(err)
			totalPrice = 0
		}
		area, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(totalPrice)*10000/float64(unitPrice)), 64)

		res.Items = append(res.Items, model.HouseDetail{
			Id:          id,
			Region:      region,
			SubRegion:   subRegion,
			Name:        name,
			Community:   community,
			CommunityId: communityInt,
			TotalPrice:  totalPrice,
			UnitPrice:   unitPrice,
			Area:        area,
			HouseType:   houseInfo.HouseType,
			Toward:      houseInfo.Toward,
			Floor:       houseInfo.Floor,
			BuildYear:   houseInfo.BuildYear,
			BuildType:   houseInfo.BuildType,
			Villa:       houseInfo.Villa,
			Date:        time.Now(),
		})
	})
	return res
}

func getHouseInfo(houseInfoStr string) model.HouseInfo {
	infoArr := strings.Split(houseInfoStr, "|")
	infoLen := len(infoArr)
	var houseInfo model.HouseInfo
	if infoLen == 7 {
		houseInfo = model.HouseInfo{
			HouseType: infoArr[0],
			Area:      infoArr[1],
			Toward:    infoArr[2],
			Level:     infoArr[3],
			Floor:     infoArr[4],
			BuildYear: infoArr[5],
			BuildType: infoArr[6],
		}
	} else if infoLen == 6 {
		houseInfo = model.HouseInfo{
			HouseType: infoArr[0],
			Area:      infoArr[1],
			Toward:    infoArr[2],
			Level:     infoArr[3],
			Floor:     infoArr[4],
			//BuildYear: "",
			BuildType: infoArr[5],
		}
	} else if infoLen == 8 {
		houseInfo = model.HouseInfo{
			HouseType: infoArr[0],
			Area:      infoArr[1],
			Toward:    infoArr[2],
			Level:     infoArr[3],
			Floor:     infoArr[4],
			BuildYear: infoArr[5],
			BuildType: infoArr[6],
			Villa:     infoArr[7], //别墅
		}
	} else if infoLen == 4 || infoLen == 5 {
		//车位的话 只看前三个数据得了
		houseInfo = model.HouseInfo{
			HouseType: infoArr[0],
			Area:      infoArr[1],
			Toward:    infoArr[2],
		}
	} else {
		log.Println("houseInfo Parse Err:")
		log.Println(infoArr)
		os.Exit(100)
	}
	return houseInfo
}

func GetNumFromStr(str string) int {
	numRegexp := regexp.MustCompile(`\d+`)
	params := numRegexp.FindAllString(str, -1)
	num, _ := strconv.Atoi(params[0])
	return num
}

func GetFloatFromStr(str string) float64 {
	numRegexp := regexp.MustCompile(`\d+\.?\d*`)
	params := numRegexp.FindAllString(str, -1)
	float, _ := strconv.ParseFloat(params[0], 64)
	return float
}
