package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func main() {
	 // Check if file is provided as an argument
	 if len(os.Args) < 2 {
        fmt.Println("Please provide a file with a list of addresses as an argument.")
        return
    }

    // Open file for reading
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a http client with custom transport that skips certificate verification errors
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Loop through each line in the file
	for scanner.Scan() {
		address := scanner.Text()

		// Send a request to the address and check the certificate parameters
		resp, err := client.Get("https://" + address)
		if err != nil {
			fmt.Printf("%s: %s\n", address, err)
			continue
		}
		defer resp.Body.Close()

		cert := resp.TLS.PeerCertificates[0]
		fmt.Printf("Address: %s\n", address)
		fmt.Printf("Common Name: %s\n", cert.Subject.CommonName)
		fmt.Printf("Issuer: %s\n", cert.Issuer.CommonName)
		fmt.Printf("Valid from: %s to %s\n", cert.NotBefore, cert.NotAfter)
		fmt.Printf("Signature algorithm: %s\n", cert.SignatureAlgorithm)
		fmt.Printf("Key algorithm: %s\n", cert.PublicKeyAlgorithm)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
}
