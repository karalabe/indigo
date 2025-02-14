package atproto

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// schema: com.atproto.admin.updateAccountEmail

func init() {
}

type AdminUpdateAccountEmail_Input struct {
	Account string `json:"account" cborgen:"account"`
	Email   string `json:"email" cborgen:"email"`
}

func AdminUpdateAccountEmail(ctx context.Context, c *xrpc.Client, input *AdminUpdateAccountEmail_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.admin.updateAccountEmail", nil, input, nil); err != nil {
		return err
	}

	return nil
}
