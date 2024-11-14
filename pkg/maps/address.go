package maps

import (
	"context"
	"fmt"

	"cloud.google.com/go/maps/addressvalidation/apiv1/addressvalidationpb"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/genproto/googleapis/type/postaladdress"
)

func (c *Client) GetCoordinates(ctx context.Context, address string, city string, state string, country string) (*latlng.LatLng, error) {
	// TODO: Cache result somewhere

	if address == "" {
		return nil, fmt.Errorf("address required")
	}

	if city == "" {
		return nil, fmt.Errorf("city required")
	}

	if state == "" {
		return nil, fmt.Errorf("state required")
	}

	if country == "" {
		return nil, fmt.Errorf("country required")
	}

	res, err := c.addrClient.ValidateAddress(ctx, &addressvalidationpb.ValidateAddressRequest{
		Address: &postaladdress.PostalAddress{
			RegionCode:         country,
			AdministrativeArea: state,
			Locality:           city,
			AddressLines: []string{
				address,
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to validate address: %w", err)
	}

	return res.GetResult().GetGeocode().GetLocation(), nil
}
