package main

import (
	"io/ioutil"
	"encoding/json"
 	models "cf/models"
	"fmt"
	"net/http"
	"log"
	"os"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	maxCount := "1"
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
		Status string `json:"status"`
		Result models.Actions `json:"result"`
	}
	json.Unmarshal(body, &resJson)
	
	SaveToFile(resJson.Result, "data")
	// fmt.Println(resJson.Result)

	// mongo 
	// connect
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("can't connect to database")
	} 
	coll := client.Database("CF-RSS").Collection("recent-actions")
	var  toInterfacefrombody []interface{}
	for _ , action := range body {
		toInterfacefrombody = append(toInterfacefrombody, action)
	}
	_ , err3 := coll.InsertMany(context.TODO(), toInterfacefrombody)

	if err3 != nil {
		log.Println(err3)
		log.Println("can't save data")
	} else {
		fmt.Println("done")
	}

}