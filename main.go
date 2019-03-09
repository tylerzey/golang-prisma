package main

import (
	"context"
	"fmt"

	prisma "github.com/tylerzey/golang-prisma/generated/prisma-client"
)

func main() {
	client := prisma.New(nil)
	ctx := context.TODO()

	// Create a new user
	name := "Tyler"
	newUser, err := client.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created new user: %+v\n", newUser)

	users, err := client.Users(nil).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users)
}
