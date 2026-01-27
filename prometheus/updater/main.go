package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Hosts struct {
	Hosts []Host `xml:"host"`
}

type Host struct {
	Address string `xml:"address"`
}

const (
	hostsFile   = "../hosts.xml"
	targetsFile = "../targets.yml"
	promReload  = "http://localhost:9090/-/reload"
)

func main() {
	for {
		updateTargets()
		time.Sleep(10 * time.Second)
	}
}

func updateTargets() {
	xmlData, err := os.ReadFile(hostsFile)
	if err != nil {
		fmt.Println("Error reading XML:", err)
		return
	}

	var hosts Hosts
	if err := xml.Unmarshal(xmlData, &hosts); err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	file, err := os.Create(targetsFile)
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

	reloadPrometheus()
}

func reloadPrometheus() {
	req, err := http.NewRequest("POST", promReload, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Println("Error creating reload request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error reloading Prometheus:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Prometheus reloaded successfully")
	} else {
		fmt.Println("Prometheus reload failed, status:", resp.Status)
	}
}
