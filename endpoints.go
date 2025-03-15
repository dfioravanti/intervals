package intervals

import (
	"time"
)

// endpointType is used to represent if an endpoint is plus or minus infinity or if it is bounded
// so for example in the interval ]-inf, 3] the left endpoint will be represented with the minusInf endpointType.
type endpointType string

var (
	finiteType        = endpointType("finite")
	plusInfiniteType  = endpointType("+inf")
	minusInfiniteType = endpointType("-inf")
)

var (
	minusInfinite = newMinusInfiniteEndpoint()
	plusInfinite  = newPlusInfiniteEndpoint()
)

// endpoint represents the endpoint of an open or closed interval.
// A endpoint can be either open, which means that the corresponding interval does not contain the endpoint,
// or closed, which means that the interval contains the endpoint.
// A endpoint can be finite or infinite
type endpoint struct {
	Value        time.Time    `json:"Value"`
	EndpointType endpointType `json:"Type"`
	Included     bool         `json:"Included"`
}

// newOpenEndpoint returns a new open, i.e. not included in the interval, endpoint
func newOpenEndpoint(v time.Time) endpoint {
	return endpoint{
		Value:        v,
		EndpointType: finiteType,
		Included:     false,
	}
}

// newCloseEndpoint returns a new open, i.e. included in the interval, endpoint
func newCloseEndpoint(v time.Time) endpoint {
	return endpoint{
		Value:        v,
		EndpointType: finiteType,
		Included:     true,
	}
}

// newMinusInfiniteEndpoint returns a minus infinite endpoint, i.e. an endpoint that is smaller than any other endpoints
func newMinusInfiniteEndpoint() endpoint {
	return endpoint{
		EndpointType: minusInfiniteType,
		Included:     false,
	}
}

// newPlusInfiniteEndpoint returns a plus infinite endpoint, i.e. an endpoint that is bigger than any other endpoints
func newPlusInfiniteEndpoint() endpoint {
	return endpoint{
		EndpointType: plusInfiniteType,
		Included:     false,
	}
}

// String returns the endpoint in string format
func (e endpoint) String() string {
	if e.EndpointType == minusInfiniteType {
		return "-inf"
	}

	if e.EndpointType == plusInfiniteType {
		return "+inf"
	}

	return e.Value.Format(time.RFC3339)
}

// IsFinite returns true if the endpoint is finite and false if it is infinite
func (e endpoint) IsFinite() bool {
	return e.EndpointType == finiteType
}

// IsOpen returns true if the endpoint should not be included in the interval that it belongs to, false otherwise
func (e endpoint) IsOpen() bool {
	return !e.Included
}

// IsClosed returns true if the endpoint should be included in the interval that it belongs to, false otherwise
func (e endpoint) IsClosed() bool {
	return e.Included
}

// Equal returns true if two endpoints are identical, false otherwise
func (e endpoint) Equal(e2 endpoint) bool {
	return e.Value == e2.Value && e.Included == e2.Included && e.EndpointType == e2.EndpointType
}

// Before returns true if and endpoint e is before the endpoint e2
// The MinusInfinite endpoint is before all the other endpoints except for itself
// The PlusInfinite endpoint is after all other endpoints
func (e endpoint) Before(e2 endpoint) bool {
	if e2.EndpointType == minusInfiniteType {
		return false
	}

	if e.EndpointType == minusInfiniteType {
		return true
	}

	if e.EndpointType == plusInfiniteType {
		return false
	}

	if e2.EndpointType == plusInfiniteType {
		return true
	}

	// if both are included then the case of e == e2 means that that e is not before e2
	if e.Included && e2.Included {
		return e.Value.Before(e2.Value) && !(e.Value.Equal(e2.Value))
	}

	// if e is included but e2 is not then it means that e == e2 means that e is before e2
	if e.Included && !e2.Included {
		return e.Value.Before(e2.Value) || e.Value.Equal(e2.Value)
	}

	// The other two cases are covered by this
	return e.Value.Before(e2.Value)
}

// After returns true if and endpoint e is after the endpoint e2
// The MinusInfinite endpoint is before all the other endpoints except for itself
// The PlusInfinite endpoint is after all other endpoints
func (e endpoint) After(e2 endpoint) bool {
	if e2.EndpointType == plusInfiniteType {
		return false
	}

	if e.EndpointType == plusInfiniteType {
		return true
	}

	if e2.EndpointType == minusInfiniteType {
		return true
	}

	if e.EndpointType == minusInfiniteType {
		return false
	}

	// if both are included then the case of e == e2 means that that e is not after e2
	if e.Included && e2.Included {
		return e.Value.After(e2.Value) && !(e.Value.Equal(e2.Value))
	}

	// if e is included but e2 is not then it means that e == e2 means that e is after e2
	if e.Included && !e2.Included {
		return e.Value.After(e2.Value) || e.Value.Equal(e2.Value)
	}

	// The other two cases are covered by this
	return e.Value.After(e2.Value)
}
