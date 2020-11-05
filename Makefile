PROTO_SRC=proto
PROTO_DEST=evt-model

generate:
	protoc -I=${PROTO_SRC}/model --go_out=${PROTO_DEST} ${PROTO_SRC}/model/*


produce:
	go run evt-producer/main.go


consume:
	go run evt-consumer/main.go