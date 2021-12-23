.PHONY: cover coverage

cover:
	go test -v -coverprofile=coverage.out github.com/deanorhan/biscuit/...
	go tool cover -func=coverage.out
	rm coverage.out

cover-html:
	go test -coverprofile=coverage.out github.com/deanorhan/biscuit/...
	go tool cover -html=coverage.out
	rm coverage.out