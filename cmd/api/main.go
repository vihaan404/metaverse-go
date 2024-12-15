package main

type config struct {
	port int
}

type application struct {
	config config
}

func main() {
	app := application{
		config: config{
			port: 4000,
		},
	}
	err := app.serve()
	if err != nil {
		panic(err)
	}
}
