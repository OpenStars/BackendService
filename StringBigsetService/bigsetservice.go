package StringBigsetService

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	transports "github.com/OpenStars/BackendService/StringBigsetService/transportsv2"
)

var reconnect = true
var mureconnect sync.Mutex

type StringBigsetService struct {
	host string
	port string
	sid  string
}

func NewClient(etcdEndpoints []string, sid string, defaultEndpointsHost string, defaultEndpointPort string) Client {
	log.Println("Init StringBigset Service sid", sid, "address", defaultEndpointsHost+":"+defaultEndpointPort)
	client := transports.GetBsGenericClient(defaultEndpointsHost, defaultEndpointPort)
	if client == nil || client.Client == nil {
		return nil
	}
	stringbs := &StringBigsetService{
		host: defaultEndpointsHost,
		port: defaultEndpointPort,
		sid:  sid,
	}
	return stringbs
}

func (m *StringBigsetService) TotalStringKeyCount() (r int64, err error) {

	client := transports.GetBsGenericClient(m.host, m.port)

	if client == nil || client.Client == nil {
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err = client.Client.(*generic.TStringBigSetKVServiceClient).TotalStringKeyCount(ctx)
	if err != nil {
		transports.ServiceDisconnect(client)
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)

	return r, nil
}

func (m *StringBigsetService) GetListKey(fromIndex int64, count int32) ([]string, error) {

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	timeOut := count/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetListKey(ctx, fromIndex, count)

	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)

	var listKey []string
	for _, item := range r {
		listKey = append(listKey, string(item))
	}
	return listKey, nil
}

func (m *StringBigsetService) BsMultiPutBsItem(lsItem []*generic.TBigsetItem) (failedItem []*generic.TBigsetItem, err error) {
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return lsItem, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	timeOut := len(lsItem)/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiPutBsItem(ctx, lsItem)

	if err != nil {
		transports.ServiceDisconnect(client)
		return lsItem, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if r.Error != generic.TErrorCode_EGood {
		return lsItem, nil
	}
	// if r.Ok == false {
	// 	err := errors.New("Can not write bskey: " + bskey + " itemkey: " + itemKey)
	// 	transports.ServiceDisconnect(client)
	// 	return false, err
	// }
	return r.FailedPutbsItem, nil
}

func (m *StringBigsetService) BsPutItem(bskey string, itemKey, itemVal string) (bool, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	// log.Println("BsPutItem host", m.host+":"+m.port)
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(ctx, generic.TStringKey(bskey), &generic.TItem{
		Key:   []byte(itemKey),
		Value: []byte(itemVal),
	})

	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if r.Error != generic.TErrorCode_EGood {
		return false, nil
	}
	// if r.Ok == false {
	// 	err := errors.New("Can not write bskey: " + bskey + " itemkey: " + itemKey)
	// 	transports.ServiceDisconnect(client)
	// 	return false, err
	// }
	return true, nil

}

func (m *StringBigsetService) BsRangeQuery(bskey string, startKey string, endKey string) ([]*generic.TItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, generic.TStringKey(bskey), generic.TItemKey(startKey), generic.TItemKey(endKey))
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)

	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsRangeQueryAll(bskey string) ([]*generic.TItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, generic.TStringKey(bskey), nil, generic.TItemKey("z"))
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)

	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

// BsRangeQueryByPage get >= startkey && <= endkey có chia page theo begin and end
func (m *StringBigsetService) BsRangeQueryByPage(bskey string, startKey, endKey string, begin, end int64) ([]*generic.TItem, int64, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)

	if client == nil || client.Client == nil {

		return nil, -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, generic.TStringKey(bskey), generic.TItemKey(startKey), generic.TItemKey(endKey))
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, -1, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)

	if r.Items != nil && r.Items.Items != nil && len(r.Items.Items) > 0 { // pagination
		if begin < 0 {
			begin = 0
		}
		if end > int64(len(r.Items.Items)) {
			end = int64(len(r.Items.Items))
		}
		total := int64(len(r.Items.Items))
		r.Items.Items = r.Items.Items[begin:end]

		return r.Items.Items, total, nil
	}

	return nil, 0, nil
}

func (m *StringBigsetService) BsGetItem(bskey string, itemkey string) (*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	// fmt.Printf("[BsGetItem] get client host = %s, %s, key = %s, %s \n", m.host, m.port, bskey, itemkey)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetItem(ctx, generic.TStringKey(bskey), generic.TItemKey(itemkey))
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer transports.BackToPool(client)
	if r.Error != generic.TErrorCode_EGood || r.Item == nil || r.Item.Key == nil {
		return nil, nil
	}
	return r.Item, nil
}

func (m *StringBigsetService) GetTotalCount(bskey string) (int64, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetTotalCount(ctx, generic.TStringKey(bskey))

	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)

	if r <= 0 {
		return 0, nil
	}
	return r, nil

}

func (m *StringBigsetService) GetBigSetInfoByName(bskey string) (*generic.TStringBigSetInfo, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetBigSetInfoByName(ctx, generic.TStringKey(bskey))
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer transports.BackToPool(client)

	if rs.Info == nil {
		return nil, nil
	}
	return rs.Info, nil

}

func (m *StringBigsetService) RemoveAll(bskey string) (bool, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Client.(*generic.TStringBigSetKVServiceClient).RemoveAll(ctx, generic.TStringKey(bskey))
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer transports.BackToPool(client)
	return true, nil
}
func (m *StringBigsetService) CreateStringBigSet(bskey string) (*generic.TStringBigSetInfo, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).CreateStringBigSet(ctx, generic.TStringKey(bskey))
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer transports.BackToPool(client)

	if rs.Info == nil {
		return nil, nil
	}
	return rs.Info, nil
}

func (m *StringBigsetService) BsGetSlice(bskey string, fromPos int32, count int32) ([]*generic.TItem, error) {
	if count == 0 {
		return nil, nil
	}
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	timeOut := count/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSlice(ctx, generic.TStringKey(bskey), fromPos, count)
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceR(bskey string, fromPos int32, count int32) ([]*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	timeOut := count/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceR(ctx, generic.TStringKey(bskey), fromPos, count)
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsRemoveItem(bskey string, itemkey string) (bool, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ok, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRemoveItem(ctx, generic.TStringKey(bskey), generic.TItemKey(itemkey))
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return ok, nil
}

func (m *StringBigsetService) BsMultiPut(bskey string, lsItems []*generic.TItem) (bool, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	itemset := &generic.TItemSet{
		Items: lsItems,
	}
	timeOut := len(lsItems)/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiPut(ctx, generic.TStringKey(bskey), itemset, false, false)
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if rs.Error != generic.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *StringBigsetService) BsMultiRemoveBsItem(listItems []*generic.TBigsetItem) (listFailedRemove []*generic.TBigsetItem, err error) {
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	timeOut := len(listItems)/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiRemoveBsItem(ctx, listItems)
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if rs.Error != generic.TErrorCode_EGood {
		return rs.FailedRemovebsItem, nil
	}

	return rs.FailedRemovebsItem, nil
}

func (m *StringBigsetService) BsGetSliceFromItem(bskey string, fromKey string, count int32) ([]*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	timeOut := count/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItem(ctx, generic.TStringKey(bskey), generic.TItemKey(fromKey), count)
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceFromItemR(bskey string, fromKey string, count int32) ([]*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {

		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	timeOut := count/1000 + 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItemR(ctx, generic.TStringKey(bskey), generic.TItemKey(fromKey), count)
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NeewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}
	return rs.Items.Items, nil
}
