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

// enable git commit comment list
var _ client.GitCommitCommentLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListGitCommitComment(ctx context.Context, option metav1alpha1.GitCommitOption, listOption metav1alpha1.ListOptions) (metav1alpha1.GitCommitCommentList, error) {
	return metav1alpha1.GitCommitCommentList{}, ErrNotImplemented
}

// enable git commit comment creation
var _ client.GitCommitCommentCreator = &TemplatePlugin{}

func (t *TemplatePlugin) CreateGitCommitComment(ctx context.Context, payload metav1alpha1.CreateCommitCommentPayload) (metav1alpha1.GitCommitComment, error) {
	return metav1alpha1.GitCommitComment{}, nil
}

// enable git pull request comment list
var _ client.GitPullRequestCommentLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListPullRequestComment(ctx context.Context, option metav1alpha1.GitPullRequestOption, listOption metav1alpha1.ListOptions) (metav1alpha1.GitPullRequestNoteList, error) {
	return metav1alpha1.GitPullRequestNoteList{}, ErrNotImplemented
}

// enable git pull request comment creation
var _ client.GitPullRequestCommentCreator = &TemplatePlugin{}

func (t *TemplatePlugin) CreatePullRequestComment(ctx context.Context, option metav1alpha1.CreatePullRequestCommentPayload) (metav1alpha1.GitPullRequestNote, error) {
	return metav1alpha1.GitPullRequestNote{}, ErrNotImplemented
}
