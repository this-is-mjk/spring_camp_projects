package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	// store "CF-RSS/store"
	model "CF-RSS/pkg/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "fmt"
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
func registerUser(c *gin.Context, userData *mongo.Collection) {
    var registrationData model.UserSignUp
    if err := c.ShouldBindJSON(&registrationData); err != nil {
        c.Error(err)
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }
    // fmt.Println(registrationData)
    _ , err := userData.InsertOne(context.TODO(), registrationData)
    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusOK, err)
        fmt.Println("can't save new entry")
    }
    c.JSON(http.StatusOK, "registration done")
}
func loginUser(c *gin.Context, userData *mongo.Collection) {
    var SignInData model.SignInData
    if err := c.ShouldBindJSON(&SignInData); err != nil {
        fmt.Println("here")
        c.AbortWithStatus(http.StatusBadRequest)
        c.Error(err)
        return
    }
    var savedUser model.UserSignUp
    err := userData.FindOne(context.TODO(),  bson.D{{"email", SignInData.Email}}).Decode(&savedUser)
    if err !=nil {
        if err == mongo.ErrNoDocuments {
            c.AbortWithStatus(404)
            return
		}
        log.Fatal(err)
    }
    fmt.Println(savedUser.Password, SignInData.Password)
    if savedUser.Password != SignInData.Password {
        c.AbortWithStatus(250)
        return
    }
    fmt.Println("found")
    c.IndentedJSON(http.StatusOK, savedUser)
}
func Subscribe(c *gin.Context, userData *mongo.Collection){
    var savedUser model.UserSignUp
    var SubscribeRequest model.SubscribeRequest
    if err := c.ShouldBindJSON(&SubscribeRequest); err != nil {
        c.Error(err)
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }
    filter := bson.D{
        {"email", SubscribeRequest.Email},
        {"password", SubscribeRequest.Password},
    }
    update := bson.D{
        {"$addToSet", bson.D{
            {"subscriptions", SubscribeRequest.BlogId},
        }},
    }
    if err := userData.FindOne(context.TODO(),  filter).Decode(&savedUser); err !=nil {
        if err == mongo.ErrNoDocuments {
            c.AbortWithStatus(250)
            return
		}
        log.Fatal(err)
    }
    userData.UpdateOne(context.TODO(), filter, update)
    return
}
func Unsubscribe(c *gin.Context, userData *mongo.Collection){
    var SubscribeRequest model.SubscribeRequest
    if err := c.ShouldBindJSON(&SubscribeRequest); err != nil {
        c.Error(err)
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }
    filter := bson.D{
        {"email", SubscribeRequest.Email},
        {"password", SubscribeRequest.Password},
    }
    update := bson.D{
        {"$pull", bson.D{
            {"subscriptions", SubscribeRequest.BlogId},
        }},
    }
    result, err:= userData.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    if result.MatchedCount == 0 {
        // No matching documents found, handle accordingly
        c.AbortWithStatus(http.StatusNotFound)
        return
    }
    return


}
func userRecentAction(c *gin.Context, userData *mongo.Collection, reader *mongo.Collection) {
    var SignInData model.SignInData
    if err := c.ShouldBindJSON(&SignInData); err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        c.Error(err)
        return
    }
    var savedUser model.UserSignUp
    err := userData.FindOne(context.TODO(),  bson.D{{"email", SignInData.Email}}).Decode(&savedUser)
    if err !=nil {
        if err == mongo.ErrNoDocuments {
            c.AbortWithStatus(404)
            return
		}
        log.Fatal(err)
    }
    fmt.Println(savedUser.Password, SignInData.Password)
    if savedUser.Password != SignInData.Password {
        c.AbortWithStatus(250)
        return
    }
    var results model.Actions
    cur, err := reader.Find(context.TODO(), bson.D{{"blogentry.id", bson.D{{"$in", savedUser.Subscriptions}}}})
    for cur.Next(context.TODO()) {
        //Create a value into which the single document can be decoded
        var elem model.Action
        err := cur.Decode(&elem)
        fmt.Println(elem)
        if err != nil {
            log.Fatal(err)
        }
        results =append(results, elem)
        
    }
    if err != nil {
        log.Fatal(err)
    }
    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }
    c.IndentedJSON(http.StatusOK, results)
}











// func authenticateReq(c *gin.Context, userData *mongo.Collection) {
//     // read cookie
//     var username, err1 = c.Cookie("userName")
//     var password, err2 = c.Cookie("password")
//     if (err1 != nil || err2 != nil) {
//       c.AbortWithStatus(300)  
//     }
//     loginUser(c, user)
   
// }