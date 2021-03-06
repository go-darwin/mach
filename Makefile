GO_TAGS := osusergo,netgo
GO_LDFLAGS := -s -w "-extldflags=-static"
GO_INSTALLSUFFIX := netgo
GO_FLAGS = -trimpath -tags='${GO_TAGS}' -ldflags='${GO_LDFLAGS}' -installsuffix='${GO_INSTALLSUFFIX}'

GO_TEST ?= ${TOOLS_BIN}/gotestsum --
GO_TEST_FLAGS ?= -race -count=1
GO_TEST_FUNC ?= .
GO_TEST_PACKAGE ?= ./...
GO_LINT_PACKAGE ?= ./...

TOOLS_DIR := ${CURDIR}/tools
TOOLS_BIN := ${TOOLS_DIR}/bin
TOOLS := $(shell cd ${TOOLS_DIR} && go list -v -x -f '{{ join .Imports " " }}' -tags=tools)

JOBS := $(shell getconf _NPROCESSORS_CONF)
ifeq ($(CIRCLECI),true)
ifeq (${GO_OS},linux)
	# https://circleci.com/changelog#container-cgroup-limits-now-visible-inside-the-docker-executor
	JOBS := $(shell echo $$(($$(cat /sys/fs/cgroup/cpu/cpu.shares) / 1024)))
	GO_TEST_FLAGS+=-p=${JOBS} -cpu=${JOBS}
endif
endif


##@ fmt, lint

.PHONY: fmt
fmt: tools/asmfmt  ## Run asmfmt.
	$(call target)
	find . -type f -name '*.s' -not -path './vendor/*' | xargs -P ${JOBS} ${TOOLS_BIN}/asmfmt -w

.PHONY: lint
lint: lint/asmvet  ## Run all linter.

.PHONY: lint/asmvet
lint/asmvet: tools/asmvet  ## Run asmvet.
	$(call target)
	go vet -vettool=${TOOLS_BIN}/asmvet ${GO_LINT_PACKAGE}


##@ test

define go_test
${GO_TEST} $(strip ${GO_FLAGS}) ${GO_TEST_FLAGS} -run=${GO_TEST_FUNC} ${GO_TEST_PACKAGE}
endef

.PHONY: test
test: tools/gotestsum  ## Run test.
	$(call go_test)

.PHONY: coverage
coverage: GO_TEST_FLAGS+=-covermode=atomic -coverpkg=./... -coverprofile=coverage.out
coverage: tools/gotestsum  ## Run test and collect coverages.
	$(call go_test)


##@ tools

.PHONY: tools
tools: tools/''  ## Install tools

tools/%: ${TOOLS_DIR}/go.mod ${TOOLS_DIR}/go.sum
	cd tools; \
	  for t in ${TOOLS}; do \
			if [ -z '$*' ] || [ $$(basename $$t) = '$*' ]; then \
				echo "Install $$t ..." >&2; \
				GOBIN=${TOOLS_BIN} CGO_ENABLED=0 go install -mod=mod ${GO_FLAGS} "$${t}"; \
			fi \
	  done


##@ clean

.PHONY: clean
clean:  ## Cleanups binaries and extra files in the package.
	$(call target)
	@$(RM) -rf *.out *.test *.prof trace.txt **/_obj ${TOOLS_BIN}


# ----------------------------------------------------------------------------
##@ help

.PHONY: help
help:  ## Show make target help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[33m<target>\033[0m\n"} /^[a-zA-Z_0-9\/_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: env/% env
env:  ## Print the value of MAKEFILE_VARIABLE. Use `make env/MAKEFILE_VARIABLE`.
env/%:
	@echo $($*)
