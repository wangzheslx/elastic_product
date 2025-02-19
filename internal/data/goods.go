package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"product/internal/biz"
	"strconv"

	"errors"

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
func (g *greeterRepo) SearchGoods(context.Context, *biz.SearchGoodsRequest) (*biz.SearchGoodsResponse, error) {

	panic("unimplemented")
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
