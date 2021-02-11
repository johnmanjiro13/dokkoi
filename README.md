# dokkoi
![](https://github.com/johnmanjiro13/dokkoi/workflows/test%20and%20build/badge.svg?branch=master)

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
$ TOKEN=<YOUR_TOKEN> API_KEY=<YOUR_API_KEY> ENGINE_ID=<YOUR_ENGINE_ID> dokkoi
```

With flags:
```
$ dokkoi --token <YOUR_TOKEN> --api_key <YOUR_API_KEY> --engine_id <YOUR_ENGINE_ID>
```

You can see a help for dokkoi's commands with `dokkoi help` command.
