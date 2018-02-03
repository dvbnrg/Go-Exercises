package main

import (
	"fmt"
	// Import the Radix.v2 redis package.
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	// Establish a connection to the Redis server listening on port 6379 of the
	// local machine. 6379 is the default port, so unless you've already
	// changed the Redis configuration file this should work.
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	// Importantly, use defer to ensure the connection is always properly
	// closed before exiting the main() function.
	defer conn.Close()

	// Send our command across the connection. The first parameter to Cmd()
	// is always the name of the Redis command (in this example HMSET),
	// optionally followed by any necessary arguments (in this example the
	// key, followed by the various hash fields and values).
	resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	// Check the Err field of the *Resp object for any errors.

	if resp.Err != nil {
		log.Fatal(resp.Err)
	}

	err = conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8).Err

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Electric Ladyland added!")

	// Issue a HGET command to retrieve the title for a specific album, and use
	// the Str() helper method to convert the reply to a string.
	title, err := conn.Cmd("HGET", "album:1", "title").Str()
	if err != nil {
		log.Fatal(err)
	}

	// Similarly, get the artist and convert it to a string.
	artist, err := conn.Cmd("HGET", "album:1", "artist").Str()
	if err != nil {
		log.Fatal(err)
	}

	// And the price as a float64...
	price, err := conn.Cmd("HGET", "album:1", "price").Float64()
	if err != nil {
		log.Fatal(err)
	}

	// And the number of likes as an integer.
	likes, err := conn.Cmd("HGET", "album:1", "likes").Int()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)
}
