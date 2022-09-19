# katanomi-plugin-sample

## How To Start

### 修改项目配置

* [ ] `sample-plugin/Chart.yaml` 修改插件名称、版本号、描述

* [ ] `integrationclasses/sample-plugin.yaml` 把 sample-plugin 换成正确的工具名称

* [ ] `sample-plugin/values.yaml`
    * [ ] `registry.address` 修改插件镜像仓库地址
    * [ ] `images.plugin.repository` 修改插件镜像名称
    * [ ] `images.plugin.image` 修改插件镜像 Tag
    * [ ] `ko.image` 如果使用 ko 部署，修改为正确的 cmd 路径

* [ ] `cmd/sample-plugin/main.go`
  * [ ] plugin := &templatepkg.TemplatePlugin{} 换成正确的插件
  * [ ] `sharedmain.App("sample-plugin")` 把 sample-plugin 换成正确的插件名称
  * [ ] `PluginAttributes(plugin, system.FilesPath()+"/attributes/sample-plugin.yaml") sample-plugin.yaml` 修改插件自定义属性文件

### 调试

检查 charts 语法是否正确

```shell
cd charts/sample-plugin
helm lint
```

检查 charts values 是否设置正确

```shell
cd charts/sample-plugin
helm isntall test ./ --dry-run 
```

检查输出的 yaml 是否符合预期

### 本地部署

```shell
C=sample-plugin make deploy
```
