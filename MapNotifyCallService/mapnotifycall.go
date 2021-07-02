package MapNotifyCallService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/MapNotifyCallService/mapnoitfycall/thrift/gen-go/OpenStars/Common/MapNotifyCallKV"
	"github.com/OpenStars/BackendService/MapNotifyCallService/mapnoitfycall/transports"
	// "github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type mapnotifycallservice struct {
	host string
	port string
	sid  string
	// epm         GoEndpointBackendManager.EndPointManagerIf
	// etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func (m *mapnotifycallservice) PutData(pubkey string, token string) error {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetTMapNotifyCallKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Can not connect to model")
	}

	_, err := client.Client.(*MapNotifyCallKV.TMapNotifyKVServiceClient).PutData(context.Background(), pubkey, token)
	defer client.BackToPool()
	return err
}

func (m *mapnotifycallservice) GetTokenByPubkey(pubkey string) (string, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetTMapNotifyCallKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to model")
	}

	r, err := client.Client.(*MapNotifyCallKV.TMapNotifyKVServiceClient).GetTokenByPubkey(context.Background(), pubkey)
	if err != nil {
		return "", err
	}
	defer client.BackToPool()
	if r.ErrorCode != MapNotifyCallKV.TErrorCode_EGood {
		return "", errors.New("Get token by pubkey errors " + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}

func (m *mapnotifycallservice) GetPubkeyByToken(token string) (string, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetTMapNotifyCallKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return "", errors.New("Can not connect to model")
	}

	r, err := client.Client.(*MapNotifyCallKV.TMapNotifyKVServiceClient).GetPubkeyByToken(context.Background(), token)
	if err != nil {
		return "", err
	}
	defer client.BackToPool()
	if r.ErrorCode != MapNotifyCallKV.TErrorCode_EGood {
		return "", errors.New("Get phonenubmer by pubkey errors " + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}

// func (m *mapnotifycallservice) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
// 	m.host = ep.Host
// 	m.port = ep.Port
// 	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
// }

func NewMapNotifyCallService(etcdServers []string, sid, host, port string) Client {
	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	// log.Println("Load endpoit ", serviceID, "err", err.Error())
	// 	log.Println("Init Local MapNotifyCall sid:", defaultEnpoint.ServiceID, "host:", defaultEnpoint.Host, "port:", defaultEnpoint.Port)
	// 	return &mapnotifycallservice{
	// 		host: defaultEnpoint.Host,
	// 		port: defaultEnpoint.Port,
	// 		sid:  defaultEnpoint.ServiceID,
	// 	}
	// }
	// sv := &mapnotifycallservice{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handlerEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd MapNoitfyCall sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv
	sv := &mapnotifycallservice{
		host: host,
		port: port,
		sid:  sid,
		// etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}
	// if sv.etcdManager == nil {
	// 	return nil
	// }
	// err := sv.etcdManager.SetDefaultEntpoint(serviceID, defaultEndpoint.Host, defaultEndpoint.Port)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
	// 	return nil
	// }
	// sv.etcdManager.GetAllEndpoint(serviceID)
	return sv
}
