package ElasticSearchService

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

type client struct {
	esclient *elasticsearch.Client
	url      string
}

func (m *client) Index(indexName, docID, documentJson string) (bool, error) {
	req := esapi.IndexRequest{
		Index:      indexName,
		Body:       strings.NewReader(documentJson),
		DocumentID: docID,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), m.esclient)
	if err != nil {
		return false, err
	}
	if res.IsError() {
		return false, errors.New("Error Indexing document " + res.Status())
	}
	defer res.Body.Close()
	resultBody, err := ioutil.ReadAll(res.Body)
	var resultBodyMap map[string]interface{}
	err = json.Unmarshal(resultBody, &resultBodyMap)
	if err != nil {
		return false, errors.New("Response unvalid")
	}
	if resultBodyMap["_id"] == nil {
		return false, errors.New("Index document error " + string(resultBody))
	}
	// log.Println("[ESINFO] Index document response", string(resultBody))
	return true, nil
}

func (m *client) Search(indexName string, query map[string]interface{}) (rawResult []byte, err error) {
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}
	res, err := m.esclient.Search(
		m.esclient.Search.WithContext(context.Background()),
		m.esclient.Search.WithIndex(indexName),
		m.esclient.Search.WithBody(&buf),
		m.esclient.Search.WithTrackTotalHits(true),
		m.esclient.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resultBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if string(resultBody) == "" {
		return nil, errors.New("NOT FOUND")
	}
	return resultBody, nil

}

func (m *client) Delete(indexName string, docID string) (bool, error) {
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: docID,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), m.esclient)
	if err != nil {
		return false, err
	}
	if res.IsError() {
		return false, errors.New("Error Indexing document " + res.Status())
	}
	log.Println("[ESINFO] Index document response", res.String())
	return true, nil
}

func (m *client) Get(indexName string, id string) (rawResult []byte, err error) {
	req := esapi.GetRequest{
		Index:      indexName,
		DocumentID: id,
	}
	res, err := req.Do(context.Background(), m.esclient)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New("Error Get document " + res.String())
	}
	defer res.Body.Close()
	resultBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if string(resultBody) == "" {
		return nil, errors.New("NOT FOUND")
	}
	return resultBody, nil
}

func (m *client) Update(indexName string, id string, documentJson string) (bool, error) {

	req := esapi.UpdateRequest{
		Index:      indexName,
		DocumentID: id,
		Body:       strings.NewReader(documentJson),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), m.esclient)
	if err != nil {
		return false, err
	}
	if res.IsError() {
		return false, errors.New("Error Update document " + res.Status())
	}
	log.Println("[ESINFO] Index document response", res.String())
	return true, nil
}

func (m *client) DeteleIndex(indexName string) (bool, error) {

	req := esapi.DeleteRequest{
		Index: indexName,
	}
	res, err := req.Do(context.Background(), m.esclient)
	if err != nil {
		return false, err
	}
	if res.IsError() {
		return false, errors.New("Error delete document " + res.Status())
	}
	log.Println("[ESINFO] delete document response", res.String())
	return true, nil
}

func ParseResultToDocuments(rawResult []byte) ([]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(rawResult, &result)
	if err != nil {
		return nil, err
	}
	if result["hits"] == nil {
		return nil, errors.New("NOT FOUND")
	}
	hits := result["hits"].(map[string]interface{})
	if hits == nil {
		return nil, errors.New("NOT FOUND")
	}
	if hits["hits"] == nil {
		return nil, errors.New("NOT FOUND")
	}
	histhist := hits["hits"].([]interface{})
	if histhist == nil {
		return nil, errors.New("NOT FOUND")
	}
	var listDoc []interface{}
	for _, h := range histhist {
		hi := h.(map[string]interface{})
		if hi == nil {
			continue
		}
		if hi["_source"] == nil {
			continue
		}
		listDoc = append(listDoc, hi["_source"].(interface{}))
	}
	if len(listDoc) == 0 {
		return nil, errors.New("NOT FOUND")
	}
	return listDoc, nil
}
