default: build

GO ?= go
BUILDTAGS ?= exclude_graphdriver_btrfs btrfs_noversion exclude_graphdriver_devicemapper containers_image_openpgp
BUILDFLAGS := -mod=vendor $(BUILDFLAGS)

.PHONY: build
build:
	$(GO) build $(BUILDFLAGS) -tags "$(BUILDTAGS)" -o bin/netkit

.PHONY: test
test: test-uml
	# $(GO) test -p 1 $(BUILDFLAGS) -tags "$(BUILDTAGS)" ./pkg/netkit ./driver/uml ./driver/podman ./util/topsort

.PHONY: test-uml
test-uml:
	UML_ORIG_UID=$(shell id -u) \
	UML_ORIG_EUID=$(shell echo $EUID) \
	UML_ORIG_GID=$(shell id -g) \
		unshare -mUr $(GO) test -p 1 $(BUILDFLAGS) -tags "$(BUILDTAGS)" ./driver/uml

.PHONY: vendor
vendor:
	GO111MODULE=on $(GO) mod tidy
	GO111MODULE=on $(GO) mod vendor
	GO111MODULE=on $(GO) mod verify
	rm vendor/github.com/containers/storage/pkg/unshare/unshare_cgo.go
