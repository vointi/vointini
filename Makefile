APPNAME?=vointini

VERSION := $(shell cat VERSION)
BUILD := $(shell git rev-parse v$(VERSION))
BUILDDATE := $(shell git log -1 --format=%aI v$(VERSION))
LDFLAGS := -trimpath -ldflags "-s -w -X=meta.VERSION=$(VERSION) -X=meta.BUILD=$(BUILD) -X=meta.BUILDDATE=$(BUILDDATE)"
RELEASETMPDIR := $(shell mktemp -d -t ${APPNAME}-rel-XXXXXX)
APPANDVER := ${APPNAME}-$(VERSION)
RELEASETMPAPPDIR := $(RELEASETMPDIR)/$(APPANDVER)

# https://golang.org/doc/install/source#environment
LINUX_ARCHS := amd64 arm arm64

genschema:
	pg_dump --format p --schema-only --if-exists --clean --no-tablespaces --no-privileges --no-owner --port 5432 --host 127.0.0.1 --username vointini vointini --file backend/storage/postgres/schema/schema.sql

dbbackup:
	pg_dump --format p --if-exists --clean --no-tablespaces --no-privileges --no-owner --port 5432 --host 127.0.0.1 --username vointini vointini --file backup.sql

buildjs:
	go run ./cmd/build
	cd frontend/templates && npm run build

build: buildjs
	go build -o ./bin/server ./cmd/server

run:
	./bin/server

linux-build: buildjs
	@for arch in $(LINUX_ARCHS); do \
	  echo "GNU/Linux build... $$arch"; \
	  CGO_ENABLED=0 GOOS=linux GOARCH=$$arch go build $(LDFLAGS) -o ./bin/linux-$$arch/${APPNAME}-server ./cmd/server ; \
	done

shasums:
	@echo "Checksumming..."
	@pushd "release/${VERSION}" && shasum -a 256 ${APPNAME}-${VERSION}-* > $(APPANDVER).sha256sums

# Compress executables
upx-pack:
	@upx $(UPXFLAGS) ./bin/linux-amd64/${APPNAME}-server
	@upx $(UPXFLAGS) ./bin/linux-arm/${APPNAME}-server

# Copy common files to release directory
# Creates $(APPNAME)-$(VERSION) directory prefix where everything will be copied by compress-$OS targets
copycommon:
	echo "Copying common files to temporary release directory '$(RELEASETMPAPPDIR)'.."; \
	mkdir -p "$(RELEASETMPAPPDIR)/bin" && \
	cp "./LICENSE" "$(RELEASETMPAPPDIR)" && \
	cp "./README.md" "$(RELEASETMPAPPDIR)" && \
	cp "./backend/storage/postgres/schema/schema.sql" "$(RELEASETMPAPPDIR)/pg-schema.sql" && \
	cp "./release/config.json" "$(RELEASETMPAPPDIR)/config.dist.json" && \
	mkdir --parents "$(PWD)/release/${VERSION}"

# Compress files: GNU/Linux
compress-linux:
	@for arch in $(LINUX_ARCHS); do \
	  echo "GNU/Linux tar... $$arch"; \
	  cp -v "$(PWD)/bin/linux-$$arch/${APPNAME}-server" "$(RELEASETMPAPPDIR)/bin"; \
	  cd "$(RELEASETMPDIR)"; \
	  tar --numeric-owner --owner=0 --group=0 -zcvf "$(PWD)/release/${VERSION}/$(APPANDVER)-linux-$$arch.tar.gz" . ; \
	  rm "$(RELEASETMPAPPDIR)/bin/${APPNAME}-server"; \
	done

# Move all to temporary directory and compress with common files
compress-everything: copycommon compress-linux
	@echo "$@ ..."
	rm -rf "$(RELEASETMPDIR)/*"

# Linux distributions
release-ldistros: ldistro-arch
	@echo "Linux distros release done..."

release: linux-build upx-pack compress-everything shasums
	@echo "release done..."

# Distro: Arch linux - https://www.archlinux.org/
# Generates multi-arch PKGBUILD
ldistro-arch:
	pushd release/linux/arch && go run . -version ${VERSION}

.PHONY: all clean test default