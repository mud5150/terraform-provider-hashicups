TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=edu
NAME=hashicups
BINARY=terraform-provider-${NAME}
OS_ARCH=linux_amd64
VERSION=0.2


default: install

build:
	go build -o ${BINARY}

winbuild:
	go build -o ${BINARY}.exe

remove:
	rm -f ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build remove clean_example
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

wininstall: OS_ARCH=windows_amd64
wininstall: INSTALL_PATH=%APPDATA%\terraform.d\plugins\${HOSTNAME}\${NAMESPACE}\${NAME}\${VERSION}\${OS_ARCH}
wininstall: winbuild
	if not exist ${INSTALL_PATH} md ${INSTALL_PATH}
	move ${BINARY}.exe ${INSTALL_PATH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

clean_example:
	rm -rf ./examples/.terraform
	rm -f ./examples/.terraform.lock.hcl