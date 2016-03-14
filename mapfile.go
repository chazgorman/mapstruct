package main

import (
	"bytes"
	//	"encoding/json"
	"log"
	"os"

	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj"
	"github.com/geo-data/mapfile/tokens"
)

func main() {
	mapfile := os.Args[1]
	tokens, err := tokens.TokenizeMap(mapfile)
	if err != nil {
		log.Fatal(err)
	}

	var map_ *mapobj.Map
	if map_, err = mapobj.New(tokens); err != nil {
		panic(err)
		log.Fatal(err)
	}

	/*b, err := json.Marshal(map_)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")*/

	var out bytes.Buffer
	enc := encoding.NewMapfileEncoder(&out)
	if err = enc.Encode(map_); err != nil {
		log.Fatal(err)
	}

	if _, err = out.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
