.PHONY: build
build:
	@go build ./...


.PHONY: test
test:
	@go test ./...


.PHONY: coverage
coverage:
	@go test -cover ./...


.PHONY: coverage-report
coverage-report:
	@go test -coverprofile=/tmp/coverage.txt ./...
	@go tool cover -html=/tmp/coverage.txt


.PHONY: lint
lint:
	@gofmt -l .
	@go vet ./...


.PHONY: docs-server
docs-server:
	@echo "Documentation @ http://127.0.0.1:6060"
	@godoc -http=:6060


.PHONY: find_todo
find_todo:
	@grep --color=always --include=\*.go -PnRe '(//|/*).*TODO' ./ || true


.PHONY: find_fixme
find_fixme:
	@grep --color=always --include=\*.go -PnRe '(//|/*).*FIXME' ./ || true


.PHONY: find_xxx
find_xxx:
	@grep --color=always --include=\*.go -PnRe '(//|/*).*XXX' ./ || true


.PHONY: clean
clean:
	@# @rm -f file
	@echo "Removing files"


.PHONY: count
count:
	@echo "Lines of code:"
	@find . -type f -name "*.go" | xargs wc -l
