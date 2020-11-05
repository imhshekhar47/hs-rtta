package model

import (
	"fmt"
	"math/rand"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestBinarySerDe(t *testing.T) {
	buyCall := &TradeCall{
		Action: CallType_BUY,
		Stock:  "AMZN",
		Units:  int32(rand.Intn(10)),
		Price:  float32(rand.Intn(80)) + rand.Float32(),
	}

	bytes, err := proto.Marshal(buyCall)

	if err != nil {
		t.Fail()
	}

	emptyCall := &TradeCall{}

	err = proto.Unmarshal(bytes, emptyCall)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("Out: %s", emptyCall.Action)

}
