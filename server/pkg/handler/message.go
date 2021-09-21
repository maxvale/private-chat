package handler

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	pb "server/pkg/proto"
)

type MessageHandler interface {
	HandleMessage(w http.ResponseWriter, r *http.Request)
}

type messageHandlerImpl struct {
	lastMessage string
}

func NewMessageHandler() *messageHandlerImpl {
	return &messageHandlerImpl{lastMessage: ""}
}

func (h *messageHandlerImpl) HandleMessage(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrappedErr := fmt.Errorf("error reading request body %w", err)
		log.Print(wrappedErr)
		http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		return

	}
	m := &pb.Message{}
	if err := proto.Unmarshal(bodyBytes, m); err != nil {
		wrappedErr := fmt.Errorf("error parsing message %w", err)
		log.Print(wrappedErr)
		http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		return
	}
	h.lastMessage = m.String()
	log.Printf("message delivered %s", h.lastMessage)
	w.WriteHeader(http.StatusCreated)
}
