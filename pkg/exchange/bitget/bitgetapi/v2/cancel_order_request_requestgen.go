// Code generated by "requestgen -method POST -responseType .APIResponse -responseDataField Data -url /api/v2/spot/trade/cancel-order -type CancelOrderRequest -responseDataType .CancelOrder"; DO NOT EDIT.

package bitgetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/puper/bbgo/pkg/exchange/bitget/bitgetapi"
	"net/url"
	"reflect"
	"regexp"
)

func (c *CancelOrderRequest) Symbol(symbol string) *CancelOrderRequest {
	c.symbol = symbol
	return c
}

func (c *CancelOrderRequest) OrderId(orderId string) *CancelOrderRequest {
	c.orderId = &orderId
	return c
}

func (c *CancelOrderRequest) ClientOrderId(clientOrderId string) *CancelOrderRequest {
	c.clientOrderId = &clientOrderId
	return c
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (c *CancelOrderRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (c *CancelOrderRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check symbol field -> json key symbol
	symbol := c.symbol

	// assign parameter of symbol
	params["symbol"] = symbol
	// check orderId field -> json key orderId
	if c.orderId != nil {
		orderId := *c.orderId

		// assign parameter of orderId
		params["orderId"] = orderId
	} else {
	}
	// check clientOrderId field -> json key clientOid
	if c.clientOrderId != nil {
		clientOrderId := *c.clientOrderId

		// assign parameter of clientOrderId
		params["clientOid"] = clientOrderId
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (c *CancelOrderRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := c.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if c.isVarSlice(_v) {
			c.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (c *CancelOrderRequest) GetParametersJSON() ([]byte, error) {
	params, err := c.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (c *CancelOrderRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (c *CancelOrderRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (c *CancelOrderRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (c *CancelOrderRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (c *CancelOrderRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := c.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

// GetPath returns the request path of the API
func (c *CancelOrderRequest) GetPath() string {
	return "/api/v2/spot/trade/cancel-order"
}

// Do generates the request object and send the request object to the API endpoint
func (c *CancelOrderRequest) Do(ctx context.Context) (*CancelOrder, error) {

	params, err := c.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	var apiURL string

	apiURL = c.GetPath()

	req, err := c.client.NewAuthenticatedRequest(ctx, "POST", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := c.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse bitgetapi.APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}

	type responseValidator interface {
		Validate() error
	}
	validator, ok := interface{}(apiResponse).(responseValidator)
	if ok {
		if err := validator.Validate(); err != nil {
			return nil, err
		}
	}
	var data CancelOrder
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
