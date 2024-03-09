package main

import (
	store "CF-RSS/pkg/store"
	worker "CF-RSS/pkg/worker"
	server "CF-RSS/pkg/server"
	"sync"
	// "fmt"
)
func DataCollection() {
	// maxCount := "100"
	// // first time fetch
	// dataAsInterfaceArray := store.ConvertToInterfaceArray(store.Fetch(maxCount))
	// create connection
	mongoClient := store.OpenConnectionWithMongoDB()
	// fmt.Println(mongoClient)
	// store data for first time
	// store.StoreRecentActionsInTheDatabase( mongoClient , dataAsInterfaceArray)
	// starting worker to take recent actions in fixed time
	worker.PerformWork(mongoClient)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go DataCollection()
	go server.CreateRoutes()

	wg.Wait()

}