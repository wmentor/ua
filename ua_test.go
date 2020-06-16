package ua

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttp(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<html>
<body>
<h1>Hello, World!</h1>
</body>
</html>`)

	}))
	defer ts.Close()

	client := New()

	resp, err := client.Request("GET", ts.URL, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp == nil || resp.StatusCode != 200 {
		t.Fatal("response failes")
	}

	resp, err = client.Request("POST", ts.URL, map[string]string{"X-Request-Id": "123"}, strings.NewReader("test"))
	if err != nil {
		t.Fatal(err)
	}

	if resp == nil || resp.StatusCode != 200 {
		t.Fatal("response failes")
	}
}
