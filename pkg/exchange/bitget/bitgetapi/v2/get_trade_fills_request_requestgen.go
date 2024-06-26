// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Data -url /api/v2/spot/trade/fills -type GetTradeFillsRequest -responseDataType []Trade"; DO NOT EDIT.

package bitgetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/puper/bbgo/pkg/exchange/bitget/bitgetapi"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (s *GetTradeFillsRequest) Symbol(symbol string) *GetTradeFillsRequest {
	s.symbol = symbol
	return s
}

func (s *GetTradeFillsRequest) Limit(limit string) *GetTradeFillsRequest {
	s.limit = &limit
	return s
}

func (s *GetTradeFillsRequest) IdLessThan(idLessThan string) *GetTradeFillsRequest {
	s.idLessThan = &idLessThan
	return s
}

func (s *GetTradeFillsRequest) StartTime(startTime time.Time) *GetTradeFillsRequest {
	s.startTime = &startTime
	return s
}

func (s *GetTradeFillsRequest) EndTime(endTime time.Time) *GetTradeFillsRequest {
	s.endTime = &endTime
	return s
}

func (s *GetTradeFillsRequest) OrderId(orderId string) *GetTradeFillsRequest {
	s.orderId = &orderId
	return s
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (s *GetTradeFillsRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}
	// check symbol field -> json key symbol
	symbol := s.symbol

	// assign parameter of symbol
	params["symbol"] = symbol
	// check limit field -> json key limit
	if s.limit != nil {
		limit := *s.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}
	// check idLessThan field -> json key idLessThan
	if s.idLessThan != nil {
		idLessThan := *s.idLessThan

		// assign parameter of idLessThan
		params["idLessThan"] = idLessThan
	} else {
	}
	// check startTime field -> json key startTime
	if s.startTime != nil {
		startTime := *s.startTime

		// assign parameter of startTime
		// convert time.Time to milliseconds time stamp
		params["startTime"] = strconv.FormatInt(startTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check endTime field -> json key endTime
	if s.endTime != nil {
		endTime := *s.endTime

		// assign parameter of endTime
		// convert time.Time to milliseconds time stamp
		params["endTime"] = strconv.FormatInt(endTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check orderId field -> json key orderId
	if s.orderId != nil {
		orderId := *s.orderId

		// assign parameter of orderId
		params["orderId"] = orderId
	} else {
	}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (s *GetTradeFillsRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (s *GetTradeFillsRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := s.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if s.isVarSlice(_v) {
			s.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (s *GetTradeFillsRequest) GetParametersJSON() ([]byte, error) {
	params, err := s.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (s *GetTradeFillsRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (s *GetTradeFillsRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (s *GetTradeFillsRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (s *GetTradeFillsRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (s *GetTradeFillsRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := s.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

// GetPath returns the request path of the API
func (s *GetTradeFillsRequest) GetPath() string {
	return "/api/v2/spot/trade/fills"
}

// Do generates the request object and send the request object to the API endpoint
func (s *GetTradeFillsRequest) Do(ctx context.Context) ([]Trade, error) {

	// no body params
	var params interface{}
	query, err := s.GetQueryParameters()
	if err != nil {
		return nil, err
	}

	var apiURL string

	apiURL = s.GetPath()

	req, err := s.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := s.client.SendRequest(req)
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
	var data []Trade
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return data, nil
}
