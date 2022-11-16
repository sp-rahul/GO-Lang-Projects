package main

// https://github.com/rapito/go-shopify

import (
	"fmt"

	"github.com/rapito/go-shopify/shopify"
)

func main() {
	//1st= store name, 2nd = API key, 3rd = Admin API access token (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)
	shop := shopify.New("SP-SHOPING-ROOM", "0467848da62ce15a71093fba27dd75f0", "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")

	result, _ := shop.Get("products")

	fmt.Println(string(result))
}
