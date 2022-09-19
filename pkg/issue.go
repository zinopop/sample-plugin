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

// enable issue list
var _ client.IssueLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListIssues(ctx context.Context, params metav1alpha1.IssueOptions, option metav1alpha1.ListOptions) (*metav1alpha1.IssueList, error) {
	return nil, ErrNotImplemented
}

// enable issue get
var _ client.IssueGetter = &TemplatePlugin{}

func (t *TemplatePlugin) GetIssue(ctx context.Context, params metav1alpha1.IssueOptions, option metav1alpha1.ListOptions) (*metav1alpha1.Issue, error) {
	return nil, ErrNotImplemented
}

// enable issue branch list
var _ client.IssueBranchLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListIssueBranches(ctx context.Context, params metav1alpha1.IssueOptions, option metav1alpha1.ListOptions) (*metav1alpha1.BranchList, error) {
	return nil, ErrNotImplemented

}

// enable issue branch creation
var _ client.IssueBranchCreator = &TemplatePlugin{}

func (t *TemplatePlugin) CreateIssueBranch(ctx context.Context, params metav1alpha1.IssueOptions, payload metav1alpha1.Branch) (*metav1alpha1.Branch, error) {
	return nil, ErrNotImplemented

}

// enable issue branch deletion
var _ client.IssueBranchDeleter = &TemplatePlugin{}

func (t *TemplatePlugin) DeleteIssueBranch(ctx context.Context, params metav1alpha1.IssueOptions, option metav1alpha1.ListOptions) error {
	return ErrNotImplemented
}

// enable issue attributes get
var _ client.IssueAttributeGetter = &TemplatePlugin{}

func (t *TemplatePlugin) GetIssueAttribute(ctx context.Context, params metav1alpha1.IssueOptions, option metav1alpha1.ListOptions) (*metav1alpha1.Attribute, error) {
	return nil, ErrNotImplemented
}
