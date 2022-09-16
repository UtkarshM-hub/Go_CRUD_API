package routes

import (
	"context"
	"net/http"
	"time"

	getcollection "github.com/Go_CRUD_API/Collection"
	database "github.com/Go_CRUD_API/Database"
	model "github.com/Go_CRUD_API/Model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOne(c *gin.Context){
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	DB:=database.ConnectDB()
	postCollection:=getcollection.GetCollection(DB,"Post")

	postId:=c.Param("postId")
	var post model.Post

	defer cancel()

	if err:=c.BindJSON(&post);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"err"})
	}

	edited:=bson.M{"title":post.Title,"article":post.Article}

	objId,_:=primitive.ObjectIDFromHex(postId)
	result,err:=postCollection.UpdateOne(ctx,bson.M{"id":objId},bson.M{"$set":edited})

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"err1"})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success!","Data":result})

}