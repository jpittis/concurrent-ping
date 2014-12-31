package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type task interface {
	process()
	done()
}

type factory interface {
	make(line string) task
}

func run(f factory, n int) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.make(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("error when reading os.Stdin", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.done()
	}
}

type getFactory struct{}

func (g *getFactory) make(line string) task {
	return &get{url: line}
}

type get struct {
	url    string
	status string
	err    error
}

func (g *get) process() {
	resp, err := http.Get("http://" + g.url)
	if err != nil {
		g.err = err
		g.status = "Error"
	} else {
		g.status = resp.Status
	}
}

func (g *get) done() {
	fmt.Printf("%s | %s\n", g.status, g.url)
}

func main() {
	run(&getFactory{}, 100)
}
