package usecase_mocks

import (
	domain "Task_Manager/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserUsecase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginUser provides a mock function with given fields: ctx, user
func (_m *UserUsecase) LoginUser(ctx context.Context, user domain.User) (int, error, string) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 int
	var r1 error
	var r2 string
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (int, error, string)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) int); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.User) string); ok {
		r2 = rf(ctx, user)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// RegisterUser provides a mock function with given fields: ctx, user
func (_m *UserUsecase) RegisterUser(ctx context.Context, user domain.User) (int, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (int, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) int); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}