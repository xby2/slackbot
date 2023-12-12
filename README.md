# X by 2 Slackbot


## Design

![image](./design.jpg)

## Docker build

Create and push docker tag `xby2/slackbot`

1. Docker build `docker build -t xby2/slackbot .`

## Docker run

1. Run container on specified port `docker run -d --name assistbot-0 -e SLACKBOT_TOKEN='xoxb-414980923588-416254894241-KcaAQfecMcW4S6cUJ6FyY3uV' -e SERVICE_ADDRESS='http://localhost:6005/eksy' -e BOT_TYPE='ASSISTANT' -e SERVICE_PORT='5001' -p 5001:5001 xby2/slackbot:firstbot`

## Running locally

Running code will need the following:
* Slack bot token with proper permissions
* API endpoint access (will still run without, just returns error)

1. Update go modules `go mod tidy`
2. Run `go run slackbot.go`

