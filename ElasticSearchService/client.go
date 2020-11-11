package ElasticSearchService

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Location struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}

func ConvertToMapInterface(stu interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(stu).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.Type.String() == "map[string]string" {
			for k, v := range valueField.Interface().(map[string]string) {
				if k == "uid" {
					vi, _ := strconv.ParseInt(v, 10, 64)
					result[k] = vi
					continue
				}
				if k == "location" {
					var locinfo Location
					json.Unmarshal([]byte(v), &locinfo)
					result[k] = locinfo
				}
				result[k] = v
			}
			continue
		}

		valuedatabytes, _ := json.Marshal(valueField.Interface())
		valueDataString := string(valuedatabytes)

		if valueDataString == "0" || valueDataString == "" || valueDataString == "null" || valueDataString == "\"\"" {
			continue
		}
		tags := strings.Split(typeField.Tag.Get("json"), ",")
		if len(tags) == 0 {
			continue
		}

		result[tags[0]] = valueField.Interface()
	}
	return result
}

func ToJson(stu interface{}) string {
	mapData := ConvertToMapInterface(stu)
	databytes, _ := json.Marshal(mapData)
	if databytes == nil {
		return ""
	}
	return string(databytes)
}

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

type ShardsInfo struct {
	Total      int64 `json:"total,omitempty"`
	Successful int64 `json:"successful,omitempty"`
	Skipped    int64 `json:"skipped,omitempty"`
	Failed     int64 `json:"failed,omitempty"`
}

type TotalHits struct {
	Value    int64  `json:"value,omitempty"`
	Relation string `json:"relation,omitempty"`
}
type HitElement struct {
	Index  string      `json:"_index,omitempty"`
	Type   string      `json:"_type,omitempty"`
	Score  float32     `json:"_score,omitempty"`
	Source interface{} `json:"_source,omitempty"`
}
type HitsInfo struct {
	Total    *TotalHits    `json:"total,omitempty"`
	MaxScore *float64      `json:"max_score,omitempty"`
	Hits     []*HitElement `json:"hits,omitempty"`
}

type AggsResult struct {
	Took    int64       `json:"took,omitempty"`
	TimeOut bool        `json:"time_out,omitempty"`
	Shards  *ShardsInfo `json:"_shards,omitempty"`
	Hits    HitsInfo    `json:"hits,omitempty"`
}
