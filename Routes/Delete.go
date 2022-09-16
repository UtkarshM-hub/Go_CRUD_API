package routes

import (
	"context"
	"net/http"
	"time"

	getcollection "github.com/Go_CRUD_API/Collection"
	database "github.com/Go_CRUD_API/Database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOne(c *gin.Context){
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	DB:=database.ConnectDB()
	postCollection:=getcollection.GetCollection(DB,"Post")

	postId:=c.Param("postId")

	defer cancel()

	objId,_:=primitive.ObjectIDFromHex(postId)
	result,err:=postCollection.DeleteOne(ctx,bson.M{"id":objId})

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}


	c.JSON(http.StatusCreated,gin.H{"message":"Article was deleted!","Data":result})

}