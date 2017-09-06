package shorty

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	yaml "gopkg.in/yaml.v2"
)

var (
	config     = &Config{}
	configLock = new(sync.RWMutex)
)

type Config struct {
	sync.RWMutex
	Redirects map[string]string
}

// LoadConfig loads the config and watches for a reload signal
func LoadConfig() error {
	if err := readConfig(); err != nil {
		return err
	}
	log.Println("loaded config")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			if err := readConfig(); err != nil {
				log.Println("error reloading config", err)
			} else {
				log.Println("reloaded config")
			}
		}
	}()

	return nil
}

func readConfig() error {
	configLock.Lock()
	defer configLock.Unlock()

	data, err := ioutil.ReadFile("redirects.yml")
	if err != nil {
		return err
	}

	temp := &Config{}
	if err = yaml.Unmarshal(data, temp); err != nil {
		return err
	}

	config = temp
	return nil
}
