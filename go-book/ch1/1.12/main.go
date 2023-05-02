package main

import (
	"log"
	"net/http"
	"strconv"
)

const port = "8080"

func main() {
	if err := http.ListenAndServe(":"+port, http.HandlerFunc(lissajousHandler)); err != nil {
		log.Fatalf("localhost:%s server err: %v\n", port, err)
	}
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parsing url params: %v", err)
	}

	opts := NewGifOpts()
	for k, v := range r.Form {
		switch k {
		case "cycles":
			if cycles, err := strconv.Atoi(v[0]); err == nil {
				opts.Cycles = cycles
			}
		case "res":
			if res, err := strconv.ParseFloat(v[0], 32); err == nil {
				opts.Res = res
			}
		case "size":
			if size, err := strconv.Atoi(v[0]); err == nil {
				opts.Size = size
			}
		case "nframes":
			if nframes, err := strconv.Atoi(v[0]); err == nil {
				opts.NFrames = nframes
			}
		case "delay":
			if delay, err := strconv.Atoi(v[0]); err == nil {
				opts.Delay = delay
			}
		}
	}

	lissajousWithGIFOpts(w, opts)
}
