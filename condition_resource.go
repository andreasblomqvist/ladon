package ladon

import (
	"errors"
	"strings"
)

// ResourceCondition is a condition which is fulfilled if the request's subject is equal to the given value string
type ResourceCondition struct{}

// Fulfills returns true if the request's resouce contains the given value string
func (c *ResourceCondition) Fulfills(value interface{}, r *Request) bool {

	operator, ok := r.Context["resourceFilterRule"]
	if !ok {
		//Default to equal
		operator = "equals"
	}

	filter, ok := r.Context["resourceFilter"]
	if !ok {
		//Default to equal
		panic(errors.New("missing resourceFilter"))
	}

	filterArr := strings.Split(filter.(string), ":")
	filterKey := filterArr[0]
	filterValue := filterArr[1] // TODO: SPLIT ON , or delimiter

	// Split the resource string based on ":"
	// take the value after the value matching the "resourceFilter"
	resourceArr := strings.Split(r.Resource, ":")

	var resourceValue *string
	for i := range resourceArr {
		if resourceArr[i] == filterKey {
			resourceValue = &resourceArr[i+1] // TODO: SPLIT ON , or delimiter
		}
	}
	if resourceValue == nil {
		return false
	}

	switch operator {
	case "equal":
		return (strings.Compare(*resourceValue, filterValue) == 0)
	case "contains":
		return (strings.Contains(*resourceValue, filterValue))

	}

	return false
}

// GetName returns the condition's name.
func (c *ResourceCondition) GetName() string {
	return "ResourceCondition"
}
