.PHONY: bi-directional-streaming

bi-directional-streaming:
        
	protoc -I protos/ protos/numbers.proto --go_out=plugins=grpc:protos/
