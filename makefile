test:
	go test -v -cover --short ./...

test-profiler:
	go test -coverprofile=c.out ./...

integration-test:
	go test -v -cover ./...

compose-up:
	docker compose up -d

compose-down:
	docker compose down