// Copyright (c) 2021 Cisco Systems, Inc. and its affiliates
// All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DatasetsApiController binds http requests to an api service and writes the service results to the http response
type DatasetsApiController struct {
	service DatasetsApiServicer
}

// NewDatasetsApiController creates a default api controller
func NewDatasetsApiController(s DatasetsApiServicer) Router {
	return &DatasetsApiController{service: s}
}

// Routes returns all of the api route for the DatasetsApiController
func (c *DatasetsApiController) Routes() Routes {
	return Routes{
		{
			"CreateDataset",
			strings.ToUpper("Post"),
			"/{user}/datasets",
			c.CreateDataset,
		},
		{
			"GetAllDatasets",
			strings.ToUpper("Get"),
			"/datasets",
			c.GetAllDatasets,
		},
		{
			"GetDataset",
			strings.ToUpper("Get"),
			"/{user}/datasets/{datasetId}",
			c.GetDataset,
		},
		{
			"GetDatasets",
			strings.ToUpper("Get"),
			"/{user}/datasets",
			c.GetDatasets,
		},
		{
			"UpdateDataset",
			strings.ToUpper("Put"),
			"/{user}/datasets/{datasetId}",
			c.UpdateDataset,
		},
	}
}

// CreateDataset - Create meta info for a new dataset.
func (c *DatasetsApiController) CreateDataset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	datasetInfo := &DatasetInfo{}
	if err := json.NewDecoder(r.Body).Decode(&datasetInfo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.CreateDataset(r.Context(), user, *datasetInfo)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAllDatasets - Get the meta info on all the datasets
func (c *DatasetsApiController) GetAllDatasets(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.GetAllDatasets(r.Context(), limit)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetDataset - Get dataset meta information
func (c *DatasetsApiController) GetDataset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	datasetId := params["datasetId"]

	result, err := c.service.GetDataset(r.Context(), user, datasetId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetDatasets - Get the meta info on all the datasets owned by user
func (c *DatasetsApiController) GetDatasets(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	user := params["user"]

	limit, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.GetDatasets(r.Context(), user, limit)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateDataset - Update meta info for a given dataset
func (c *DatasetsApiController) UpdateDataset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	datasetId := params["datasetId"]

	datasetInfo := &DatasetInfo{}
	if err := json.NewDecoder(r.Body).Decode(&datasetInfo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.UpdateDataset(r.Context(), user, datasetId, *datasetInfo)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
