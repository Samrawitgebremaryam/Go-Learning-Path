// Code generated by mockery v2.44.1. DO NOT EDIT.

package inframocks

import (
	jwt "github.com/dgrijalva/jwt-go"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// SecurityService is an autogenerated mock type for the SecurityService type
type SecurityService struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: hash, password
func (_m *SecurityService) ComparePassword(hash string, password string) bool {
	ret := _m.Called(hash, password)

	if len(ret) == 0 {
		panic("no return value specified for ComparePassword")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hash, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateToken provides a mock function with given fields: id, email, User_type
func (_m *SecurityService) CreateToken(id primitive.ObjectID, email string, User_type string) (string, error) {
	ret := _m.Called(id, email, User_type)

	if len(ret) == 0 {
		panic("no return value specified for CreateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, string, string) (string, error)); ok {
		return rf(id, email, User_type)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, string, string) string); ok {
		r0 = rf(id, email, User_type)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID, string, string) error); ok {
		r1 = rf(id, email, User_type)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashPassword provides a mock function with given fields: password
func (_m *SecurityService) HashPassword(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: tokenstr
func (_m *SecurityService) ValidateToken(tokenstr string) (*jwt.Token, error) {
	ret := _m.Called(tokenstr)

	if len(ret) == 0 {
		panic("no return value specified for ValidateToken")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(tokenstr)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(tokenstr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenstr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSecurityService creates a new instance of SecurityService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSecurityService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SecurityService {
	mock := &SecurityService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
