package exporter

import (
	_ "embed"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed testdata/success.xml
var xml_data string

func TestCollectorWithMockServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(xml_data))
	}))
	defer ts.Close()

	//w := httptest.NewRecorder()

	tests := []struct {
		name    string
		url     string
		success bool
	}{
		{"Fetch XML data successfully", ts.URL, true},
		{"Timeout", "http://localhost/", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := fetchData(ts.URL)
			if err != nil {
				t.Errorf("%s: Expected result %t, got: %v", test.name, test.success, err)
			}
		})
	}
}
