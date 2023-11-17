package domain

type ImageMapper struct {
}

func (m *ImageMapper) MapToDomain(source ImageRequest) Image {

	return Image{
		Name: source.Name,
	}
}

func (m *ImageMapper) MapToDetailedResponse(source Image) ImageDetailedResponse {
	return ImageDetailedResponse{
		Id:   source.Id,
		Name: source.Name,
	}
}

func (m *ImageMapper) MapToHeaderResponse(source Image) ImageHeaderResponse {

	return ImageHeaderResponse{
		Id:   source.Id,
		Name: source.Name,
	}
}

func (m *ImageMapper) MapToHeaderResponseSlice(source []Image) []ImageHeaderResponse {
	var targetSlice []ImageHeaderResponse
	for _, image := range source {
		targetSlice = append(targetSlice, ImageHeaderResponse{Id: image.Id, Name: image.Name})
	}

	return targetSlice
}
