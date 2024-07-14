package service

import (
	"context"

	"github.com/thiagohmm/gRPCEstudo/internal/database"
	"github.com/thiagohmm/gRPCEstudo/internal/pb"
)

type CatefgoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CatefgoryService {
	return &CatefgoryService{CategoryDB: categoryDB}
}

func (c *CatefgoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
