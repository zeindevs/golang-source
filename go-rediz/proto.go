package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSET    = "set"
	CommandGET    = "get"
	CommandHello  = "hello"
	CommandClient = "client"
)

type Command interface{}

type SetCommand struct {
	key, val []byte
}

type GetCommand struct {
	key []byte
}

type HelloCommand struct {
	value string
}

type ClientCommand struct {
	value string
}

func parseCommand(raw string) (Command, error) {
	rd := resp.NewReader(bytes.NewBufferString(raw))

	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if v.Type() == resp.Array {
			for _, val := range v.Array() {
				switch val.String() {
				case CommandClient:
					cmd := ClientCommand{
						value: v.Array()[1].String(),
					}
					return cmd, nil
				case CommandGET:
					if len(v.Array()) != 2 {
						return nil, fmt.Errorf("invalid number of variables of GET command")
					}
					cmd := GetCommand{key: v.Array()[1].Bytes()}
					return cmd, nil
				case CommandSET:
					if len(v.Array()) != 3 {
						return nil, fmt.Errorf("invalid number of variables of SET command")
					}
					cmd := SetCommand{
						key: v.Array()[1].Bytes(),
						val: v.Array()[2].Bytes(),
					}
					return cmd, nil
				case CommandHello:
					cmd := HelloCommand{
						value: v.Array()[1].String(),
					}
					return cmd, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("invalid or unknown command received: %s", raw)
}

func respWriteMap(m map[string]string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	rw := resp.NewWriter(buf)
	for k, v := range m {
		rw.WriteString(k)
		rw.WriteString(":" + v)
	}
	return buf.Bytes()
}
