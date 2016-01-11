// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"services"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void createCity(string token, string title,  coords)")
	fmt.Fprintln(os.Stderr, "   getNearBy(Coordinate coordinate, i32 limit)")
	fmt.Fprintln(os.Stderr, "   getAllCities()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := services.NewCitySvcClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "createCity":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "CreateCity requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg191 := flag.Arg(3)
		mbTrans192 := thrift.NewTMemoryBufferLen(len(arg191))
		defer mbTrans192.Close()
		_, err193 := mbTrans192.WriteString(arg191)
		if err193 != nil {
			Usage()
			return
		}
		factory194 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt195 := factory194.GetProtocol(mbTrans192)
		containerStruct2 := services.NewCitySvcCreateCityArgs()
		err196 := containerStruct2.ReadField3(jsProt195)
		if err196 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.Coords
		value2 := argvalue2
		fmt.Print(client.CreateCity(value0, value1, value2))
		fmt.Print("\n")
		break
	case "getNearBy":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetNearBy requires 2 args")
			flag.Usage()
		}
		arg197 := flag.Arg(1)
		mbTrans198 := thrift.NewTMemoryBufferLen(len(arg197))
		defer mbTrans198.Close()
		_, err199 := mbTrans198.WriteString(arg197)
		if err199 != nil {
			Usage()
			return
		}
		factory200 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt201 := factory200.GetProtocol(mbTrans198)
		argvalue0 := services.NewCoordinate()
		err202 := argvalue0.Read(jsProt201)
		if err202 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err203 := (strconv.Atoi(flag.Arg(2)))
		if err203 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.GetNearBy(value0, value1))
		fmt.Print("\n")
		break
	case "getAllCities":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetAllCities requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetAllCities())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
