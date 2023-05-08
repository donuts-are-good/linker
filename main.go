package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

type NetworkConnection struct {
	Proto          string
	LocalIP        string
	LocalPort      string
	RemoteIP       string
	RemotePort     string
	ProgramCommand string
}

func listNetworkConnections() ([]NetworkConnection, error) {

	// `lsof` to get network connections
	cmd := exec.Command("lsof", "-i", "-P", "-n", "-F", "cna")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// read the output line by line
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	scanner.Split(bufio.ScanLines)

	var connections []NetworkConnection
	var connection NetworkConnection

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "p"):

			connection = NetworkConnection{}
			connection.ProgramCommand = strings.TrimPrefix(line, "p")
		case strings.HasPrefix(line, "c"):

			connection.Proto = strings.TrimPrefix(line, "c")
		case strings.HasPrefix(line, "n"):
			addr := strings.TrimPrefix(line, "n")
			parts := strings.Split(addr, "->")
			if len(parts) != 2 {
				continue
			}
			localAddr := strings.Split(parts[0], ":")
			remoteAddr := strings.Split(parts[1], ":")

			if len(localAddr) != 2 || len(remoteAddr) != 2 {
				continue
			}

			// assign the local and remote ip:port info
			connection.LocalIP = localAddr[0]
			connection.LocalPort = localAddr[1]
			connection.RemoteIP = remoteAddr[0]
			connection.RemotePort = remoteAddr[1]

			connections = append(connections, connection)
		}
	}

	// get the program name for this pid
	for i := range connections {
		connections[i].ProgramCommand, err = getBinaryName(connections[i].ProgramCommand)
		if err != nil {
			return nil, err
		}
	}

	return connections, nil
}

func getBinaryName(pid string) (string, error) {

	// use ps to cross reference the pid
	cmd := exec.Command("ps", "-p", pid, "-o", "comm=")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// clean it up
	binaryName := strings.TrimSpace(string(output))
	return filepath.Base(binaryName), nil
}

func displayNetworkConnections() {
	connections, err := listNetworkConnections()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// try our best to align the otuput
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Proto\t\t\t\tLocal IP\tLocal Port\tRemote IP\tRemote Port\tBinary Name")
	fmt.Fprintln(w, "-----------------------------------------------------------------------------")
	for _, conn := range connections {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", conn.Proto, conn.LocalIP, conn.LocalPort, conn.RemoteIP, conn.RemotePort, conn.ProgramCommand)
	}
	w.Flush()
}

func main() {
	displayNetworkConnections()
}
