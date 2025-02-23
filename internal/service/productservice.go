package service

import (
	"context"

	pb "product/api/product/v1"
	"product/internal/biz"

	"google.golang.org/protobuf/types/known/structpb"
)

type ProductServiceService struct {
	pb.UnimplementedProductServiceServer
	gc *biz.GoodsUsecase
}

func NewProductServiceService(gc *biz.GoodsUsecase) *ProductServiceService {
	return &ProductServiceService{gc: gc}
}

func (s *ProductServiceService) CreateGoods(ctx context.Context, req *pb.CreateGoodsRequest) (*pb.CreateGoodsResponse, error) {
	bizReq := &biz.CreateGoodsRequest{
		Name: req.Name,
		Tags: req.Tags,
		Type: req.Type,

		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}
	resp, err := s.gc.CreateGoods(ctx, bizReq)
	if err != nil {
		return nil, err
	}

	return &pb.CreateGoodsResponse{
		Id: resp.ID,
	}, nil
}
func (s *ProductServiceService) GoodsList(ctx context.Context, req *pb.GoodsFilterRequest) (*pb.GoodsListResponse, error) {
	bizReq := &biz.GoodsFilterRequest{
		Keywords: req.Keywords,
		Type:     req.Type,
		MinPrice: req.MinPrice,
		MaxPrice: req.MaxPrice,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	resp, err := s.gc.GoodsList(ctx, bizReq)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.GoodsInfoResponse, 0)
	for _, goods := range resp.List {
		list = append(list, &pb.GoodsInfoResponse{
			Id:          goods.ID,
			Name:        goods.Name,
			Tags:        goods.Tags,
			Type:        goods.Type,
			Description: goods.Description,
			Price:       goods.Price,
			Quantity:    goods.Quantity,
		})
	}
	return &pb.GoodsListResponse{
		Total: resp.Total,
		List:  list,
	}, nil
}
func (s *ProductServiceService) UpdateGoods(ctx context.Context, req *pb.UpdateGoodsRequest) (*pb.UpdateGoodsResponse, error) {
	return &pb.UpdateGoodsResponse{}, nil
}
func (s *ProductServiceService) DeleteGoods(ctx context.Context, req *pb.DeleteGoodsRequest) (*pb.DeleteGoodsResponse, error) {
	return &pb.DeleteGoodsResponse{}, nil
}
func (s *ProductServiceService) SearchGoods(ctx context.Context, req *pb.SearchGoodsRequest) (*pb.SearchGoodsResponse, error) {
	bizReq := &biz.SearchGoodsRequest{
		Query: structToMap(req.Query),
		Page:  req.Page,
		Size:  req.PageSize,
	}
	resp, err := s.gc.SearchGoods(ctx, bizReq)
	if err != nil {
		return nil, err
	}
	results := make([]*pb.GoodsInfoResponse, 0)
	for _, goods := range resp.Results {
		results = append(results, &pb.GoodsInfoResponse{
			Id:          goods.ID,
			Name:        goods.Name,
			Tags:        goods.Tags,
			Type:        goods.Type,
			Description: goods.Description,
			Price:       goods.Price,
			Quantity:    goods.Quantity,
		})
	}
	return &pb.SearchGoodsResponse{
		Results: results,
	}, nil
}
func (s *ProductServiceService) AutocompleteSearch(ctx context.Context, req *pb.AutoCompleteRequest) (*pb.AutoCompleteResponse, error) {
	return &pb.AutoCompleteResponse{}, nil
}

func structToMap(s *structpb.Struct) map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.AsMap()
}
