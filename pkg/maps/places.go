package maps

import (
	"context"
	"fmt"

	"cloud.google.com/go/maps/places/apiv1/placespb"
	"github.com/googleapis/gax-go/v2/callctx"
	"google.golang.org/genproto/googleapis/type/latlng"
)

const (
	placesFieldMask = "places.displayName,places.types"
)

func (c *Client) GetNearbyRestaurants(ctx context.Context, coords *latlng.LatLng) error {
	req := &placespb.SearchNearbyRequest{
		RegionCode:           "US",
		IncludedPrimaryTypes: []string{"restaurant"},
		MaxResultCount:       20,
		RankPreference:       placespb.SearchNearbyRequest_DISTANCE,
		LocationRestriction: &placespb.SearchNearbyRequest_LocationRestriction{
			Type: &placespb.SearchNearbyRequest_LocationRestriction_Circle{
				Circle: &placespb.Circle{
					Center: coords,
					// In meters
					Radius: 5000,
				},
			},
		},
	}

	ctx = callctx.SetHeaders(ctx, callctx.XGoogFieldMaskHeader, placesFieldMask)

	res, err := c.placesClient.SearchNearby(ctx, req)

	if err != nil {
		return fmt.Errorf("failed to search nearby: %w", err)
	}

	for _, place := range res.GetPlaces() {
		fmt.Println("Place:", place.GetDisplayName().GetText())

		for _, t := range place.GetTypes() {
			if t == "restaurant" {
				continue
			}

			if displayName := restaurantTypes[t]; displayName != "" {
				fmt.Println("  -", displayName)
			}
		}
	}

	return nil
}
