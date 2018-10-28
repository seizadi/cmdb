// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/seizadi/cmdb/pkg/pb/contacts.proto

package pb // import "github.com/seizadi/cmdb/pkg/pb"

import options "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import query "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/infobloxopen/protoc-gen-atlas-validate/options"
import _ "google.golang.org/genproto/protobuf/field_mask"

// Reference imports to suppress errors if they are not otherwise used.

var ContactsMethodsRequireFilteringValidation = map[string]map[string]options.FilteringOption{
	"/api.contacts.Profiles/List": map[string]options.FilteringOption{
		"id":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":  options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"notes": options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH}, ValueType: options.QueryValidate_STRING},
	},
	"/api.contacts.Groups/List": map[string]options.FilteringOption{
		"id":         options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":       options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"notes":      options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH}, ValueType: options.QueryValidate_STRING},
		"profile_id": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.contacts.Groups/ListByProfile": map[string]options.FilteringOption{
		"id":         options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":       options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"notes":      options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH}, ValueType: options.QueryValidate_STRING},
		"profile_id": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.contacts.Contacts/List": map[string]options.FilteringOption{
		"id":                   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"first_name":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"middle_name":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"last_name":            options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"primary_email":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"notes":                options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH}, ValueType: options.QueryValidate_STRING},
		"home_address.address": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.city":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.state":   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.zip":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.country": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.address": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.city":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.state":   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.zip":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.country": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"profile_id":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.contacts.Contacts/ListByProfile": map[string]options.FilteringOption{
		"id":                   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"first_name":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"middle_name":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"last_name":            options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"primary_email":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"notes":                options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH}, ValueType: options.QueryValidate_STRING},
		"home_address.address": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.city":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.state":   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.zip":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.country": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.address": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.city":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.state":   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.zip":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.country": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"profile_id":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.contacts.Contacts/ListByGroup": map[string]options.FilteringOption{
		"id":                   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"first_name":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"middle_name":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"last_name":            options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"primary_email":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"notes":                options.FilteringOption{Deny: []options.QueryValidate_FilterOperator{options.QueryValidate_MATCH}, ValueType: options.QueryValidate_STRING},
		"home_address.address": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.city":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.state":   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.zip":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"home_address.country": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.address": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.city":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.state":   options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.zip":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"work_address.country": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"profile_id":           options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
}
var ContactsMethodsRequireSortingValidation = map[string][]string{
	"/api.contacts.Profiles/List": []string{
		"id",
		"name",
	},
	"/api.contacts.Groups/List": []string{
		"id",
		"name",
		"profile_id",
	},
	"/api.contacts.Groups/ListByProfile": []string{
		"id",
		"name",
		"profile_id",
	},
	"/api.contacts.Contacts/List": []string{
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"primary_email",
		"home_address.address",
		"home_address.city",
		"home_address.state",
		"home_address.zip",
		"home_address.country",
		"work_address.address",
		"work_address.city",
		"work_address.state",
		"work_address.zip",
		"work_address.country",
		"profile_id",
	},
	"/api.contacts.Contacts/ListByProfile": []string{
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"primary_email",
		"home_address.address",
		"home_address.city",
		"home_address.state",
		"home_address.zip",
		"home_address.country",
		"work_address.address",
		"work_address.city",
		"work_address.state",
		"work_address.zip",
		"work_address.country",
		"profile_id",
	},
	"/api.contacts.Contacts/ListByGroup": []string{
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"primary_email",
		"home_address.address",
		"home_address.city",
		"home_address.state",
		"home_address.zip",
		"home_address.country",
		"work_address.address",
		"work_address.city",
		"work_address.state",
		"work_address.zip",
		"work_address.country",
		"profile_id",
	},
}
var ContactsMethodsRequireFieldSelectionValidation = map[string][]string{
	"/api.contacts.Profiles/Read": {
		"id",
		"name",
		"notes",
	},
	"/api.contacts.Profiles/List": {
		"id",
		"name",
		"notes",
	},
	"/api.contacts.Groups/Read": {
		"id",
		"name",
		"notes",
		"profile_id",
	},
	"/api.contacts.Groups/List": {
		"id",
		"name",
		"notes",
		"profile_id",
	},
	"/api.contacts.Groups/ListByProfile": {
		"id",
		"name",
		"notes",
		"profile_id",
	},
	"/api.contacts.Contacts/Read": {
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"primary_email",
		"notes",
		"emails.address",
		"emails",
		"home_address.address",
		"home_address.city",
		"home_address.state",
		"home_address.zip",
		"home_address.country",
		"home_address",
		"work_address.address",
		"work_address.city",
		"work_address.state",
		"work_address.zip",
		"work_address.country",
		"work_address",
		"profile_id",
		"nicknames.value",
		"nicknames",
		"groups",
	},
	"/api.contacts.Contacts/List": {
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"primary_email",
		"notes",
		"emails.address",
		"emails",
		"home_address.address",
		"home_address.city",
		"home_address.state",
		"home_address.zip",
		"home_address.country",
		"home_address",
		"work_address.address",
		"work_address.city",
		"work_address.state",
		"work_address.zip",
		"work_address.country",
		"work_address",
		"profile_id",
		"nicknames.value",
		"nicknames",
		"groups",
	},
	"/api.contacts.Contacts/ListByProfile": {
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"primary_email",
		"notes",
		"emails.address",
		"emails",
		"home_address.address",
		"home_address.city",
		"home_address.state",
		"home_address.zip",
		"home_address.country",
		"home_address",
		"work_address.address",
		"work_address.city",
		"work_address.state",
		"work_address.zip",
		"work_address.country",
		"work_address",
		"profile_id",
		"nicknames.value",
		"nicknames",
		"groups",
	},
	"/api.contacts.Contacts/ListByGroup": {
		"id",
		"first_name",
		"middle_name",
		"last_name",
		"primary_email",
		"notes",
		"emails.address",
		"emails",
		"home_address.address",
		"home_address.city",
		"home_address.state",
		"home_address.zip",
		"home_address.country",
		"home_address",
		"work_address.address",
		"work_address.city",
		"work_address.state",
		"work_address.zip",
		"work_address.country",
		"work_address",
		"profile_id",
		"nicknames.value",
		"nicknames",
		"groups",
	},
}

func ContactsValidateFiltering(methodName string, f *query.Filtering) error {
	info, ok := ContactsMethodsRequireFilteringValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFiltering(f, info)
}
func ContactsValidateSorting(methodName string, s *query.Sorting) error {
	info, ok := ContactsMethodsRequireSortingValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateSorting(s, info)
}
func ContactsValidateFieldSelection(methodName string, s *query.FieldSelection) error {
	info, ok := ContactsMethodsRequireFieldSelectionValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFieldSelection(s, info)
}