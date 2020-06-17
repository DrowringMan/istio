// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config/host"
	"istio.io/istio/pkg/config/labels"
)

type ServiceDiscovery struct {
	GetIstioServiceAccountsStub        func(*model.Service, []int) []string
	getIstioServiceAccountsMutex       sync.RWMutex
	getIstioServiceAccountsArgsForCall []struct {
		arg1 *model.Service
		arg2 []int
	}
	getIstioServiceAccountsReturns struct {
		result1 []string
	}
	getIstioServiceAccountsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetProxyServiceInstancesStub        func(*model.Proxy) ([]*model.ServiceInstance, error)
	getProxyServiceInstancesMutex       sync.RWMutex
	getProxyServiceInstancesArgsForCall []struct {
		arg1 *model.Proxy
	}
	getProxyServiceInstancesReturns struct {
		result1 []*model.ServiceInstance
		result2 error
	}
	getProxyServiceInstancesReturnsOnCall map[int]struct {
		result1 []*model.ServiceInstance
		result2 error
	}
	GetProxyWorkloadLabelsStub        func(*model.Proxy) (labels.Collection, error)
	getProxyWorkloadLabelsMutex       sync.RWMutex
	getProxyWorkloadLabelsArgsForCall []struct {
		arg1 *model.Proxy
	}
	getProxyWorkloadLabelsReturns struct {
		result1 labels.Collection
		result2 error
	}
	getProxyWorkloadLabelsReturnsOnCall map[int]struct {
		result1 labels.Collection
		result2 error
	}
	GetServiceStub        func(host.Name) (*model.Service, error)
	getServiceMutex       sync.RWMutex
	getServiceArgsForCall []struct {
		arg1 host.Name
	}
	getServiceReturns struct {
		result1 *model.Service
		result2 error
	}
	getServiceReturnsOnCall map[int]struct {
		result1 *model.Service
		result2 error
	}
	InstancesByPortStub        func(*model.Service, int, labels.Collection) ([]*model.ServiceInstance, error)
	instancesByPortMutex       sync.RWMutex
	instancesByPortArgsForCall []struct {
		arg1 *model.Service
		arg2 int
		arg3 labels.Collection
	}
	instancesByPortReturns struct {
		result1 []*model.ServiceInstance
		result2 error
	}
	instancesByPortReturnsOnCall map[int]struct {
		result1 []*model.ServiceInstance
		result2 error
	}
	ServicesStub        func() ([]*model.Service, error)
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct {
	}
	servicesReturns struct {
		result1 []*model.Service
		result2 error
	}
	servicesReturnsOnCall map[int]struct {
		result1 []*model.Service
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ServiceDiscovery) GetIstioServiceAccounts(arg1 *model.Service, arg2 []int) []string {
	var arg2Copy []int
	if arg2 != nil {
		arg2Copy = make([]int, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getIstioServiceAccountsMutex.Lock()
	ret, specificReturn := fake.getIstioServiceAccountsReturnsOnCall[len(fake.getIstioServiceAccountsArgsForCall)]
	fake.getIstioServiceAccountsArgsForCall = append(fake.getIstioServiceAccountsArgsForCall, struct {
		arg1 *model.Service
		arg2 []int
	}{arg1, arg2Copy})
	fake.recordInvocation("GetIstioServiceAccounts", []interface{}{arg1, arg2Copy})
	fake.getIstioServiceAccountsMutex.Unlock()
	if fake.GetIstioServiceAccountsStub != nil {
		return fake.GetIstioServiceAccountsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getIstioServiceAccountsReturns
	return fakeReturns.result1
}

func (fake *ServiceDiscovery) GetIstioServiceAccountsCallCount() int {
	fake.getIstioServiceAccountsMutex.RLock()
	defer fake.getIstioServiceAccountsMutex.RUnlock()
	return len(fake.getIstioServiceAccountsArgsForCall)
}

func (fake *ServiceDiscovery) GetIstioServiceAccountsCalls(stub func(*model.Service, []int) []string) {
	fake.getIstioServiceAccountsMutex.Lock()
	defer fake.getIstioServiceAccountsMutex.Unlock()
	fake.GetIstioServiceAccountsStub = stub
}

func (fake *ServiceDiscovery) GetIstioServiceAccountsArgsForCall(i int) (*model.Service, []int) {
	fake.getIstioServiceAccountsMutex.RLock()
	defer fake.getIstioServiceAccountsMutex.RUnlock()
	argsForCall := fake.getIstioServiceAccountsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ServiceDiscovery) GetIstioServiceAccountsReturns(result1 []string) {
	fake.getIstioServiceAccountsMutex.Lock()
	defer fake.getIstioServiceAccountsMutex.Unlock()
	fake.GetIstioServiceAccountsStub = nil
	fake.getIstioServiceAccountsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *ServiceDiscovery) GetIstioServiceAccountsReturnsOnCall(i int, result1 []string) {
	fake.getIstioServiceAccountsMutex.Lock()
	defer fake.getIstioServiceAccountsMutex.Unlock()
	fake.GetIstioServiceAccountsStub = nil
	if fake.getIstioServiceAccountsReturnsOnCall == nil {
		fake.getIstioServiceAccountsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getIstioServiceAccountsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *ServiceDiscovery) GetProxyServiceInstances(arg1 *model.Proxy) ([]*model.ServiceInstance, error) {
	fake.getProxyServiceInstancesMutex.Lock()
	ret, specificReturn := fake.getProxyServiceInstancesReturnsOnCall[len(fake.getProxyServiceInstancesArgsForCall)]
	fake.getProxyServiceInstancesArgsForCall = append(fake.getProxyServiceInstancesArgsForCall, struct {
		arg1 *model.Proxy
	}{arg1})
	fake.recordInvocation("GetProxyServiceInstances", []interface{}{arg1})
	fake.getProxyServiceInstancesMutex.Unlock()
	if fake.GetProxyServiceInstancesStub != nil {
		return fake.GetProxyServiceInstancesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getProxyServiceInstancesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ServiceDiscovery) GetProxyServiceInstancesCallCount() int {
	fake.getProxyServiceInstancesMutex.RLock()
	defer fake.getProxyServiceInstancesMutex.RUnlock()
	return len(fake.getProxyServiceInstancesArgsForCall)
}

func (fake *ServiceDiscovery) GetProxyServiceInstancesCalls(stub func(*model.Proxy) ([]*model.ServiceInstance, error)) {
	fake.getProxyServiceInstancesMutex.Lock()
	defer fake.getProxyServiceInstancesMutex.Unlock()
	fake.GetProxyServiceInstancesStub = stub
}

func (fake *ServiceDiscovery) GetProxyServiceInstancesArgsForCall(i int) *model.Proxy {
	fake.getProxyServiceInstancesMutex.RLock()
	defer fake.getProxyServiceInstancesMutex.RUnlock()
	argsForCall := fake.getProxyServiceInstancesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServiceDiscovery) GetProxyServiceInstancesReturns(result1 []*model.ServiceInstance, result2 error) {
	fake.getProxyServiceInstancesMutex.Lock()
	defer fake.getProxyServiceInstancesMutex.Unlock()
	fake.GetProxyServiceInstancesStub = nil
	fake.getProxyServiceInstancesReturns = struct {
		result1 []*model.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) GetProxyServiceInstancesReturnsOnCall(i int, result1 []*model.ServiceInstance, result2 error) {
	fake.getProxyServiceInstancesMutex.Lock()
	defer fake.getProxyServiceInstancesMutex.Unlock()
	fake.GetProxyServiceInstancesStub = nil
	if fake.getProxyServiceInstancesReturnsOnCall == nil {
		fake.getProxyServiceInstancesReturnsOnCall = make(map[int]struct {
			result1 []*model.ServiceInstance
			result2 error
		})
	}
	fake.getProxyServiceInstancesReturnsOnCall[i] = struct {
		result1 []*model.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) GetProxyWorkloadLabels(arg1 *model.Proxy) (labels.Collection, error) {
	fake.getProxyWorkloadLabelsMutex.Lock()
	ret, specificReturn := fake.getProxyWorkloadLabelsReturnsOnCall[len(fake.getProxyWorkloadLabelsArgsForCall)]
	fake.getProxyWorkloadLabelsArgsForCall = append(fake.getProxyWorkloadLabelsArgsForCall, struct {
		arg1 *model.Proxy
	}{arg1})
	fake.recordInvocation("GetProxyWorkloadLabels", []interface{}{arg1})
	fake.getProxyWorkloadLabelsMutex.Unlock()
	if fake.GetProxyWorkloadLabelsStub != nil {
		return fake.GetProxyWorkloadLabelsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getProxyWorkloadLabelsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ServiceDiscovery) GetProxyWorkloadLabelsCallCount() int {
	fake.getProxyWorkloadLabelsMutex.RLock()
	defer fake.getProxyWorkloadLabelsMutex.RUnlock()
	return len(fake.getProxyWorkloadLabelsArgsForCall)
}

func (fake *ServiceDiscovery) GetProxyWorkloadLabelsCalls(stub func(*model.Proxy) (labels.Collection, error)) {
	fake.getProxyWorkloadLabelsMutex.Lock()
	defer fake.getProxyWorkloadLabelsMutex.Unlock()
	fake.GetProxyWorkloadLabelsStub = stub
}

func (fake *ServiceDiscovery) GetProxyWorkloadLabelsArgsForCall(i int) *model.Proxy {
	fake.getProxyWorkloadLabelsMutex.RLock()
	defer fake.getProxyWorkloadLabelsMutex.RUnlock()
	argsForCall := fake.getProxyWorkloadLabelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServiceDiscovery) GetProxyWorkloadLabelsReturns(result1 labels.Collection, result2 error) {
	fake.getProxyWorkloadLabelsMutex.Lock()
	defer fake.getProxyWorkloadLabelsMutex.Unlock()
	fake.GetProxyWorkloadLabelsStub = nil
	fake.getProxyWorkloadLabelsReturns = struct {
		result1 labels.Collection
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) GetProxyWorkloadLabelsReturnsOnCall(i int, result1 labels.Collection, result2 error) {
	fake.getProxyWorkloadLabelsMutex.Lock()
	defer fake.getProxyWorkloadLabelsMutex.Unlock()
	fake.GetProxyWorkloadLabelsStub = nil
	if fake.getProxyWorkloadLabelsReturnsOnCall == nil {
		fake.getProxyWorkloadLabelsReturnsOnCall = make(map[int]struct {
			result1 labels.Collection
			result2 error
		})
	}
	fake.getProxyWorkloadLabelsReturnsOnCall[i] = struct {
		result1 labels.Collection
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) GetService(arg1 host.Name) (*model.Service, error) {
	fake.getServiceMutex.Lock()
	ret, specificReturn := fake.getServiceReturnsOnCall[len(fake.getServiceArgsForCall)]
	fake.getServiceArgsForCall = append(fake.getServiceArgsForCall, struct {
		arg1 host.Name
	}{arg1})
	fake.recordInvocation("GetService", []interface{}{arg1})
	fake.getServiceMutex.Unlock()
	if fake.GetServiceStub != nil {
		return fake.GetServiceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getServiceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ServiceDiscovery) GetServiceCallCount() int {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return len(fake.getServiceArgsForCall)
}

func (fake *ServiceDiscovery) GetServiceCalls(stub func(host.Name) (*model.Service, error)) {
	fake.getServiceMutex.Lock()
	defer fake.getServiceMutex.Unlock()
	fake.GetServiceStub = stub
}

func (fake *ServiceDiscovery) GetServiceArgsForCall(i int) host.Name {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	argsForCall := fake.getServiceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ServiceDiscovery) GetServiceReturns(result1 *model.Service, result2 error) {
	fake.getServiceMutex.Lock()
	defer fake.getServiceMutex.Unlock()
	fake.GetServiceStub = nil
	fake.getServiceReturns = struct {
		result1 *model.Service
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) GetServiceReturnsOnCall(i int, result1 *model.Service, result2 error) {
	fake.getServiceMutex.Lock()
	defer fake.getServiceMutex.Unlock()
	fake.GetServiceStub = nil
	if fake.getServiceReturnsOnCall == nil {
		fake.getServiceReturnsOnCall = make(map[int]struct {
			result1 *model.Service
			result2 error
		})
	}
	fake.getServiceReturnsOnCall[i] = struct {
		result1 *model.Service
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) InstancesByPort(arg1 *model.Service, arg2 int, arg3 labels.Collection) ([]*model.ServiceInstance, error) {
	fake.instancesByPortMutex.Lock()
	ret, specificReturn := fake.instancesByPortReturnsOnCall[len(fake.instancesByPortArgsForCall)]
	fake.instancesByPortArgsForCall = append(fake.instancesByPortArgsForCall, struct {
		arg1 *model.Service
		arg2 int
		arg3 labels.Collection
	}{arg1, arg2, arg3})
	fake.recordInvocation("InstancesByPort", []interface{}{arg1, arg2, arg3})
	fake.instancesByPortMutex.Unlock()
	if fake.InstancesByPortStub != nil {
		return fake.InstancesByPortStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.instancesByPortReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ServiceDiscovery) InstancesByPortCallCount() int {
	fake.instancesByPortMutex.RLock()
	defer fake.instancesByPortMutex.RUnlock()
	return len(fake.instancesByPortArgsForCall)
}

func (fake *ServiceDiscovery) InstancesByPortCalls(stub func(*model.Service, int, labels.Collection) ([]*model.ServiceInstance, error)) {
	fake.instancesByPortMutex.Lock()
	defer fake.instancesByPortMutex.Unlock()
	fake.InstancesByPortStub = stub
}

func (fake *ServiceDiscovery) InstancesByPortArgsForCall(i int) (*model.Service, int, labels.Collection) {
	fake.instancesByPortMutex.RLock()
	defer fake.instancesByPortMutex.RUnlock()
	argsForCall := fake.instancesByPortArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ServiceDiscovery) InstancesByPortReturns(result1 []*model.ServiceInstance, result2 error) {
	fake.instancesByPortMutex.Lock()
	defer fake.instancesByPortMutex.Unlock()
	fake.InstancesByPortStub = nil
	fake.instancesByPortReturns = struct {
		result1 []*model.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) InstancesByPortReturnsOnCall(i int, result1 []*model.ServiceInstance, result2 error) {
	fake.instancesByPortMutex.Lock()
	defer fake.instancesByPortMutex.Unlock()
	fake.InstancesByPortStub = nil
	if fake.instancesByPortReturnsOnCall == nil {
		fake.instancesByPortReturnsOnCall = make(map[int]struct {
			result1 []*model.ServiceInstance
			result2 error
		})
	}
	fake.instancesByPortReturnsOnCall[i] = struct {
		result1 []*model.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) Services() ([]*model.Service, error) {
	fake.servicesMutex.Lock()
	ret, specificReturn := fake.servicesReturnsOnCall[len(fake.servicesArgsForCall)]
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct {
	}{})
	fake.recordInvocation("Services", []interface{}{})
	fake.servicesMutex.Unlock()
	if fake.ServicesStub != nil {
		return fake.ServicesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.servicesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ServiceDiscovery) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *ServiceDiscovery) ServicesCalls(stub func() ([]*model.Service, error)) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = stub
}

func (fake *ServiceDiscovery) ServicesReturns(result1 []*model.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 []*model.Service
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) ServicesReturnsOnCall(i int, result1 []*model.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	if fake.servicesReturnsOnCall == nil {
		fake.servicesReturnsOnCall = make(map[int]struct {
			result1 []*model.Service
			result2 error
		})
	}
	fake.servicesReturnsOnCall[i] = struct {
		result1 []*model.Service
		result2 error
	}{result1, result2}
}

func (fake *ServiceDiscovery) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getIstioServiceAccountsMutex.RLock()
	defer fake.getIstioServiceAccountsMutex.RUnlock()
	fake.getProxyServiceInstancesMutex.RLock()
	defer fake.getProxyServiceInstancesMutex.RUnlock()
	fake.getProxyWorkloadLabelsMutex.RLock()
	defer fake.getProxyWorkloadLabelsMutex.RUnlock()
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	fake.instancesByPortMutex.RLock()
	defer fake.instancesByPortMutex.RUnlock()
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ServiceDiscovery) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ model.ServiceDiscovery = new(ServiceDiscovery)
