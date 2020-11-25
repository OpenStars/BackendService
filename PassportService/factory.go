package PassportService

func NewPassPortService(etcdServers []string, serviceID, defaulHost, defaultPort string) PassportService {

	reportitem := &ppassportservice{
		host: defaulHost,
		port: defaultPort,
		sid:  serviceID,
	}

	return reportitem
}
