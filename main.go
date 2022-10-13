package main

import (
	"context"
	"fmt"
)

func main() {
	c, err := New()
	if err != nil {
		panic(err)
	}

	res := search(c, "aws.")
	err = writeToJSONFile(res)
	if err != nil {
		panic(err)
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
