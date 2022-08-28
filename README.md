# dokkoi
![](https://github.com/johnmanjiro13/dokkoi/workflows/test%20and%20build/badge.svg?branch=main)

dokkoi is a friendly discord bot.

## Installation
```
$ go install github.com/johnmanjiro13/dokkoi
```

## Usage
You must set some environment variables.
* DISCORD_TOKEN : a token of discord bot
* CUSTOMSEARCH_API_KEY : an api key of google custom search api
* CUSTOMSEARCH_ENGINE_ID : an engine id of google custom search api

```
$ dokkoi
```

You can see a help for dokkoi's commands with sending `dokkoi help` message on discord.

## Development
You can run dokkoi on the docker.
```
$ cp docker-compose.yml.sample docker-compose.yml
# after setting environment variables
$ docker-compose up
```
