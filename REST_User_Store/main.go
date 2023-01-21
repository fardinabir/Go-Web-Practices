package main

func main() {
	intialMigration() // DB migration (dbConn.go)
	handleRequests()  //  Routers (httpRouter.go)
}
