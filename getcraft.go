package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func pullGetCraftInfluencersData(db *sql.DB) {
	urlPath := "https://getcraft.com/api/search/results?"
	totalCollapse := 0
	limit := 20
	skip := 20
	page := 1

	queryStrings := url.Values{}
	queryStrings.Add("timeZone", "Asia/Jakarta")
	queryStrings.Add("sessionId", "62f22caf93c860005da52ccc")
	queryStrings.Add("limit", strconv.Itoa(limit))
	queryStrings.Add("skip", strconv.Itoa(skip))
	queryStrings.Add("agg_term_limit", "200")
	queryStrings.Add("categoryIds[]", "59774c2082dd76d441d7cddd")
	queryStrings.Add("query", "getcraft-all-creators")
	queryStrings.Add("collapse", "{\"field\":\"profileId\",\"inner_hits\":{\"name\":\"services\",\"size\":100}}")
	queryStrings.Add("filter", "{\"term\":{\"availability\":true}}")

	getSinglePageInfluencersDataFromGetCraft(urlPath, queryStrings)

	for skip <= totalCollapse || totalCollapse == 0 {
		res := getSinglePageInfluencersDataFromGetCraft(urlPath, queryStrings)
		totalCollapse = res.TotalCollapse
		jsonResult, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		queryRes, err := db.Exec("INSERT INTO pages VALUES ($1, $2)", page, jsonResult)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Println(queryRes)

		skip += limit
		queryStrings.Set("skip", strconv.Itoa(skip))
		page += 1
		fmt.Println(page)
	}
}

func getSinglePageInfluencersDataFromGetCraft(urlPath string, queryStrings url.Values) GetCraftResponse {
	urlPath += queryStrings.Encode()

	res, err := http.Get(urlPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	getCraftResponse := GetCraftResponse{}

	json.Unmarshal(bodyBytes, &getCraftResponse)

	return getCraftResponse
}

func returnCSV(db *sql.DB) {
	newFormatData := []InfluencerCSV{}
	getCraftData := GetCraftResponse{}
	responseString := ""
	page := 1
	loop := true
	csvresult := [][]string{}
	csvFile, err := os.Create("influencers.csv")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer csvFile.Close()

	csvFileWriter := csv.NewWriter(csvFile)
	defer csvFileWriter.Flush()

	for loop {
		resDB := db.QueryRow("SELECT * FROM pages WHERE page = $1", page).Scan(&page, &responseString)
		if resDB != nil {
			loop = false
		}

		err := json.Unmarshal([]byte(responseString), &getCraftData)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		for _, v := range getCraftData.Products {
			tempString := []string{}

			temp := InfluencerCSV{}
			temp.Name = v.ProfileName
			temp.City = v.City
			temp.Country = v.Country
			temp.Occupation = strings.Join(v.Occupation, ", ")
			temp.InstagramURL = v.InstagramURL
			temp.LinkedInURL = v.LinkedinURL
			temp.TwitterURL = v.TwitterURL
			temp.YoutubeURL = v.YoutubeURL
			temp.YoutubeSubscribers = v.YoutubeFollowers
			temp.TwitterFollowers = v.TwitterFollowers
			temp.InstagramFollowers = v.IgFollowers
			temp.Cpv = v.Cpv
			temp.Views = v.Viewers

			for _, vpackage := range v.InnerHits.Services.Hits.Hits {
				temp.Packages += vpackage.Source.Name + "="
				temp.Packages += strconv.Itoa(vpackage.Source.Price) + ","
			}
			newFormatData = append(newFormatData, temp)

			tempString = append(
				tempString,
				temp.Name,
				temp.City,
				temp.Country,
				temp.Occupation,
				temp.InstagramURL,
				temp.LinkedInURL,
				temp.TwitterURL,
				temp.YoutubeURL,
				strconv.Itoa(temp.YoutubeSubscribers),
				strconv.Itoa(temp.TwitterFollowers),
				strconv.Itoa(temp.InstagramFollowers),
				strconv.Itoa(temp.Cpv),
				strconv.Itoa(temp.Views),
				temp.Packages,
			)

			csvresult = append(csvresult, tempString)

			err := csvFileWriter.Write(tempString)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}

		page += 1
	}
}
