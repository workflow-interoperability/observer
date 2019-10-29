package controller

import (
	"context"

	observer "github.com/workflow-interoperability/observer/grpc"
)

// ObserverServer implements Observer grpc service
type ObserverServer int

// GetProperties get properties of current server
func (o ObserverServer) GetProperties(ctx context.Context, in *observer.GetPropertiesRq) (*observer.GetPropertiesRs, error) {
	return &observer.GetPropertiesRs {
		Key: Key,
	}, nil
}

// StateChanged receive state change request from instance
func (o ObserverServer) StateChanged(ctx context.Context, in *observer.StateChangedRq) (*observer.StateChangedRs, error) {
	
}
