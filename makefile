JS_FOLDER = ./js_grpc
GO_FOLDER = ./go_grpc

PROTOC = /opt/homebrew/bin/protoc
PROTOC_GEN_GO = /Users/khisby/go/bin/protoc-gen-go
PROTOC_GEN_GRPC = $(JS_FOLDER)/node_modules/.bin/protoc-gen-ts
PROTOC_GEN_GRPC_PLUGIN = /opt/homebrew/bin/grpc_tools_node_protoc_plugin

JS_OUT = ./$(JS_FOLDER)/grpc
GO_OUT = ./$(GO_FOLDER)/grpc

PROTOS_PATH = ./contract
vpath %.proto $(PROTOS_PATH)

$(GO_OUT)/:
	mkdir -p $(GO_OUT)

$(JS_OUT)/:
	mkdir -p $(JS_OUT)

.PHONY: gen-proto
gen-proto: build-go build-js 
	echo "Generate Proto Done..."

.PHONY: build-go
build-go: $(GO_OUT)/
	for file in $(PROTOS_PATH)/*.proto; do \
		filename=$$(basename $$file); \
		dirname=$$(echo $$filename | cut -f 1 -d '.'); \
		target_dir=$(GO_OUT)/$$(echo $$dirname); \
		mkdir -p $$target_dir; \
		$(PROTOC) -I=$(PROTOS_PATH) --go_out=plugins=grpc:$(GO_OUT)/$$dirname $$file; \
	done

.PHONY: build-js
build-js: $(JS_OUT)/
	for file in $(PROTOS_PATH)/*.proto; do \
		filename=$$(basename $$file); \
		dirname=$$(echo $$filename | cut -f 1 -d '.'); \
		target_dir=$(JS_OUT)/$$(echo $$dirname); \
		mkdir -p $$target_dir; \
		cd $(JS_FOLDER) && grpc_tools_node_protoc -I=$(PROTOS_PATH) --js_out=import_style=commonjs,binary:grpc/$$dirname --grpc_out=generate_package_definition:grpc/$$dirname --plugin=protoc-gen-grpc=$(PROTOC_GEN_GRPC_PLUGIN) --proto_path ../$(PROTOS_PATH) $$filename; \
		$(PROTOC) --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --ts_out=grpc_js:grpc/$$dirname --proto_path ../$(PROTOS_PATH) $$filename; \
	done

.PHONY: clean
clean:
	rm -rf $(GO_OUT)
	rm -rf $(JS_OUT)

.PHONY: run-go
run-go: 	
	cd $(GO_FOLDER) && go run main.go

.PHONY: run-js
run-js:
	cd $(JS_FOLDER) && npm start