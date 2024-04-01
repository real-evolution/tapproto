#!make

# generate api protobuf files
proto:
	@buf generate --verbose \
	              --include-imports \
	              --include-wkt proto \
	              --config proto/buf.yaml \
	              --template proto/buf.gen.yaml

.PHONY: proto
