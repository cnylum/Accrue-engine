//go:build js && wasm

// Package main builds the SDK as a WebAssembly module.
// Build with: GOOS=js GOARCH=wasm go build -o accrue.wasm ./wasm/
package main

import (
	"context"
	"encoding/json"
	"syscall/js"

	"github.com/cnylum/accrue-engine/sdk/client"
	"github.com/cnylum/accrue-engine/sdk/types"
)

var defaultClient *client.Client

func main() {
	js.Global().Set("accrueInit", js.FuncOf(accrueInit))
	js.Global().Set("accruePlaceOrder", js.FuncOf(accruePlaceOrder))
	js.Global().Set("accrueGetBalance", js.FuncOf(accrueGetBalance))

	// Block forever — WASM module stays alive.
	select {}
}

func accrueInit(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return jsError("accrueInit requires baseURL and apiKey")
	}
	defaultClient = client.New(client.Config{
		BaseURL: args[0].String(),
		APIKey:  args[1].String(),
	})
	return nil
}

func accruePlaceOrder(_ js.Value, args []js.Value) any {
	if len(args) < 4 {
		return jsError("accruePlaceOrder requires symbol, side, orderType, quantity")
	}
	order, err := defaultClient.PlaceOrder(context.Background(), types.PlaceOrderRequest{
		Symbol:    args[0].String(),
		Side:      args[1].String(),
		OrderType: args[2].String(),
		Quantity:  args[3].String(),
	})
	if err != nil {
		return jsError(err.Error())
	}
	return jsJSON(order)
}

func accrueGetBalance(_ js.Value, _ []js.Value) any {
	balance, err := defaultClient.GetBalance(context.Background())
	if err != nil {
		return jsError(err.Error())
	}
	return jsJSON(balance)
}

func jsError(msg string) any {
	return js.ValueOf(map[string]any{"error": msg})
}

func jsJSON(v any) any {
	b, err := json.Marshal(v)
	if err != nil {
		return jsError(err.Error())
	}
	var result map[string]any
	json.Unmarshal(b, &result)
	return js.ValueOf(result)
}
