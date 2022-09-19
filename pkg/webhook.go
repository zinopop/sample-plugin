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
	"context"

	cloudevent "github.com/cloudevents/sdk-go/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/katanomi/pkg/plugin/client"
	"knative.dev/pkg/apis"

	metav1alpha1 "github.com/katanomi/pkg/apis/meta/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

var _ client.WebhookRegister = &TemplatePlugin{}

func (t *TemplatePlugin) CreateWebhook(ctx context.Context, spec metav1alpha1.WebhookRegisterSpec, secret corev1.Secret) (result metav1alpha1.WebhookRegisterStatus, err error) {
	err = ErrNotImplemented
	return
}

func (t *TemplatePlugin) UpdateWebhook(ctx context.Context, spec metav1alpha1.WebhookRegisterSpec, secret corev1.Secret) (result metav1alpha1.WebhookRegisterStatus, err error) {
	err = ErrNotImplemented
	return
}

func (t *TemplatePlugin) DeleteWebhook(ctx context.Context, spec metav1alpha1.WebhookRegisterSpec, secret corev1.Secret) (err error) {
	err = ErrNotImplemented
	return
}

func (h *TemplatePlugin) ListWebhooks(ctx context.Context, uri apis.URL, secret corev1.Secret) ([]metav1alpha1.WebhookRegisterStatus, error) {
	return []metav1alpha1.WebhookRegisterStatus{}, ErrNotImplemented
}

func (t *TemplatePlugin) IsSameResource(ctx context.Context, i, j metav1alpha1.ResourceURI) bool {
	// customize if necessary
	return i.URI.String() == j.URI.String()
}

var _ client.WebhookReceiver = &TemplatePlugin{}

// ReceiveWebhook basic webhook receiver
// actually it just receives cloud events and returns
func (t *TemplatePlugin) ReceiveWebhook(ctx context.Context, req *restful.Request, secret string) (event cloudevent.Event, err error) {
	err = ErrNotImplemented
	return
}
