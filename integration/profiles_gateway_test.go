// +build integration

package integration

import (
	"github.com/bitly/go-simplejson"
	"net/http"
	"testing"
)

// TestCreateProfile_REST uses the REST gateway to create a new profile and
// ensure the JSON response matches is expected
// 1. Create a profile entry with CreateResource ()
// 2. Ensure the JSON fields are expected
// 3. Read created profile and ensure the JSON fields are expected
func TestCreateProfile_REST(t *testing.T) {
	dbTest.Reset(t)
	profile := map[string]interface{}{
		"name": "work", "notes": "profile for work-related topics",
	}
	id, createJSON := CreateResource(t, profile, "profiles")

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "profile id",
				json:   json.GetPath("result", "id"),
				expect: `"atlas-contacts-app/profiles/1"`,
			},
			{
				name:   "profile notes",
				json:   json.GetPath("result", "notes"),
				expect: `"profile for work-related topics"`,
			},
			{
				name:   "profile name",
				json:   json.GetPath("result", "name"),
				expect: `"work"`,
			},
			{
				name:   "success response",
				json:   json.GetPath("success"),
				expect: `{"code":"OK","status":200}`,
			},
		}
		return tests
	}
	for _, test := range testCases(&createJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
	readJSON := ReadResource(t, "profiles", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestReadProfile_REST uses the REST gateway to create a new profile and
// then read that profile
// 1. Create a profile entry with CreateResource () and receive id
// 2. Read created profile with ReadResource() by id and ensure the JSON fields are expected
func TestReadProfile_REST(t *testing.T) {
	dbTest.Reset(t)
	profile := map[string]interface{}{
		"name": "personal", "notes": "profile for personal matters",
	}
	id, _ := CreateResource(t, profile, "profiles")

	readJSON := ReadResource(t, "profiles", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "profile id",
				json:   json.GetPath("result", "id"),
				expect: `"atlas-contacts-app/profiles/1"`,
			},
			{
				name:   "profile notes",
				json:   json.GetPath("result", "notes"),
				expect: `"profile for personal matters"`,
			},
			{
				name:   "profile name",
				json:   json.GetPath("result", "name"),
				expect: `"personal"`,
			},
			{
				name:   "success response",
				json:   json.GetPath("success"),
				expect: `{"code":"OK","status":200}`,
			},
		}
		return tests
	}
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestUpdateProfile_REST uses the REST gateway to create a new profile and
// then read that profile
// 1. Create a profile entry with CreateResource () and receive id
// 2. Update the profile with UpdateResource () by id and body with new values
// 3. Ensure the response from UpdateResource () has expected fields
// 4. Read updated profile with ReadResource () by id and ensure the JSON fields are expected
func TestUpdateProfile_REST(t *testing.T) {
	dbTest.Reset(t)
	profile := map[string]interface{}{
		"name": "photography", "notes": "profile to show my photography portfolio",
	}
	id, _ := CreateResource(t, profile, "profiles")

	updated := map[string]interface{}{
		"name":  "woodworking",
		"notes": "profile to show my woodworking portfolio",
	}
	updateJSON := UpdateResource(t, updated, "profiles", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "profile id",
				json:   json.GetPath("result", "id"),
				expect: `"atlas-contacts-app/profiles/1"`,
			},
			{
				name:   "profile name",
				json:   json.GetPath("result", "name"),
				expect: `"woodworking"`,
			},
			{
				name:   "profile notes",
				json:   json.GetPath("result", "notes"),
				expect: `"profile to show my woodworking portfolio"`,
			},
			{
				name:   "success response",
				json:   json.GetPath("success"),
				expect: `{"code":"OK","status":200}`,
			},
		}
		return tests
	}
	for _, test := range testCases(&updateJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
	readJSON := ReadResource(t, "profiles", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestDeleteProfile_REST uses the REST gateway to create a new profile and
// ensures it can get deleted
// 1. Create a profile entry with CreateResource () and receive id
// 2. Delete created profile with DeleteResource () by id and ensure the JSON fields are expected
// 3. Attempt to read deleted profile with ReadResourceWithStatus() by id and ensure the JSON fields are expected
func TestDeleteProfile_REST(t *testing.T) {
	dbTest.Reset(t)
	profile := map[string]interface{}{
		"name": "school", "notes": "profile for academic-related content",
	}
	id, _ := CreateResource(t, profile, "profiles")
	deleteJSON := DeleteResource(t, "profiles", id)

	t.Run("success response", func(t *testing.T) {
		ValidateJSONSchema(t, &deleteJSON, `{"success":{"code":"OK","status":200}}`)
	})
	readJSON := ReadResourceWithStatus(t, "profiles", id, http.StatusNotFound)

	var tests = Tests{
		{
			name:   "check message",
			json:   readJSON.GetPath("error", "message"),
			expect: `"record not found"`,
		},
		{
			name:   "check status",
			json:   readJSON.GetPath("error", "status"),
			expect: `404`,
		},
		{
			name:   "check code",
			json:   readJSON.GetPath("error", "code"),
			expect: `"NOT_FOUND"`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestListProfiles_REST uses the REST gateway to create two profiles and
// ensures can read this profiles
// 1. Create two profiles with CreateResource ()
// 2. Read all the profiles with ListResource () and ensure the JSON fields are expected
func TestListProfiles_REST(t *testing.T) {
	dbTest.Reset(t)
	first := map[string]interface{}{
		"name": "cooking", "notes": "profile for cooking projects",
	}
	CreateResource(t, first, "profiles")
	second := map[string]interface{}{
		"name": "family", "notes": "profile for family information",
	}
	CreateResource(t, second, "profiles")

	listJSON := ListResource(t, "profiles")

	var tests = Tests{
		{
			name:   "first profile",
			json:   listJSON.Get("results").GetIndex(0),
			expect: `{"id":"atlas-contacts-app/profiles/1","name":"cooking","notes":"profile for cooking projects"}`,
		},
		{
			name:   "second profile",
			json:   listJSON.Get("results").GetIndex(1),
			expect: `{"id":"atlas-contacts-app/profiles/2","name":"family","notes":"profile for family information"}`,
		},
		{
			name:   "success response",
			json:   listJSON.GetPath("success"),
			expect: `{"code":"OK","status":200}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestPartiallyUpdateProfile_REST uses the REST gateway to create a new profile, update few fields in it and
// then read that profile
// 1. Create a profile entry with CreateResource () and receive id
// 2. Update the profile with PartiallyUpdateResource () by id and body with new fields
// 3. Ensure the response from PartiallyUpdateResource () has expected fields
// 4. Read updated profile with ReadResource () by id and ensure the JSON fields are expected
func TestPartiallyUpdateProfile_REST(t *testing.T) {
	dbTest.Reset(t)
	//create a new profile
	profile := map[string]interface{}{
		"name":  "hobby",
		"notes": "profile for my hobby",
	}
	id, _ := CreateResource(t, profile, "profiles")

	updatedFields := map[string]interface{}{
		"name": "leisure",
	}
	updateJSON := PartiallyUpdateResource(t, updatedFields, "profiles", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "profile id",
				json:   json.GetPath("result", "id"),
				expect: `"` + ApplicationID + "/profiles/" + id + `"`,
			},
			{
				name:   "profile name",
				json:   json.GetPath("result", "name"),
				expect: `"leisure"`,
			},
			{
				name:   "profile notes",
				json:   json.GetPath("result", "notes"),
				expect: `"profile for my hobby"`,
			},
			{
				name:   "success response",
				json:   json.GetPath("success"),
				expect: `{"code":"OK","status":200}`,
			},
		}
		return tests
	}
	for _, test := range testCases(&updateJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
	readJSON := ReadResource(t, "profiles", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

//					Negative tests

// TestReadNonExistProfile_REST attempts to read a non-existent profile from the application
// 1. Attempt to get the non-existent profile with incorrect ID by ReadResourceWithStatus ()
// 2. Ensure the response from ReadResourceWithStatus () has expected fields and detailed error message
func TestReadNonExistProfile_REST(t *testing.T) {
	dbTest.Reset(t)
	id := "1"
	readJSON := ReadResourceWithStatus(t, "profiles", id, http.StatusNotFound)

	var tests = Tests{
		{
			name:   "check message",
			json:   readJSON.GetPath("error", "message"),
			expect: `"record not found"`,
		},
		{
			name:   "check status",
			json:   readJSON.GetPath("error", "status"),
			expect: `404`,
		},
		{
			name:   "check code",
			json:   readJSON.GetPath("error", "code"),
			expect: `"NOT_FOUND"`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestPartiallyUpdateNonExistProfile_REST attempts to update field for a non-existent profile
// 1. Attempt to update field with incorrect ID for the non-existent profile by PartiallyUpdateResourceWithStatus ()
// 2. Ensure the response from PartiallyUpdateResourceWithStatus () has expected fields and detailed error message
func TestPartiallyUpdateNonExistProfile_REST(t *testing.T) {
	dbTest.Reset(t)
	updatedFields := map[string]interface{}{
		"name": "Super profile",
	}
	id := "1"
	updateJSON := PartiallyUpdateResourceWithStatus(t, updatedFields, "profiles", id, http.StatusNotFound)

	var tests = Tests{
		{
			name:   "check message",
			json:   updateJSON.GetPath("error", "message"),
			expect: `"record not found"`,
		},
		{
			name:   "check status",
			json:   updateJSON.GetPath("error", "status"),
			expect: `404`,
		},
		{
			name:   "check code",
			json:   updateJSON.GetPath("error", "code"),
			expect: `"NOT_FOUND"`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}
