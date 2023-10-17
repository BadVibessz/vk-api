package main

import (
	"context"
	"fmt"
	http "net/http"
	"vk-api"
)

//func ContextClient(ctx context.Context) *http.Client {
//	if ctx != nil {
//		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
//			return hc
//		}
//	}
//	if appengineClientHook != nil {
//		return appengineClientHook(ctx)
//	}
//	return http.DefaultClient
//}

func main() {

	h := http.Client{}
	client := vkapi.Client{
		Http:       &h,
		BaseURL:    "http://google.com",
		Retry:      false,
		RetryCount: 0}

	ctx := context.Background()

	resp, err := client.Post(ctx, "", nil)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	println(resp)
}
