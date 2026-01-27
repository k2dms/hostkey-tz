package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type Hosts struct {
	Hosts []Host `xml:"host"`
}

type Host struct {
	Address string `xml:"address"`
}

func main() {
	for {
		updateTargets()
		time.Sleep(10 * time.Second)
	}
}

func updateTargets() {
	xmlData, err := os.ReadFile("../hosts.xml")
	if err != nil {
		fmt.Println("Error reading XML:", err)
		return
	}

	var hosts Hosts
	if err := xml.Unmarshal(xmlData, &hosts); err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	file, err := os.Create("../targets.yml")
	if err != nil {
		fmt.Println("Error writing targets:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "- targets:")
	for _, h := range hosts.Hosts {
		fmt.Fprintf(file, "  - %s:9100\n", h.Address)
	}
	fmt.Fprintln(file, "  labels:")
	fmt.Fprintln(file, "    job: virtualization")

	fmt.Println("targets.yml updated")
}
