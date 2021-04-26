package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mayankshah1607/kubectl-flagger/pkg/flagger"
)

const usage = "```Usage:\n/flagger [action] [canary name] [canary namespace]\n" +
	`Available actions:
promote
rollback

Example:
/flagger promote payments consumer
` + "```"

func parseCommand(cmd string) (*command, error) {
	s := strings.Split(cmd, " ")
	log.Println(s)
	if len(s) != 3 {
		return &command{}, fmt.Errorf("invalid command format.\n%s", usage)
	}

	if s[0] != string(promote) && s[0] != string(rollback) {
		return &command{}, fmt.Errorf("invalid action \"%s\"\n%s", s[0], usage)
	}

	return &command{
		Action:          Action(s[0]),
		CanaryName:      s[1],
		CanaryNamespace: s[2],
	}, nil
}

func SendSlackRespnose(slackURL, message string) error {
	slackMsg := &Response{
		ResponseType: "in_channel",
		Text:         message,
	}

	b, err := json.Marshal(slackMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %s", err.Error())
	}

	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message on slack: %s", err.Error())
	}
	defer resp.Body.Close()
	return nil
}

func PerformAction(s *Request, loadTesterNs string) error {
	cmd, err := parseCommand(s.Text)
	if err != nil {
		return fmt.Errorf("failed to parse command: %s", err.Error())
	}

	switch cmd.Action {
	case help:
		err = SendSlackRespnose(s.ResponseURL, usage)
		if err != nil {
			log.Printf("failed to send slack response: %s\n", err.Error())
		}
		return nil
	case promote:
		err := flagger.Promote(cmd.CanaryName, cmd.CanaryNamespace, loadTesterNs, 5)
		if err != nil {
			return fmt.Errorf("failed to promote: %s", err.Error())
		}
		responseMessage := fmt.Sprintf("<@%s> *successfully promoted* `%s/canary/%s`",
			s.UserID, cmd.CanaryName, cmd.CanaryNamespace)
		err = SendSlackRespnose(s.ResponseURL, responseMessage)
		if err != nil {
			log.Printf("failed to send slack response: %s\n", err.Error())
		}
		return nil
	case rollback:
		err := flagger.Rollback(cmd.CanaryName, cmd.CanaryNamespace, loadTesterNs, 5)
		if err != nil {
			return fmt.Errorf("failed to rollback: %s", err.Error())
		}
		responseMessage := fmt.Sprintf("<@%s> *successfully rolledback* `%s/canary/%s`",
			s.UserID, cmd.CanaryName, cmd.CanaryNamespace)
		err = SendSlackRespnose(s.ResponseURL, responseMessage)
		if err != nil {
			log.Printf("failed to send slack response: %s\n", err.Error())
		}
		return nil
	default:
		return fmt.Errorf("invalid action")
	}
}
