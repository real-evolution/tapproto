# required env variables
## LANGUAGE: language to generate code for
ifndef LANGUAGE
	$(error LANGUAGE is not set)
endif
## OUTPUT: output directory
ifndef OUTPUT
	$(error OUTPUT is not set)
endif
## COMPONENT: component to generate code for
ifndef COMPONENT
	$(error COMPONENT is not set)
endif

# protoc binary
PROTOC ?= protoc

# grpc plugin
GRPCPLUGIN ?= $(shell which grpc_$(LANGUAGE)_plugin)

# proto include directory.
PROTOINCLUDE ?= /usr/local/include:third_party/google_apis

# compilation flags
FLAGS+= --proto_path=rbt/$(COMPONENT):$(PROTOINCLUDE)
FLAGS+= --$(LANGUAGE)_out=$(OUTPUT)

# code generation options
ifeq ($(LANGUAGE),go)
	GEN_OPTS += paths=source_relative

	FLAGS+= --$(LANGUAGE)-grpc_out=$(OUTPUT)
	FLAGS+= --$(LANGUAGE)_opt=$(GEN_OPTS) 
  FLAGS+= --$(LANGUAGE)-grpc_opt=$(GEN_OPTS) 
else
	FLAGS+=	--plugin=$(GRPCPLUGIN)
endif

# protobuf files to compile
PROTOS:= $(shell find rbt/$(COMPONENT) -type f -name '*.proto')

# abort if no proto files found
ifeq ($(PROTOS),)
	$(error no proto files found in rbt/$(COMPONENT))
endif

# generate code from selected component protobuf files
gen: $(PROTOS)
	@echo "[+] compiling proto files in 'rbt/$(COMPONENT)' to '$(OUTPUT)'";
	@mkdir -p $(OUTPUT)
	@$(PROTOC) $(FLAGS) $(PROTOS)

# show help message
help:
	@echo "Usage: $(MAKE) gen LANGUAGE=<lang> OUTPUT=<output_dir> COMPONENT=<component>"

.PHONY: help
