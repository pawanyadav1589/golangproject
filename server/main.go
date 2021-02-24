
package main

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}
}

func main() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	fmt.Println(os.Getenv("DB_NAME"))

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("SESSIONTIMEOUT"))

	seed.Load(server.DB)

	server.Run(os.Getenv("SERVER_PORT"))

}