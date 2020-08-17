[![GoDoc](https://godoc.org/github.com/william20111/go-smartsheet?status.svg)](http://godoc.org/github.com/william20111/go-smartsheet) [![Go Report Card](https://goreportcard.com/badge/github.com/william20111/go-smartsheet)](https://goreportcard.com/report/github.com/william20111/go-smartsheet) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE) [![codecov](https://codecov.io/gh/william20111/go-smartsheet/branch/master/graph/badge.svg)](https://codecov.io/gh/william20111/go-smartsheet)
# go-smartsheet

go-smartsheet is a [go](https://golang.org/) client library for the [Smartsheet 2.0 API](https://smartsheet-platform.github.io/api-docs/).

### Using client library

```go
package main

import (
	"fmt"
	"github.com/william20111/go-smartsheet/pkg/smartsheet"
	"log"
)

func main() {
	options := smartsheet.ClientOptions{}
	client := smartsheet.NewSmartsheetClient(&options)
	sheet, err := client.GetSheet("1234567891011")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sheet.Rows)
}
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License

[Apache 2](http://www.apache.org/licenses/LICENSE-2.0)
