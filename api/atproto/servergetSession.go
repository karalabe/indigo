package atproto

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// schema: com.atproto.server.getSession

func init() {
}

type ServerGetSession_Output struct {
	Did    string  `json:"did" cborgen:"did"`
	Email  *string `json:"email,omitempty" cborgen:"email,omitempty"`
	Handle string  `json:"handle" cborgen:"handle"`
}

func ServerGetSession(ctx context.Context, c *xrpc.Client) (*ServerGetSession_Output, error) {
	var out ServerGetSession_Output
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.server.getSession", nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
