package slack

type Action string

const (
	promote  Action = "promote"
	rollback Action = "rollback"
	help     Action = "help"
)

// Command defines the structure of a command recieved from Slack
type command struct {
	CanaryName      string
	CanaryNamespace string
	Action          Action
}

// Request parses the form body from a slack request to a struct
type Request struct {
	Token               string `schema:"token"` // deprecated - DONT USE
	TeamID              string `schema:"team_id"`
	TeamDomain          string `schema:"team_domain"`
	IsEnterpriseInstall string `schema:"is_enterprise_install"`
	EnterpriseID        string `schema:"enterprise_id"`
	EnterpriseName      string `schema:"enterprise_name"`
	ChannelID           string `schema:"channel_id"`
	ChannelName         string `schema:"channel_name"`
	UserID              string `schema:"user_id"`
	UserName            string `schema:"user_name"`
	Command             string `schema:"command"`
	Text                string `schema:"text"`
	ResponseURL         string `schema:"response_url"`
	TriggerID           string `schema:"trigger_id"`
	APIApp              string `schema:"api_app_id"`
}

// Response is the response that will be sent to the Slack URL
type Response struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
}
