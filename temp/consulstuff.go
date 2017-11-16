package temp

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

func doSomethingWithConsul() {
	c, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		log.Panicln("Error creating consul client", err)
	}

	kv := c.KV()

	// keys, _, err := kv.Keys("", "", nil)
	// if err != nil {
	// 	log.Panicln("error getting list of keys")
	// }

	// log.Println("Retrieved keys.", keys)

	// w, err := kv.Put(&consulapi.KVPair{
	// 	Key:   "configuration/security/token",
	// 	Value: []byte("This is value that I'd like to put into storage"),
	// }, nil)

	// if err != nil {
	// 	log.Panicln("Error storing value", err.Error())
	// }

	// log.Println("Value stored.", w.RequestTime)

	go func() {
		currentIndex := uint64(0)
		for {
			currentIndex = getKeyValueWithIndex(kv, currentIndex)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to exit...")
	reader.ReadString('\n')
}

func getKeyValueWithIndex(kv *consulapi.KV, index uint64) uint64 {
	keyToGet := "myTestKey"
	pair, meta, err := kv.Get(keyToGet, &consulapi.QueryOptions{
		WaitIndex: index,
	})

	if err != nil {
		log.Panicln("Error read from KV", err.Error(), err)
	}

	if pair != nil && meta != nil {
		log.Println("Received value from key/value store: ", string(pair.Value), "meta: ", meta.LastIndex)
		return meta.LastIndex
	}

	log.Println("Check performed. Found nil for pair or meta")
	time.Sleep(1 * time.Second)
	return 0
}
