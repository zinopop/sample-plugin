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

// enables git branch list

var _ client.GitBranchLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListGitBranch(ctx context.Context, branchOption metav1alpha1.GitBranchOption, option metav1alpha1.ListOptions) (metav1alpha1.GitBranchList, error) {
	return metav1alpha1.GitBranchList{}, ErrNotImplemented
}

// enables git branch get
var _ client.GitBranchGetter = &TemplatePlugin{}

func (t *TemplatePlugin) GetGitBranch(ctx context.Context, repoOption metav1alpha1.GitRepo, branch string) (metav1alpha1.GitBranch, error) {
	return metav1alpha1.GitBranch{}, ErrNotImplemented
}

// enables git branch creation
var _ client.GitBranchCreator = &TemplatePlugin{}

func (t *TemplatePlugin) CreateGitBranch(ctx context.Context, payload metav1alpha1.CreateBranchPayload) (metav1alpha1.GitBranch, error) {
	return metav1alpha1.GitBranch{}, ErrNotImplemented
}
