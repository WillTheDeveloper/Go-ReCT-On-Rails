package env

import (
	"os"
)

func setEnvironment()
{
	os.Setenv(port, 8080) //PORT OF THE SERVER
	os.Setenv(host, "localhost") //HOST ADDRESS OF THE SERVER

}