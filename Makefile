#!make

# generate api protobuf files
proto:
	buf generate --include-imports --include-wkt proto

.PHONY: proto
