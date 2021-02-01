# dokkoi
![](https://github.com/johnmanjiro13/dokkoi/workflows/build/badge.svg?branch=master)

dokkoi is a friendly discord bot.

## Installation
```
$ go get github.com/johnmanjiro13/dokkoi
```

## usage
You must set a token of discord bot, an api key and engine id of google custom search.
You can set them either environment variables or flags.

With environment variables:
```
$ TOKEN=<YOUR_TOKEN> API_KEY=<YOUR_API_KEY> ENGINE_ID=<YOUR_ENGINE_ID> go run main.go handler.go
```

With flags:
```
$ go run main.go handler.go --token <YOUR_TOKEN> --api_key <YOUR_API_KEY> --engine_id <YOUR_ENGINE_ID>
```

You can execute dokkoi's commands with `dokkoi` prefix in discord like `dokkoi echo hoge`
