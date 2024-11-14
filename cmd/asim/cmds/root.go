package cmds

import (
	"context"
	"fmt"
	"log"

	"github.com/evertras/address-simulator/pkg/maps"
	"github.com/spf13/cobra"

	places "cloud.google.com/go/maps/places/apiv1"
)

func doPlaces(ctx context.Context) error {
	placesClient, err := places.NewClient(ctx)

	if err != nil {
		return fmt.Errorf("failed to create places client: %v", err)
	}

	defer func() {
		err := placesClient.Close()
		if err != nil {
			log.Printf("failed to close places client: %v", err)
		}
	}()

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "asim",
	Short: "Address Simulator (asim) lists some basic life information at a given address for the purposes of choosing a spot to move to.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := maps.New(cmd.Context())

		if err != nil {
			return fmt.Errorf("failed to create maps client: %w", err)
		}

		defer c.Close()

		coords, err := c.GetCoordinates(cmd.Context(), config.Address.Search, config.Address.City, config.Address.State, config.Address.Country)

		if err != nil {
			return fmt.Errorf("failed to get place ID: %w", err)
		}

		fmt.Println("Coordinates:", coords.Latitude, coords.Longitude)

		err = c.GetNearbyRestaurants(cmd.Context(), coords)

		if err != nil {
			return fmt.Errorf("failed to get nearby restaurants: %w", err)
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
