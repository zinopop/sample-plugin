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
	"fmt"

	"github.com/katanomi/pkg/plugin/client"

	"github.com/katanomi/pkg/apis/meta/v1alpha1"
)

func (t *TemplatePlugin) AuthCheck(ctx context.Context, options v1alpha1.AuthCheckOptions) (*v1alpha1.AuthCheck, error) {
	// meta include tool base url and version
	meta := client.ExtraMeta(ctx)
	fmt.Println(meta)

	// auth include tool base url and version
	auth := client.ExtractAuth(ctx)
	fmt.Println(auth)

	// checking auth against tool here

	return nil, ErrNotImplemented
}

func (t *TemplatePlugin) AuthToken(ctx context.Context) (*v1alpha1.AuthToken, error) {
	// meta include tool base url and version
	meta := client.ExtraMeta(ctx)
	fmt.Println(meta)

	// auth include tool base url and version
	auth := client.ExtractAuth(ctx)
	fmt.Println(auth)

	// refresh auth token and return with new one here

	return nil, ErrNotImplemented
}
