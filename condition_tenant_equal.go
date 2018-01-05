package ladon

import (
	"net/url"
)

// EqualTenantCondition is a condition which is fulfilled if the request's subject is equal to the given value string
type EqualTenantCondition struct{}

// Fulfills returns true if the request's resouce contains the given value string
func (c *EqualTenantCondition) Fulfills(value interface{}, r *Request) bool {
	s, ok := value.(string)

	//[^&?]*?=[^&?]*

	//retrives all the params from the url
	u, err := url.Parse(r.Resource)
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	tenantIds := m["tenantIds"][0]

	return ok && s == tenantIds
}

// GetName returns the condition's name.
func (c *EqualTenantCondition) GetName() string {
	return "EqualTenantCondition"
}
