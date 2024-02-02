package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	// store "CF-RSS/store"
	model "CF-RSS/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func getBlogs(c *gin.Context, reader *mongo.Collection) {
	//find records
    //pass these options to the Find method
    findOptions := options.Find()
    //Set the limit of the number of record to find
    //findOptions.SetLimit(10)
    //Define an array in which you can store the decoded documents
    var results model.Actions
    //Passing the bson.D{{}} as the filter matches  documents in the collection
    cur, err := reader.Find(context.TODO(), bson.D{{}}, findOptions)
    if err !=nil {
        log.Fatal(err)
    }
	//Finding multiple documents returns a cursor
    //Iterate through the cursor allows us to decode documents one at a time

    for cur.Next(context.TODO()) {
        //Create a value into which the single document can be decoded
        var elem model.Action
        err := cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }
        results =append(results, elem)
    }
    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }
    //Close the cursor once finished
    cur.Close(context.TODO())
	// showing on chrome
    c.IndentedJSON(http.StatusOK, results)
}