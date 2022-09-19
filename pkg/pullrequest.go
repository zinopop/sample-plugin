/*
Copyright 2022 The Katanomi Authors.

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

	metav1alpha1 "github.com/katanomi/pkg/apis/meta/v1alpha1"
)

func (t *TemplatePlugin) ListGitPullRequest(ctx context.Context, option metav1alpha1.GitPullRequestListOption, listOption metav1alpha1.ListOptions) (metav1alpha1.GitPullRequestList, error) {
	return metav1alpha1.GitPullRequestList{}, ErrNotImplemented
}

func (t *TemplatePlugin) GetGitPullRequest(ctx context.Context, option metav1alpha1.GitPullRequestOption) (metav1alpha1.GitPullRequest, error) {
	return metav1alpha1.GitPullRequest{}, ErrNotImplemented
}

func (t *TemplatePlugin) CreatePullRequest(ctx context.Context, payload metav1alpha1.CreatePullRequestPayload) (metav1alpha1.GitPullRequest, error) {
	return metav1alpha1.GitPullRequest{}, ErrNotImplemented
}
