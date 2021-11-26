package ElasticSearchService

type Client interface {
	Index(indexName string, docID string, documentJson string) (bool, error)

	IndexBulk(indexName string, docIDField string, bulkDocumentJson string) (bool, error)

	Search(indexName string, query map[string]interface{}) (rawResult []byte, err error)

	Delete(indexName string, docID string) (bool, error)

	Get(indexName string, docID string) (r interface{}, err error)

	Update(indexName string, docID string, documentJson string) (bool, error)

	SearchRawString(indexName, rawQuery string) (rawResult [][]byte, total int64, err error)
}
