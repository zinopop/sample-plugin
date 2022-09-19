{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}

{{- define "plugin.image" -}}
{{- $registryAddress :=  .Values.registry.address -}}
{{- $repositoryName := .Values.images.plugin.repository -}}
{{- $tag := .Values.images.plugin.tag -}}
{{- if .Values.ko.enable }}
{{- printf "%s" .Values.ko.image -}}
{{- else }}
{{- printf "%s/%s:%s" $registryAddress $repositoryName $tag -}}
{{- end }}
{{- end -}}

{{- define "plugin.imagePolicy" -}}
{{- if .Values.ko.enable }}
{{- printf "IfNotPresent" -}}
{{- else -}}
{{- printf "Always" -}}
{{- end }}
{{- end -}}

{{- define "plugin.tolerations" -}}
tolerations:
- effect: NoSchedule
  key: no-cpaas-pod
  operator: Equal
  value: "true"
{{- end -}}
