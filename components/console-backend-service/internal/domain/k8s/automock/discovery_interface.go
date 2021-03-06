// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import openapi_v2 "github.com/googleapis/gnostic/OpenAPIv2"
import rest "k8s.io/client-go/rest"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import version "k8s.io/apimachinery/pkg/version"

// DiscoveryInterface is an autogenerated mock type for the DiscoveryInterface type
type DiscoveryInterface struct {
	mock.Mock
}

// OpenAPISchema provides a mock function with given fields:
func (_m *DiscoveryInterface) OpenAPISchema() (*openapi_v2.Document, error) {
	ret := _m.Called()

	var r0 *openapi_v2.Document
	if rf, ok := ret.Get(0).(func() *openapi_v2.Document); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi_v2.Document)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RESTClient provides a mock function with given fields:
func (_m *DiscoveryInterface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// ServerGroups provides a mock function with given fields:
func (_m *DiscoveryInterface) ServerGroups() (*v1.APIGroupList, error) {
	ret := _m.Called()

	var r0 *v1.APIGroupList
	if rf, ok := ret.Get(0).(func() *v1.APIGroupList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.APIGroupList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerGroupsAndResources provides a mock function with given fields:
func (_m *DiscoveryInterface) ServerGroupsAndResources() ([]*v1.APIGroup, []*v1.APIResourceList, error) {
	ret := _m.Called()

	var r0 []*v1.APIGroup
	if rf, ok := ret.Get(0).(func() []*v1.APIGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.APIGroup)
		}
	}

	var r1 []*v1.APIResourceList
	if rf, ok := ret.Get(1).(func() []*v1.APIResourceList); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*v1.APIResourceList)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ServerPreferredNamespacedResources provides a mock function with given fields:
func (_m *DiscoveryInterface) ServerPreferredNamespacedResources() ([]*v1.APIResourceList, error) {
	ret := _m.Called()

	var r0 []*v1.APIResourceList
	if rf, ok := ret.Get(0).(func() []*v1.APIResourceList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.APIResourceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerPreferredResources provides a mock function with given fields:
func (_m *DiscoveryInterface) ServerPreferredResources() ([]*v1.APIResourceList, error) {
	ret := _m.Called()

	var r0 []*v1.APIResourceList
	if rf, ok := ret.Get(0).(func() []*v1.APIResourceList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.APIResourceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerResources provides a mock function with given fields:
func (_m *DiscoveryInterface) ServerResources() ([]*v1.APIResourceList, error) {
	ret := _m.Called()

	var r0 []*v1.APIResourceList
	if rf, ok := ret.Get(0).(func() []*v1.APIResourceList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.APIResourceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerResourcesForGroupVersion provides a mock function with given fields: groupVersion
func (_m *DiscoveryInterface) ServerResourcesForGroupVersion(groupVersion string) (*v1.APIResourceList, error) {
	ret := _m.Called(groupVersion)

	var r0 *v1.APIResourceList
	if rf, ok := ret.Get(0).(func(string) *v1.APIResourceList); ok {
		r0 = rf(groupVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.APIResourceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupVersion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerVersion provides a mock function with given fields:
func (_m *DiscoveryInterface) ServerVersion() (*version.Info, error) {
	ret := _m.Called()

	var r0 *version.Info
	if rf, ok := ret.Get(0).(func() *version.Info); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
