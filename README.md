# X by 2 Slackbot


## Design

![image](./design.jpg)

## Docker

Create and push docker tag `xby2/slackbot`

1. Docker build `docker build -t xby2/slackbot .`

## Running locally

Running code will need the following:
* Slack bot token with proper permissions
* API endpoint access (will still run without, just returns error)

1. Update go modules `go mod tidy`
2. Run `go run slackbot.go`

