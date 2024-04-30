# go-rest-client
Golang http client wrapper

## Usage


#### Basic usage

```go
package main

import (
	"fmt"
	"github.com/sitnikovik/go-rest-client/request"
	"github.com/sitnikovik/go-rest-client/rest_client"
)

func main() {
	req := request.NewRequest().
		SetUrl("https://www.google.com").
		AsGet()

	client := rest_client.NewRestClient()
	bb, err := client.SendAndDecode(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bb))
}
```

#### Wait for error code

```go
package main

import (
	"fmt"
	"github.com/sitnikovik/go-rest-client/request"
	"github.com/sitnikovik/go-rest-client/rest_client"
)

func main() {
	req := request.NewRequest().
		SetUrl("https://software.hixie.ch/utilities/cgi/test-tools/http-error?status=500+Internal+Server+Error").
		SetStatusExpected(500)
		AsGet()
	
    client := rest_client.NewRestClient()
	bb, err := client.SendAndDecode(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bb)) // Responded with: 500 Internal Server Error
}
```

#### With JSON response decoding

```go
package main

import (
	"fmt"
	"github.com/sitnikovik/go-rest-client/request"
	"github.com/sitnikovik/go-rest-client/rest_client"
)

func main() {
	req := request.NewRequest().
		SetUrl("https://mocki.io/v1/e597d386-ef58-45fa-81ad-dd9a62a066f7").
		AsGet()

	v := struct {
		Employee struct {
			Name           string `json:"name"`
			Age            int    `json:"age"`
			Salary         int    `json:"salary"`
			SalaryCurrency string `json:"salaryCurrency"`
		} `json:"employee"`
	}{}
	err := rest_client.NewRestJsonClient().SendAndDecode(req, &v)

	if err != nil {
		fmt.Printf("Error: %v but want nil", err)
		return
	}
	
	// Expected: {Employee:{Name:John Doe Age:20 Salary:10000 SalaryCurrency:USD}}
	fmt.Printf("%+v", v)
}
```