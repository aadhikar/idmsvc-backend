##
# This file adds rules to generate code from the
# event message specification
#
# This is based on the repository below:
# https://github.com/RedHatInsights/playbook-dispatcher
#
# https://github.com/atombender/go-jsonschema
##

GOJSONSCHEMA := $(BIN)/gojsonschema
# see: https://github.com/atombender/go-jsonschema/issues/32
# v0.9.0 breaks on 'additionalProperties'
GOJSONSCHEMA_VERSION := v0.8.0

EVENTS := introspect_request

EVENT_SCHEMA_DIR := $(PROJECT_DIR)/internal/api/event
EVENT_MESSAGE_DIR := $(PROJECT_DIR)/api/event

SCHEMA_YAML_FILES := $(wildcard $(EVENT_MESSAGE_DIR)/*.event.yaml)
SCHEMA_JSON_FILES := $(patsubst $(EVENT_MESSAGE_DIR)/%.event.yaml,$(EVENT_SCHEMA_DIR)/%.event.json,$(wildcard $(EVENT_MESSAGE_DIR)/*.event.yaml))


.PHONY: install-gojsonschema
install-gojsonschema: $(GOJSONSCHEMA)

# This rule install gojsonschema into your system
# go install github.com/atombender/go-jsonschema/cmd/gojsonschema
$(GOJSONSCHEMA): $(BIN)
	GOBIN="$(dir $(CURDIR)/$@)" go install "github.com/atombender/go-jsonschema/cmd/gojsonschema@$(GOJSONSCHEMA_VERSION)"
