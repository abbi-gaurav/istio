RELEASE?=0.0.7
APP?=mygrpcadapter
GOOS?=darwin
GOARCH?=amd64
DOCKER_ACCOUNT?=gabbi
CONTAINER_IMAGE?=${DOCKER_ACCOUNT}/${APP}

clean:
	rm -rf ${APP}

vet:
	GO111MODULE=on go vet .

test:
	GO111MODULE=on go test -v -race ./...

build: clean vet test
	GO111MODULE=on CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo -v -o ${APP} cmd/main.go

build-image: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push-image: build-image
	docker push $(CONTAINER_IMAGE):$(RELEASE)