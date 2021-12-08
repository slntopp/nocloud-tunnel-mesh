package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	pb "github.com/slntopp/nocloud-tunnel-mesh/pkg/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("TUNNEL_HOST", "localhost:8080")
	viper.SetDefault("SECURE", false)
}

func main() {

	host := viper.GetString("TUNNEL_HOST")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	if viper.GetBool("SECURE") {
		cred := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		opts[0] = grpc.WithTransportCredentials(cred)
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
