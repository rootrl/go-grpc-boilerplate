.PHONY: proto
proto:
	@protoc -I . \
		--go_out=. --go_opt=paths=source_relative \
	       	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(shell find protos -name "*.proto")
