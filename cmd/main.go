package main

import (
	"fmt"
	"log"

	klipper "kommander/klipper"
	serial "kommander/serialcom"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	port, err := serial.NewSerialStreamer("/tmp/printer", 250000)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	read, err := port.Reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read %v\n", read)
	apiCmd := r.Group("/api/")
	{
		apiCmd.GET("/cmd/ext:id", klipper.GetExtruderTemp(port))
		apiCmd.GET("/cmd/bed:id", klipper.GetBedTemp(port))
		apiCmd.POST("/cmd/ext:id/:temp", klipper.SetExtruderTemp)
		apiCmd.POST("/cmd/bed:id/:temp", klipper.SetBedTemp)
	}
	r.GET("/setup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "setup",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8090") // listen and serve on 0.0.0.0:8080
}
