package EndpointsManager

//EnpointManagerIf interface of enpoint manager
type EnpointManagerIf interface {
	GetEndpoint(serviceID string) (host, port string, err error)
	GetAllEndpoint(serviceID string) ([]*Endpoint, error)
	SetDefaultEntpoint(serviceID, host, port string) (err error)
}
