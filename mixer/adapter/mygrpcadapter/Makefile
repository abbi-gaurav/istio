RELEASE?=0.0.1
APP?=mygrpcadapter
GOOS?=darwin
GOARCH?=amd64
DOCKER_ACCOUNT?=gabbi
CONTAINER_IMAGE?=${DOCKER_ACCOUNT}/${APP}

clean:
	rm -rf ${APP}

vet:
	go vet .

test:
	go test -v -race ./...

build: clean vet test
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo -v -o ${APP} cmd/main.go

build-image: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push-image: build-image
	docker push $(CONTAINER_IMAGE):$(RELEASE)