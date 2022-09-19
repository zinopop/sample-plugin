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

import "github.com/katanomi/pkg/plugin/client"

// enables attributes
var _ client.PluginAttributes = &TemplatePlugin{}

func (t *TemplatePlugin) SetAttribute(key string, values ...string) {
	if t.attributes == nil {
		t.attributes = make(map[string][]string)
	}
	t.attributes[key] = values
}

func (t *TemplatePlugin) GetAttribute(key string) []string {
	if values, exist := t.attributes[key]; exist {
		return values
	}

	return []string{}
}

func (t *TemplatePlugin) Attributes() map[string][]string {
	if t.attributes == nil {
		return make(map[string][]string)
	}
	return t.attributes
}
