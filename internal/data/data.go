package data

import (
	"product/internal/conf"

	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGoodsRepo, NewElasticsearch)

// Data .
type Data struct {
	// TODO wrapped database client
	EsClient *es.Client
	logger   *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, es *es.Client) (*Data, func(), error) {

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		logger:   log.NewHelper(logger),
		EsClient: es,
	}, cleanup, nil
}

func NewElasticsearch(c *conf.Data) *es.Client {
	client, err := es.NewClient(es.Config{
		Addresses: []string{c.Elastic.Addr},
	})
	if err != nil {
		panic(err)
	}
	pingResp, err := client.Ping(client.Ping.WithPretty(), client.Ping.WithHuman())
	if err != nil {
		panic(err)
	}
	log.Info("Elasticsearch ping response: %s", pingResp.String())
	infoResp, err := client.Info(client.Info.WithHuman())
	if err != nil {
		panic(err)
	}
	log.Info("Elasticsearch info response: %s", infoResp.String())

	return client

}
