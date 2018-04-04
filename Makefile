BINDIR := $(CURDIR)/bin
MAINDIR := bosh-alicloud-cpi
MAINFILE := $(CURDIR)/src/$(MAINDIR)/main/main.go
EXECUTABLE := $(BINDIR)/alicloud_cpi

GOPATH := $(CURDIR)

COMMIT = $(shell git rev-parse --short HEAD)

GO_OPTION ?=
ifeq ($(VERBOSE), 1)
GO_OPTIONS += -v
endif

# TODO add local link invocation
BUILD_OPTIONS = -a -ldflags "-X main.GitCommit=\"$(COMMIT)\""

all: clean build

clean:
	rm -f ${EXECUTABLE}
	rm -f bin/bosh-alicloud-cpi.tgz

build:
	mkdir -p $(BINDIR)
	go build $(GO_OPTIONS) $(BUILD_OPTIONS) -o ${EXECUTABLE} $(MAINFILE)

testdeps:
	go install bosh-alicloud-cpi/vendor/github.com/onsi/ginkgo/ginkgo

test: testdeps
	ginkgo -r -skipPackage integration src/bosh-alicloud-cpi

integration: testdeps
	ginkgo -r src/bosh-alicloud-cpi/integration

testintci: testdeps
	ginkgo src/bosh-alicloud-cpi/integration -slowSpecThreshold=500 -progress -nodes=3 -randomizeAllSpecs -randomizeSuites $(GINKGO_ARGS) -v

create-release: build
	bosh create-release --force --tarball=bin/bosh-alicloud-cpi.tgz

# upload-release: create-release
#    scp bin/bosh-alicloud-cpi.tgz root@${CPI_UPLOAD_HOST}:/root
#    rm -f bin/bosh-alicloud-cpi.tgz


