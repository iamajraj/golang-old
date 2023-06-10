package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"It worked!",
		})
	})

	r.POST("/api", func(ctx *gin.Context) {
		b, err := ioutil.ReadAll(ctx.Request.Body)
		var x map[string]interface{}
		err = json.Unmarshal(b, &x);
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":"Failed",
			})
		}
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":"Failed",
			})
		}
		ctx.JSON(200, gin.H{
			"d": x,
			"status":"success",
		})
		fmt.Println(x)

	})

	r.Run(":3000")
}