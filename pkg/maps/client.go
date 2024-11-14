package maps

import (
	"context"
	"fmt"

	addressvalidation "cloud.google.com/go/maps/addressvalidation/apiv1"
	places "cloud.google.com/go/maps/places/apiv1"
)

type Client struct {
	addrClient   *addressvalidation.Client
	placesClient *places.Client
}

func New(ctx context.Context) (*Client, error) {
	addrClient, err := addressvalidation.NewClient(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to create address validation client: %v", err)
	}

	placesClient, err := places.NewClient(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to create places client: %v", err)
	}

	return &Client{
		addrClient:   addrClient,
		placesClient: placesClient,
	}, nil
}

func (c *Client) Close() {
	c.addrClient.Close()
	c.placesClient.Close()
}
