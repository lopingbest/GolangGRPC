package main

import (
	"context"
	"fmt"
	"github.com/lopinhbest/GolangGRPC/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Galih",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println("Blog has been created: %v", createBlogRes)
	blogId := createBlogRes.GetBlog().GetId()

	//read Blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5f1e7f8e-d8e2-4f0f-b8e3-f8f8f8f8f8f8"})
	if err2 != nil {
		fmt.Printf("Error happend while reading: %v \n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happend while reading: %v \n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v \n", readBlogRes)

	//update Blog
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error happend while updating: %v \n", updateErr)
	}
	fmt.Printf("Blog was updated: %v \n", updateRes)

}
