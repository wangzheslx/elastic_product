package data

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"product/internal/biz"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// AutocompleteSearch implements biz.GoodsRepo.
func (g *greeterRepo) AutocompleteSearch(context.Context, *biz.AutoCompleteRequest) (*biz.AutoCompleteResponse, error) {
	panic("unimplemented")
}

//create Index

// func (g *greeterRepo) CreateGoods(context.Context, *biz.CreateGoodsRequest) (*biz.CreateGoodsResponse, error) {
// 	fmt.Println("create goods")
// 	dsl := bytes.NewBufferString(
// 		`{
//   "settings": {
//     "number_of_shards": 3,  // 设置为 3，根据你的数据量和性能需求调整
//     "number_of_replicas": 0  // 单机部署不需要副本
//   },
//   "mappings": {
//     "properties": {
//       "name": {
//         "type": "text",
//         "fields": {
//           "keyword": {
//             "type": "keyword",
//             "ignore_above": 256
//           }
//         }
//       },
//       "tags": {
//         "type": "text",
//         "fields": {
//           "keyword": {
//             "type": "keyword",
//             "ignore_above": 256
//           }
//         }
//       },
//       "type": {
//         "type": "keyword"
//       },
//       "description": {
//         "type": "text"
//       },
//       "price": {
//         "type": "double"
//       },
//       "quantity": {
//         "type": "integer"
//       }
//     }
//   }
// }`)
// 	createindices := g.data.EsClient.Indices.Create
// 	resp, err := createindices("goods", createindices.WithBody(dsl))
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(resp)
// 	return nil, err
// }

// CreateGoods implements biz.GoodsRepo.
func (g *greeterRepo) CreateGoods(ctx context.Context, req *biz.CreateGoodsRequest) (*biz.CreateGoodsResponse, error) {
	var reqmap = make(map[string]interface{})
	reqmap["name"] = req.Name
	reqmap["tags"] = req.Tags
	reqmap["type"] = req.Type
	reqmap["description"] = req.Description
	reqmap["price"] = req.Price
	reqmap["quantity"] = req.Quantity

	restr, err := json.Marshal(reqmap)
	if err != nil {
		return nil, err
	}
	doc := bytes.NewBufferString(
		//封装字符串
		// `{
		// 	"name": "iPhone 13",
		// 	"tags": "手机,苹果",
		// 	"type": "电子产品",
		// 	"description": "苹果iPhone 13手机，6.1英寸屏幕，4G网络，64GB存储",
		// 	"price": 8999.0,
		// 	"quantity": 100
		//   }`
		string(restr),
	)
	//index 创建随机索引
	// index := g.data.EsClient.Index
	// resp, err := index("mygoods", doc, index.WithPretty())
	// if err != nil {
	// 	return nil, err
	// }
	create := g.data.EsClient.Create
	ids := req.Quantity
	idStrs := strconv.Itoa(int(ids))
	resp, err := create("mygoods", idStrs, doc, create.WithPretty())
	if err != nil {
		return nil, err
	}

	g.log.Info("****************", resp)
	fmt.Println(resp)
	var respmap map[string]interface{}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &respmap)
	if err != nil {
		return nil, err
	}
	if respmap["error"] != nil {
		//返回错误信息

		if reason, ok := respmap["error"].(map[string]interface{}); ok {
			return nil, errors.New(reason["reason"].(string))

		}
		return nil, errors.New("change errr")

	}

	fmt.Println("###################\n", respmap)
	fmt.Println(respmap["_id"])

	idStr := respmap["_id"].(string)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	g.log.Info("******************", id)
	return &biz.CreateGoodsResponse{
		ID: id,
	}, err
}

// DeleteGoods implements biz.GoodsRepo.
func (g *greeterRepo) DeleteGoods(context.Context, *biz.DeleteGoodsRequest) (*biz.DeleteGoodsResponse, error) {
	panic("unimplemented")
}

// GoodsList implements biz.GoodsRepo.
func (g *greeterRepo) GoodsList(context.Context, *biz.GoodsFilterRequest) (*biz.GoodsListResponse, error) {

	panic("unimplemented")
}

