.PHONY: test mockgen

test:
	go vet ./...
	go test -race -cover ./...

mockgen:
	mockgen -destination command/mock_command/mock_command.go github.com/johnmanjiro13/dokkoi/command CustomSearchRepository,ScoreRepository
