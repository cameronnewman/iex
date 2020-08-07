/*
Package iex is a simple client library for iextrading.com API. Package supports a Deep lookup.
*/
package iex

// Simple example:
//
//	iexClient := iex.NewIEX(&http.Client{
//		Timeout: time.Second * 10,
//	})
//
//	// Print out details for Apple Inc.
//	fmt.Println(iexClient.GetDEEP("APPL"))
//
