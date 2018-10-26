// +build integration

package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bitly/go-simplejson"
)

type Tests []struct {
	name   string
	json   *simplejson.Json
	expect string
}

// TestCreateContact_REST uses the REST gateway to create a new contact and
// ensure the JSON response matches is expected
// 1. Create a contact entry with CreateResource ()
// 2. Ensure the JSON fields are expected
// 3. Read created contact and ensure the JSON fields are expected
func TestCreateContact_REST(t *testing.T) {
	dbTest.Reset(t)
	contact := map[string]interface{}{
		"first_name":    "Steven",
		"middle_name":   "James",
		"last_name":     "McKubernetes",
		"primary_email": "test@test.com",
		"notes":         "set sail at sunrise",
	}
	id, createJSON := CreateResource(t, contact, "contacts")

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "contact first name",
				json:   json.GetPath("result", "first_name"),
				expect: `"Steven"`,
			},
			{
				name:   "contact middle name",
				json:   json.GetPath("result", "middle_name"),
				expect: `"James"`,
			},
			{
				name:   "contact last",
				json:   json.GetPath("result", "last_name"),
				expect: `"McKubernetes"`,
			},
			{
				name:   "contact notes",
				json:   json.GetPath("result", "notes"),
				expect: `"set sail at sunrise"`,
			},
			{
				name:   "contact primary email",
				json:   json.GetPath("result", "primary_email"),
				expect: `"test@test.com"`,
			},
			{
				name:   "contact email list",
				json:   json.GetPath("result", "emails"),
				expect: `[{"address":"test@test.com"}]`,
			},
			{
				name:   "contact home address list",
				json:   json.GetPath("result", "home_address"),
				expect: `null`,
			},
			{
				name:   "contact work address list",
				json:   json.GetPath("result", "work_address"),
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
	for _, test := range testCases(&createJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
	readJSON := ReadResource(t, "contacts", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestReadContact_REST uses the REST gateway to create a new contact and
// then read that contact
// 1. Create a contact entry with CreateResource () and receive id
// 2. Read created contact with ReadResource () by id and ensure the JSON fields are expected
func TestReadContact_REST(t *testing.T) {
	dbTest.Reset(t)
	contact := map[string]interface{}{
		"first_name":    "Wilfred",
		"middle_name":   "Wallace",
		"last_name":     "O'Docker",
		"primary_email": "test@test.com",
		"notes":         "build the containers at dusk",
	}
	id, _ := CreateResource(t, contact, "contacts")
	readJSON := ReadResource(t, "contacts", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "contact first name",
				json:   json.GetPath("result", "first_name"),
				expect: `"Wilfred"`,
			},
			{
				name:   "contact middle name",
				json:   json.GetPath("result", "middle_name"),
				expect: `"Wallace"`,
			},
			{
				name:   "contact last",
				json:   json.GetPath("result", "last_name"),
				expect: `"O'Docker"`,
			},
			{
				name:   "contact notes",
				json:   json.GetPath("result", "notes"),
				expect: `"build the containers at dusk"`,
			},
			{
				name:   "contact primary email",
				json:   json.GetPath("result", "primary_email"),
				expect: `"test@test.com"`,
			},
			{
				name:   "contact email list",
				json:   json.GetPath("result", "emails"),
				expect: `[{"address":"test@test.com"}]`,
			},
			{
				name:   "contact home_address list",
				json:   json.GetPath("result", "home_address"),
				expect: `null`,
			},
			{
				name:   "contact work_address list",
				json:   json.GetPath("result", "work_address"),
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
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestInvalidEmail_REST attempts to create a new contact that has an invalid
// email address.
// 1. Use the REST API to create a new contact with an invalid email
// 2. Ensure the HTTP response status code is 400
// 3. Ensure the HTTP response has a detailed error message
func TestInvalidEmail_REST(t *testing.T) {
	t.Skip("Temporarily skipping the test...")
	dbTest.Reset(t)
	contact := map[string]interface{}{
		"primary_email": "invalid-email-address",
	}
	resDelete, err := MakeRequestWithDefaults(
		http.MethodPost, "http://localhost:8080/v1/contacts",
		contact,
	)
	if err != nil {
		t.Fatalf("unable to create contact %v", err)
	}
	ValidateResponseCode(t, resDelete, http.StatusBadRequest)
	deleteJSON, err := simplejson.NewFromReader(resDelete.Body)
	if err != nil {
		t.Fatalf("unable to unmarshal response json: %v", err)
	}
	var tests = []struct {
		name   string
		json   *simplejson.Json
		expect string
	}{
		{
			name:   "check response code",
			json:   deleteJSON.GetPath("error", "code"),
			expect: `"INVALID_ARGUMENT"`,
		},
		{
			name:   "check http status",
			json:   deleteJSON.GetPath("error", "status"),
			expect: `400`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestDeleteContact_REST uses the REST gateway to create a new contact and
// ensures it can get deleted
// 1. Create a contact entry with CreateResource () and receive id
// 2. Delete created contact with DeleteResource () by id and ensure the JSON fields are expected
// 3. Attempt to read deleted contact with ReadResourceWithStatus() by id and ensure the JSON fields are expected
func TestDeleteContact_REST(t *testing.T) {
	dbTest.Reset(t)
	contact := map[string]interface{}{
		"primary_email": "test@test.com",
	}
	id, _ := CreateResource(t, contact, "contacts")
	deleteJSON := DeleteResource(t, "contacts", id)

	t.Run("success response", func(t *testing.T) {
		ValidateJSONSchema(t, &deleteJSON, `{"success":{"code":"OK","status":200}}`)
	})
	readJSON := ReadResourceWithStatus(t, "contacts", id, http.StatusNotFound)
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

// TestUpdateContacts_REST uses the REST gateway to create a new contact and
// then read that contact
// 1. Create a contact entry with CreateResource () and receive id
// 2. Update the contact with UpdateResource () by id and body with new values
// 3. Ensure the response from UpdateResource () has expected fields
// 4. Read updated contact with ReadResource () by id and ensure the JSON fields are expected
func TestUpdateContacts_REST(t *testing.T) {
	dbTest.Reset(t)
	contact := map[string]interface{}{
		"first_name":    "Agatha",
		"middle_name":   "Mary Clarissa",
		"last_name":     "Christie",
		"primary_email": "Agatha_Christie@test.com",
		"notes":         "author of The Murder of Roger Ackroyd",
	}
	id, _ := CreateResource(t, contact, "contacts")

	updated := map[string]interface{}{
		"first_name":    "William",
		"middle_name":   "John",
		"last_name":     "Shakespeare",
		"primary_email": "William_Shakespeare@test.com",
		"notes":         "author of Romeo and Juliet",
	}
	updateJSON := UpdateResource(t, updated, "contacts", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "contact first name",
				json:   json.GetPath("result", "first_name"),
				expect: `"William"`,
			},
			{
				name:   "contact middle name",
				json:   json.GetPath("result", "middle_name"),
				expect: `"John"`,
			},
			{
				name:   "contact last",
				json:   json.GetPath("result", "last_name"),
				expect: `"Shakespeare"`,
			},
			{
				name:   "contact notes",
				json:   json.GetPath("result", "notes"),
				expect: `"author of Romeo and Juliet"`,
			},
			{
				name:   "contact primary email",
				json:   json.GetPath("result", "primary_email"),
				expect: `"William_Shakespeare@test.com"`,
			},
			{
				name:   "contact email list",
				json:   json.GetPath("result", "emails"),
				expect: `[{"address":"William_Shakespeare@test.com"}]`,
			},
			{
				name:   "contact home address list",
				json:   json.GetPath("result", "home_address"),
				expect: `null`,
			},
			{
				name:   "contact work address list",
				json:   json.GetPath("result", "work_address"),
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
	readJSON := ReadResource(t, "contacts", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestListContacts_REST uses the REST gateway to create two contacts and
// ensures can read this contacts
// 1. Create two contacts with CreateResource ()
// 2. Read all the contacts with ListResource () and ensure the JSON fields are expected
func TestListContacts_REST(t *testing.T) {
	dbTest.Reset(t)
	first := map[string]interface{}{
		"first_name":    "Alexander",
		"middle_name":   "Sergeyevich",
		"last_name":     "Pushkin",
		"primary_email": "Alexander_Pushkin@test.com",
		"notes":         "author of Eugene Onegin",
	}
	CreateResource(t, first, "contacts")
	second := map[string]interface{}{
		"first_name":    "Leo",
		"middle_name":   "Nikolayevich",
		"last_name":     "Tolstoy",
		"primary_email": "Leo_Tolstoy@test.com",
		"notes":         "author of War and Peace",
	}
	CreateResource(t, second, "contacts")

	listJSON := ListResource(t, "contacts")

	var tests = Tests{
		{
			name:   "first contact",
			json:   listJSON.Get("results").GetIndex(0),
			expect: `{"emails":[{"address":"Alexander_Pushkin@test.com"}],"first_name":"Alexander","id":"atlas-contacts-app/contacts/1","last_name":"Pushkin","middle_name":"Sergeyevich","notes":"author of Eugene Onegin","primary_email":"Alexander_Pushkin@test.com"}`,
		},
		{
			name:   "second contact",
			json:   listJSON.Get("results").GetIndex(1),
			expect: `{"emails":[{"address":"Leo_Tolstoy@test.com"}],"first_name":"Leo","id":"atlas-contacts-app/contacts/2","last_name":"Tolstoy","middle_name":"Nikolayevich","notes":"author of War and Peace","primary_email":"Leo_Tolstoy@test.com"}`,
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

// TestPartiallyUpdateContacts_REST uses the REST gateway to create a new contact, update few fields in it and
// then read that contact
// 1. Create a contact entry with CreateResource () and receive id
// 2. Update the contact with PartiallyUpdateResource () by id and body with new fields
// 3. Ensure the response from PartiallyUpdateResource () has expected fields
// 4. Read updated contact with ReadResource () by id and ensure the JSON fields are expected
func TestPartiallyUpdateContacts_REST(t *testing.T) {
	dbTest.Reset(t)
	//create a new contact
	contact := map[string]interface{}{
		"first_name":    "Joanne",
		"middle_name":   "Kathleen",
		"last_name":     "Rowling",
		"primary_email": "Joanne_Rowling@test.com",
		"notes":         "author of Harry Potter",
	}
	id, _ := CreateResource(t, contact, "contacts")

	updatedFields := map[string]interface{}{
		"first_name":  "J.",
		"middle_name": "K.",
	}
	updateJSON := PartiallyUpdateResource(t, updatedFields, "contacts", id)

	testCases := func(json *simplejson.Json) Tests {
		var tests = Tests{
			{
				name:   "contact id",
				json:   json.GetPath("result", "id"),
				expect: `"` + ApplicationID + "/contacts/" + id + `"`,
			},
			{
				name:   "contact first name",
				json:   json.GetPath("result", "first_name"),
				expect: `"J."`,
			},
			{
				name:   "contact middle name",
				json:   json.GetPath("result", "middle_name"),
				expect: `"K."`,
			},
			{
				name:   "contact last name",
				json:   json.GetPath("result", "last_name"),
				expect: `"Rowling"`,
			},
			{
				name:   "contact notes",
				json:   json.GetPath("result", "notes"),
				expect: `"author of Harry Potter"`,
			},
			{
				name:   "contact primary email",
				json:   json.GetPath("result", "primary_email"),
				expect: `"Joanne_Rowling@test.com"`,
			},
			{
				name:   "contact home address list",
				json:   json.GetPath("result", "home_address"),
				expect: `null`,
			},
			{
				name:   "contact work address list",
				json:   json.GetPath("result", "work_address"),
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
	readJSON := ReadResource(t, "contacts", id)
	for _, test := range testCases(&readJSON) {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

// TestSendSMSContacts_REST uses the REST gateway to send SMS to contact and
// ensure the status code is OK
// 1. Create contact with CreateResource ()
// 2. Send SMS by GET request to /contacts/{id}/sms
// 3. Check the status code to ensure sent sms was ok
func TestSendSMSContacts_REST(t *testing.T) {
	dbTest.Reset(t)
	//create a new contacts
	contact := map[string]interface{}{
		"first_name":    "Stephen",
		"middle_name":   "Edwin",
		"last_name":     "King",
		"primary_email": "Stephen_King.com",
		"notes":         "author of The Shining",
	}
	id, _ := CreateResource(t, contact, "contacts")

	//sending sms
	messageSMS := map[string]interface{}{
		"message": "Hello",
	}
	resSendSMS, err := MakeRequestWithDefaults(
		http.MethodPost,
		fmt.Sprintf("http://localhost:8080/v1/contacts/%s/sms", id),
		messageSMS,
	)
	if err != nil {
		t.Fatalf("unable to send sms to contact: %v", err)
	}
	smsJSON, err := simplejson.NewFromReader(resSendSMS.Body)
	if err != nil {
		t.Fatalf("unable to unmarshal json response: %v", err)
	}
	var tests = []struct {
		name   string
		json   *simplejson.Json
		expect string
	}{
		{
			name:   "success response",
			json:   smsJSON.GetPath("success"),
			expect: `{"code":"OK","status":200}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ValidateJSONSchema(t, test.json, test.expect)
		})
	}
}

func TestHomeWorkAddresses_REST(t *testing.T) {
	dbTest.Reset(t)

	home_address := map[string]string{
		"city":    "Portsmouth",
		"state":   "United Kingdom",
		"country": "England",
	}
	work_address := map[string]string{
		"city":    "London",
		"state":   "United Kingdom",
		"country": "England",
	}
	first := map[string]interface{}{
		"first_name":    "Charles",
		"middle_name":   "John",
		"last_name":     "Dickens",
		"primary_email": "Charles_Dickens@test.com",
		"notes":         "author of Oliver Twist",
		"home_address":  home_address,
		"work_address":  work_address,
	}
	CreateResource(t, first, "contacts")

	home_address2 := map[string]string{
		"city":    "Moscow",
		"state":   "-",
		"country": "Russia",
	}
	work_address2 := map[string]string{
		"city":    "Stavropol",
		"state":   "-",
		"country": "Russia",
	}
	second := map[string]interface{}{
		"first_name":    "Mikhail",
		"middle_name":   "Yuryevich",
		"last_name":     "Lermontov",
		"primary_email": "Mikhail_Lermontov@test.com",
		"notes":         "author of Mtsyri",
		"home_address":  home_address2,
		"work_address":  work_address2,
	}
	CreateResource(t, second, "contacts")

	home_address3 := map[string]string{
		"city":    "Taganrog",
		"state":   "-",
		"country": "Russia",
	}
	work_address3 := map[string]string{
		"city":    "St Petersburg",
		"state":   "-",
		"country": "Russia",
	}
	third := map[string]interface{}{
		"first_name":    "Anton",
		"middle_name":   "Pavlovich",
		"last_name":     "Chekhov",
		"primary_email": "Anton_Chekhov@test.com",
		"notes":         "author of Kashtanka",
		"home_address":  home_address3,
		"work_address":  work_address3,
	}
	CreateResource(t, third, "contacts")

	home_address4 := map[string]string{
		"city":    "Moscow",
		"state":   "-",
		"country": "Russia",
	}
	work_address4 := map[string]string{
		"city":    "St Petersburg",
		"state":   "-",
		"country": "Russia",
	}
	fourth := map[string]interface{}{
		"first_name":    "Fyodor",
		"middle_name":   "Mikhailovich",
		"last_name":     "Dostoevsky",
		"primary_email": "Fyodor_Dostoevsky@test.com",
		"notes":         "author of Crime and Punishment",
		"home_address":  home_address4,
		"work_address":  work_address4,
	}
	CreateResource(t, fourth, "contacts")
	listJSON := ListResource(t, `contacts?_filter=home_address.city=='Moscow'%20or%20work_address.city=='St%20Petersburg'%20and%20home_address.country=='Russia'&_order_by=home_address.city,work_address.city`)
	var tests = Tests{
		{
			name:   "with filtering and sorting",
			json:   listJSON.Get("results").GetIndex(0),
			expect: `{"emails":[{"address":"Mikhail_Lermontov@test.com"}],"first_name":"Mikhail","home_address":{"city":"Moscow","country":"Russia","state":"-"},"id":"atlas-contacts-app/contacts/2","last_name":"Lermontov","middle_name":"Yuryevich","notes":"author of Mtsyri","primary_email":"Mikhail_Lermontov@test.com","work_address":{"city":"Stavropol","country":"Russia","state":"-"}}`,
		},
		{
			name:   "with filtering and sorting",
			json:   listJSON.Get("results").GetIndex(1),
			expect: `{"emails":[{"address":"Fyodor_Dostoevsky@test.com"}],"first_name":"Fyodor","home_address":{"city":"Moscow","country":"Russia","state":"-"},"id":"atlas-contacts-app/contacts/4","last_name":"Dostoevsky","middle_name":"Mikhailovich","notes":"author of Crime and Punishment","primary_email":"Fyodor_Dostoevsky@test.com","work_address":{"city":"St Petersburg","country":"Russia","state":"-"}}`,
		},
		{
			name:   "with filtering and sorting",
			json:   listJSON.Get("results").GetIndex(2),
			expect: `{"emails":[{"address":"Anton_Chekhov@test.com"}],"first_name":"Anton","home_address":{"city":"Taganrog","country":"Russia","state":"-"},"id":"atlas-contacts-app/contacts/3","last_name":"Chekhov","middle_name":"Pavlovich","notes":"author of Kashtanka","primary_email":"Anton_Chekhov@test.com","work_address":{"city":"St Petersburg","country":"Russia","state":"-"}}`,
		},
		{
			name:   "with filtering and sorting",
			json:   listJSON.Get("results").GetIndex(3),
			expect: `null`,
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

// ValidateResponseCode checks the http status of a given request and will
// fail the current test if it doesn't match the expected status code
func ValidateResponseCode(t *testing.T, res *http.Response, expected int) {
	if expected != res.StatusCode {
		t.Errorf("validation error: unexpected http response status: have %d; want %d",
			res.StatusCode, expected,
		)
	}
}

// ValidateJSONSchema ensures a given json field matches an expcted json
// string
func ValidateJSONSchema(t *testing.T, json *simplejson.Json, expected string) {
	if json == nil {
		t.Fatalf("validation error: json schema for is nil")
	}
	encoded, err := json.Encode()
	if err != nil {
		t.Fatalf("validation error: unable to encode expected json: %v", err)
	}
	if actual := string(encoded); actual != expected {
		t.Errorf("actual json schema does not match expected schema: have %s; want %v",
			actual, expected,
		)
	}
}

//					Negative tests

// TestReadNonExistContacts_REST attempts to read a non-existent contact from the application
// 1. Attempt to get the non-existent contact with incorrect ID by ReadResourceWithStatus ()
// 2. Ensure the response from ReadResourceWithStatus () has expected fields and detailed error message
func TestReadNonExistContact_REST(t *testing.T) {
	dbTest.Reset(t)
	id := "1"
	readJSON := ReadResourceWithStatus(t, "contacts", id, http.StatusNotFound)

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

// TestPartiallyUpdateNonExistContact_REST attempts to update field for a non-existent contact
// 1. Attempt to update field with incorrect ID for the non-existent contact by PartiallyUpdateResourceWithStatus ()
// 2. Ensure the response from PartiallyUpdateResourceWithStatus () has expected fields and detailed error message
func TestPartiallyUpdateNonExistContact_REST(t *testing.T) {
	dbTest.Reset(t)
	updatedFields := map[string]interface{}{
		"first_name": "Jack",
	}
	id := "1"
	updateJSON := PartiallyUpdateResourceWithStatus(t, updatedFields, "contacts", id, http.StatusNotFound)

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
