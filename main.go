package main

import (
	"context"
	"fmt"
)

func main() {
	var res []string
	c, err := New()
	if err != nil {
		panic(err)
	}

	aResp := &Resp{}
	err = c.GetAWSMetricsList(context.TODO(), aResp, "aws.")
	for _, mm := range aResp.Metrics {
		restemp := search(c, mm)
		err = writeToJSONFile(restemp)
		res = append(res, restemp...)
	}

	fmt.Println(len(res))
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
