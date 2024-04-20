// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Data -url /api/spot/v1/trade/fills -type GetFillsRequest -responseDataType .ServerTime"; DO NOT EDIT.

package bitgetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/puper/bbgo/pkg/types"
	"net/url"
	"reflect"
	"regexp"
)

func (g *GetFillsRequest) After(after string) *GetFillsRequest {
	g.after = &after
	return g
}

func (g *GetFillsRequest) Before(before string) *GetFillsRequest {
	g.before = &before
	return g
}

func (g *GetFillsRequest) Limit(limit string) *GetFillsRequest {
	g.limit = &limit
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetFillsRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetFillsRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check after field -> json key after
	if g.after != nil {
		after := *g.after

		// assign parameter of after
		params["after"] = after
	} else {
	}
	// check before field -> json key before
	if g.before != nil {
		before := *g.before

		// assign parameter of before
		params["before"] = before
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetFillsRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetFillsRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetFillsRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetFillsRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetFillsRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetFillsRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetFillsRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (g *GetFillsRequest) Do(ctx context.Context) (*types.MillisecondTimestamp, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/api/spot/v1/trade/fills"

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data types.MillisecondTimestamp
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
