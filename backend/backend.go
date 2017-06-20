package backend

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/spike-force-1-bacon-evaluators/neo4bacon/api"
	"golang.org/x/net/context"
)

// Get calls backend to retrieve a new restaurant ranking
func Get() (*api.RestaurantList, error) {
	// Neo4bacon backend
	backendPort := "neo4bacon:50051"

	conn, err := grpc.Dial(backendPort, grpc.WithInsecure())
	if err != nil {
		return &api.RestaurantList{}, fmt.Errorf("could not connect to backend %s: %s", backendPort, err)
	}
	defer conn.Close()

	client := api.NewNeo4BaconClient(conn)

	restaurantList, err := client.List(context.Background(), &api.Empty{})
	if err != nil {
		return &api.RestaurantList{}, fmt.Errorf("could not retrieve restaurant's list: %s", err)
	}
	return restaurantList, nil
}
