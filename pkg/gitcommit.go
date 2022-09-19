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

	"github.com/katanomi/pkg/plugin/client"

	metav1alpha1 "github.com/katanomi/pkg/apis/meta/v1alpha1"
)

// enable git commit list
var _ client.GitCommitLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListGitCommit(ctx context.Context, option metav1alpha1.GitCommitListOption, listOption metav1alpha1.ListOptions) (metav1alpha1.GitCommitList, error) {
	return metav1alpha1.GitCommitList{}, ErrNotImplemented
}

// enable git commit get
var _ client.GitCommitGetter = &TemplatePlugin{}

func (t *TemplatePlugin) GetGitCommit(ctx context.Context, option metav1alpha1.GitCommitOption) (metav1alpha1.GitCommit, error) {
	return metav1alpha1.GitCommit{}, ErrNotImplemented
}

// enable git commit status list
var _ client.GitCommitStatusLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListGitCommitStatus(ctx context.Context, option metav1alpha1.GitCommitOption, listOption metav1alpha1.ListOptions) (metav1alpha1.GitCommitStatusList, error) {
	return metav1alpha1.GitCommitStatusList{}, nil
}

// enable git commit status creation
var _ client.GitCommitStatusCreator = &TemplatePlugin{}

func (t *TemplatePlugin) CreateGitCommitStatus(ctx context.Context, payload metav1alpha1.CreateCommitStatusPayload) (metav1alpha1.GitCommitStatus, error) {
	return metav1alpha1.GitCommitStatus{}, ErrNotImplemented
}
