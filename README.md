# dokkoi
![](https://github.com/johnmanjiro13/dokkoi/workflows/test%20and%20build/badge.svg?branch=master)

dokkoi is a friendly discord bot.

## Installation
```
$ go get github.com/johnmanjiro13/dokkoi
```

## Usage
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

## Development
### Run with docker-compose
You must copy `docker-compose.yaml.example` to `docker-compose.yaml` and fill in some environment variables.
```
$ cp docker-compose.yaml.example docker-compose.yaml # fill in environment variables
$ docker-compose up
```
### DB migration
We use [sql-migrate](https://github.com/rubenv/sql-migrate) for the db migration.

`dbconfig.yml` is a config file for local development.

If you want to add other credentials, you can use `local_dbconfig.yml` which is gitignored.
```
$ cp dbconfig.yml local_dbconfig.yml
# You can use local_dbconfig.yml with -config option
$ sql-migrate status -profile='production' -config='local_dbconfig.yml`
```

Creating new migration file
```
$ sql-migrate new <file-name>
```
Migration
```
# apply
$ sql-migrate up
# with specify profile and config file
$ sql-migrate up -profile='development' -config='local_dbconfig.yml'

# rollback
$ sql-migrate down
# with specify profile and config file
$ sql-migrate down -profile='development' -config='local_dbconfig.yml'
```
