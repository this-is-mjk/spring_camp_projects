package main

import (
	"io/ioutil"
	"encoding/json"
	// "fmt"
	"net/http"
	"log"
	"os"
)
type C struct {
	Id int `json:"id"`
	CreationTimeSeconds int `json:"creationTimeSeconds"`
	CommentatorHandle string `json:"commentatorHandle"`
	Locale string `json:"locale"`
	Text string `json:"text"`
	ParentCommentId int `json:"parentCommentId"`
	Rating int `json:"rating"`
}
type BE struct {
	OriginalLocale string `json:"originalLocale"`
	CreationTimeSeconds int `json:"creationTimeSeconds"`
	Rating int  `json:"rating"`
	AuthorHandle string `json:"authorHandle"`
	ModificationTimeSeconds int `json:"modificationTimeSeconds"`
	Id int `json:"id"`
	Title string `json:"title"`
	Locale string `json:"locale"`
	Tags []string `json:"comment"`
}

type Actions []Action

type Action struct {
	Time int `json:"timeSeconds"`
	Blockentry BE `json:"blogEntry"`
	Comment  C `json:"comment"`
}

func SaveToFile(data interface{}, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}


func main() {
	maxCount := "100"
	// requesting
	url := "https://codeforces.com/api/recentActions?maxCount=" + maxCount
	resp, err1 := http.Get(url)
	// checking error
	if err1 != nil {
		log.Fatal("Error in request process")
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalln(err2)
	 }
	defer resp.Body.Close()

	var resJson struct {
		Result Actions `json:"result"`
	}
	json.Unmarshal(body, &resJson)
	SaveToFile(resJson.Result, "data")
	// fmt.Println(resJson.Result)
	


	

}