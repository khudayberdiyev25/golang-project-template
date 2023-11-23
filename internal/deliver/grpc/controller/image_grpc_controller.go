package controller

import (
	"context"
	"golang-project-template/internal/deliver/grpc/mapper"
	"golang-project-template/internal/deliver/grpc/stub"
	"golang-project-template/internal/domain"
)

type ImageServerImpl struct {
	stub.UnimplementedImageServiceServer
	UseCase domain.ImageUseCase
	Mapper  *mapper.ImageGrpcMapper
}

func (i *ImageServerImpl) Create(ctx context.Context, grpcRequest *stub.ImageRequest) (*stub.ImageId, error) {
	request := i.Mapper.MapToDomainFromRequest(grpcRequest)
	response, err := i.UseCase.Create(request)
	if err != nil {
		return nil, err
	}
	return i.Mapper.MapToImageIdFromDomainId(int64(response)), nil
}

func (i *ImageServerImpl) FindAll(ctx context.Context, grpcRequest *stub.ImageSpecification) (*stub.ImageHeaderListResponse, error) {
	name := grpcRequest.GetName()

	response, err := i.UseCase.Filter(name)
	if err != nil {
		return nil, err
	}

	return i.Mapper.MapToHeaderSliceFromDomainHeaderSlice(response), nil
}

func (i *ImageServerImpl) GetOne(ctx context.Context, grpcRequest *stub.ImageKey) (*stub.ImageDetailedReply, error) {
	image, err := i.UseCase.GetOne(grpcRequest.GetKey())
	if err != nil {
		return nil, err
	}
	return i.Mapper.MapToDetailedReply(image), nil
}

func (i *ImageServerImpl) DeleteOne(ctx context.Context, grpcRequest *stub.ImageKey) (*stub.ImageEmptyReply, error) {
	err := i.UseCase.DeleteByIdOrName(grpcRequest.GetKey())
	if err != nil {
		return nil, err
	}
	return &stub.ImageEmptyReply{}, nil
}

func (i *ImageServerImpl) Delete(context.Context, *stub.ImageSpecification) (*stub.ImagePruneReply, error) {
	response, err := i.UseCase.DeleteAllUnusedOnes()
	if err != nil {
		return nil, err
	}
	return i.Mapper.MapToImagePruneReply(response), nil
}
