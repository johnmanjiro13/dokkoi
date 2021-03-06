package google

import (
	"math/rand"

	"github.com/spf13/viper"
	"google.golang.org/api/customsearch/v1"

	"github.com/johnmanjiro13/dokkoi/command"
)

func init() {
	viper.BindEnv("customsearch.api.key", "CUSTOMSEARCH_API_KEY")
	viper.BindEnv("customsearch.engine.id", "CUSTOMSEARCH_ENGINE_ID")

	viper.SetDefault("customsearch.api.key", "")
	viper.SetDefault("customsearch.engine.id", "")
}

const imageNum = 5

type customSearchRepository struct {
	svc      *customsearch.Service
	engineID string
}

func NewCustomSearchRepository(service *customsearch.Service, engineID string) command.CustomSearchRepository {
	return &customSearchRepository{
		svc:      service,
		engineID: engineID,
	}
}

func (r *customSearchRepository) SearchImage(query string) (*customsearch.Result, error) {
	search := r.svc.Cse.List().Cx(r.engineID).
		SearchType("image").
		Num(imageNum).
		Q(query).
		Start(1)
	res, err := search.Do()
	if err != nil {
		return nil, err
	}
	images := res.Items
	if len(images) <= 0 {
		return nil, command.ErrImageNotFound
	}
	return images[rand.Intn(len(images))], nil
}
