PROJECT_PATH=/go/src/github.com/rikonor/dns-auto-update

all: build package clean
arm: build_arm package clean

build:
	docker run \
		-- rm \
		-v ${PWD}:${PROJECT_PATH} \
                -w ${PROJECT_PATH} \
                golang:alpine go build

build_arm:
	docker run \
		--rm \
		-v ${PWD}:${PROJECT_PATH} \
		-w ${PROJECT_PATH} \
		-e GOOS=linux -e GOARCH=arm -e GOARM=6 \
		golang:alpine go build

package:
	docker build -t dns-auto-update .

clean:
	rm dns-auto-update

