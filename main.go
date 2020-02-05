// steaks-cli/main.go
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/polosate/product-service/proto/product"
)

const (
	defaultFilename = "product.json"
)

func parseFile(file string) (*pb.Product, error) {
	var product *pb.Product
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &product)
	return product, err
}

func Run() {
	service := micro.NewService(micro.Name("steaks.cli.product"))
	service.Init()

	client := pb.NewProductServiceClient("steaks.product.service", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	product, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateProduct(context.Background(), product)

	if err != nil {
		log.Fatalf("Could not create a product: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetProducts(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list products: %v", err)
	}
	for _, v := range getAll.Products {
		log.Println(v)
	}
}

func main() {
	Run()
}
