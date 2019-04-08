package main

import (
	"fmt"
	"log"
	"strings"

	serial "kommander/serialcom"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GetBedTemp - returns bed temp
func GetBedTemp(p serial.SerialStreamer) gin.HandlerFunc {
	//
	fn := func(c *gin.Context) {
		bedID := c.Param("id")
		_, err := p.Writer.WriteString("M105\n")
		out, err := p.Reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
		//temp := strings.Split(out
		c.JSON(200, gin.H{
			"bed" + bedID: out,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetExtruderTemp - returns bed temp
func GetExtruderTemp(p serial.SerialStreamer) gin.HandlerFunc {
	//
	fn := func(c *gin.Context) {
		extID := c.Param("id")
		_, err := p.Writer.WriteString("M105\n")
		err = p.Writer.Flush()
		out, err := p.Reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(strings.Split(out, " ")[1][3:])
		// temp := strings.TrimRight((strings.Split(string(out), ":")[1]), "\n\u0000\u0000\u0000")
		c.JSON(200, gin.H{
			"ext" + extID: out,
		})
	}
	return gin.HandlerFunc(fn)
}

// SetExtruderTemp - returns extruder temp
func SetExtruderTemp(c *gin.Context) {
	// do stuff
	extID := c.Param("id")
	extTEMP := c.Param("temp")
	fmt.Println(extID, extTEMP)
	c.JSON(200, gin.H{
		"ext" + extID: extTEMP,
	})

}

// SetBedTemp - returns extruder temp
func SetBedTemp(c *gin.Context) {
	// do stuff
	bedID := c.Param("id")
	bedTEMP := c.Param("temp")
	fmt.Println(bedID, bedTEMP)
	c.JSON(200, gin.H{
		"bed" + bedID: bedTEMP,
	})

}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	port, err := serial.NewSerialStreamer("/tmp/printer", 250000)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	read, err := port.Reader.ReadString("// Klipper state: Ready\n")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read %v\n", read)
	apiCmd := r.Group("/api/")
	{
		apiCmd.GET("/cmd/ext:id", GetExtruderTemp(port))
		apiCmd.GET("/cmd/bed:id", GetBedTemp(port))
		apiCmd.POST("/cmd/ext:id/:temp", SetExtruderTemp)
		apiCmd.POST("/cmd/bed:id/:temp", SetBedTemp)
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
