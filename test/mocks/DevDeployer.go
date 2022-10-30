// Copyright 2022 Giuseppe De Palma, Matteo Trentin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DevDeployer is an autogenerated mock type for the DevDeployer type
type DevDeployer struct {
	mock.Mock
}

// CreateFLNetwork provides a mock function with given fields: ctx
func (_m *DevDeployer) CreateFLNetwork(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PullCoreImage provides a mock function with given fields: ctx
func (_m *DevDeployer) PullCoreImage(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PullPromImage provides a mock function with given fields: ctx
func (_m *DevDeployer) PullPromImage(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PullWorkerImage provides a mock function with given fields: ctx
func (_m *DevDeployer) PullWorkerImage(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveCoreContainer provides a mock function with given fields: ctx
func (_m *DevDeployer) RemoveCoreContainer(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveFLNetwork provides a mock function with given fields: ctx
func (_m *DevDeployer) RemoveFLNetwork(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePromContainer provides a mock function with given fields: ctx
func (_m *DevDeployer) RemovePromContainer(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveWorkerContainer provides a mock function with given fields: ctx
func (_m *DevDeployer) RemoveWorkerContainer(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Setup provides a mock function with given fields: ctx, coreImg, workerImg
func (_m *DevDeployer) Setup(ctx context.Context, coreImg string, workerImg string) error {
	ret := _m.Called(ctx, coreImg, workerImg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, coreImg, workerImg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartCore provides a mock function with given fields: ctx
func (_m *DevDeployer) StartCore(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartProm provides a mock function with given fields: ctx
func (_m *DevDeployer) StartProm(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartWorker provides a mock function with given fields: ctx
func (_m *DevDeployer) StartWorker(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDevDeployer interface {
	mock.TestingT
	Cleanup(func())
}

// NewDevDeployer creates a new instance of DevDeployer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDevDeployer(t mockConstructorTestingTNewDevDeployer) *DevDeployer {
	mock := &DevDeployer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
