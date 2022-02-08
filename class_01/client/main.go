package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_learn/class_01/pb"
	"log"
	"time"
)

type ProductInfo struct {
	name  string
	desc  string
	price float32
}

func main() {
	client, err := grpc.Dial("localhost:8008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer client.Close()

	c := pb.NewProductInfoClient(client)
	p := ProductInfo{
		name:  "iPhone-13",
		desc:  "Pro-Max",
		price: 6927.55,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx, &pb.Product{
		Name:  p.name,
		Desc:  p.desc,
		Price: p.price,
	})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}

	log.Printf("Product: %v", product.String())
}
