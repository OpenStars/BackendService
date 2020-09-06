package SorlClientService

import solr "github.com/rtt/Go-Solr"

type Client interface {
	AddDataToSolr(solrID int64, solrData map[string]interface{}) (err error)
	UpdateDataSolr(query string, solrData map[string]interface{}) (err error)
	SelectByQueryString(query string) (result []solr.Document, total int64, err error)
	SelectByQueryParams(query string) (result []solr.Document, total int64, err error)
	DeleteDataByIndex(index int64) bool
	DeleteAllData() bool
}
