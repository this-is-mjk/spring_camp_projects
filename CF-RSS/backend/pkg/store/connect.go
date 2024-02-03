package store

import (
	"log"
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	model "CF-RSS/pkg/model"
)
var data struct {
	Status string `json:"status"`
	Result model.Actions `json:"result"`
}
var time struct {
	Timestamp int64 `bson:"timeseconds"`
}
func OpenConnectionWithMongoDB() *mongo.Client {
	const uri = "mongodb://localhost:27017"
	client, err1 := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err1 != nil {
		fmt.Println("can't connect to database")
	}
	return client
}
func StoreRecentActionsInTheDatabase(m *mongo.Client,  dataAsInterfaceArray []interface{}) {
	coll := m.Database("CF-RSS").Collection("recent-actions-final")
	_ , err3 := coll.InsertMany(context.TODO(), dataAsInterfaceArray)
	if err3 != nil {
		log.Println(err3)
		log.Fatal("can't save intital data")
	}
	log.Println("saved the initial data")
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
func Fetch(maxCount string) model.Actions {
	url := "https://codeforces.com/api/recentActions?maxCount=" + maxCount
	client := &http.Client{}
    req, err1 := http.NewRequest("GET", url , nil)
    if err1 != nil {
        log.Fatalln("error making http request:", err1)
    }

    resp, err2 := client.Do(req)
    if err2 != nil {
		log.Fatalln("Error sending request to server:", err2)
    }
    body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(body, &data)
	// fmt.Printf("%T\n", data.Result)
    // SaveToFile(data.Result, "data")
	return data.Result
	// var  dataAsInterfaceArray []interface{}
	// for _ , action := range data.Result {
	// 	dataAsInterfaceArray = append(dataAsInterfaceArray, action)
	// }
	// return dataAsInterfaceArray
	
}
func ConvertToInterfaceArray(data model.Actions) []interface{} {
	var  dataAsInterfaceArray []interface{}
	for _ , action := range data {
		dataAsInterfaceArray = append(dataAsInterfaceArray, action)
	}
	return dataAsInterfaceArray
}
func GetMaxTimeStamp(m *mongo.Client) int64 {
	coll := m.Database("CF-RSS").Collection("recent-actions-final")
	// finding max
	projection := bson.M{"timeseconds": 1}
	sort := bson.D{{Key: "timeseconds", Value: -1}}
	options := options.FindOne().SetProjection(projection).SetSort(sort)
	err := coll.FindOne(context.TODO(), bson.D{}, options).Decode(&time)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No documents found. creating new one")
			return 0
		} else {
			log.Fatal(err)
		}
	}
	return time.Timestamp
}