package domain

type ContainerMapper struct {
}

func (c *ContainerMapper) MapToDomain(source *ContainerRequest) *Container {

	return &Container{
		Name:    source.Name,
		ImageId: source.ImageId,
		Command: source.Command,
	}
}

func (c *ContainerMapper) MapToHeaderResponse(source *Container) *ContainerHeaderResponse {

	return &ContainerHeaderResponse{
		Id:      source.Id,
		Name:    source.Name,
		ImageId: source.ImageId,
		Command: source.Command,
		Created: source.Created,
		Status:  source.Status,
	}
}

func (c *ContainerMapper) MapToHeaderResponseSlice(sourceSlice *[]Container) *[]ContainerHeaderResponse {
	var targetSlice []ContainerHeaderResponse
	for _, source := range *sourceSlice {
		targetSlice = append(targetSlice, *c.MapToHeaderResponse(&source))
	}
	return &targetSlice
}

func (c *ContainerMapper) MapToDetailedResponse(source *Container) *ContainerDetailedResponse {

	return &ContainerDetailedResponse{
		Id:      source.Id,
		Name:    source.Name,
		ImageId: source.ImageId,
		Command: source.Command,
		Created: source.Created,
		Status:  source.Status,
	}
}
