package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Users(t *testing.T) {
	expectedURI := "/list"
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.RequestURI != expectedURI {
			t.Fatalf("Expected URI to be ''%s' but got '%s'", expectedURI, request.RequestURI)
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("{\"results\":[\"testUser1\",\"anotherUser\"]}"))
	}))
	defer testServer.Close()

	client := New(testServer.URL)
	users, err := client.Users(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 2 {
		t.Fatalf("Expected %d users but got %d", 2, len(users))
	}
}

func Test_Devices(t *testing.T) {
	expectedURI := "/list?user=test"
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.RequestURI != expectedURI {
			t.Fatalf("Expected URI to be ''%s' but got '%s'", expectedURI, request.RequestURI)
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("{\"results\":[\"testDevice1\",\"testDevice2\"]}"))
	}))
	defer testServer.Close()

	client := New(testServer.URL)
	devices, err := client.Devices(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}

	if len(devices) != 2 {
		t.Fatalf("Expected %d devices but got %d", 2, len(devices))
	}
}

func Test_Locations(t *testing.T) {

}

func Test_Version(t *testing.T) {
	expectedURI := "/version"
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.RequestURI != expectedURI {
			t.Fatalf("Expected URI to be ''%s' but got '%s'", expectedURI, request.RequestURI)
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("{\"version\":\"1.3.3.7\"}"))
	}))
	defer testServer.Close()

	client := New(testServer.URL)
	version, err := client.Version(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	expectedVersion := "1.3.3.7"

	if version.Version != expectedVersion {
		t.Fatalf("Expected version to be %s but got %s", expectedVersion, version.Version)
	}
}
