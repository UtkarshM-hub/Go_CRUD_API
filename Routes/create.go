package routes

import (
	getcollection "github.com/Go_CRUD_API/Collection"
	database "github.com/Go_CRUD_API/Database"
	model "github.com/Go_CRUD_API/Model"
	"context"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *gin.Context){
	var DB=database.ConnectDB()
	var postCollection=getcollection.GetCollection(DB,"Post")

	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	post:=new(model.Post)
	
	defer cancel()
	
	if err:=c.BindJSON(&post);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"Message"})
		log.Fatal("err")
		return
	}
	
	postPayload:=model.Post{
		ID: primitive.NewObjectID(),
		Title: post.Title,
		Article: post.Article,
	}

	result,err:=postCollection.InsertOne(ctx,postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Message1"})
		return
	}

	
	c.JSON(200,gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}
