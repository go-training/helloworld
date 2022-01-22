NAME            = helloworld
REPO            = qaas/$(NAME)
REGISTRY        = m3274968.gra5.container-registry.ovh.net
TAG             = latest

all: 

build-binary:
	CGO_ENABLED=0 go build -o $(NAME)

test:
	go test

run:
	go run main.go

build: build-binary
	docker build -t $(REGISTRY)/$(REPO):$(TAG) .

run-image:
	docker run --name $(NAME) --rm $(REGISTRY)/$(REPO):$(TAG)

push:
	docker push $(REGISTRY)/$(REPO):$(TAG)
