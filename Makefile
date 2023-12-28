#!make

# generate api protobuf files
proto:
	buf generate proto

.PHONY: proto
