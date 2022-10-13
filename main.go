package main

import (
	"context"
	"flag"
	"fmt"
)

func main() {
	metricName := flag.String("metric", "aws.", "metrics name prefix, eg: aws. or cpu.")
	flag.Parse()

	c, err := New()
	if err != nil {
		panic(err)
	}

	res := search(c, *metricName)
	if len(res) > 0 {
		err = writeToJSONFile(res, *metricName)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("There are %d metrics found for %s metric prefix\n", len(res), *metricName)
}

func search(c *Client, metric string) []string {
	var res []string
	cResp := &Resp{}

	err := c.GetAWSMetricsList(context.TODO(), cResp, metric)
	if err != nil {
		panic(err)
	}

	for _, mt := range cResp.Metrics {
		if mt[len(mt)-1:] != "." && !contains(res, mt) {
			res = append(res, mt)
		} else {
			if mt[len(mt)-1:] == "." {
				res = append(res, search(c, mt)...)
			}
		}
	}

	return res
}
