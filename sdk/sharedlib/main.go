// Package main builds the SDK as a C shared library (.so/.dylib/.dll).
// Build with: go build -buildmode=c-shared -o libaccrue.so ./sharedlib/
package main

// #include <stdlib.h>
import "C"
import (
	"context"
	"encoding/json"
	"unsafe"

	"github.com/cnylum/accrue-engine/sdk/client"
	"github.com/cnylum/accrue-engine/sdk/types"
)

var defaultClient *client.Client

//export AccrueInit
func AccrueInit(baseURL, apiKey *C.char) {
	defaultClient = client.New(client.Config{
		BaseURL: C.GoString(baseURL),
		APIKey:  C.GoString(apiKey),
	})
}

//export AccruePlaceOrder
func AccruePlaceOrder(symbol, side, orderType, quantity, limitPrice *C.char) *C.char {
	return wrap(func() (any, error) {
		return defaultClient.PlaceOrder(context.Background(), types.PlaceOrderRequest{
			Symbol:     C.GoString(symbol),
			Side:       C.GoString(side),
			OrderType:  C.GoString(orderType),
			Quantity:   C.GoString(quantity),
			LimitPrice: C.GoString(limitPrice),
		})
	})
}

//export AccrueGetBalance
func AccrueGetBalance() *C.char {
	return wrap(func() (any, error) {
		return defaultClient.GetBalance(context.Background())
	})
}

//export AccrueFree
func AccrueFree(p *C.char) {
	C.free(unsafe.Pointer(p))
}

// wrap executes fn with panic recovery and returns a JSON C string.
// The JSON has either {"data":...} or {"error":"..."}.
// Caller must free the returned string with AccrueFree.
func wrap(fn func() (any, error)) *C.char {
	defer func() {
		if r := recover(); r != nil {
			// Panic recovery — never crash the host process.
		}
	}()

	result, err := fn()
	if err != nil {
		b, _ := json.Marshal(map[string]string{"error": err.Error()})
		return C.CString(string(b))
	}

	b, err := json.Marshal(map[string]any{"data": result})
	if err != nil {
		b, _ = json.Marshal(map[string]string{"error": err.Error()})
		return C.CString(string(b))
	}
	return C.CString(string(b))
}

func main() {}
