// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// LoadBalancerStrategyApplyConfiguration represents an declarative configuration of the LoadBalancerStrategy type for use
// with apply.
type LoadBalancerStrategyApplyConfiguration struct {
	Scope               *v1.LoadBalancerScope                             `json:"scope,omitempty"`
	AllowedSourceRanges []v1.CIDR                                         `json:"allowedSourceRanges,omitempty"`
	ProviderParameters  *ProviderLoadBalancerParametersApplyConfiguration `json:"providerParameters,omitempty"`
	DNSManagementPolicy *v1.LoadBalancerDNSManagementPolicy               `json:"dnsManagementPolicy,omitempty"`
}

// LoadBalancerStrategyApplyConfiguration constructs an declarative configuration of the LoadBalancerStrategy type for use with
// apply.
func LoadBalancerStrategy() *LoadBalancerStrategyApplyConfiguration {
	return &LoadBalancerStrategyApplyConfiguration{}
}

// WithScope sets the Scope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scope field is set to the value of the last call.
func (b *LoadBalancerStrategyApplyConfiguration) WithScope(value v1.LoadBalancerScope) *LoadBalancerStrategyApplyConfiguration {
	b.Scope = &value
	return b
}

// WithAllowedSourceRanges adds the given value to the AllowedSourceRanges field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedSourceRanges field.
func (b *LoadBalancerStrategyApplyConfiguration) WithAllowedSourceRanges(values ...v1.CIDR) *LoadBalancerStrategyApplyConfiguration {
	for i := range values {
		b.AllowedSourceRanges = append(b.AllowedSourceRanges, values[i])
	}
	return b
}

// WithProviderParameters sets the ProviderParameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProviderParameters field is set to the value of the last call.
func (b *LoadBalancerStrategyApplyConfiguration) WithProviderParameters(value *ProviderLoadBalancerParametersApplyConfiguration) *LoadBalancerStrategyApplyConfiguration {
	b.ProviderParameters = value
	return b
}

// WithDNSManagementPolicy sets the DNSManagementPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSManagementPolicy field is set to the value of the last call.
func (b *LoadBalancerStrategyApplyConfiguration) WithDNSManagementPolicy(value v1.LoadBalancerDNSManagementPolicy) *LoadBalancerStrategyApplyConfiguration {
	b.DNSManagementPolicy = &value
	return b
}
