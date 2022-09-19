/*
Copyright 2021 The Katanomi Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"fmt"
	"os"

	"github.com/katanomi/pkg/plugin/client"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/network"
	"knative.dev/pkg/system"
)

var _ client.PluginRegister = &TemplatePlugin{}

func (t *TemplatePlugin) GetIntegrationClassName() string {
	// TODO: change this to the real integration class name
	return "template"
}

func (t *TemplatePlugin) GetAddressURL() *apis.URL {
	svcNamespace := system.Namespace()
	svcName := os.Getenv("SERVICE_NAME")
	urlStr := fmt.Sprintf("http://%s", network.GetServiceHostname(svcName, svcNamespace))
	url, _ := apis.ParseURL(urlStr)
	return url
}

func (t *TemplatePlugin) GetWebhookURL() (*apis.URL, bool) {
	webhookAddr := os.Getenv("WEBHOOK_ADDRESS")
	var url *apis.URL
	if webhookAddr != "" {
		url, _ = apis.ParseURL(webhookAddr)
	}
	return url, url != nil
}

func (t *TemplatePlugin) GetSupportedVersions() []string {
	// TODO: change this to the real supported versions
	return []string{
		"v0.1.0",
		//"online",
	}
}

func (t *TemplatePlugin) GetSecretTypes() []string {
	// TODO: change this to the real supported auth
	return []string{
		string(corev1.SecretTypeBasicAuth),
	}
}

func (t *TemplatePlugin) GetReplicationPolicyTypes() []string {
	// TODO: change this to the real supported versions
	return []string{
		"Manual",
		// "AttributeBased",
		// "Private",
	}
}

func (t *TemplatePlugin) GetResourceTypes() []string {
	// TODO: change this to the real supported resource types
	return []string{
		"ImageRepository",
	}
}

func (n *TemplatePlugin) GetAllowEmptySecret() []string {
	return []string{"false"}
}
