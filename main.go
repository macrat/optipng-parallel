package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func Worker(in, out chan string) {
	for target := range in {
		before, err := os.Stat(target)
		if err != nil {
			out <- fmt.Sprintf("%s: %s", target, err)
			continue
		}

		err = exec.Command("optipng", append(os.Args[1:], target)...).Run()
		if err != nil {
			out <- fmt.Sprintf("%s: %s", target, err)
			continue
		}

		after, err := os.Stat(target)
		if err != nil {
			out <- fmt.Sprintf("%s: %s", target, err)
			continue
		}

		out <- fmt.Sprintf("%s: %dB -> %dB (%d%%)", target, before.Size(), after.Size(), after.Size()*100/before.Size())
	}
}

func main() {
	out := make(chan string, 1000)
	defer close(out)
	in := make(chan string, 1000)
	defer close(in)

	for i := 0; i < runtime.NumCPU(); i++ {
		go Worker(in, out)
	}

	total := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		total++
		in <- scan.Text()
	}

	for count := 1; count <= total; count++ {
		log.Printf("%d/%d(%d%%): %s", count, total, count*100/total, <-out)
	}
}
