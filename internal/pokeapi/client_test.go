package pokeapi_test

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
    
    "github.com/wdrg22/pokedex/internal/pokeapi"
)

func TestNewClient(t *testing.T) {
    client := pokeapi.NewClient(5 * time.Second)
    if client == nil {
        t.Error("expected client to not be nil")
    }
}

func TestGetLocationAreas(t *testing.T) {
    // Create test server with mock response
    mockResponse := pokeapi.LocationAreaResponse{
        Count: 2,
        Next: strPtr("https://example.com/next"),
        Previous: strPtr("https://example.com/prev"),
        Results: []struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        }{
            {Name: "test-location-1", URL: "https://example.com/1"},
            {Name: "test-location-2", URL: "https://example.com/2"},
        },
    }

    // Create test server
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(mockResponse)
    }))
    defer ts.Close()

    // Create client
    client := pokeapi.NewClient(5 * time.Second)

    // Test first request (should hit API)
    resp1, err := client.GetLocationAreas(ts.URL)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    // Verify response data
    if resp1.Count != 2 {
        t.Errorf("expected Count=2, got %d", resp1.Count)
    }
    if *resp1.Next != "https://example.com/next" {
        t.Errorf("expected Next=https://example.com/next, got %s", *resp1.Next)
    }
    if len(resp1.Results) != 2 {
        t.Errorf("expected 2 results, got %d", len(resp1.Results))
    }
    if resp1.Results[0].Name != "test-location-1" {
        t.Errorf("expected first location name=test-location-1, got %s", resp1.Results[0].Name)
    }

    // Modify server response to verify we get cached response
    mockResponse.Count = 999

    // Test second request (should hit cache)
    resp2, err := client.GetLocationAreas(ts.URL)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    // Verify we got the cached response, not the modified one
    if resp2.Count != 2 {
        t.Errorf("expected cached Count=2, got %d", resp2.Count)
    }
}

func TestGetLocationAreasError(t *testing.T) {
    // Create test server that returns an error
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError)
    }))
    defer ts.Close()

    client := pokeapi.NewClient(5 * time.Second)
    _, err := client.GetLocationAreas(ts.URL)
    if err == nil {
        t.Error("expected error for bad response, got nil")
    }
}

// Helper function to create string pointer
func strPtr(s string) *string {
    return &s
}
