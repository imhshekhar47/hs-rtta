package common

import (
	"encoding/json"
	"log"

	"google.golang.org/protobuf/proto"
)

// SerializeToJSON : serialize to json
func SerializeToJSON(message proto.Message) ([]byte, error) {
	bytes, err := json.Marshal(message)

	if err != nil {
		log.Println("Could not serialize")
	}

	return bytes, nil
}

// DeSerializeFromJSON : deserialize from json
func DeSerializeFromJSON(bytes []byte, message *proto.Message) error {
	return json.Unmarshal(bytes, message)
}
