package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

// Declaring portRange and host global variable.
var portRange int
var host string

func argparse() {
	ptrng := flag.Int("port", 1024, "Specify Port Range for Scan.")
	hst := flag.String("target", "scanme.nmap.org", "Specify the target hostname or IP.")
	flag.Parse()
	portRange = int(*ptrng)
	host = string(*hst)
}

func banner() {
	banner := ` _______          __   _________                     
 \      \   _____/  |_/   _____/ ____ _____    ____  
 /   |   \_/ __ \   __\_____  \_/ ___\\__  \  /    \ 
/    |    \  ___/|  | /        \  \___ / __ \|   |  \
\____|__  /\___  >__|/_______  /\___  >____  /___|  /
        \/     \/            \/     \/     \/     \/ 

`
	fmt.Printf("\033[1;032m%s\033[0m", banner)
}

func scanner(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	banner()
	argparse()
	fmt.Printf("            Target: [ \033[1;032m%s\033[0m ]\n\n", host)
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go scanner(ports, results)
	}

	go func() {
		for i := 0; i <= portRange; i++ {
			ports <- i
		}
	}()

	for i := 0; i < portRange; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf(" [\033[032m+\033[0m] Port %d Open\n", port)
	}
}
