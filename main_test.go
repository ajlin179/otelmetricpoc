package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
)

func TestCounter(t *testing.T) {
	defer initMetricProvider()()

	ctx := context.Background()
	meter := global.Meter("herapoc-demo-client-meter")

	requestCount, _ := meter.SyncInt64().Counter(
		"herapoc_demo_client_request_counts",
		instrument.WithDescription("The number of requests processed"),
	)

	commonLabels := []attribute.KeyValue{
		attribute.String("method", "repl"),
		attribute.String("client", "cli"),
	}

	for i := 1; i <= 2500; i++ {
		requestCount.Add(ctx, 1, commonLabels...)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Counter:==>" + strconv.Itoa(i))
	}

}

func TestCounterWithVariableDimention(t *testing.T) {
	defer initMetricProvider()()

	ctx := context.Background()
	meter := global.Meter("herapoc-demo-client-meter")

	requestCount, _ := meter.SyncInt64().Counter(
		"herapoc_demo_mdvar_counts",
		instrument.WithDescription("The number of requests processed"),
	)

	commonLabels := []attribute.KeyValue{
		attribute.String("method", "repl"),
		attribute.String("client", "cli"),
	}

	for i := 1; i <= 2500; i++ {

		min := 0
		max := 1500
		commonLabelsLocal := append(commonLabels, attribute.String("sqlhash", strconv.Itoa(rand.Intn(max-min)+min)))

		requestCount.Add(ctx, 1, commonLabelsLocal...)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Counter:==>" + strconv.Itoa(i))
	}

}

func TestHistogram(t *testing.T) {
	defer initMetricProvider()()

	ctx := context.Background()
	meter := global.Meter("herapoc-demo-client-meter")

	// Recorder metric example
	requestLatency, _ := meter.SyncFloat64().Histogram(
		"herapoc_demo_request_latency",
		instrument.WithDescription("The latency of requests processed"),
	)

	commonLabels := []attribute.KeyValue{
		attribute.String("method", "repl"),
		attribute.String("client", "cli"),
	}

	for i := 1; i <= 2500; i++ {
		min := 1
		max := 15
		requestLatency.Record(ctx, float64(rand.Intn(max-min)+min), commonLabels...)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Counter:==>" + strconv.Itoa(i))
	}

}

func TestHistogramWithExplicitBoundry(t *testing.T) {
	defer initMetricProvider()()

	ctx := context.Background()
	meter := global.Meter("herapoc-demo-client-meter")

	// Recorder metric example
	requestLatency, _ := meter.SyncFloat64().Histogram(
		"herapoc_demo_ExplicitBoundry_request_latency",
		instrument.WithDescription("The latency of requests processed"),
	)

	commonLabels := []attribute.KeyValue{
		attribute.String("method", "repl"),
		attribute.String("client", "cli"),
	}

	for i := 1; i <= 2500; i++ {
		min := 1
		max := 15
		requestLatency.Record(ctx, float64(rand.Intn(max-min)+min), commonLabels...)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Counter:==>" + strconv.Itoa(i))
	}

}
