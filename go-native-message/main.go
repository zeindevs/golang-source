package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	logFile, err := os.OpenFile("message.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	for {
		var length uint32
		err := binary.Read(os.Stdin, binary.LittleEndian, &length)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading length: %v\n", err)
			return
		}

		data := make([]byte, length)
		_, err = io.ReadFull(os.Stdin, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading data: %v\n", err)
			return
		}

		var msg Message
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshaling data: %v\n", err)
			return
		}

		logEntry := fmt.Sprintf("%s - Received: %s\n", time.Now().Format(time.RFC3339), msg.Text)
		if _, err := logFile.WriteString(logEntry); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to log file: %v\n", err)
		}

		response := Message{Text: "Received: " + msg.Text}
		respond(response)
	}
}

func respond(msg Message) {
	response, err := json.Marshal(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling response: %v\n", err)
		return
	}

	binary.Write(os.Stdout, binary.LittleEndian, uint32(len(response)))
	os.Stdout.Write(response)
}
