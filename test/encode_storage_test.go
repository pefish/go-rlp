package test

import (
	bytes2 "bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/pefish/go-rlp"
)

type Test struct {
	Str string
	Uint64_ uint64
}

func ExampleEncodeString() {
	bytes, err := rlp.EncodeToBytes(`123`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", bytes)


	bytes, err = json.Marshal(`123`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X", bytes)

	// Output:
	//
}

func ExampleStorageEncode() {
	test := Test{
		Str: `123`,
		Uint64_: 234,
	}
	bytes, err := rlp.EncodeToBytes(test)
	if err != nil {
		panic(err)
	}

	var test1 Test
	err = rlp.DecodeBytes(bytes, &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RLP: %x  ", bytes)
	fmt.Printf("%#v\n", test1)

	testBuf := bytes2.Buffer{}
	gobEncoder := gob.NewEncoder(&testBuf)
	err = gobEncoder.Encode(test)
	if err != nil {
		panic(err)
	}

	gobDecoder := gob.NewDecoder(bytes2.NewReader(testBuf.Bytes()))
	var test2 Test
	err = gobDecoder.Decode(&test2)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`GOB: %x  `, testBuf.Bytes())
	fmt.Printf("%#v\n", test2)


	bytes1, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}

	var test3 Test
	err = json.Unmarshal(bytes1, &test3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %x  ", bytes1)
	fmt.Printf("%#v\n", test3)

	// Output:
	// RLP: c68331323381ea  test.Test{Str:"123", Uint64_:0xea}
	// GOB: 26ff81030101045465737401ff820001020103537472010c00010755696e7436345f01060000000bff82010331323301ffea00  test.Test{Str:"123", Uint64_:0xea}
	// JSON: 7b22537472223a22313233222c2255696e7436345f223a3233347d  test.Test{Str:"123", Uint64_:0xea}
}
