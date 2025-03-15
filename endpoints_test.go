package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestString checks if the output of Endpoint.String() is what we expect
func TestStringTime(t *testing.T) {
	testCases := []struct {
		name string
		got  endpoint
		want string
	}{
		{
			name: "finite open time endpoint",
			got:  newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: "2022-11-02T01:02:03Z",
		},
		{
			name: "finite close time endpoint",
			got:  newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: "2022-11-02T01:02:03Z",
		},
		{
			name: "-inf time endpoint",
			got:  minusInfinite,
			want: "-inf",
		},
		{
			name: "+inf time endpoint",
			got:  plusInfinite,
			want: "+inf",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.got.String()
			assert.Equal(t, tc.want, result)
		})
	}
}

// TestIsFinite checks if the output of Endpoint.IsFinite() is what we expect
func TestIsFinite(t *testing.T) {
	testCases := []struct {
		name string
		got  endpoint
		want bool
	}{
		{
			name: "finite open time endpoint",
			got:  newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "finite close time endpoint",
			got:  newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "-inf time endpoint",
			got:  minusInfinite,
			want: false,
		},
		{
			name: "+inf time endpoint",
			got:  plusInfinite,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.got.IsFinite()
			assert.Equal(t, tc.want, result)
		})
	}
}

// TestIsOpen checks if the output of Endpoint.IsOpen() is what we expect
func TestIsOpen(t *testing.T) {
	testCases := []struct {
		name string
		got  endpoint
		want bool
	}{
		{
			name: "finite open time endpoint",
			got:  newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "finite close time endpoint",
			got:  newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "-inf time endpoint",
			got:  minusInfinite,
			want: true,
		},
		{
			name: "+inf time endpoint",
			got:  plusInfinite,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.got.IsOpen()
			assert.Equal(t, tc.want, result)
		})
	}
}

// TestIsClosed checks if the output of Endpoint.IsClosed() is what we expect
func TestIsClosed(t *testing.T) {
	testCases := []struct {
		name string
		got  endpoint
		want bool
	}{
		{
			name: "finite open time endpoint",
			got:  newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "finite close time endpoint",
			got:  newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "-inf time endpoint",
			got:  minusInfinite,
			want: false,
		},
		{
			name: "+inf time endpoint",
			got:  plusInfinite,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.got.IsClosed()
			assert.Equal(t, tc.want, result)
		})
	}
}

// TestEqual checks if the output of Endpoint.Equal() is what we expect
func TestEqual(t *testing.T) {
	testCases := []struct {
		name string
		i    endpoint
		j    endpoint
		want bool
	}{
		{
			name: "plus inf is equal to itself",
			i:    plusInfinite,
			j:    plusInfinite,
			want: true,
		},
		{
			name: "minus inf is equal to itself",
			i:    minusInfinite,
			j:    minusInfinite,
			want: true,
		},
		{
			name: "open endpoint is equal to itself",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "close endpoint is equal to itself",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "close endpoint is not equal to itself open",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "open endpoint is not equal to itself closed",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "different open endpoints are not equal",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-03T01:02:03Z")),
			want: false,
		},
		{
			name: "finite endpoint is not equal to plus infinite",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    plusInfinite,
			want: false,
		},
		{
			name: "finite endpoint is not equal to minus infinite",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    minusInfinite,
			want: false,
		},
		{
			name: "plus infinite not equal to minus infinite",
			i:    plusInfinite,
			j:    minusInfinite,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.i.Equal(tc.j)
			assert.Equal(t, tc.want, result)
		})
	}
}

// TestBefore checks if the output of Endpoint.Before() is what we expect
func TestBefore(t *testing.T) {
	testCases := []struct {
		name string
		i    endpoint
		j    endpoint
		want bool
	}{
		{
			name: "plus inf is not before to itself",
			i:    plusInfinite,
			j:    plusInfinite,
			want: false,
		},
		{
			name: "minus inf is not before to itself",
			i:    minusInfinite,
			j:    minusInfinite,
			want: false,
		},
		{
			name: "minus inf is before to plus inf",
			i:    minusInfinite,
			j:    plusInfinite,
			want: true,
		},
		{
			name: "plus inf is not before to minus inf",
			i:    plusInfinite,
			j:    minusInfinite,
			want: false,
		},
		{
			name: "mins inf is before a finite open endpoint",
			i:    minusInfinite,
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "mins inf is before a finite close endpoint",
			i:    minusInfinite,
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "mins inf is before a finite open endpoint",
			i:    minusInfinite,
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "mins inf is before a finite close endpoint",
			i:    minusInfinite,
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "mins inf is not before a finite open endpoint",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    minusInfinite,
			want: false,
		},
		{
			name: "mins inf is not before a finite close endpoint",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    minusInfinite,
			want: false,
		},
		{
			name: "plus inf is not before a finite open endpoint",
			i:    plusInfinite,
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "plus inf is not before a finite close endpoint",
			i:    plusInfinite,
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "a finite open endpoint is before plus inf",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    plusInfinite,
			want: true,
		},
		{
			name: "a finite close endpoint is before plus inf",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    plusInfinite,
			want: true,
		},
		{
			name: "a finite close endpoint not is before itself",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "a finite close endpoint is before itself open",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "a finite open endpoint is before a closed one with a different value",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:00:03Z")),
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.i.Before(tc.j)
			assert.Equal(t, tc.want, result)
		})
	}
}

// TestAfter checks if the output of Endpoint.After() is what we expect
func TestAfter(t *testing.T) {
	testCases := []struct {
		name string
		i    endpoint
		j    endpoint
		want bool
	}{
		{
			name: "plus inf is after itself",
			i:    plusInfinite,
			j:    plusInfinite,
			want: false,
		},
		{
			name: "minus inf is not after itself",
			i:    minusInfinite,
			j:    minusInfinite,
			want: true,
		},
		{
			name: "minus inf is not after plus inf",
			i:    minusInfinite,
			j:    plusInfinite,
			want: false,
		},
		{
			name: "plus inf is after minus inf",
			i:    plusInfinite,
			j:    minusInfinite,
			want: true,
		},
		{
			name: "mins inf is not after a finite open endpoint",
			i:    minusInfinite,
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "mins inf is not after a finite close endpoint",
			i:    minusInfinite,
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "mins inf is not a finite open endpoint",
			i:    minusInfinite,
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "mins inf is before a finite close endpoint",
			i:    minusInfinite,
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "mins inf is after a finite open endpoint",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    minusInfinite,
			want: true,
		},
		{
			name: "mins inf is after a finite close endpoint",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    minusInfinite,
			want: true,
		},
		{
			name: "plus inf is after a finite open endpoint",
			i:    plusInfinite,
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "plus inf is after a finite close endpoint",
			i:    plusInfinite,
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "a finite open endpoint is not after plus inf",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    plusInfinite,
			want: false,
		},
		{
			name: "a finite close endpoint is not after plus inf",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    plusInfinite,
			want: false,
		},
		{
			name: "a finite close endpoint is not after itself",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: false,
		},
		{
			name: "a finite close endpoint is after itself open",
			i:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			j:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
		{
			name: "a finite open endpoint is after a closed one with a different value",
			i:    newOpenEndpoint(mustParseRFC3339("2022-11-02T01:03:03Z")),
			j:    newCloseEndpoint(mustParseRFC3339("2022-11-02T01:02:03Z")),
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.i.After(tc.j)
			assert.Equal(t, tc.want, result)
		})
	}
}
