package main

import (
	"flag"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

// config vars, to be manipulated via command line flags
var (
	carbon     string
	prefix     string
	metrics    int
	workers    int
	tasks      int
	datapoints int
	cacheConns bool
)

type Worker struct {
	ID         int
	Addr       string
	Connection net.Conn
}

func (w *Worker) Flush(taskID int, metricNames []string) error {
	if w.Connection == nil {
		conn, err := net.Dial("tcp", w.Addr)
		if err != nil {
			return err
		}
		w.Connection = conn
	}

	if !cacheConns {
		defer func() {
			if w.Connection != nil {
				w.Connection.Close()
				w.Connection = nil
			}
		}()
	}

	epoch := time.Now().Unix()
	for _, name := range metricNames {
		err := carbonate(w.Connection, name, rand.Intn(1000), epoch, datapoints)
		if err != nil {
			w.Connection = nil
			return err
		}
	}

	log.Printf("agent %d: flushed %d metrics for task %d\n", w.ID, len(metricNames), taskID)
	return nil
}

func launchAgent(wg *sync.WaitGroup, q chan int, id, n int, addr, prefix string) {
	defer wg.Done()

	w := &Worker{
		ID:   id,
		Addr: addr,
	}

	for {
		taskID, ok := <-q
		if !ok {
			return
		}
		metricNames := genMetricNames(prefix, taskID, n)
		if err := w.Flush(taskID, metricNames); err != nil {
			log.Printf("agent %d: %s\n", w.ID, err)
		}
	}
}

func init() {
	flag.StringVar(&carbon, "carbon", "localhost:2003", "address of carbon host")
	flag.StringVar(&prefix, "prefix", "haggar", "prefix for metrics")
	flag.IntVar(&metrics, "metrics", 10000, "number of metrics for each agent to hold")
	flag.IntVar(&workers, "workers", 100, "max number of workers to run concurrently")
	flag.IntVar(&datapoints, "datapoints", 1, "number of datapoints each metrics")
	flag.IntVar(&tasks, "tasks", 100, "number of tasks that will pass to woker")
	flag.BoolVar(&cacheConns, "cache_connections", false, "if set, keep connections open between flushes (default: false)")
}

func main() {
	flag.Parse()

	log.Printf("master: pid %d\n", os.Getpid())

	q := make(chan int, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go launchAgent(&wg, q, i, metrics, carbon, prefix)
	}

	for i := 0; i < tasks; i++ {
		q <- i
	}
	close(q)

	wg.Wait()
}
