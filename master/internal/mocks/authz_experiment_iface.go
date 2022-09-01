// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	bun "github.com/uptrace/bun"

	mock "github.com/stretchr/testify/mock"

	model "github.com/determined-ai/determined/master/pkg/model"

	projectv1 "github.com/determined-ai/determined/proto/pkg/projectv1"
)

// ExperimentAuthZ is an autogenerated mock type for the ExperimentAuthZ type
type ExperimentAuthZ struct {
	mock.Mock
}

// CanCreateExperiment provides a mock function with given fields: curUser, proj, e
func (_m *ExperimentAuthZ) CanCreateExperiment(curUser model.User, proj *projectv1.Project, e *model.Experiment) error {
	ret := _m.Called(curUser, proj, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *projectv1.Project, *model.Experiment) error); ok {
		r0 = rf(curUser, proj, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanDeleteExperiment provides a mock function with given fields: curUser, e
func (_m *ExperimentAuthZ) CanDeleteExperiment(curUser model.User, e *model.Experiment) error {
	ret := _m.Called(curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment) error); ok {
		r0 = rf(curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanEditExperiment provides a mock function with given fields: curUser, e
func (_m *ExperimentAuthZ) CanEditExperiment(curUser model.User, e *model.Experiment) error {
	ret := _m.Called(curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment) error); ok {
		r0 = rf(curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanEditExperimentsMetadata provides a mock function with given fields: curUser, e
func (_m *ExperimentAuthZ) CanEditExperimentsMetadata(curUser model.User, e *model.Experiment) error {
	ret := _m.Called(curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment) error); ok {
		r0 = rf(curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanForkFromExperiment provides a mock function with given fields: curUser, e
func (_m *ExperimentAuthZ) CanForkFromExperiment(curUser model.User, e *model.Experiment) error {
	ret := _m.Called(curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment) error); ok {
		r0 = rf(curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanGetExperiment provides a mock function with given fields: curUser, e
func (_m *ExperimentAuthZ) CanGetExperiment(curUser model.User, e *model.Experiment) (bool, error) {
	ret := _m.Called(curUser, e)

	var r0 bool
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment) bool); ok {
		r0 = rf(curUser, e)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.User, *model.Experiment) error); ok {
		r1 = rf(curUser, e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CanGetExperimentArtifacts provides a mock function with given fields: curUser, e
func (_m *ExperimentAuthZ) CanGetExperimentArtifacts(curUser model.User, e *model.Experiment) error {
	ret := _m.Called(curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment) error); ok {
		r0 = rf(curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanPreviewHPSearch provides a mock function with given fields: curUser
func (_m *ExperimentAuthZ) CanPreviewHPSearch(curUser model.User) error {
	ret := _m.Called(curUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User) error); ok {
		r0 = rf(curUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsCheckpointGCPolicy provides a mock function with given fields: curUser, e
func (_m *ExperimentAuthZ) CanSetExperimentsCheckpointGCPolicy(curUser model.User, e *model.Experiment) error {
	ret := _m.Called(curUser, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment) error); ok {
		r0 = rf(curUser, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsMaxSlots provides a mock function with given fields: curUser, e, slots
func (_m *ExperimentAuthZ) CanSetExperimentsMaxSlots(curUser model.User, e *model.Experiment, slots int) error {
	ret := _m.Called(curUser, e, slots)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment, int) error); ok {
		r0 = rf(curUser, e, slots)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsPriority provides a mock function with given fields: curUser, e, priority
func (_m *ExperimentAuthZ) CanSetExperimentsPriority(curUser model.User, e *model.Experiment, priority int) error {
	ret := _m.Called(curUser, e, priority)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment, int) error); ok {
		r0 = rf(curUser, e, priority)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetExperimentsWeight provides a mock function with given fields: curUser, e, weight
func (_m *ExperimentAuthZ) CanSetExperimentsWeight(curUser model.User, e *model.Experiment, weight float64) error {
	ret := _m.Called(curUser, e, weight)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User, *model.Experiment, float64) error); ok {
		r0 = rf(curUser, e, weight)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterExperimentLabelsQuery provides a mock function with given fields: curUser, proj, query
func (_m *ExperimentAuthZ) FilterExperimentLabelsQuery(curUser model.User, proj *projectv1.Project, query *bun.SelectQuery) (*bun.SelectQuery, error) {
	ret := _m.Called(curUser, proj, query)

	var r0 *bun.SelectQuery
	if rf, ok := ret.Get(0).(func(model.User, *projectv1.Project, *bun.SelectQuery) *bun.SelectQuery); ok {
		r0 = rf(curUser, proj, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bun.SelectQuery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.User, *projectv1.Project, *bun.SelectQuery) error); ok {
		r1 = rf(curUser, proj, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterExperimentsQuery provides a mock function with given fields: curUser, proj, query
func (_m *ExperimentAuthZ) FilterExperimentsQuery(curUser model.User, proj *projectv1.Project, query *bun.SelectQuery) (*bun.SelectQuery, error) {
	ret := _m.Called(curUser, proj, query)

	var r0 *bun.SelectQuery
	if rf, ok := ret.Get(0).(func(model.User, *projectv1.Project, *bun.SelectQuery) *bun.SelectQuery); ok {
		r0 = rf(curUser, proj, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bun.SelectQuery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.User, *projectv1.Project, *bun.SelectQuery) error); ok {
		r1 = rf(curUser, proj, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExperimentAuthZ interface {
	mock.TestingT
	Cleanup(func())
}

// NewExperimentAuthZ creates a new instance of ExperimentAuthZ. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExperimentAuthZ(t mockConstructorTestingTNewExperimentAuthZ) *ExperimentAuthZ {
	mock := &ExperimentAuthZ{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
