// Code generated by "requestgen -method GET -url /futures/data/globalLongShortAccountRatio -type FuturesGlobalLongShortAccountRatioRequest -responseType []FuturesGlobalLongShortAccountRatio"; DO NOT EDIT.

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

func (f *FuturesGlobalLongShortAccountRatioRequest) Symbol(symbol string) *FuturesGlobalLongShortAccountRatioRequest {
	f.symbol = symbol
	return f
}

func (f *FuturesGlobalLongShortAccountRatioRequest) Period(period types.Interval) *FuturesGlobalLongShortAccountRatioRequest {
	f.period = period
	return f
}

func (f *FuturesGlobalLongShortAccountRatioRequest) Limit(limit uint64) *FuturesGlobalLongShortAccountRatioRequest {
	f.limit = &limit
	return f
}

func (f *FuturesGlobalLongShortAccountRatioRequest) StartTime(startTime time.Time) *FuturesGlobalLongShortAccountRatioRequest {
	f.startTime = &startTime
	return f
}

func (f *FuturesGlobalLongShortAccountRatioRequest) EndTime(endTime time.Time) *FuturesGlobalLongShortAccountRatioRequest {
	f.endTime = &endTime
	return f
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (f *FuturesGlobalLongShortAccountRatioRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (f *FuturesGlobalLongShortAccountRatioRequest) GetParameters() (map[string]interface{}, error) {
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
func (f *FuturesGlobalLongShortAccountRatioRequest) GetParametersQuery() (url.Values, error) {
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
func (f *FuturesGlobalLongShortAccountRatioRequest) GetParametersJSON() ([]byte, error) {
	params, err := f.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (f *FuturesGlobalLongShortAccountRatioRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (f *FuturesGlobalLongShortAccountRatioRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (f *FuturesGlobalLongShortAccountRatioRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (f *FuturesGlobalLongShortAccountRatioRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (f *FuturesGlobalLongShortAccountRatioRequest) GetSlugsMap() (map[string]string, error) {
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

func (f *FuturesGlobalLongShortAccountRatioRequest) Do(ctx context.Context) ([]FuturesGlobalLongShortAccountRatio, error) {

	// empty params for GET operation
	var params interface{}
	query, err := f.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/futures/data/globalLongShortAccountRatio"

	req, err := f.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := f.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []FuturesGlobalLongShortAccountRatio
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
