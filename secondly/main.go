package main

import (
	"sync"
	"bufio"
	"os"
	"log"
	"fmt"
	"net/http"
	"flag"
)

var c int

type task interface {
	process()
	output()
}

type factory interface {
	create (line string) task
}

func run(f factory) {
	var wg sync.WaitGroup
	in :=  make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}

		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range in {
				t.process()
				out <- t
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.output()
	}
}

type HttpTask struct {
	url string
	ok bool
}

func (t *HttpTask) process() {
	if t.url == "" {
		t.ok = false
		return
	}
	resp, err := http.Get(t.url)

	if err != nil {
		t.ok = false
		return
	}

	if resp.StatusCode == http.StatusOK {
		t.ok = true
		return
	}

	t.ok = false
}

type Factory struct {}

func (f *Factory) create(line string) task {
	h := &HttpTask{}
	h.url = line
	return h
}

func (t *HttpTask) output() {
	fmt.Printf("%s  %t\n", t.url, t.ok)
}

func main() {
	count := flag.Int("count", 10, "Number of concurrent task being processed")
	flag.Parse()
	c = *count

	f := &Factory{}
	run(f)
}
