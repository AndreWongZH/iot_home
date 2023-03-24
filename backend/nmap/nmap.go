package nmap

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
)

type Nmap struct {
	execPath string
	command  *exec.Cmd
	ipAddr   string
}

// A gin route that respond http with a list of ip addresses
func DiscoverNetworkDevices(ctx *gin.Context) {
	// scan for network devices
	nmap, err := initNmap()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	ipList, err := nmap.findAllDevices()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	fmt.Println(ipList)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ipList,
	})
}

func initNmap() (*Nmap, error) {
	var nmap Nmap

	if ipAddr, err := getLocalIpAddr(); err == nil {
		nmap.ipAddr = ipAddr
	} else {
		return nil, err
	}

	path, err := exec.LookPath("nmap")
	if err != nil {
		fmt.Println("error finding path to nmap")
		fmt.Println("is nmap installed?")
		return nil, errors.New("error finding path to nmap")
	}

	fmt.Println(path)
	nmap.execPath = path

	nmap.command = exec.Command(nmap.execPath, "-sn", nmap.ipAddr)

	return &nmap, nil
}

func (nmap *Nmap) findAllDevices() ([]string, error) {
	stdout, err := nmap.command.StdoutPipe()
	if err != nil {
		return nil, errors.New("unable to execute command")
	}

	if err := nmap.command.Start(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	b, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := nmap.command.Wait(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	ipList, err := parseOutput(string(b))
	if err != nil {
		return nil, err
	}

	return ipList, nil
}

func parseOutput(stdout string) ([]string, error) {
	// regex, err := regexp.Compile(`Nmap scan report for .* \(.*\)`)
	// if err != nil {
	// 	return nil, errors.New("compile regex error")
	// }

	ipRegex, err := regexp.Compile(`[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*`)
	if err != nil {
		return nil, errors.New("compile regex error")
	}

	// stringList := regex.FindAllString(stdout, -1)
	ipList := ipRegex.FindAllString(stdout, -1)
	if ipList == nil {
		return nil, errors.New("no ip address found")
	}

	return ipList, nil
}

// only returns the first valid ipv4 address
// if no valid ipv4 address found, will throw an error
func getLocalIpAddr() (string, error) {
	var ipList []*net.IPNet

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", errors.New("error acquiring ip addresses")
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipList = append(ipList, ipnet)
			}
		}
	}

	if len(ipList) == 0 {
		return "", errors.New("ipList is empty")
	}

	if len(ipList) > 1 {
		log.Println("multiple ip address detected")
		log.Println("using: ", ipList[0].String())
	}

	return ipList[0].String(), nil
}
