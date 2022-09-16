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

func GetOne(c *gin.Context){
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	DB:=database.ConnectDB()
	postCollection:=getcollection.GetCollection(DB,"Post")

	postId:=c.Param("postId")
	var result model.Post

	defer cancel()

	objId,_:=primitive.ObjectIDFromHex(postId)
	err:=postCollection.FindOne(ctx,bson.M{"id":objId}).Decode(&result)

	// interface keywork means the value could be anything
	// res:=map[string]interface{}{"data":result}

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":err})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success!","Data":result})

}