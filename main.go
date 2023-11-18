package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var targetIP string
	var customPayload string
	var intervalStr string

	// Get target IP address from the user
	fmt.Print("Enter the target IP address: ")
	fmt.Scanln(&targetIP)

	// Get custom payload from the user
	fmt.Print("Enter custom payload: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		customPayload = scanner.Text()
	}

	// Get packet interval from the user
	fmt.Print("Enter the packet interval (e.g., 1s, 500ms): ")
	fmt.Scanln(&intervalStr)

	// Parse the interval duration
	packetInterval, err := time.ParseDuration(intervalStr)
	if err != nil {
		fmt.Printf("Error parsing packet interval: %v\n", err)
		os.Exit(1)
	}

	// Construct the UDP address
	destAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", targetIP, 8080))
	if err != nil {
		fmt.Printf("Error resolving UDP address: %v\n", err)
		os.Exit(1)
	}

	// Infinite loop to send UDP packets at regular intervals
	for {
		// Dial UDP connection
		conn, err := net.DialUDP("udp", nil, destAddr)
		if err != nil {
			fmt.Printf("Error dialing UDP connection: %v\n", err)
			os.Exit(1)
		}

		// Send UDP packet with custom payload
		_, err = conn.Write([]byte(customPayload))
		if err != nil {
			fmt.Printf("Error sending UDP packet: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("UDP packet sent to %s:%d with custom payload: %s\n", targetIP, 8080, customPayload)

		// Close the connection
		conn.Close()

		// Sleep for the specified interval before sending the next packet
		time.Sleep(packetInterval)
	}
}
