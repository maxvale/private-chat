package main

import (
	"bufio"
	"bytes"
	pb "client/pkg/proto"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var message string
	messageSend := &pb.Message{}
	cl := retryablehttp.NewClient()
	cl.RetryMax = 1
	client := cl.StandardClient()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Your message-> ")
		message, _ = reader.ReadString('\n')
		message = strings.ReplaceAll(message, "\n", "")
		fmt.Println(message)
		if message == "stop" {
			break
		}
		messageSend.Name = "Client"
		messageSend.Body = message
		bodyBytes, err := proto.Marshal(messageSend)
		if err != nil {
			log.Printf("Cannot parse message %e", err)
			continue
		}
		buffer := bytes.NewBuffer(bodyBytes)
		req, err := http.NewRequest("POST", "http://localhost:8080/message/", buffer)
		if err != nil {
			log.Printf("Cannot create request %e", err)
			continue
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error while do request %e", err)
			continue
		}
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Status)
	}
}
