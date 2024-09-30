// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	models "MentalHealthCare/models"

	mock "github.com/stretchr/testify/mock"
)

// ComplaintRepository is an autogenerated mock type for the ComplaintRepository type
type ComplaintRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: complaintReq
func (_m *ComplaintRepository) Create(complaintReq models.ComplaintRequest) (models.Complaint, error) {
	ret := _m.Called(complaintReq)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 models.Complaint
	var r1 error
	if rf, ok := ret.Get(0).(func(models.ComplaintRequest) (models.Complaint, error)); ok {
		return rf(complaintReq)
	}
	if rf, ok := ret.Get(0).(func(models.ComplaintRequest) models.Complaint); ok {
		r0 = rf(complaintReq)
	} else {
		r0 = ret.Get(0).(models.Complaint)
	}

	if rf, ok := ret.Get(1).(func(models.ComplaintRequest) error); ok {
		r1 = rf(complaintReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ComplaintRepository) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ComplaintRepository) GetAll() ([]models.Complaint, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.Complaint
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Complaint, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Complaint); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Complaint)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *ComplaintRepository) GetByID(id string) (models.Complaint, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 models.Complaint
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Complaint, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Complaint); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Complaint)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: complaintReq, id
func (_m *ComplaintRepository) Update(complaintReq models.ComplaintRequest, id string) (models.Complaint, error) {
	ret := _m.Called(complaintReq, id)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.Complaint
	var r1 error
	if rf, ok := ret.Get(0).(func(models.ComplaintRequest, string) (models.Complaint, error)); ok {
		return rf(complaintReq, id)
	}
	if rf, ok := ret.Get(0).(func(models.ComplaintRequest, string) models.Complaint); ok {
		r0 = rf(complaintReq, id)
	} else {
		r0 = ret.Get(0).(models.Complaint)
	}

	if rf, ok := ret.Get(1).(func(models.ComplaintRequest, string) error); ok {
		r1 = rf(complaintReq, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func NewComplaintRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ComplaintRepository {
	mock := &ComplaintRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}