
# This how we want to name the binary output
MAIN="cmd/command/main.go"
BINARY="cmd/bin/tubadic"
VER_FILE="version.txt"

# These are the values we want to pass for VERSION and BUILD
BUILD_TS := $(shell date +%FT%T%z)

# Parse file
PREV_VERSION := `[ -f ${VER_FILE} ] && cat ${VER_FILE} || echo 0.0.0`
MAJOR := $(shell echo $(PREV_VERSION) | cut -f1 -d.)
MINOR := $(shell echo $(PREV_VERSION) | cut -f2 -d.)
BUILD := $(shell echo $(PREV_VERSION) | cut -f3 -d.)

ADD_MAJOR := $(shell echo ${MAJOR}+1 | bc).0.0
ADD_MINOR := $(MAJOR).$(shell echo ${MINOR}+1 | bc).0
ADD_BUILD := $(MAJOR).$(MINOR).$(shell echo ${BUILD}+1 | bc)

.PHONY: clean major minor build

update_ver_file=$(shell echo $(1) > $(VER_FILE))

# Builds the project
major:
	$(MAKE) make_linux LDFLAGS="-X main.Version=$(ADD_MAJOR) -X main.Build=$(BUILD_TS)"
	$(call update_ver_file,$(ADD_MAJOR))
minor:
	$(MAKE) make_linux LDFLAGS="-X main.Version=$(ADD_MINOR) -X main.Build=$(BUILD_TS)"
	$(call update_ver_file,$(ADD_MINOR))

build:
	$(MAKE) make_linux LDFLAGS="-X main.Version=$(ADD_BUILD) -X main.Build=$(BUILD_TS)"
	$(call update_ver_file,$(ADD_BUILD))

make_linux: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BINARY) $(MAIN)

make_native: clean
	go build ${LDFLAGS} -o ${BINARY}"_native" ${MAIN}

# Remove temporary files
clean:
	go clean
