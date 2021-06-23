/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"wwwin-github.cisco.com/fledge/fledge/cmd/controller"
)

// DesignsApiService is a service that implents the logic for the DesignsApiServicer
// This service should implement the business logic for every endpoint for the DesignsApi API.
// Include any external packages or services that will be required by this service.
type DesignsApiService struct {
}

// NewDesignsApiService creates a default api service
func NewDesignsApiService() DesignsApiServicer {
	return &DesignsApiService{}
}

// GetDesigns - Get list of all the designs created by the user.
func (s *DesignsApiService) GetDesigns(ctx context.Context, user string, limit int32) (ImplResponse, error) {
	zap.S().Debugf("get list of designs for user:%s | limit:%d", user, limit)

	designList, err := controller.GetDesigns(user, limit)

	if err != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("get list of designs request failed")
	} else {
		return Response(http.StatusOK, designList), nil
	}
}
