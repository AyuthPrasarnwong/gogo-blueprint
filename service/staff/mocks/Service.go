// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import domainstaff "github.com/devit-tel/gogo-blueprint/domain/staff"
import goerror "github.com/devit-tel/goerror"
import mock "github.com/stretchr/testify/mock"
import staff "github.com/devit-tel/gogo-blueprint/service/staff"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateStaff provides a mock function with given fields: ctx, input
func (_m *Service) CreateStaff(ctx context.Context, input *staff.CreateStaffInput) (*domainstaff.Staff, goerror.Error) {
	ret := _m.Called(ctx, input)

	var r0 *domainstaff.Staff
	if rf, ok := ret.Get(0).(func(context.Context, *staff.CreateStaffInput) *domainstaff.Staff); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domainstaff.Staff)
		}
	}

	var r1 goerror.Error
	if rf, ok := ret.Get(1).(func(context.Context, *staff.CreateStaffInput) goerror.Error); ok {
		r1 = rf(ctx, input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(goerror.Error)
		}
	}

	return r0, r1
}

// GetStaffsByCompany provides a mock function with given fields: ctx, input
func (_m *Service) GetStaffsByCompany(ctx context.Context, input *staff.GetStaffsByCompanyInput) ([]*domainstaff.Staff, goerror.Error) {
	ret := _m.Called(ctx, input)

	var r0 []*domainstaff.Staff
	if rf, ok := ret.Get(0).(func(context.Context, *staff.GetStaffsByCompanyInput) []*domainstaff.Staff); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domainstaff.Staff)
		}
	}

	var r1 goerror.Error
	if rf, ok := ret.Get(1).(func(context.Context, *staff.GetStaffsByCompanyInput) goerror.Error); ok {
		r1 = rf(ctx, input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(goerror.Error)
		}
	}

	return r0, r1
}

// UpdateStaff provides a mock function with given fields: ctx, input
func (_m *Service) UpdateStaff(ctx context.Context, input *staff.UpdateStaffInput) (*domainstaff.Staff, goerror.Error) {
	ret := _m.Called(ctx, input)

	var r0 *domainstaff.Staff
	if rf, ok := ret.Get(0).(func(context.Context, *staff.UpdateStaffInput) *domainstaff.Staff); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domainstaff.Staff)
		}
	}

	var r1 goerror.Error
	if rf, ok := ret.Get(1).(func(context.Context, *staff.UpdateStaffInput) goerror.Error); ok {
		r1 = rf(ctx, input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(goerror.Error)
		}
	}

	return r0, r1
}