package kommander

import (
	"bufio"
	"fmt"
	"io"

	serial "github.com/jacobsa/go-serial/serial"
)

// SerialStreamer - struct to stream to serial
type SerialStreamer struct {
	serialPort io.ReadWriteCloser
	Reader     *bufio.Reader
	Writer     *bufio.Writer
}

// NewSerialStreamer - creates a streamer
func NewSerialStreamer(port string, baud uint) (SerialStreamer, error) {
	c := serial.OpenOptions{PortName: port, BaudRate: baud, MinimumReadSize: 4, StopBits: 1, DataBits: 8}
	s := SerialStreamer{}
	var err error
	s.serialPort, err = serial.Open(c)
	if err != nil {
		fmt.Printf("Could not open serial port with error: %v", err)
	}
	s.Reader = bufio.NewReader(s.serialPort)
	s.Writer = bufio.NewWriter(s.serialPort)
	return s, err
}

// func serialReader(reader *bufio.Reader) result {
// 	c, err := reader.ReadBytes('\n')
// 	if err != nil {
// 		return result{"serial-error", fmt.Sprintf("%s", err)}
// 	}
// 	b := string(c)
// 	if b == "ok\r\n" {
// 		return result{"ok", ""}
// 	} else if len(b) >= 5 && b[:5] == "error" {
// 		return result{"error", b[6 : len(b)-1]}
// 	} else if len(b) >= 5 && b[:5] == "alarm" {
// 		return result{"alarm", b[6 : len(b)-1]}
// 	} else {
// 		return result{"info", b[:len(b)-1]}
// 	}
// }

// func (s *SerialStreamer) handleRes(str string) {
// 	// Look for a response
// 	res := serialReader(s.reader)

// 	switch res.level {
// 	case "error":
// 		panic(fmt.Sprintf("Received error from CNC: %s, block: %s", res.message, str))
// 	case "alarm":
// 		panic(fmt.Sprintf("Received alarm from CNC: %s, block: %s", res.message, str))
// 	case "info":
// 		fmt.Printf("\nReceived info from CNC: %s\n", res.message)
// 	default:
// 	}
// }

// Init - smth
// func (s *SerialStreamer) Init() {
// 	s.Write = func(str string) {
// 		str += "\n"

// 		_, err := s.writer.WriteString(str)
// 		if err != nil {
// 			panic(fmt.Sprintf("Error while sending data: %s", err))
// 		}
// 		err = s.writer.Flush()
// 		if err != nil {
// 			panic(fmt.Sprintf("Error while flushing writer: %s", err))
// 		}
// 		s.handleRes(str)
// 	}
// 	//s.GrblGenerator.Init()
// }

// Connect to a serial port at the given path and baudrate
// func (s *SerialStreamer) Connect(name string, baud int) error {
// 	c := &serial.Config{Name: name, Baud: baud}
// 	var err error
// 	s.serialPort, err = serial.OpenPort(c)
// 	if err != nil {
// 		return err
// 	}

// 	s.reader = bufio.NewReader(s.serialPort)
// 	s.writer = bufio.NewWriter(s.serialPort)

// 	for {
// 		c, err := s.reader.ReadBytes('\n')
// 		m := string(c)
// 		if len(m) == 26 && m[:5] == "Grbl " && m[9:] == " ['$' for help]\r\n" {
// 			fmt.Printf("Grbl version %s initialized\n", m[5:9])
// 			break
// 		} else if m == "\r\n" {
// 			continue
// 		}

// 		if err != nil {
// 			return errors.New("Unable to detect initialized GRBL")
// 		}
// 	}

// 	return nil
// }
