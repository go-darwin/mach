TOOLS_DIR := ${CURDIR}/tools
TOOLS_BIN := ${TOOLS_DIR}/bin
TOOLS := $(shell cd ${TOOLS_DIR} && go list -v -x -f '{{ join .Imports " " }}' -tags=tools)

GO_LINT_PACKGAGES ?= $(shell go list ./...)

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
fmt: tools/asmfmt  ## Run goimports and asmfmt.
	$(call target)
	find . -type f -name '*.s' -not -path './vendor/*' | xargs -P ${JOBS} ${TOOLS_BIN}/asmfmt -w

.PHONY: lint/asmvet
lint/asmvet: tools/asmvet  ## Run asmvet.
	$(call target)
	go vet -vettool=${TOOLS_BIN}/asmvet ${GO_LINT_PACKGAGES}

##@ tools

.PHONY: tools
tools: tools/bin/''  ## Install tools

tools/%:  ## install an individual dependent tool
	@${MAKE} tools/bin/$* 1>/dev/null

tools/bin/%: ${TOOLS_DIR}/go.mod ${TOOLS_DIR}/go.sum
	@cd tools; \
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
	@$(RM) -rf ./bin *.out *.test *.prof trace.txt **/_obj ${TOOLS_BIN}


# ----------------------------------------------------------------------------
##@ help

.PHONY: help
help:  ## Show make target help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[33m<target>\033[0m\n"} /^[a-zA-Z_0-9\/_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: env/% env
env:  ## Print the value of MAKEFILE_VARIABLE. Use `make env/MAKEFILE_VARIABLE`.
env/%:
	@echo $($*)
