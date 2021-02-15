package main

import (
	"encoding/json"
	"fmt"
	"nibar/internal"
)

func main() {
	if b, err := pprint(*internal.RunStatus()); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", b)
	}

}

func pprint(i interface{}) ([]byte, error) {
	return json.MarshalIndent(i, "", "  ")
}
