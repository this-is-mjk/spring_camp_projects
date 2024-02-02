package main

import (
	store "CF-RSS/store"
	worker "CF-RSS/worker"
	server "CF-RSS/server"
	"sync"
	// "fmt"
)
func DataCollection() {
	maxCount := "100"
	// first time fetch
	dataAsInterfaceArray := store.ConvertToInterfaceArray(store.Fetch(maxCount))
	// create connection
	mongoClient := store.OpenConnectionWithMongoDB()
	// fmt.Println(mongoClient)
	// store data for first time
	store.StoreRecentActionsInTheDatabase( mongoClient , dataAsInterfaceArray)
	// starting worker to take recent actions in every 5 min
	worker.PerformWork(mongoClient)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go DataCollection()
	go server.CreateRoutes()

	wg.Wait()

}