// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2beta3 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/client/clientset/versioned/typed/config/v2beta3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeApisixV2beta3 struct {
	*testing.Fake
}

func (c *FakeApisixV2beta3) ApisixClusterConfigs() v2beta3.ApisixClusterConfigInterface {
	return &FakeApisixClusterConfigs{c}
}

func (c *FakeApisixV2beta3) ApisixConsumers(namespace string) v2beta3.ApisixConsumerInterface {
	return &FakeApisixConsumers{c, namespace}
}

func (c *FakeApisixV2beta3) ApisixPluginConfigs(namespace string) v2beta3.ApisixPluginConfigInterface {
	return &FakeApisixPluginConfigs{c, namespace}
}

func (c *FakeApisixV2beta3) ApisixRoutes(namespace string) v2beta3.ApisixRouteInterface {
	return &FakeApisixRoutes{c, namespace}
}

func (c *FakeApisixV2beta3) ApisixTlses(namespace string) v2beta3.ApisixTlsInterface {
	return &FakeApisixTlses{c, namespace}
}

func (c *FakeApisixV2beta3) ApisixUpstreams(namespace string) v2beta3.ApisixUpstreamInterface {
	return &FakeApisixUpstreams{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApisixV2beta3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
