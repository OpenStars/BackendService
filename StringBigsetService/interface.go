package StringBigsetService

import "github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"

type Client interface {
	TotalStringKeyCount() (r int64, err error)
	GetListKey(fromIndex int64, count int32) ([]string, error)
	BsPutItem(bskey string, itemKey string, itemVal string) (bool, error)
	BsGetItem(bskey string, itemkey string) (*generic.TItem, error)
	GetTotalCount(bskey string) (int64, error)
	BsMultiPut(bskey string, lsItems []*generic.TItem) (bool, error)
	BsGetSlice(bskey string, fromPos int32, count int32) ([]*generic.TItem, error)
	BsRemoveItem(bskey string, itemkey string) (bool, error)
	GetBigSetInfoByName(bskey string) (*generic.TStringBigSetInfo, error)
	CreateStringBigSet(bskey string) (*generic.TStringBigSetInfo, error)
	BsRangeQuery(bskey string, startKey string, endKey string) ([]*generic.TItem, error)
	BsRangeQueryByPage(bskey string, startKey string, endKey string, begin, end int64) ([]*generic.TItem, int64, error)
	BsGetSliceFromItem(bskey string, fromKey string, count int32) ([]*generic.TItem, error)
	BsGetSliceFromItemR(bskey string, fromKey string, count int32) ([]*generic.TItem, error)
	RemoveAll(bskey string) (bool, error)
	BsGetSliceR(bskey string, fromPos int32, count int32) ([]*generic.TItem, error)
	BsMultiPutBsItem(lsItem []*generic.TBigsetItem) (failedItem []*generic.TBigsetItem, err error)
	BsMultiRemoveBsItem(listItems []*generic.TBigsetItem) (listFailedRemove []*generic.TBigsetItem, err error)
	BsRangeQueryAll(bskey string) ([]*generic.TItem, error)
}
