
## RLP 压缩算法

### Quick Start

```golang
package test

import (
	"fmt"
	"github.com/pefish/go-rlp"
)

func ExampleEncode() {
	test := Test{
		Str: `123`,
		Uint64_: 234,
	}
	bytes, err := rlp.EncodeToBytes(test)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", bytes)

	var test1 Test
	err = rlp.DecodeBytes(bytes, &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", test1)
	// Output:
	// C68331323381EA
	// test.Test{Str:"123", Uint64_:0xea}
}

```

