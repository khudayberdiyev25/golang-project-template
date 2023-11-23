package mapper

import (
	"golang-project-template/internal/deliver/grpc/stub"
	"golang-project-template/internal/domain"
)

type ImageGrpcMapper struct {
}

func (i *ImageGrpcMapper) MapToDomainFromRequest(source *stub.ImageRequest) *domain.ImageRequest {

	return &domain.ImageRequest{
		Name:    source.GetName(),
		RepoTag: source.GetRepoTag(),
	}
}

func (i *ImageGrpcMapper) MapToImageIdFromDomainId(id int64) *stub.ImageId {

	return &stub.ImageId{Value: int32(id)}
}

func (i *ImageGrpcMapper) MapToHeaderReplyFromDomainResponse(source *domain.ImageHeaderResponse) *stub.ImageHeaderReply {

	return &stub.ImageHeaderReply{
		Id:      int32(source.Id),
		Name:    source.Name,
		RepoTag: source.RepoTag,
	}
}

func (i *ImageGrpcMapper) MapToHeaderSliceFromDomainHeaderSlice(sourceSlice *[]domain.ImageHeaderResponse) *stub.ImageHeaderListResponse {
	var targetSlice []*stub.ImageHeaderReply
	for _, source := range *sourceSlice {
		response := i.MapToHeaderReplyFromDomainResponse(&source)
		targetSlice = append(targetSlice, response)
	}

	return &stub.ImageHeaderListResponse{Images: targetSlice}
}

func (i *ImageGrpcMapper) MapToDetailedReply(source *domain.ImageDetailedResponse) *stub.ImageDetailedReply {

	return &stub.ImageDetailedReply{
		Id:      int32(source.Id),
		Name:    source.Name,
		RepoTag: source.RepoTag,
	}
}

func (i *ImageGrpcMapper) MapToImagePruneReply(source string) *stub.ImagePruneReply {

	return &stub.ImagePruneReply{Content: source}
}
