{{/*
Expand the name of the chart.
*/}}
{{- define "cmdb.name" -}}
{{- default .Chart.Name .Values.cmdb.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "cmdb.fullname" -}}
{{- if .Values.cmdb.fullnameOverride }}
{{- .Values.cmdb.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.cmdb.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "cmdb.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "cmdb.labels" -}}
helm.sh/chart: {{ include "cmdb.chart" . }}
{{ include "cmdb.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "cmdb.selectorLabels" -}}
app.kubernetes.io/name: {{ include "cmdb.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "cmdb.serviceAccountName" -}}
{{- if .Values.cmdb.serviceAccount.create }}
{{- default (include "cmdb.fullname" .) .Values.cmdb.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.cmdb.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "sslmode" -}}
{{- if .Values.postgresql.tls.enabled }}
{{- print "enable" }}
{{- else }}
{{- print "disable" }}
{{- end }}
{{- end }}

{{- define "portspec" -}}
  {{- $portname := .portname -}}
  {{- range $port := .ports -}}
    {{- if eq $port.name $portname -}}
      {{- $port.port -}}
    {{- end -}}
  {{- end -}}
{{- end -}}
