package main

/*
HackType
A fun mini recon tool built while learning Go.

This version asks the user to type a domain directly in the terminal.
*/

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {

	var target string

	// Ask the user for a domain
	fmt.Print("Enter target domain: ")
	fmt.Scanln(&target)

	fmt.Println("\n[+] Target locked:", target)
	time.Sleep(time.Second)

	// Step 1: Find the IP address
	resolveIP(target)

	fmt.Println("\n[+] Running Recon Scan...\n")

	// Step 2: Scan ports
	scanPorts(target)

	// Step 3: Check web server info
	checkHTTP(target)

	fmt.Println("\n[✓] Scan Complete.")
}

// Resolve domain to IP
func resolveIP(domain string) {

	fmt.Println("[+] Resolving IP...")
	time.Sleep(time.Second)

	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Could not resolve IP")
		return
	}

	for _, ip := range ips {
		fmt.Println("[✓] IP Found:", ip)
	}
}

// Scan common ports
func scanPorts(domain string) {

	ports := []string{"80", "443", "22", "8080"}

	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()

			address := net.JoinHostPort(domain, p)

			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err != nil {
				fmt.Println("Port", p, "-> CLOSED")
				return
			}

			conn.Close()
			fmt.Println("Port", p, "-> OPEN")

		}(port)
	}

	wg.Wait()
}

// Check HTTP info
func checkHTTP(domain string) {

	url := "http://" + domain

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP check failed")
		return
	}

	fmt.Println("\nServer:", resp.Header.Get("Server"))
	fmt.Println("Status:", resp.Status)
}