package main

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	client := influxdb2.NewClient("http://54.219.21.24:8086", "eh2GDwqhbF-gbCEmkrAQROc3lG5efcXGCkx5XmCYm4hvpnYSk5N-8t9Y5cT6PVNWJREcR3La-f2qw8uaXiOZjg==")
	defer client.Close()

	queryAPI := client.QueryAPI("coinsky")
	query := `from(bucket: "twitter_index_ranking")
  |> range(start:-7d)
  |> filter(fn: (r) => r["_measurement"] == "topic")
  |> filter(fn: (r) => r["_field"] == "BTC")
  |> aggregateWindow(every: 4h, fn: last, createEmpty: false)`

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	if result.Err() != nil {
		fmt.Printf("query parsing error: %\n", result.Err().Error())
	}

}
