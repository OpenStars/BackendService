package mock

import (
	"testing"

	generic "github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"go.uber.org/mock/gomock"
)

func TestBigset(t *testing.T) {
	ctr := gomock.NewController(t)
	mockBigset := NewMockClient(ctr)
	mockBigset.EXPECT().BsPutItem("TEST", "TESTKEY", "TESTVAL").Return(true, nil)
	mockBigset.EXPECT().BsGetItem("TEST", "TESTKEY").Return(&generic.TItem{
		Key:   []byte("TESTKEY"),
		Value: []byte("TESTVAL"),
	}, nil)

	ok, err := mockBigset.BsPutItem("TEST", "TESTKEY", "TESTVAL")
	if err != nil {
		t.Error("put item err", err)
	}
	if !ok {
		t.Error("put item not ok")
	}
	item, err := mockBigset.BsGetItem("TEST", "TESTKEY")
	if err != nil {
		t.Error("Get item err", err)
	}
	if item == nil {
		t.Error("not found item")
	}
	if string(item.Key) != "TESTKEY" && string(item.Value) != "TESTVAL" {
		t.Error("value item get wrong")
	}
}
