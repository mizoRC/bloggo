package api

import (
	DB "../db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type Post struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Title         string             `json:"title" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Markdown      string             `json:"markdown" binding:"required"`
	Date          time.Time
	FormattedDate string
}

type PostPreview struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Title         string             `json:"title" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Date          time.Time
	FormattedDate string
}

type Fields struct {
	ID            int `bson:"_id"`
	Title         int
	Description   int
	Markdown      int
	Date          int
	FormattedDate int
}

func GetPosts(c *gin.Context) {
	// Here's an array in which you can store the decoded documents
	var posts []*PostPreview

	collection := DB.MongoClient.Database("bloggo").Collection("posts")

	projection := Fields{
		ID:            1,
		Title:         1,
		Description:   1,
		Markdown:      1,
		Date:          1,
		FormattedDate: 1,
	}

	//"$project": {"title": 1, "description": 1, "date": 1, "formatteddate": 1}

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetProjection(projection)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Panic(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var post PostPreview
		err := cur.Decode(&post)
		if err != nil {
			log.Panic(post)
		}

		posts = append(posts, &post)
	}

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {

}

func SavePost(c *gin.Context) {
	var post Post
	c.BindJSON(&post)
	now := time.Now()
	formattedDate := fmt.Sprintf("%02d-%02d-%d %02d:%02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())
	post.ID = primitive.NewObjectID()
	post.Date = now
	post.FormattedDate = formattedDate

	/*
		log.Printf("Title => %v", post.Title)
		log.Printf("Markdown => %v", post.Markdown)
		log.Printf("Date => %v", post.Date)
		log.Printf("FormattedDate => %v", post.FormattedDate)
	*/

	// STORE POST IN MONGO
	collection := DB.MongoClient.Database("bloggo").Collection("posts")
	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, post)
	}

	log.Printf("Inserted new post: %v", insertResult.InsertedID)

	c.JSON(http.StatusOK, post)
}
