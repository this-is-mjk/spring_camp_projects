package worker

import (
	store "CF-RSS/store"
	model "CF-RSS/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"fmt"
	"time"
)
func PerformWork(m *mongo.Client) {
	maxCount := "100"
	sleepTime := 300
	coll := m.Database("CF-RSS").Collection("recent-actions-final")
	for {
		// find max time stamp
		maxTimeStamp := store.GetMaxTimeStamp(m)
		// fetch new
		dataAsInterfaceArray := store.Fetch(maxCount)
		// check old entry or new entry
		for index , action := range dataAsInterfaceArray {
			k := []model.Action{}
			if action.TimeSeconds > maxTimeStamp {
				fmt.Println(index)
				_ , err := coll.InsertOne(context.TODO(), store.ConvertToInterfaceArray(append(k,action)))
				if err != nil {
					log.Println(err)
					log.Fatal("can't save new entry")
				}
			}
		}
		fmt.Println("addition done going to sleep")
		// sleep for t time
		time.Sleep(time.Duration(sleepTime) * time.Second)
		log.Println("starting again")
	}
}


