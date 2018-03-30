# go-pagerduty

quoinx-go-client is a [go](https://golang.org/) client library for [Quoine v2 API](https://developers.quoine.com/).
[godoc]()

## Installation

```
go get github.com/jjjjpppp/quoinex-go-client
```

### From golang libraries

```go
package main

import (
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client"
)


func main() {
	client, _ := NewClient("apiTokenID", "secret", nil) // your token and secret setup here
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  priceLevels, err := client.GetOrderBook(ctx,[productID])

  // put your code here...
}
```

## License
[MIT](https://opensource.org/licenses/mit-license.php)

## Contributing

1. Fork it ( https://github.com/jjjjpppp/quoinex-go-client/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
