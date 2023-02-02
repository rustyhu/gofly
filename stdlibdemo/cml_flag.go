package stdlibdemo

// log "github.com/sirupsen/logrus"
import (
	"flag"
	"fmt"
	"os"
)

func getEnvOrDefault(key, defaultstr string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultstr
}

func ParseFlag() {
	// log.WithFields(log.Fields{"size": 20}).Info("A walrus appears")

	confFile := flag.String("f", getEnvOrDefault("FLY_CONFIG", "config.yaml"), "configuration file")
	switchNotify := flag.Bool("n", os.Getenv("FLY_NOTIFY") == "true", "switch of notification")
	switchShow := flag.Bool("s", false, "switch of show")
	version := flag.Bool("v", false, "prints version")

	flag.Parse()

	if *switchNotify {
		fmt.Println("Flag \"-n\"! ")
	}
	if *switchShow {
		fmt.Println("Flag \"-s\"! ")
	}
	if *version {
		fmt.Println("Flag \"-v\"! ")
	}

	fmt.Println("String flag content:", *confFile)
}
