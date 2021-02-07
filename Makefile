.PHONY: test mockgen

test:
	go vet ./...
	go test -race -cover ./...

mockgen:
	mockgen -destination command/mock_google/mock_google.go github.com/johnmanjiro13/dokkoi/command CustomSearchRepository
	mockgen -destination command/mock_score/mock_score.go github.com/johnmanjiro13/dokkoi/command ScoreRepository
