package temp

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

//type KvGetInterface interface {
//	Get(key string, q *consulapi.QueryOptions) (*consulapi.KVPair, *consulapi.QueryMeta, error)
//}

//func main() {
//	c, err := consulapi.NewClient(consulapi.DefaultConfig())
//	if err != nil {
//		log.Panicln("Error creating consul client", err)
//	}
//
//	kv := c.KV()
//
//	ch := make(chan string)
//
//	go subscribeToChanges("loglevel", ch, kv)
//
//	go func() {
//		for data := range ch {
//			log.Println("Detected key change:", data)
//		}
//	}()
//
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Println("Press ENTER to exit...")
//	reader.ReadString('\n')
//}

//func subscribeToChanges(key string, ch chan string, kv KvGetInterface) {
//	currentIndex := uint64(0)
//	for {
//		pair, meta, err := kv.Get(key, &consulapi.QueryOptions{
//			WaitIndex: currentIndex,
//		})
//
//		if err != nil {
//			log.Panicln("Error read from KV", err.Error(), err)
//		}
//
//		if pair == nil || meta == nil {
//			// Query won't be blocked if key not found
//			time.Sleep(1 * time.Second)
//		} else {
//			ch <- string(pair.Value)
//			currentIndex = meta.LastIndex
//		}
//	}
//}
