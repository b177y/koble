default: build

GO ?= go
BUILDTAGS ?= exclude_graphdriver_btrfs btrfs_noversion exclude_graphdriver_devicemapper containers_image_openpgp
BUILDFLAGS := -mod=vendor $(BUILDFLAGS)

.PHONY: build
build:
	$(GO) build $(BUILDFLAGS) -tags "$(BUILDTAGS)" -o bin/netkit

.PHONY: test
test: test-uml test-podman
	# $(GO) test -p 1 $(BUILDFLAGS) -tags "$(BUILDTAGS)" ./pkg/netkit ./driver/uml ./driver/podman ./util/topsort

.PHONY: test-uml
test-uml:
	UML_ORIG_UID=$(shell id -u) \
	UML_ORIG_EUID=$(shell echo $EUID) \
	UML_ORIG_GID=$(shell id -g) \
		unshare -mUr $(GO) test -v -p 1 $(BUILDFLAGS) -tags "$(BUILDTAGS)" ./driver/uml

.PHONY: test-podman
test-podman:
	$(GO) test -v -p 1 $(BUILDFLAGS) -tags "$(BUILDTAGS)" ./driver/podman

.PHONY: vendor
vendor:
	GO111MODULE=on $(GO) mod tidy
	GO111MODULE=on $(GO) mod vendor
	GO111MODULE=on $(GO) mod verify
	rm vendor/github.com/containers/storage/pkg/unshare/unshare_cgo.go

.PHONY: manpages
manpages:
	GO111MODULE=on $(GO) run cmd/man/gen_manpages.go
	mkdir -p docs/modules/MAN/pages
	rm -f build/man/*.adoc docs/modules/MAN/pages/*.adoc
	find ./build/man -name '*.md' -exec sh -c 'echo "converting {}";kramdoc --heading-offset=-1 {}' {} \;
	cp build/man/*.adoc docs/modules/MAN/pages
	echo ".Manpages" > docs/modules/MAN/nav.adoc
	find ./docs/modules/MAN/pages -name '*.adoc' -execdir echo "* xref:{}[]" >> docs/modules/MAN/navtmp.adoc \;
	sort docs/modules/MAN/navtmp.adoc >> docs/modules/MAN/nav.adoc
	rm docs/modules/MAN/navtmp.adoc
