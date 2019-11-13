package main

import (
	"fmt"
	"io"
)

// generate n metric names
func genMetricNames(prefix string, id, n int) []string {
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = fmt.Sprintf("%s.agent.%d.metrics.%d", prefix, id, i)
	}

	return names
}

// actually write the data in carbon line format
func carbonate(w io.ReadWriteCloser, name string, value int, epoch int64, datapoints int) error {
	for i := 0; i < datapoints; i++ {
		e := epoch - int64(i*60)
		_, err := fmt.Fprintf(w, "%s %d %d\n", name, value, e)
		if err != nil {
			return err
		}
	}
	return nil
}
