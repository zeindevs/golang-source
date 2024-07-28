package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Message represents the structure of the message received from the browser extension.
type Message struct {
    Text string `json:"text"`
}

// Response represents the structure of the response sent back to the browser extension.
type Response struct {
    Reply string `json:"reply"`
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)

    for {
        // Read the message length (4 bytes, little endian)
        var length uint32
        err := binary.Read(reader, binary.LittleEndian, &length)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading message length: %v\n", err)
            return
        }

        // Read the message body
        messageBytes := make([]byte, length)
        _, err = reader.Read(messageBytes)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading message body: %v\n", err)
            return
        }

        // Parse the message
        var message Message
        err = json.Unmarshal(messageBytes, &message)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error unmarshaling message: %v\n", err)
            return
        }

        // Process the message and create a response
        response := Response{Reply: strings.ToUpper(message.Text)}
        responseBytes, err := json.Marshal(response)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error marshaling response: %v\n", err)
            return
        }

        // Write the response length (4 bytes, little endian)
        err = binary.Write(writer, binary.LittleEndian, uint32(len(responseBytes)))
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error writing response length: %v\n", err)
            return
        }

        // Write the response body
        _, err = writer.Write(responseBytes)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error writing response body: %v\n", err)
            return
        }

        writer.Flush()
    }
}

