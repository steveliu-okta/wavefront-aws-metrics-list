# wavefront-aws-metrics-list

```
chart/metrics/all?
Parameters:
q (string, optional) - filter for metric names or namespaces with the prefix specified; all top-level namespaces will be returned if this is left blank
h (string, optional) - filter for only metric names associated with the exact source name specified
l (integer, optional) - limit number of results to fetch; 1000 is the maximum; defaults to 100
p (string, optional) - last found metric name; used for pagination
```

This tool is for finding all the metrics under aws root namespace in wavefront.

https://okta.wavefront.com/metrics#_v01(q:aws.)

```bash
go run *.go
```