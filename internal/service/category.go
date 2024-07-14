package service

import (
	"context"
	"io"

	"github.com/thiagohmm/gRPCEstudo/internal/database"
	"github.com/thiagohmm/gRPCEstudo/internal/pb"
	// Add this line to import the package that defines pb.GetCategoryRequest
)

type CategoryServiceProto struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryServiceProto {
	return &CategoryServiceProto{CategoryDB: categoryDB}
}

func (c *CategoryServiceProto) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryServiceProto) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoryList []*pb.Category
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.CategoryList{Categories: categoryList}, nil

}

func (c *CategoryServiceProto) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryServiceProto) CreateCateoryStream(stream pb.CategoryService_CreateCateoryStreamServer) error {
	categories := []*pb.Category{}
	for {
		categor, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CategoryList{Categories: categories})
		}
		if err != nil {
			return err
		}
		category, err := c.CategoryDB.Create(categor.Name, categor.Description)
		if err != nil {
			return err
		}
		categories = append(categories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})

	}
}

func (c *CategoryServiceProto) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		categoryDB, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categoryResponse := &pb.Category{
			Id:          categoryDB.ID,
			Name:        categoryDB.Name,
			Description: categoryDB.Description,
		}
		if err := stream.Send(categoryResponse); err != nil {
			return err
		}
	}
}
