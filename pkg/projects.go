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

	"github.com/katanomi/pkg/apis/meta/v1alpha1"
	"github.com/katanomi/pkg/plugin/client"
)

// enables project listing
var _ client.ProjectLister = &TemplatePlugin{}

// ListProjects lists projects in a plugin
func (t *TemplatePlugin) ListProjects(ctx context.Context, option v1alpha1.ListOptions) (list *v1alpha1.ProjectList, err error) {
	err = ErrNotImplemented
	return
}

// enables project creation
var _ client.ProjectCreator = &TemplatePlugin{}

// CreateProject create a project in a plugin
func (t *TemplatePlugin) CreateProject(ctx context.Context, project *v1alpha1.Project) (result *v1alpha1.Project, err error) {
	err = ErrNotImplemented
	return
}
