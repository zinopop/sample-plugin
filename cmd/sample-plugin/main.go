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

package main

import (
	corev1alpha1 "github.com/katanomi/core/pkg/apis/core/v1alpha1"
	corecontrollers "github.com/katanomi/core/pkg/controllers/core"
	"github.com/katanomi/core/pkg/system"
	corewebhook "github.com/katanomi/core/pkg/webhook"
	"github.com/katanomi/pkg/sharedmain"
	"gomod.alauda.cn/katanomi-plugin-sample/pkg"
	templatepkg "gomod.alauda.cn/katanomi-plugin-sample/pkg"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	eventingv1 "knative.dev/eventing/pkg/apis/eventing/v1"
	messagingv1 "knative.dev/eventing/pkg/apis/messaging/v1"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	// eventing
	utilruntime.Must(messagingv1.AddToScheme(scheme))
	utilruntime.Must(eventingv1.AddToScheme(scheme))

	utilruntime.Must(corev1alpha1.AddToScheme(scheme))
}

func main() {
	plugin := &templatepkg.TemplatePlugin{}
	sharedmain.App("sample-plugin").
		Scheme(scheme).
		Log().
		Tracing(nil).
		Profiling().
		Plugins(
			// adds plugin
			plugin,
		).
		PluginAttributes(plugin, system.FilesPath()+"/attributes/sample-plugin.yaml").
		Webservices(
			&corewebhook.WebhookHandler{
				WebhookReceiver: plugin,
			},
		).
		Controllers(

			// basic IntegrationClass controller
			&corecontrollers.IntegrationClassReconciler{
				PluginRegister: plugin,
			},
			// webhook register
			&corecontrollers.WebhookReconciler{
				WebhookRegister: plugin,
			},
			&corecontrollers.GitTriggerReconciler{
				GitTriggerRegister: plugin,
			},
			// core
			&corecontrollers.TaskInstaller{
				ImportTypes: pkg.GetImportersForPlugin(),
			},
		).
		APIDocs().
		Run()
}
