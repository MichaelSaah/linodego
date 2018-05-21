package main

/**
 * The tests in the examples directory demontrate use and test the library
 * in a real-use setting
 *
 * cd examples && go test -test.v
 */

import (
	"fmt"
	"log"
	"strings"

	golinode "github.com/chiefy/go-linode"
)

func ExampleListTypes_all() {
	types, err := linodeClient.ListTypes(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID contains class:", strings.Index(types[0].ID, types[0].Class) > -1)
	fmt.Println("Plan has Ram:", types[0].Memory > 0)

	// Output:
	// ID contains class: true
	// Plan has Ram: true
}

// ExampleGetType_missing demonstrates the Error type, which allows inspecting
// the request and response.  Error codes will be the HTTP status code,
// or sub-100 for errors before the request was issued.
func ExampleGetType_missing() {
	_, err := linodeClient.GetType("missing-type")
	if err != nil {
		if v, ok := err.(*golinode.Error); ok {
			fmt.Println("Request was:", v.Response.Request.URL)
			fmt.Println("Response was:", v.Response.Status)
			fmt.Println("Error was:", v)
		}
	}

	// Output:
	// Request was: https://api.linode.com/v4/linode/types/missing-type
	// Response was: 404 NOT FOUND
	// Error was: [404] Not found
}

// ExampleListKernels_all Demonstrates how to list all Linode Kernels.  Paginated
// responses are automatically traversed and concatenated when the ListOptions are nil
func ExampleListKernels_all() {
	kernels, err := linodeClient.ListKernels(nil)
	if err != nil {
		log.Fatal(err)
	}

	// The Linode API default pagination size is 100.
	fmt.Println("Fetched > 100:", len(kernels) > 100)

	// Output:
	// Fetched > 100: true
}

func ExampleListKernels_allWithOpts() {
	filterOpt := golinode.NewListOptions(0, "")
	kernels, err := linodeClient.ListKernels(filterOpt)
	if err != nil {
		log.Fatal(err)
	}

	// The Linode API default pagination size is 100.
	fmt.Println("Fetched > 100:", len(kernels) > 100)
	fmt.Println("Fetched Results/100 pages:", filterOpt.Pages > filterOpt.Results/100)
	fmt.Println("Fetched all results:", filterOpt.Results == len(kernels))

	// Output:
	// Fetched > 100: true
	// Fetched Results/100 pages: true
	// Fetched all results: true

}

func ExampleListKernels_filtered() {
	filterOpt := golinode.ListOptions{Filter: "{\"label\":\"Recovery - Finnix (kernel)\"}"}
	kernels, err := linodeClient.ListKernels(&filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	for _, kern := range kernels {
		fmt.Println(kern.ID, kern.Label)
	}

	// Unordered output:
	// linode/finnix Recovery - Finnix (kernel)
	// linode/finnix-legacy Recovery - Finnix (kernel)
}

func ExampleListKernels_page1() {
	filterOpt := golinode.NewListOptions(1, "")
	kernels, err := linodeClient.ListKernels(filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	// The Linode API default pagination size is 100.
	fmt.Println("Fetched == 100:", len(kernels) == 100)
	fmt.Println("Results > 100:", filterOpt.Results > 100)
	fmt.Println("Pages > 1:", filterOpt.Pages > 1)
	k := kernels[len(kernels)-1]
	fmt.Println("Kernel Version in Label:", strings.Index(k.Label, k.Version) > -1)
	fmt.Println("Kernel Version in ID:", strings.Index(k.ID, k.Label) > -1)

	// Output:
	// Fetched == 100: true
	// Results > 100: true
	// Pages > 1: true
	// Kernel Version in Label: true
	// Kernel Version in ID: true
}

func ExampleListImages_all() {
	filterOpt := golinode.NewListOptions(0, "")
	images, err := linodeClient.ListImages(filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched Results/100 pages:", filterOpt.Pages > filterOpt.Results/100)
	fmt.Println("Fetched all results:", filterOpt.Results == len(images))

	// Output:
	// Fetched Results/100 pages: true
	// Fetched all results: true

}

// ExampleListImages_notfound demonstrates that an empty slice is returned,
// not an error, when a filter matches no results.
func ExampleListImages_notfound() {
	filterOpt := golinode.ListOptions{Filter: "{\"label\":\"not-found\"}"}
	images, err := linodeClient.ListImages(&filterOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Images with Label 'not-found':", len(images))

	// Output:
	// Images with Label 'not-found': 0
}

// ExampleListImages_notfound demonstrates that an error is returned by
// the API and golinode when an invalid filter is provided
func ExampleListImages_badfilter() {
	filterOpt := golinode.ListOptions{Filter: "{\"foo\":\"bar\"}"}
	images, err := linodeClient.ListImages(&filterOpt)
	if err == nil {
		log.Fatal(err)
	}
	fmt.Println("Error given on bad filter:", err)
	fmt.Println("Images on bad filter:", images) // TODO: nil would be better here

	// Output:
	// Error given on bad filter: [400] [X-Filter] Cannot filter on foo
	// Images on bad filter: []
}

func ExampleListLongviewSubscriptions_page1() {
	pageOpt := golinode.ListOptions{PageOptions: &golinode.PageOptions{Page: 1}}
	subscriptions, err := linodeClient.ListLongviewSubscriptions(&pageOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Longview Subscription Types:", len(subscriptions))

	// Output:
	// Longview Subscription Types: 4
}
