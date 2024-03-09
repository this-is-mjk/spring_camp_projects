package worker

import (
	store "CF-RSS/pkg/store"
	// model "CF-RSS/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"fmt"
	"time"
)
func PerformWork(m *mongo.Client) {
	maxCount := "5"
	sleepTime := 100
	coll := m.Database("CF-RSS").Collection("recent-actions-final")
	for {
		// find max time stamp
		maxTimeStamp := store.GetMaxTimeStamp(m)
		// if data is lost somehow create new
		if maxTimeStamp == 0 {
			dataAsInterfaceArray := store.ConvertToInterfaceArray(store.Fetch(maxCount))
			store.StoreRecentActionsInTheDatabase(m , dataAsInterfaceArray)
			continue
		}
		// fetch new
		dataAsInterfaceArray := store.Fetch(maxCount)
		// check old entry or new entry
		for index , action := range dataAsInterfaceArray {
			if action.TimeSeconds > maxTimeStamp {
				fmt.Println(index)
				_ , err := coll.InsertOne(context.TODO(), action)
				if err != nil {
					fmt.Println(err)
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


