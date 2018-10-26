{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "chart-app.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "chart-app.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "imagespec" -}}
  {{- $containername := .containername -}}
  {{- range $container := .containers -}}
    {{- if eq $container.containername $containername -}}
      {{- printf "%s:%s"  $container.imagerepo  $container.imagetag -}}
    {{- end -}}
  {{- end -}}
{{- end -}}


{{- define "imagepull" -}}
  {{- $containername := .containername -}}
  {{- range $container := .containers -}}
    {{- if eq $container.containername $containername -}}
      {{- $container.imagepullpolicy -}}
    {{- end -}}
  {{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "chart-app.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name for the postgres requirement.
*/}}
{{- define "chart-app.postgresql.fullname" -}}
{{- $postgresContext := dict "Values" .Values.postgresql "Release" .Release "Chart" (dict "Name" "postgresql") -}}
{{ template "postgresql.fullname" $postgresContext }}
{{- end -}}

{{- define "portspec" -}}
  {{- $portname := .portname -}}
  {{- range $port := .ports -}}
    {{- if eq $port.name $portname -}}
      {{- $port.port -}}
    {{- end -}}
  {{- end -}}
{{- end -}}

{{- define "hostName" -}}
  {{- range $host :=  .Values.ingress.hosts }}
      {{- printf "%s"  $host -}}
  {{- end -}}
{{- end -}}
