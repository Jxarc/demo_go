package main

import (
	api "github.com/axel526/jikkosoft/src/controllers"
)

const APIPort = "8090"

func main() { api.InitAPI(APIPort) }
