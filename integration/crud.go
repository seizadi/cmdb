// +build integration

package integration

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"
	"strings"
	"testing"
)

// CreateResourceWithStatus uses the REST gateway to:
// 1. Create a new object with a POST request by the url and map with values as body
// 2. Compare response code from the response with the statusCode
// 3. Unmarshal the JSON from response into a simplejson struct and returns it
// 4. Return ID of created object
func CreateResourceWithStatus(t *testing.T, created map[string]interface{}, resource string, statusCode int) (id string, json simplejson.Json) {
	resCreate, err := MakeRequestWithDefaults(
		http.MethodPost,
		fmt.Sprintf("http://%s/v1/%s", GatewayAddress, resource),
		created,
	)
	if err != nil {
		t.Fatalf("unable to create object: %v", err)
	}
	ValidateResponseCode(t, resCreate, statusCode)
	createJSON, err := simplejson.NewFromReader(resCreate.Body)
	if err != nil {
		t.Fatalf("unable to marshal json response: %v", err)
	}
	id, err = createJSON.GetPath("result", "id").String()
	if err != nil {
		t.Fatalf("unable to get object id from response json: %v", err)
	}
	id = strings.TrimPrefix(id, fmt.Sprintf("%s/%s/", ApplicationID, resource))
	return id, *createJSON
}

func CreateResource(t *testing.T, created map[string]interface{}, resource string) (id string, json simplejson.Json) {
	statusCode := http.StatusOK
	return CreateResourceWithStatus(t, created, resource, statusCode)
}

// ReadResourceWithStatus uses the REST gateway to:
// 1. Read the object with a GET request by the url
// 2. Compare response code from the response with the statusCode
// 3. Unmarshal the JSON from response into a simplejson struct and returns it
func ReadResourceWithStatus(t *testing.T, resource, id string, statusCode int) (json simplejson.Json) {
	resRead, err := MakeRequestWithDefaults(
		http.MethodGet,
		fmt.Sprintf("http://%s/v1/%s/%s", GatewayAddress, resource, id),
		nil,
	)
	if err != nil {
		t.Fatalf("unable to get object: %v", err)
	}
	ValidateResponseCode(t, resRead, statusCode)
	readJSON, err := simplejson.NewFromReader(resRead.Body)
	if err != nil {
		t.Fatalf("unable to marshal json response: %v", err)
	}
	return *readJSON
}

func ReadResource(t *testing.T, resource, id string) (json simplejson.Json) {
	statusCode := http.StatusOK
	return ReadResourceWithStatus(t, resource, id, statusCode)
}

// UpdateResourceWithStatus uses the REST gateway to:
// 1. Update the object with a PUT request by the url and map with values as body
// 2. Compare response code from the response with the statusCode
// 3. Unmarshal the JSON from response into a simplejson struct and returns it
func UpdateResourceWithStatus(t *testing.T, updated map[string]interface{}, resource, id string, statusCode int) (json simplejson.Json) {
	resUpdate, err := MakeRequestWithDefaults(
		http.MethodPut,
		fmt.Sprintf("http://%s/v1/%s/%s", GatewayAddress, resource, id),
		updated,
	)
	if err != nil {
		t.Fatalf("unable to update object: %v", err)
	}
	ValidateResponseCode(t, resUpdate, statusCode)
	updateJSON, err := simplejson.NewFromReader(resUpdate.Body)
	if err != nil {
		t.Fatalf("unable to unmarshal json response: %v", err)
	}
	return *updateJSON
}

func UpdateResource(t *testing.T, updated map[string]interface{}, resource, id string) (json simplejson.Json) {
	statusCode := http.StatusOK
	return UpdateResourceWithStatus(t, updated, resource, id, statusCode)
}

// DeleteResourceWithStatus uses the REST gateway to:
// 1. Delete the object with a DELETE request by the url
// 2. Compare response code from the response with the statusCode
// 3. Unmarshal the JSON from response into a simplejson struct and returns it
func DeleteResourceWithStatus(t *testing.T, resource, id string, statusCode int) (json simplejson.Json) {
	resDelete, err := MakeRequestWithDefaults(
		http.MethodDelete,
		fmt.Sprintf("http://%s/v1/%s/%s", GatewayAddress, resource, id),
		nil,
	)
	if err != nil {
		t.Fatalf("unable to delete object: %v", err)
	}
	ValidateResponseCode(t, resDelete, statusCode)
	deleteJSON, err := simplejson.NewFromReader(resDelete.Body)
	if err != nil {
		t.Fatalf("unable to marshal json response: %v", err)
	}
	return *deleteJSON
}

func DeleteResource(t *testing.T, resource, id string) (json simplejson.Json) {
	statusCode := http.StatusOK
	return DeleteResourceWithStatus(t, resource, id, statusCode)
}

// ListResourceWithStatus uses the REST gateway to:
// 1. Read all the objects with a GET request by the url
// 2. Compare response code from the response with the statusCode
// 3. Unmarshal the JSON from response into a simplejson struct and returns it
func ListResourceWithStatus(t *testing.T, resource string, statusCode int) (json simplejson.Json) {
	resList, err := MakeRequestWithDefaults(
		http.MethodGet,
		fmt.Sprintf("http://%s/v1/%s", GatewayAddress, resource),
		nil,
	)
	if err != nil {
		t.Fatalf("unable to list objects %v", err)
	}
	ValidateResponseCode(t, resList, statusCode)
	listJSON, err := simplejson.NewFromReader(resList.Body)
	if err != nil {
		t.Fatalf("unable to unmarshal json response: %v", err)
	}
	return *listJSON
}

func ListResource(t *testing.T, resource string) (json simplejson.Json) {
	statusCode := http.StatusOK
	return ListResourceWithStatus(t, resource, statusCode)
}

// PartiallyUpdateResourceWithStatus uses the REST gateway to:
// 1. Update some fields in the object with a PATCH request by the url and map with changing fields as body
// 2. Compare response code from the response with the statusCode
// 3. Unmarshal the JSON from response into a simplejson struct and returns it
func PartiallyUpdateResourceWithStatus(t *testing.T, updatedFields map[string]interface{}, resource, id string, statusCode int) (json simplejson.Json) {
	resUpdateField, err := MakeRequestWithDefaults(
		http.MethodPatch,
		fmt.Sprintf("http://%s/v1/%s/%s", GatewayAddress, resource, id),
		updatedFields,
	)
	if err != nil {
		t.Fatalf("unable to update group: %v", err)
	}
	ValidateResponseCode(t, resUpdateField, statusCode)
	updateJSON, err := simplejson.NewFromReader(resUpdateField.Body)
	if err != nil {
		t.Fatalf("unable to unmarshal json response: %v", err)
	}
	return *updateJSON
}

func PartiallyUpdateResource(t *testing.T, updated map[string]interface{}, resource, id string) (json simplejson.Json) {
	statusCode := http.StatusOK
	return PartiallyUpdateResourceWithStatus(t, updated, resource, id, statusCode)
}