// SearchGoods implements biz.GoodsRepo.
func (g *greeterRepo) SearchGoods(ctx context.Context, req *biz.SearchGoodsRequest) (*biz.SearchGoodsResponse, error) {
	mustConditions := []map[string]interface{}{}
	for field, value := range req.Query {
		if value != nil {
			if values, ok := value.([]interface{}); ok {
				shouldConditions := []map[string]interface{}{}
				for _, v := range values {
					if str, ok := v.(string); ok && len(str) > 0 {
						if len(str) >= 2 && str[0] == '/' && str[len(str)-1] == '/' {
							shouldConditions = append(shouldConditions, map[string]interface{}{
								"regexp": map[string]interface{}{
									field: str[1 : len(str)-1],
								},
							})
						} else if containsWildcard(str) {
							shouldConditions = append(shouldConditions, map[string]interface{}{
								"wildcard": map[string]interface{}{
									field: str,
								},
							})
						} else if len(str) >= 1 && str[0] == '#' {
							shouldConditions = append(shouldConditions, map[string]interface{}{
								"match": map[string]interface{}{
									field: str[1:], // 去掉开头的 #
								},
							})
						} else {
							shouldConditions = append(shouldConditions, map[string]interface{}{
								"term": map[string]interface{}{
									field: str,
								},
							})
						}
					}
				}
				if len(shouldConditions) > 0 {
					mustConditions = append(mustConditions, map[string]interface{}{
						"bool": map[string]interface{}{
							"should":               shouldConditions,
							"minimum_should_match": 1,
						},
					})
				}
			} else if str, ok := value.(string); ok && len(str) > 0 {
				if len(str) >= 2 && str[0] == '/' && str[len(str)-1] == '/' {
					mustConditions = append(mustConditions, map[string]interface{}{
						"regexp": map[string]interface{}{
							field: str[1 : len(str)-1],
						},
					})
				} else if containsWildcard(str) {
					mustConditions = append(mustConditions, map[string]interface{}{
						"wildcard": map[string]interface{}{
							field: str,
						},
					})
				} else if len(str) >= 1 && str[0] == '#' {
					mustConditions = append(mustConditions, map[string]interface{}{
						"match": map[string]interface{}{
							field: str[1:], // 去掉开头的 #
						},
					})
				} else {
					mustConditions = append(mustConditions, map[string]interface{}{
						"term": map[string]interface{}{
							field: str,
						},
					})
				}
			}
		}
	}
	esQuery := map[string]interface{}{
		"from": (req.Page - 1) * req.Size,
		"size": req.Size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": mustConditions,
			},
		},
	}

	// 动态添加价格筛选条件
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		priceRange := map[string]interface{}{}
		if req.MinPrice > 0 {
			priceRange["gte"] = req.MinPrice
		}
		if req.MaxPrice > 0 {
			priceRange["lte"] = req.MaxPrice
		}
		esQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			esQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"range": map[string]interface{}{
					"price": priceRange,
				},
			},
		)
	}

	// 动态添加排序条件
	if req.Sort != "" {
		// 动态添加排序条件
		sortOrder := "asc" // 默认升序
		if req.Sort == "desc" {
			sortOrder = "desc"
		}
		esQuery["sort"] = []map[string]interface{}{
			{
				"price": map[string]interface{}{
					"order": sortOrder,
				},
			},
		}
	}

	queryJSON, err := json.Marshal(esQuery)
	if err != nil {
		return nil, err
	}

	search := g.data.EsClient.Search
	resp, err := search(
		search.WithContext(ctx),
		search.WithIndex("mygoods_v2"),
		search.WithBody(bytes.NewReader(queryJSON)),
		search.WithTrackTotalHits(true),
		search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, fmt.Errorf("error in search response: %s", resp.String())
	}

	var searchResult map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, err
	}

	hits := searchResult["hits"].(map[string]interface{})["hits"].([]interface{})
	results := make([]*biz.GoodsInfoResponse, 0, len(hits))
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		idStr := hit.(map[string]interface{})["_id"].(string)
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}
		results = append(results, &biz.GoodsInfoResponse{
			ID:          id,
			Name:        source["name"].(string),
			Tags:        source["tags"].(string),
			Type:        source["type"].(string),
			Description: source["description"].(string),
			Price:       source["price"].(float64),
			Quantity:    int32(source["quantity"].(float64)),
		})
	}

	return &biz.SearchGoodsResponse{
		Results: results,
	}, nil
}

// containsWildcard checks if a string contains wildcard characters.
func containsWildcard(s string) bool {
	return strings.Contains(s, "*")
}

// UpdateGoods implements biz.GoodsRepo.
func (g *greeterRepo) UpdateGoods(context.Context, *biz.UpdateGoodsRequest) (*biz.UpdateGoodsResponse, error) {
	panic("unimplemented")
}

func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
