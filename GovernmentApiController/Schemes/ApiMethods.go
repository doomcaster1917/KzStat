package Schemes

import (
	"KZstat/GovernmentApiController/Tools"
	"fmt"
	"net/url"
	"reflect"
)

var baseURL = "http://taldau.stat.gov.kz/ru/Api/"

func GetIndexTreeDataMethod(resource string, reqParams interface{}) string {
	params := url.Values{}

	v := reflect.ValueOf(reqParams)
	typeOfS := v.Type()
	length := v.NumField()
	for i := 0; i < length; i++ {
		name := typeOfS.Field(i).Name
		params.Add(name, Tools.FieldStrConverter(v.Field(i)))
	}

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	return u.String()
}

func (GetIndexTreeData) GetSegmentListMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("%d")
	return requestText
}
func (GetIndexTreeData) GetPeriodListMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
func (GetIndexTreeData) GetIndexPeriodsMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
func (GetIndexTreeData) GetIndexAttributesMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
func (GetIndexTreeData) IndexFeedBranchMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
func (GetIndexTreeData) IndexFeedRegionMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
func (GetIndexTreeData) GetIndustryPeriodsMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
func (GetIndexTreeData) GetRegionPeriodsMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
func (GetIndexTreeData) SearchMethod(GetIndexTreeData) string {
	requestText := fmt.Sprintf("")
	return requestText
}
