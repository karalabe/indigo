package bsky

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// schema: app.bsky.notification.getUnreadCount

func init() {
}

type NotificationGetUnreadCount_Output struct {
	Count int64 `json:"count" cborgen:"count"`
}

func NotificationGetUnreadCount(ctx context.Context, c *xrpc.Client, seenAt string) (*NotificationGetUnreadCount_Output, error) {
	var out NotificationGetUnreadCount_Output

	params := map[string]interface{}{
		"seenAt": seenAt,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.notification.getUnreadCount", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
