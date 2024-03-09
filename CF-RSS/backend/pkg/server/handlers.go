package server

import (
	// "encoding/json"
	"log"
	"net/http"

	// "text/template/parse"

	"github.com/gin-gonic/gin"

	// store "CF-RSS/store"
	model "CF-RSS/pkg/model"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func Signup(c *gin.Context, userDataBase *mongo.Collection) {
	var registrationData model.User
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// fmt.Println(registrationData)
	_, err := userDataBase.InsertOne(context.TODO(), registrationData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, err)
		fmt.Println("failed to save new user")
		return
	}
	c.JSON(http.StatusOK, "registration done")
}
func loginUser(c *gin.Context, userDataBase *mongo.Collection) {
	var SignInData model.UserLogin
	if err := c.ShouldBindJSON(&SignInData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		c.Error(err)
		return
	}
	var savedUser model.User
	err := userDataBase.FindOne(context.TODO(), bson.D{{"username", SignInData.Username}}).Decode(&savedUser)
	fmt.Print(SignInData.Username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, "User not found")
			return
		}
		log.Fatal(err)
	}
	// fmt.Println(savedUser.Password, SignInData.Password)
	if savedUser.Password != SignInData.Password {
		c.JSON(http.StatusBadRequest, "Invalid password")
		return
	}
	fmt.Println("found")
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": SignInData.Username,
		// will expair in 4 hours
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("123456"))
	fmt.Println(tokenString, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cant create token")
		return
	}
	// set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*4, "/", "", false, false)
	c.SetCookie("Username", savedUser.Username, 3600*4, "/", "", false, false)
	c.SetCookie("Email", savedUser.Email, 3600*4, "/", "", false, false)
	c.IndentedJSON(http.StatusOK, gin.H{
		"subscriptions": savedUser.Subscriptions,
	})
}
// middleware auth function
func authenticateReq(c *gin.Context) {
	// get the token form the cooki
	fmt.Printf("Request Headers: %v\n", c.Request.Header)
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Print("here at tokenstring")
		fmt.Print(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Print(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("123456"), nil
	})
	if err != nil {
		fmt.Println("here at jwt error")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// now we have a authorized token
		//check if the token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// lets find the user
		// fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Print("auth done")
	c.Next()
}
func getBlogs(c *gin.Context, reader *mongo.Collection, userData *mongo.Collection) {
	//find records
	username, err := c.Cookie("Username")
	if err != nil {
		fmt.Print("dont contain username")
		fmt.Print(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var savedUser model.User
	err = userData.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&savedUser)
	fmt.Print(username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, "User not found")
			return
		}
		log.Fatal(err)
	}
	//pass these options to the Find method
	findOptions := options.Find()
	//Set the limit of the number of record to find
	//findOptions.SetLimit(10)
	//Define an array in which you can store the decoded documents
	var results model.Actions
	//Passing the bson.D{{}} as the filter matches  documents in the collection
	cur, err := reader.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
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
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	//Close the cursor once finished
	cur.Close(context.TODO())
	// showing on chrome
	fmt.Print("returning the data")
	c.IndentedJSON(http.StatusOK, gin.H{
		"posts": results,
		"subscriptions": savedUser.Subscriptions,
	})
	// c.IndentedJSON(http.StatusOK, gin.H{
	// 	"subscriptions": savedUser.Subscriptions,
	// })
	// sending the userSubscriptions
}
func SubscribeRequest(c *gin.Context, userData *mongo.Collection) {
	var SubscribeRequest model.SubscribeRequest
	username, err := c.Cookie("Username")
	if err != nil {
		fmt.Print("dont contain username")
		fmt.Print(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err := c.ShouldBindJSON(&SubscribeRequest); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var update bson.D
	filter := bson.D{
		{"username", username},
	}
	if SubscribeRequest.Action == true{
		update = bson.D{
			{"$addToSet", bson.D{
				{"subscriptions", SubscribeRequest.Id},
			}},
		}
	} else if SubscribeRequest.Action == false{
		update = bson.D{
			{"$pull", bson.D{
				{"subscriptions", SubscribeRequest.Id},
			}},
		}
	}
	userData.UpdateOne(context.TODO(), filter, update)
	return
}
