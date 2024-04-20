// Code generated by "requestgen -method GET -url /futures/data/topLongShortAccountRatio -type FuturesTopTraderLongShortAccountRatioRequest -responseType []FuturesTopTraderLongShortAccountRatio"; DO NOT EDIT.

package binanceapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/puper/bbgo/pkg/types"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (f *FuturesTopTraderLongShortAccountRatioRequest) Symbol(symbol string) *FuturesTopTraderLongShortAccountRatioRequest {
	f.symbol = symbol
	return f
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) Period(period types.Interval) *FuturesTopTraderLongShortAccountRatioRequest {
	f.period = period
	return f
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) Limit(limit uint64) *FuturesTopTraderLongShortAccountRatioRequest {
	f.limit = &limit
	return f
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) StartTime(startTime time.Time) *FuturesTopTraderLongShortAccountRatioRequest {
	f.startTime = &startTime
	return f
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) EndTime(endTime time.Time) *FuturesTopTraderLongShortAccountRatioRequest {
	f.endTime = &endTime
	return f
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (f *FuturesTopTraderLongShortAccountRatioRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (f *FuturesTopTraderLongShortAccountRatioRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check symbol field -> json key symbol
	symbol := f.symbol

	// assign parameter of symbol
	params["symbol"] = symbol
	// check period field -> json key period
	period := f.period

	// assign parameter of period
	params["period"] = period
	// check limit field -> json key limit
	if f.limit != nil {
		limit := *f.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}
	// check startTime field -> json key startTime
	if f.startTime != nil {
		startTime := *f.startTime

		// assign parameter of startTime
		// convert time.Time to milliseconds time stamp
		params["startTime"] = strconv.FormatInt(startTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check endTime field -> json key endTime
	if f.endTime != nil {
		endTime := *f.endTime

		// assign parameter of endTime
		// convert time.Time to milliseconds time stamp
		params["endTime"] = strconv.FormatInt(endTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (f *FuturesTopTraderLongShortAccountRatioRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := f.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if f.isVarSlice(_v) {
			f.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (f *FuturesTopTraderLongShortAccountRatioRequest) GetParametersJSON() ([]byte, error) {
	params, err := f.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (f *FuturesTopTraderLongShortAccountRatioRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := f.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (f *FuturesTopTraderLongShortAccountRatioRequest) Do(ctx context.Context) ([]FuturesTopTraderLongShortAccountRatio, error) {

	// empty params for GET operation
	var params interface{}
	query, err := f.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/futures/data/topLongShortAccountRatio"

	req, err := f.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := f.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []FuturesTopTraderLongShortAccountRatio
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
