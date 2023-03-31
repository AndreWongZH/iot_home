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

	"github.com/AndreWongZH/iothome/logger"
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
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "cannot find path to nmap")
		return nil, err
	}

	nmap.execPath = path

	nmap.command = exec.Command(nmap.execPath, "-sn", nmap.ipAddr)

	logger.SugarLog.Info("nmap initialized")

	return &nmap, nil
}

func (nmap *Nmap) findAllDevices() ([]string, error) {
	stdout, err := nmap.command.StdoutPipe()
	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "unable to execute command")
		return nil, err
	}

	if err := nmap.command.Start(); err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "nmap command cannot start")
		return nil, err
	}

	b, err := io.ReadAll(stdout)
	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "unable to read stdout")
		return nil, err
	}

	if err := nmap.command.Wait(); err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "unable to execute command")
		return nil, err
	}

	ipList, err := parseOutput(string(b))
	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "unable to parse stdout")
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
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "unable to compile regex")
		return nil, err
	}

	// stringList := regex.FindAllString(stdout, -1)
	ipList := ipRegex.FindAllString(stdout, -1)
	if ipList == nil {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "no ip address found")
		return nil, err
	}

	return ipList, nil
}

// only returns the first valid ipv4 address
// if no valid ipv4 address found, will throw an error
func getLocalIpAddr() (string, error) {
	var ipList []*net.IPNet

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "unable to get interface address")
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipList = append(ipList, ipnet)
			}
		}
	}

	if len(ipList) == 0 {
		logger.SugarLog.Errorw(err.Error(), "location", "nmap", "extra", "no interface address")
		return "", errors.New("ipList is empty")
	}

	if len(ipList) > 1 {
		log.Println("multiple ip address detected")
		log.Println("using: ", ipList[0].String())
		logger.SugarLog.Warn("multiple ip address detected, using:", ipList[0].String())
	}

	return ipList[0].String(), nil
}
