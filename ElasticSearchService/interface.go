package ElasticSearchService

type Client interface {
	Index(indexName string, docID string, documentJson string) (bool, error)

	Search(indexName string, query map[string]interface{}) (rawResult []byte, err error)

	Delete(indexName string, docID string) (bool, error)

	Get(indexName string, docID string) (rawResult []byte, err error)

	Update(indexName string, docID string, documentJson string) (bool, error)
}
