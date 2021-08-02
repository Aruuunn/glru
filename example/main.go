package main

import (
	"fmt"
	"time"

	"github.com/arunmurugan78/glru"
)

func main() {
	cache := glru.New(glru.Config{
		MaxItems: 1000,
	})

	cache.Set("Pizza", "🍕")
	cache.Set("Time", time.Now())
	cache.Set("Beer", "🍺")
	cache.Set("Struct", struct{ Name string }{Name: "Arun"})

	if pizza, err := cache.Get("Pizza"); err == nil {
		fmt.Println(pizza) // Output: 🍕
	}

	if val, err := cache.Get("Struct"); err == nil {
		fmt.Println(val) // Output: {Arun}
	}

	_, err := cache.Get("SomeKeyWhichDoesntExist")
	fmt.Println(err) // Output: key not found

	cache.Delete("Pizza")
	_, err = cache.Get("Pizza")
	fmt.Println(err) // Output: key not found
}
