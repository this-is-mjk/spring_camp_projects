package main

import (
	store "CF-RSS/store"
	worker "CF-RSS/worker"
	// "fmt"
)

func main() {
	maxCount := "100"
	// first time fetch
	dataAsInterfaceArray := store.ConvertToInterfaceArray(store.Fetch(maxCount))
	// create connection
	mongoClient := store.OpenConnectionWithMongoDB()
	// store data for first time
	store.StoreRecentActionsInTheDatabase( mongoClient , dataAsInterfaceArray)
	// starting worker to take recent actions in every 5 min
	worker.PerformWork(mongoClient) 
}