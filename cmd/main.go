package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Fatal(err)
	}
}
