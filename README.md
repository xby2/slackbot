# X by 2 Slackbot


## Design

![image](./design.jpg)

## Docker compose

There are 5 `docker-compose` build setup at `./docker-compose`
1. assistbot - helpful assistant
2. actionitemsbot - summarizes actions items
3. keypointsbot - summarizes key points
4. sentimentbot - provides overall sentiment
5. summarybot - provides abstract summary

To run, a `SLACK TOKEN` must be provided.

Run using the following command:

1. Navigate to compose directory `cd ./docker-compose/assistbot`
2. Run `docker-compose up`

## Docker build

Create and push docker tag `xby2/slackbot`

1. Docker build `docker build -t xby2/slackbot .`

## Docker run

1. Run container on specified port `docker run -d --name assistbot-0 -e SLACKBOT_TOKEN='' -e SERVICE_ADDRESS='http://localhost:6005/eksy' -e BOT_TYPE='ASSISTANT' -e SERVICE_PORT='5001' -p 5001:5001 xby2/slackbot:firstbot`

## Running locally

Running code will need the following:
* Slack bot token with proper permissions
* API endpoint access (will still run without, just returns error)

1. Update go modules `go mod tidy`
2. Run `go run slackbot.go`

