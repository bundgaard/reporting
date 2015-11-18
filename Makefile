VERSION=1.0.0
BUILDTIME=`date +%FT%T%z`

SOURCEDIR=.
GO=go
APPNAME=test_args
SOURCES := $(shell find $(SOURCEDIR) -name "*.go")
VERSIONFILE=$(SOURCEDIR)/version.go
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILDTIME}"

.DEFAULT_GOAL: $(APPNAME)

$(APPNAME): create_version $(SOURCES)
	go build ${LDFLAGS} -o $(APPNAME) $(SOURCES)

.PHONY: clean
clean:
	test -e $(APPNAME) && rm -v $(APPNAME)


create_version:
	rm -rv $(VERSIONFILE)
	@echo "package main" > $(VERSIONFILE)
	@echo "const (" >> $(VERSIONFILE)
	@echo "		Version = \"${VERSION}\"" >> $(VERSIONFILE)
	@echo "		BuildTime = \"${BUILDTIME}\"" >> $(VERSIONFILE)
	@echo ")" >> $(VERSIONFILE)

