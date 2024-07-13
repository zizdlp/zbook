package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
)

func (server *Server) LogVisitor(ctx context.Context, req *rpcs.LogVisitorRequest) (*rpcs.LogVisitorResponse, error) {
	err := server.LogRedisVisitor(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &rpcs.LogVisitorResponse{}
	return rsp, nil
}
