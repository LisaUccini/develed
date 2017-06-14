package main

import (
	"context"
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/develed/develed/config"
	srv "github.com/develed/develed/services"

	"google.golang.org/grpc"
)

var (
	cfg = flag.String("config", "/etc/develed.toml", "configuration file")
)

func main() {
	flag.Parse()
	conf, err := config.Load(*cfg)
	if err != nil {
		log.Fatalln(err)
	}

	text := "ciao"
	if flag.NArg() >= 1 {
		text = flag.Arg(0)
	}

	conn, err := grpc.Dial(conf.Textd.GRPCServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	textd := srv.NewTextdClient(conn)
	_, err = textd.Write(context.Background(), &srv.TextRequest{
		Text:      text,
		FontColor: 0xFFAABBCC,
		FontBg:    0x00112233,
	})
	if err != nil {
		log.Errorln(err)
	}
}
