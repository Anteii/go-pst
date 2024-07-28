all:
	cd cmd/properties && go run generate.go

	go install github.com/favadi/protoc-go-inject-tag@latest
	cp ~/go/bin/protoc-go-inject-tag ./

	protoc --proto_path=cmd/properties/protobufs --go_out=paths=source_relative:pkg/properties cmd/properties/protobufs/*.proto

	protoc-go-inject-tag -input="pkg/properties/*.pb.go" -remove_tag_comment

	cd pkg/properties && go generate