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

// enables artifact list
var _ client.ArtifactLister = &TemplatePlugin{}

func (t *TemplatePlugin) ListArtifacts(ctx context.Context, params metav1alpha1.ArtifactOptions, option metav1alpha1.ListOptions) (*metav1alpha1.ArtifactList, error) {
	return nil, ErrNotImplemented
}

// enables artifact creation
var _ client.ArtifactGetter = &TemplatePlugin{}

func (t *TemplatePlugin) GetArtifact(ctx context.Context, params metav1alpha1.ArtifactOptions) (*metav1alpha1.Artifact, error) {
	return nil, ErrNotImplemented
}

// enables artifact deletion
var _ client.ArtifactDeleter = &TemplatePlugin{}

func (t *TemplatePlugin) DeleteArtifact(ctx context.Context, params metav1alpha1.ArtifactOptions) error {
	return ErrNotImplemented
}

// enables project creation
var _ client.ArtifactTagDeleter = &TemplatePlugin{}

func (t *TemplatePlugin) DeleteArtifactTag(ctx context.Context, params metav1alpha1.ArtifactTagOptions) error {
	return ErrNotImplemented
}

// enables image scan
var _ client.ScanImage = &TemplatePlugin{}

func (t *TemplatePlugin) ScanImage(ctx context.Context, params metav1alpha1.ArtifactOptions) error {
	return ErrNotImplemented
}

// enables image config
var _ client.ImageConfigGetter = &TemplatePlugin{}

func (t *TemplatePlugin) GetImageConfig(ctx context.Context, params metav1alpha1.ArtifactOptions) (*metav1alpha1.ImageConfig, error) {
	return nil, ErrNotImplemented
}
