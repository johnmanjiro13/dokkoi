package google

import (
	"google.golang.org/api/customsearch/v1"
)

type customSearchRepository struct {
	svc      *customsearch.Service
	engineID string
}

type CustomSearchRepository interface {
	SearchImage(query string) (*customsearch.Search, error)
}

func NewCustomSearchRepository(service *customsearch.Service, engineID string) CustomSearchRepository {
	return &customSearchRepository{
		svc:      service,
		engineID: engineID,
	}
}

func (r *customSearchRepository) SearchImage(query string) (*customsearch.Search, error) {
	search := r.svc.Cse.List().Cx(r.engineID).SearchType("image").Num(1).Q(query)
	search.Start(1)
	return search.Do()
}
