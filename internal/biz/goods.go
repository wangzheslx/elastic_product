package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Goods 是商品模型
type Goods struct {
	ID          int64
	Name        string
	Tags        string
	Type        string
	Description string
	Price       float64
	Quantity    int32
}

// CreateGoodsRequest 是创建商品的请求消息
type CreateGoodsRequest struct {
	Name        string  // 商品名称
	Tags        string  // 商品标签
	Type        string  // 商品类型
	Description string  // 商品描述
	Price       float64 // 商品价格
	Quantity    int32   // 商品数量
	NameSuggest string  // 商品名称建议
}

// CreateGoodsResponse 是创建商品的响应消息
type CreateGoodsResponse struct {
	ID int64 // 商品ID
}

// UpdateGoodsRequest 是更新商品的请求消息
type UpdateGoodsRequest struct {
	ID          int64   // 商品ID
	Name        string  // 商品名称
	Tags        string  // 商品标签
	Type        string  // 商品类型
	Description string  // 商品描述
	Price       float64 // 商品价格
	Quantity    int32   // 商品数量
	NameSuggest string  // 商品名称建议
}

// UpdateGoodsResponse 是更新商品的响应消息
type UpdateGoodsResponse struct {
	Success bool // 更新是否成功
}

// DeleteGoodsRequest 是删除商品的请求消息
type DeleteGoodsRequest struct {
	ID int64 // 商品ID
}

// DeleteGoodsResponse 是删除商品的响应消息
type DeleteGoodsResponse struct {
	Success bool // 删除是否成功
}

// GoodsFilterRequest 是过滤商品的请求消息
type GoodsFilterRequest struct {
	Keywords string  // 关键词
	Type     string  // 商品类型
	MinPrice float64 // 最低价格
	MaxPrice float64 // 最高价格
	Page     int32   // 页码
	PageSize int32   // 每页数量
}

// GoodsListResponse 是列出商品的响应消息
type GoodsListResponse struct {
	Total int64                // 商品总数
	List  []*GoodsInfoResponse // 商品列表
}

// GoodsInfoResponse 是商品信息响应消息
type GoodsInfoResponse struct {
	ID          int64   // 商品ID
	Name        string  // 商品名称
	Tags        string  // 商品标签
	Type        string  // 商品类型
	Description string  // 商品描述
	Price       float64 // 商品价格
	Quantity    int32   // 商品数量
	NameSuggest string  // 商品名称建议
}

// SearchGoodsRequest 是搜索商品的请求消息
type SearchGoodsRequest struct {
	Query    map[string]interface{}
	Page     int32
	Size     int32
	Sort     string // 排序方式
	MaxPrice int32
	MinPrice int32
}

// SearchGoodsResponse 是搜索商品的响应消息
type SearchGoodsResponse struct {
	Results []*GoodsInfoResponse // 搜索结果
}

// AutoCompleteRequest 是自动补全请求消息
type AutoCompleteRequest struct {
	Prefix string // 前缀
}

// AutoCompleteResponse 是自动补全响应消息
type AutoCompleteResponse struct {
	Suggestions []string // 建议列表
}

// GoodsRepo 是商品仓库接口
type GoodsRepo interface {
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsResponse, error)
	GoodsList(context.Context, *GoodsFilterRequest) (*GoodsListResponse, error)
	UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsResponse, error)
	DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsResponse, error)
	SearchGoods(context.Context, *SearchGoodsRequest) (*SearchGoodsResponse, error)
	AutocompleteSearch(context.Context, *AutoCompleteRequest) (*AutoCompleteResponse, error)
}

// GoodsUsecase 是商品用例
type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

// NewGoodsUsecase 创建一个新的商品用例
func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGoods 创建一个新的商品
func (uc *GoodsUsecase) CreateGoods(ctx context.Context, req *CreateGoodsRequest) (*CreateGoodsResponse, error) {
	uc.log.WithContext(ctx).Infof("CreateGoods: %v", req)
	return uc.repo.CreateGoods(ctx, req)
}

// GoodsList 列出商品
func (uc *GoodsUsecase) GoodsList(ctx context.Context, req *GoodsFilterRequest) (*GoodsListResponse, error) {
	uc.log.WithContext(ctx).Infof("GoodsList: %v", req)
	return uc.repo.GoodsList(ctx, req)
}

// UpdateGoods 更新商品
func (uc *GoodsUsecase) UpdateGoods(ctx context.Context, req *UpdateGoodsRequest) (*UpdateGoodsResponse, error) {
	uc.log.WithContext(ctx).Infof("UpdateGoods: %v", req)
	return uc.repo.UpdateGoods(ctx, req)
}

// DeleteGoods 删除商品
func (uc *GoodsUsecase) DeleteGoods(ctx context.Context, req *DeleteGoodsRequest) (*DeleteGoodsResponse, error) {
	uc.log.WithContext(ctx).Infof("DeleteGoods: %v", req)
	return uc.repo.DeleteGoods(ctx, req)
}

// SearchGoods 搜索商品
func (uc *GoodsUsecase) SearchGoods(ctx context.Context, req *SearchGoodsRequest) (*SearchGoodsResponse, error) {
	uc.log.WithContext(ctx).Infof("SearchGoods: %v", req)
	return uc.repo.SearchGoods(ctx, req)
}

// AutocompleteSearch 自动补全搜索词
func (uc *GoodsUsecase) AutocompleteSearch(ctx context.Context, req *AutoCompleteRequest) (*AutoCompleteResponse, error) {
	uc.log.WithContext(ctx).Infof("AutocompleteSearch: %v", req)
	return uc.repo.AutocompleteSearch(ctx, req)
}
