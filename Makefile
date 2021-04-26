build:
	CGO_ENABLED=0 go build -a -o ./bin/flagger-slack-handler ./cmd