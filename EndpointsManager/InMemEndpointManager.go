package EndpointsManager

import (
	"errors"
	// "sync"
	"net"
	"time"
)

var (
	//ErrNotSetDefautEndpoint not set default enpoints
	ErrNotSetDefautEndpoint = errors.New("Not set default endpoint")
)

// type  Endpoint struct{
// 	Host string
// 	Port string
// }

type TType int

const (
	Eunknown       TType = -1
	EAnyType       TType = 0
	EHttp          TType = 1
	EThriftBinary  TType = 2
	EThriftCompact TType = 3
	EGrpc          TType = 4
	EGrpcWeb       TType = 5
)

func (t TType) String() string {
	switch t {
	case Eunknown:
		return "Eunknown"
	case EAnyType:
		return "EAnyType"
	case EHttp:
		return "Ehttp"
	case EThriftBinary:
		return "EThriftBinary"
	case EThriftCompact:
		return "EThriftCompact"
	case EGrpc:
		return "EGrpc"
	case EGrpcWeb:
		return "EGrpcWeb"
	}
	return "UnknownType"
}

func StringToTType(t string) TType {
	switch t {
	case "thrift_compact":
		return EThriftCompact
	case "thrift_binary":
		return EThriftBinary
	case "grpc":
		return EGrpc
	case "grpc_web":
		return EGrpcWeb
	default:
		return Eunknown
	}
}

func ParseProtocol(name string) TType {
	switch name {
	case "binary":
		return EThriftBinary
	case "compact":
		return EThriftCompact
	}
	return Eunknown
}

type Endpoint struct {
	Host      string
	Port      string
	Type      TType
	ServiceID string
}

func (e *Endpoint) IsGoodEndpoint() bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(e.Host, e.Port), 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func NewEndPoint(aHost string, aPort string, aType TType) *Endpoint {
	return &Endpoint{
		Host: aHost,
		Port: aPort,
		Type: aType,
	}
}

//EnpointManagerIf
type InMemEndpointManager struct {
	defaultEndpoints map[string]*Endpoint
}

//GetEndpoint get endpoint by service id
func (o *InMemEndpointManager) GetEndpoint(serviceID string) (host, port string, err error) {

	ep := o.defaultEndpoints[serviceID]
	if ep != nil {
		host = ep.Host
		port = ep.Port
		err = nil
		return
	}
	err = ErrNotSetDefautEndpoint
	return
}

//SetDefaultEntpoint set endpoint of service by id
func (o *InMemEndpointManager) SetDefaultEntpoint(serviceID, host, port string) (err error) {
	o.defaultEndpoints[serviceID] = &Endpoint{host, port, 0, ""}
	return
}

func (o *InMemEndpointManager) GetAllEndpoint(serviceID string) ([]*Endpoint, error) {
	if o.defaultEndpoints[serviceID] != nil {
		return []*Endpoint{o.defaultEndpoints[serviceID]}, nil
	}
	return nil, errors.New("Service id not existed")
}
