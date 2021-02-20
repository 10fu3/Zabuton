package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

var (
	strContentType = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func main() {
	err := fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		map1 := map[string]interface{}{
			"name": "hawksnowlog",
			"age":  20,
		}
		doJSONWrite(ctx,200,map1)

	})

	if err != nil{
		fmt.Println(err)
	}
}

func doJSONWrite(ctx *fasthttp.RequestCtx, code int, obj interface{}) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(code)
	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}