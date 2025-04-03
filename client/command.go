package client

import context "context"

type CommandService struct {
	UnimplementedCommandServer
}

func NewCommandService() *CommandService {
	return &CommandService{}
}

func (c *CommandService) Set(ctx context.Context, req *SetCmdRequest) (*SetCmdResponse, error) {
	return nil, nil
}
func (c *CommandService) Get(ctx context.Context, req *GetCmdRequest) (*GetCmdResponse, error) {
	return nil, nil
}
func (c *CommandService) Del(ctx context.Context, req *DeleteCmdRequest) (*DeleteCmdResponse, error) {
	return nil, nil
}
