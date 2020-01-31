package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/pefish/go-rlp"
	"time"
)

func ExampleEncodeEfficiency() {
	test := Test{
		Str: `123`,
		Uint64_: 234,
	}
	num := 100000


	t1 := time.Now() // get current time
	for i := 0; i < num; i++ {
		_, err := rlp.EncodeToBytes(test)
		if err != nil {
			panic(err)
		}
	}
	elapsed := time.Since(t1)
	fmt.Println("RLP elapsed: ", elapsed)


	testBuf := bytes.Buffer{}
	gobEncoder := gob.NewEncoder(&testBuf)
	t2 := time.Now()
	for i := 0; i < num; i++ {
		err := gobEncoder.Encode(test)
		if err != nil {
			panic(err)
		}
	}
	elapsed2 := time.Since(t2)
	fmt.Println("GOB elapsed: ", elapsed2)


	t3 := time.Now() // get current time
	for i := 0; i < num; i++ {
		_, err := json.Marshal(test)
		if err != nil {
			panic(err)
		}
	}
	elapsed3 := time.Since(t3)
	fmt.Println("JSON elapsed: ", elapsed3)

	// Output:
	// RLP elapsed:  37.708578ms
	// GOB elapsed:  51.050196ms
	// JSON elapsed:  36.320832ms
}
