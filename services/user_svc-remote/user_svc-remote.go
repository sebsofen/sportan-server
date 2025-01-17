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
	fmt.Fprintln(os.Stderr, "  UserCredentials createUser(string password)")
	fmt.Fprintln(os.Stderr, "  User getMe(string token)")
	fmt.Fprintln(os.Stderr, "  User getUserById(string token, string userid)")
	fmt.Fprintln(os.Stderr, "   getFriends(string token)")
	fmt.Fprintln(os.Stderr, "   getFriendRequests(string token)")
	fmt.Fprintln(os.Stderr, "  void acceptFriendRequest(string token, string userid)")
	fmt.Fprintln(os.Stderr, "  void declineFriendRequest(string token, string userid)")
	fmt.Fprintln(os.Stderr, "  void sendFriendRequest(string token, string userid)")
	fmt.Fprintln(os.Stderr, "  void setProfile(string token, Profile profile)")
	fmt.Fprintln(os.Stderr, "  Token requestToken(string username, string plain_pw)")
	fmt.Fprintln(os.Stderr, "  void setAdmin(string token, string userid, bool admin)")
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
	client := services.NewUserSvcClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "createUser":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateUser requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.CreateUser(value0))
		fmt.Print("\n")
		break
	case "getMe":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMe requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetMe(value0))
		fmt.Print("\n")
		break
	case "getUserById":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetUserById requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.GetUserById(value0, value1))
		fmt.Print("\n")
		break
	case "getFriends":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetFriends requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetFriends(value0))
		fmt.Print("\n")
		break
	case "getFriendRequests":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetFriendRequests requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetFriendRequests(value0))
		fmt.Print("\n")
		break
	case "acceptFriendRequest":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AcceptFriendRequest requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.AcceptFriendRequest(value0, value1))
		fmt.Print("\n")
		break
	case "declineFriendRequest":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DeclineFriendRequest requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.DeclineFriendRequest(value0, value1))
		fmt.Print("\n")
		break
	case "sendFriendRequest":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SendFriendRequest requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.SendFriendRequest(value0, value1))
		fmt.Print("\n")
		break
	case "setProfile":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetProfile requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg114 := flag.Arg(2)
		mbTrans115 := thrift.NewTMemoryBufferLen(len(arg114))
		defer mbTrans115.Close()
		_, err116 := mbTrans115.WriteString(arg114)
		if err116 != nil {
			Usage()
			return
		}
		factory117 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt118 := factory117.GetProtocol(mbTrans115)
		argvalue1 := services.NewProfile()
		err119 := argvalue1.Read(jsProt118)
		if err119 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.SetProfile(value0, value1))
		fmt.Print("\n")
		break
	case "requestToken":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RequestToken requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.RequestToken(value0, value1))
		fmt.Print("\n")
		break
	case "setAdmin":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "SetAdmin requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		fmt.Print(client.SetAdmin(value0, value1, value2))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
