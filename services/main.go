package services

import (
	"fmt"
	"post_api_gateway/config"
	"post_api_gateway/genproto/crud_service"
	"post_api_gateway/genproto/post_service"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	PostService() post_service.PostServiceClient
	CrudService() crud_service.CrudServiceClient
}

type grpcClients struct {
	postService post_service.PostServiceClient
	crudService crud_service.CrudServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connPostService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connCrudService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CrudServiceHost, conf.CrudServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		postService: post_service.NewPostServiceClient(connPostService),
		crudService: crud_service.NewCrudServiceClient(connCrudService),
	}, nil
}

func (g *grpcClients) PostService() post_service.PostServiceClient {
	return g.postService
}

func (g *grpcClients) CrudService() crud_service.CrudServiceClient {
	return g.crudService
}
