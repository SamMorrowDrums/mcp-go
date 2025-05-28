package client

import (
	"github.com/sammorrowdrums/mcp-go/client/transport"
	"github.com/sammorrowdrums/mcp-go/server"
)

// NewInProcessClient connect directly to a mcp server object in the same process
func NewInProcessClient(server *server.MCPServer) (*Client, error) {
	inProcessTransport := transport.NewInProcessTransport(server)
	return NewClient(inProcessTransport), nil
}
