// +build integration

package integration

import (
	"github.com/bitly/go-simplejson"
	"net/http"
	"testing"
)

// TestCreateGroup_REST uses the REST gateway to create a new group and
// ensure the JSON response matches is expected
// 1. Create a group entry with CreateResource ()
// 2. Ensure the JSON fields are expected
// 3. Read created group and ensure the JSON fields are expected
func TestCreateGroup_REST(t *testing.T) {
	dbTest.Reset(t)
	//create a new profile
	profile := map[string]interface{}{
		"name":  "work",
		"notes": "profile for work-related topics",
	}
	idProfile, _ := CreateResource(t, profile, "profiles")
	//create a new group
	group := map[string]interface{}{
		"name":       "The Beatles",
		"notes":      `People, who like The Beatles`,
		"profile_id": idProfile,
	}
	id, createGroupJSON := CreateResource(t, group, "groups")

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{

			{
				name:   "group name",
				json:   json.GetPath("result", "name"),
				expect: `"The Beatles"`,
			},
			{
				name:   "group note",
				json:   json.GetPath("result", "notes"),
				expect: `"People, who like The Beatles"`,
			},
			{
				name:   "group profile_id",
				json:   json.GetPath("result", "profile_id"),
				expect: `"` + ApplicationID + "/profiles/" + idProfile + `"`,
			},
			{
				name:   "success response",
				json:   json.GetPath("success"),
				expect: `{"code":"OK","status":200}`,
			},
		}
		return tests
	}
	for _, test := range testCases(&createGroupJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
	readJSON := ReadResource(t, "groups", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestReadGroup_REST uses the REST gateway to create a new group and
// then read that group
// 1. Create a group entry with CreateResource () and receive id
// 2. Read created group with ReadResource () by id and ensure the JSON fields are expected
func TestReadGroup_REST(t *testing.T) {
	dbTest.Reset(t)
	//create a new profile
	profile := map[string]interface{}{
		"name":  "personal",
		"notes": "profile for personal matters",
	}
	idProfile, _ := CreateResource(t, profile, "profiles")
	//create a new group
	group := map[string]interface{}{
		"name":       "Led Zeppelin",
		"notes":      `People, who like Led Zeppelin`,
		"profile_id": idProfile,
	}
	idGroup, _ := CreateResource(t, group, "groups")
	readJSON := ReadResource(t, "groups", idGroup)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "group id",
				json:   json.GetPath("result", "id"),
				expect: `"` + ApplicationID + "/groups/" + idGroup + `"`,
			},
			{
				name:   "group name",
				json:   json.GetPath("result", "name"),
				expect: `"Led Zeppelin"`,
			},
			{
				name:   "group note",
				json:   json.GetPath("result", "notes"),
				expect: `"People, who like Led Zeppelin"`,
			},
			{
				name:   "group profile_id",
				json:   json.GetPath("result", "profile_id"),
				expect: `"` + ApplicationID + "/profiles/" + idProfile + `"`,
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

// TestDeleteGroup_REST uses the REST gateway to create a new group and
// ensures it can get deleted
// 1. Create a group entry with CreateResource () and receive id
// 2. Delete created group with DeleteResource () by id and ensure the JSON fields are expected
// 3. Attempt to read deleted group with ReadResourceWithStatus () by id and ensure the JSON fields are expected
func TestDeleteGroup_REST(t *testing.T) {
	dbTest.Reset(t)
	//create group
	group := map[string]interface{}{
		"name":  "The Rolling Stones",
		"notes": `People, who like The Rolling Stones`,
	}
	id, _ := CreateResource(t, group, "groups")
	//delete group
	deleteJSON := DeleteResource(t, "groups", id)

	t.Run("success response", func(t *testing.T) {
		ValidateJSONSchema(t, &deleteJSON, `{"success":{"code":"OK","status":200}}`)
	})
	readJSON := ReadResourceWithStatus(t, "groups", id, http.StatusNotFound)
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

// TestUpdateGroup_REST uses the REST gateway to create a new group and
// then read that group
// 1. Create a group entry with CreateResource () and receive id
// 2. Update the group with UpdateResource () by id and body with new values
// 3. Ensure the response from UpdateResource () has expected fields
// 4. Read updated group with ReadResource () by id and ensure the JSON fields are expected
func TestUpdateGroup_REST(t *testing.T) {
	dbTest.Reset(t)
	//create group
	group := map[string]interface{}{
		"name":  "Pink Floyd",
		"notes": `People, who like Pink Floyd`,
	}
	id, _ := CreateResource(t, group, "groups")
	//update group
	updated := map[string]interface{}{
		"name":  "Queen",
		"notes": `People, who like Queen`,
	}
	updateJSON := UpdateResource(t, updated, "groups", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "group id",
				json:   json.GetPath("result", "id"),
				expect: `"` + ApplicationID + "/groups/" + id + `"`,
			},
			{
				name:   "group name",
				json:   json.GetPath("result", "name"),
				expect: `"Queen"`,
			},
			{
				name:   "group notes",
				json:   json.GetPath("result", "notes"),
				expect: `"People, who like Queen"`,
			},
			{
				name:   "group profile_id",
				json:   json.GetPath("result", "profile_id"),
				expect: `null`,
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
	readJSON := ReadResource(t, "groups", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}

}

// TestListGroups_REST uses the REST gateway to create two groups and
// ensures can read this groups
// 1. Create two groups with CreateResource ()
// 2. Read all the groups with ListResource () and ensure the JSON fields are expected
func TestListGroups_REST(t *testing.T) {
	dbTest.Reset(t)
	first := map[string]interface{}{
		"name":  "Nirvana",
		"notes": `People, who like Nirvana`,
	}
	second := map[string]interface{}{
		"name":  "The Doors",
		"notes": `People, who like The Doors`,
	}
	CreateResource(t, first, "groups")
	CreateResource(t, second, "groups")
	listJSON := ListResource(t, "groups")

	var tests = Tests{
		{
			name:   "first group",
			json:   listJSON.Get("results").GetIndex(0),
			expect: `{"id":"atlas-contacts-app/groups/1","name":"Nirvana","notes":"People, who like Nirvana"}`,
		},
		{
			name:   "second group",
			json:   listJSON.Get("results").GetIndex(1),
			expect: `{"id":"atlas-contacts-app/groups/2","name":"The Doors","notes":"People, who like The Doors"}`,
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

// TestPartiallyUpdateGroup_REST uses the REST gateway to create a new group, update few fields in it and
// then read that group
// 1. Create a group entry with CreateResource () and receive id
// 2. Update the group with PartiallyUpdateResource () by id and body with new fields
// 3. Ensure the response from PartiallyUpdateResource () has expected fields
// 4. Read updated group with ReadResource () by id and ensure the JSON fields are expected
func TestPartiallyUpdateGroup_REST(t *testing.T) {
	dbTest.Reset(t)
	//create a new first profile
	profile1 := map[string]interface{}{
		"name":  "private profile",
		"notes": "profile for private date",
	}
	idProfile1, _ := CreateResource(t, profile1, "profiles")
	//create a new second profile
	profile2 := map[string]interface{}{
		"name":  "public profile",
		"notes": "profile for public date",
	}
	idProfile2, _ := CreateResource(t, profile2, "profiles")
	//create a new group
	group := map[string]interface{}{
		"name":       "Guns N’Roses",
		"notes":      `People, who like Guns N’Roses`,
		"profile_id": idProfile1,
	}
	id, _ := CreateResource(t, group, "groups")

	updatedFields := map[string]interface{}{
		"name":       "U2",
		"profile_id": idProfile2,
	}
	updateJSON := PartiallyUpdateResource(t, updatedFields, "groups", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "group id",
				json:   json.GetPath("result", "id"),
				expect: `"` + ApplicationID + "/groups/" + id + `"`,
			},
			{
				name:   "group name",
				json:   json.GetPath("result", "name"),
				expect: `"U2"`,
			},
			{
				name:   "group notes",
				json:   json.GetPath("result", "notes"),
				expect: `"People, who like Guns N’Roses"`,
			},
			{
				name:   "group profile_id",
				json:   json.GetPath("result", "profile_id"),
				expect: `"` + ApplicationID + "/profiles/" + idProfile2 + `"`,
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
	readJSON := ReadResource(t, "groups", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

//				Negative tests

// TestReadNonExistGroup_REST attempts to read a non-existent group from the application
// 1. Attempt to get the non-existent group with incorrect ID by ReadResourceWithStatus ()
// 2. Ensure the response from ReadResourceWithStatus () has expected fields and detailed error message
func TestReadNonExistGroup_REST(t *testing.T) {
	dbTest.Reset(t)
	id := "1"
	readJSON := ReadResourceWithStatus(t, "groups", id, http.StatusNotFound)

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

// TestPartiallyUpdateNonExistGroup_REST attempts to update field for a non-existent group
// 1. Attempt to update field with incorrect ID for the non-existent group by PartiallyUpdateResourceWithStatus ()
// 2. Ensure the response from PartiallyUpdateResourceWithStatus () has expected fields and detailed error message
func TestPartiallyUpdateNonExistGroup_REST(t *testing.T) {
	dbTest.Reset(t)
	updatedFields := map[string]interface{}{
		"name": "Fanta",
	}
	id := "1"
	updateJSON := PartiallyUpdateResourceWithStatus(t, updatedFields, "groups", id, http.StatusNotFound)

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
