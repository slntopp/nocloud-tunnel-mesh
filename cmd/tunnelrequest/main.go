package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	pb "github.com/slntopp/nocloud-tunnel-mesh/pkg/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("TUNNEL_HOST", "localhost:8080")
	viper.SetDefault("SECURE", true)
}

func grpcClient() {

	host := viper.GetString("TUNNEL_HOST")

	var opts []grpc.DialOption
	// opts = append(opts, grpc.WithInsecure())
	// if viper.GetBool("SECURE") {
	// 	cred := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	// 	opts[0] = grpc.WithTransportCredentials(cred)
	// }

	if viper.GetBool("SECURE") {
		// Load client cert
		//cert, err := tls.LoadX509KeyPair("cert/0client.crt", "cert/0client.key")
		cert, err := tls.LoadX509KeyPair("cert/1client.crt", "cert/1client.key")
		if err != nil {
			log.Fatal("fail to LoadX509KeyPair:", err)
		}

		// // Load CA cert
		// //Certification authority, CA
		// //A CA certificate is a digital certificate issued by a certificate authority (CA), so SSL clients (such as web browsers) can use it to verify the SSL certificates sign by this CA.
		// caCert, err := ioutil.ReadFile("../cert/cacerts.cer")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// caCertPool := x509.NewCertPool()
		// caCertPool.AppendCertsFromPEM(caCert)

		// Setup HTTPS client
		config := &tls.Config{
			Certificates: []tls.Certificate{cert},
			// RootCAs:            caCertPool,
			// InsecureSkipVerify: false,
			InsecureSkipVerify: true,
		}
		// config.BuildNameToCertificate()
		cred := credentials.NewTLS(config)

		// cred := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		// opts[0] = grpc.WithTransportCredentials(cred)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatal("fail to dial:", err)
	}
	defer conn.Close()

	log.Println("Connected to server", host)

	client := pb.NewTunnelClient(conn)

	stdreader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("c2s > ")
		note, _ := stdreader.ReadString('\n')

		// if err := stream.Send(&pb.StreamData{Message: note}); err != nil {
		// 	lg.Fatal("Failed to send a note:", zap.Error(err))

		r, err := client.SendData(context.Background(), &pb.SendDataRequest{Host: "ClientZero", Message: note})
		if err != nil {
			log.Printf("could not SendData: %v", err)
		}
		log.Printf("Greeting c: %v", r.GetResult())

	}

}
func restClient() {

	stdreader := bufio.NewScanner(os.Stdin)
	fmt.Print("c2s > ")
	for stdreader.Scan() {

		// // localAddr, err := net.ResolveIPAddr("ip", "localhost")
		// localAddr, err := net.ResolveIPAddr("ip", "127.0.0.1")
		// if err != nil {
		// 	fmt.Println("Failed to read responce", err)
		// 	return
		// }

		netClient := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				// DialContext: (&net.Dialer{
				// 		LocalAddr: &net.TCPAddr{
				// 			// IP: net.ParseIP("127.0.0.1"), //localhost
				// 			// IP: net.ParseIP("localhost"),
				// 			IP:localAddr.IP,
				// 			Port: 8090,
				// 		},
				// 	Timeout:   30 * time.Second,
				// 	KeepAlive: 30 * time.Second,
				// 	DualStack: true,
				// }).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Timeout: time.Second * 30,
		}

		netClient.Transport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			// if addr == "google.com:443" {
			//     addr = "216.58.198.206:443"
			addr = "localhost:8090"
			// }
			dialer := &net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}
			return dialer.DialContext(ctx, network, addr)
		}

		// req, err := http.NewRequest("GET", "http://node.com/sometestpass/"+stdreader.Text(), nil)
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Print("pppp2s > ")
		// response, err := netClient.Do(req)
		response, err := netClient.Get("http://zero.client.net/sometestpass/" + stdreader.Text())
		// response, err := http.Get("http://localhost:8090/sometestpass/" + stdreader.Text())
		if err != nil {
			fmt.Println("Failed to get http", err)
			return
		}

		resp, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to read responce")
			return
		}

		fmt.Printf("response: %v\n", string(resp))
		fmt.Print("c2s > ")

	}
}
func main() {
	// grpcClient()

	restClient()
}
