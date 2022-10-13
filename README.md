# wavefront-aws-metrics-list
This tool is for finding all the metrics with a specific prefix in wavefront.

This API returns all non-obsolete metric names (non obsolete means that some data point with that metric name was ingested at any point in time over the last 4 weeks).

```bash
go run *.go  -metric=aws.
go run *.go  -metric=cpu.
go run *.go  -metric=ecs.containerinsights.
go run *.go  -metric=monolith.
go run *.go  -metric=GenericJMX.
```