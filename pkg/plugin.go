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
	"fmt"

	kclient "github.com/katanomi/pkg/client"
	"github.com/katanomi/pkg/plugin/client"
	"go.uber.org/zap"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var ErrNotImplemented = fmt.Errorf("not Implemented")

type TemplatePlugin struct {
	*zap.SugaredLogger

	Client     ctrlclient.Client
	attributes map[string][]string
}

// basic interface for all plugins
var _ client.Interface = &TemplatePlugin{}

// Path custom url path to access this plugin
func (t *TemplatePlugin) Path() string {
	return "template"
}

// Setup basic initialization setup for plugins
func (t *TemplatePlugin) Setup(ctx context.Context, logger *zap.SugaredLogger) error {
	// example setup method
	t.Client = kclient.Client(ctx)
	if t.Client == nil {
		return fmt.Errorf("should have a client to start this plugin")
	}
	t.SugaredLogger = logger
	return nil
}
