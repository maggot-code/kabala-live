/*
 * @Author: maggot-code
 * @Date: 2021-07-08 10:45:33
 * @LastEditors: maggot-code
 * @LastEditTime: 2021-07-08 11:30:23
 * @Description: file content
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/video", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "video",
		})
	})

	s := &http.Server{
		Addr:    ":9960",
		Handler: router,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
