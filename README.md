Simple http client

# Summary

* Use tiny Go
* Require Go version >= 1.10
* Minimum external dependencies
* Simple in usage

# Install

```
go get github.com/wmentor/ua
```

# Usage

```go
package main

import (
  "strings"
  "time"

  "github.com/wmentor/ua"
)

package main() {

  agent := ua.New()

  agent.Timeout = time.Second * 5
  agent.UserAgent = "Mozilla"
  agent.Decode = true // deconde to utf-8

  headers := map[string]string{"X-Request-Id": "12313"}
  data := strings.NewReader("content body")

  resp, err := agent.Request( "POST", "https://someurl.ru", headers, data)

  if err != nil || resp == nil {
    panic("request failed")
  }

  if resp.StatusCode != 200 {
    panic("invalid status code")
  }

  fmt.Println(string(resp.Content))
}
```