// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DirectMessager is an autogenerated mock type for the DirectMessager type
type DirectMessager struct {
	mock.Mock
}

// DirectMessage provides a mock function with given fields: recipientID, content
func (_m *DirectMessager) DirectMessage(recipientID string, content string) error {
	ret := _m.Called(recipientID, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(recipientID, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}