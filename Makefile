.PHONY: build
build:
	go build .

.PHONY: docker.build
docker.build:
	docker buildx build --platform=linux/amd64,linux/arm64 -t hjr265/mghsiac:latest $(if $(PUSH),--push,) .
