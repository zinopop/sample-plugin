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

// enables code quality get
var _ client.CodeQualityGetter = &TemplatePlugin{}

func (t *TemplatePlugin) GetCodeQuality(ctx context.Context, projectKey string) (*metav1alpha1.CodeQuality, error) {
	return nil, ErrNotImplemented
}

func (t *TemplatePlugin) GetCodeQualityOverviewByBranch(ctx context.Context, opt metav1alpha1.CodeQualityBaseOption) (*metav1alpha1.CodeQuality, error) {
	return nil, ErrNotImplemented
}

func (t *TemplatePlugin) GetCodeQualityLineCharts(ctx context.Context, opt metav1alpha1.CodeQualityLineChartOption) (*metav1alpha1.CodeQualityLineChart, error) {
	return nil, ErrNotImplemented
}

func (t *TemplatePlugin) GetOverview(ctx context.Context) (*metav1alpha1.CodeQualityProjectOverview, error) {
	return nil, ErrNotImplemented
}

func (t *TemplatePlugin) GetSummaryByTaskID(ctx context.Context, opt metav1alpha1.CodeQualityTaskOption) (*metav1alpha1.CodeQualityTaskMetrics, error) {
	return nil, ErrNotImplemented
}
