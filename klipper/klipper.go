// Package kommander klipper - handle klipper input and output
package kommander

import (
	"fmt"
	serial "kommander/serialcom"
	"strings"

	"github.com/gin-gonic/gin"
)

// KlipperInit - hanle klipper init
func KlipperInit(p serial.SerialStreamer) error {
	// handle klipper start
	out, err := p.Reader.ReadBytes()
	return nil
}

// HandleCmd - send cmd and check output
func HandleCmd(cmd string, p serial.SerialStreamer) error {
	// send cmd and check status
	_, err := p.Writer.WriteString(cmd + "\n")
	if err != nil {
		fmt.Println("Error ", err)
	}
	out, err := p.Reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error ", err)
	}
	// check if
	if out[2:] == "ok" {
		fmt.Println("Sent ok")
	}
	return err
}

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
