package SorlClientService

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	solr "github.com/rtt/Go-Solr"
)

type solrclientservice struct {
	host       string
	port       int
	collection string
	sconn      *solr.Connection
}

// AddDataToSolr add data -> solr
// solrID để sau này tìm id đó update field
func (o *solrclientservice) AddDataToSolr(solrID int64, solrData map[string]interface{}) (err error) {
	// build an update document, in this case adding two documents
	// init a connection
	solrData["id"] = solrID
	solrData["solr_id"] = solrID
	f := map[string]interface{}{
		"add": []interface{}{
			solrData,
		},
	}
	// send off the update (2nd parameter indicates we also want to commit the operation)
	_, err = o.sconn.Update(f, true)

	if err != nil {
		// logger.Error("[AddDataToSolr] solrID = %d, solrData = %v, error = %v \n", solrID, solrData, err)
		return
	}

	// logger.Info("[AddDataToSolr] add data success solrid = %d, solrData = %v \n", solrID, solrData)
	return
}

// UpdateDataSolr by query
// example: query= "q=uid:1" update data for data user id = 1
func (o *solrclientservice) UpdateDataSolr(query string, solrData map[string]interface{}) (err error) {
	// build an update document, in this case adding two documents
	// init a connection

	res, _ := o.sconn.SelectRaw(query)
	if int64(res.Results.NumFound) <= 0 {
		// logger.Info("[UpdateDataSolr] Can not found data query = %s \n", query)
		return errors.New("can not found data")
	}
	solrID := res.Results.Collection[0].Field("id").(string)
	solrData["id"] = solrID
	solrData["solr_id"] = solrID
	f := map[string]interface{}{
		"add": []interface{}{
			solrData,
		},
	}
	// send off the update (2nd parameter indicates we also want to commit the operation)
	_, err = o.sconn.Update(f, true)

	if err != nil {
		// logger.Error("[UpdateDataSolr] solrID = %d, solrData = %v, error = %v \n", solrID, solrData, err)
		return
	}
	return
}

// SelectByQueryString query data in solr
func (o *solrclientservice) SelectByQueryString(query string) (result []solr.Document, total int64, err error) {

	// perform a query
	res, err := o.sconn.SelectRaw(query)

	if err != nil {
		// logger.Info("[SelectByQueryString] err select: %v \n", err)
		return make([]solr.Document, 0), 0, err
	}

	if res != nil && res.Results != nil && res.Results.Collection != nil && len(res.Results.Collection) > 0 {
		return res.Results.Collection, int64(res.Results.NumFound), err
	}

	//logger.Info("[SelectByQueryString] len result : %d \n", len(result))
	return make([]solr.Document, 0), 0, err
}

// SelectByQueryParams query data in solr
func (o *solrclientservice) SelectByQueryParams(query string) (result []solr.Document, total int64, err error) {

	// perform a query
	mapParam := make(map[string][]string)
	mapParam["q"] = []string{query}
	res, err := o.sconn.Select(
		&solr.Query{
			Params: mapParam,
		},
	)

	if err != nil {
		// logger.Info("[SelectByQueryString] err select: %v \n", err)
		return make([]solr.Document, 0), 0, err
	}

	if res != nil && res.Results != nil && res.Results.Collection != nil && len(res.Results.Collection) > 0 {
		return res.Results.Collection, int64(res.Results.NumFound), err
	}

	//logger.Info("[SelectByQueryString] len result : %d \n", len(result))
	return make([]solr.Document, 0), 0, err
}

// DeleteDataByIndex delete data by solr id
func (o *solrclientservice) DeleteDataByIndex(index int64) bool {

	body := fmt.Sprintf("<delete><query>id:%d</query></delete>", index)

	client := &http.Client{}
	// build a new request, but not doing the POST yet
	address := fmt.Sprintf("http://%s:%s/solr/%s/update?commit=true", o.host, o.port, o.collection)
	req, err := http.NewRequest("POST", address, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
	}
	// you can then set the Header here
	// I think the content-type should be "application/xml" like json...
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	// now POST it
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	return true
}

// DeleteAllData delete all data
func (o *solrclientservice) DeleteAllData() bool {

	body := "<delete><query>*:*</query></delete>"

	client := &http.Client{}
	// build a new request, but not doing the POST yet
	address := fmt.Sprintf("http://%s:%s/solr/%s/update?commit=true", o.host, o.port, o.collection)
	req, err := http.NewRequest("POST", address, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
	}
	// you can then set the Header here
	// I think the content-type should be "application/xml" like json...
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	// now POST it
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	return true
}
