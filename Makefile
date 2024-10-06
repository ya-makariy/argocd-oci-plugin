BINARY=argocd-oci-plugin

default: build

quality:
	go vet github.com/ya-makariy/argocd-oci-plugin
	go test -race -v -coverprofile cover.out ./...

build:
	go build -buildvcs=false -o ${BINARY} .

install: build

e2e: install
	./argocd-oci-plugin