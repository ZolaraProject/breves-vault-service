{{- define "brevesvaultservice.name" -}}
{{- default "breves-vault-service" .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Helm required labels */}}
{{- define "brevesvaultservice.labels" -}}
heritage: {{ .Release.Service }}
release: {{ .Release.Name }}
chart: {{ .Chart.Name }}
app: "{{ template "brevesvaultservice.name" . }}"
{{- end -}}

{{/* matchLabels */}}
{{- define "brevesvaultservice.matchLabels" -}}
release: {{ .Release.Name }}
app: "{{ template "brevesvaultservice.name" . }}"
{{- end -}}
