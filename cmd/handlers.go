package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/grofers/flagger-slack-handler/pkg/slack"
)

func decodeSlackRequest(r *http.Request, s *slack.Request) error {
	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("internal error. failed to parse form: %s", err.Error())
	}
	decoder := schema.NewDecoder()

	err := decoder.Decode(s, r.PostForm)
	if err != nil {
		return fmt.Errorf("internal error. failed to decode form: %s", err.Error())
	}
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	slackReq := &slack.Request{}
	err := decodeSlackRequest(r, slackReq)
	if err != nil {
		log.Println(err.Error())
		responseMessage := fmt.Sprintf("<@%s> Looks like there was an issue with your action: %s",
			slackReq.UserID, err.Error())
		_ = slack.SendSlackRespnose(slackReq.ResponseURL, responseMessage)
		return
	}

	// mark request as recieved
	w.WriteHeader(http.StatusOK)

	// TODO: Add checks for username and channel

	go func() {
		err = slack.PerformAction(slackReq, loadTesterNs)
		if err != nil {
			log.Println(err.Error())
			responseMessage := fmt.Sprintf("<@%s> Looks like there was an issue with your action: %s",
				slackReq.UserID, err.Error())
			_ = slack.SendSlackRespnose(slackReq.ResponseURL, responseMessage)
			return
		}
	}()
}
