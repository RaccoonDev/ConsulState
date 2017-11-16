package temp

//import (
//	"log"
//	"testing"
//
//	consulapi "github.com/hashicorp/consul/api"
//)
//
///*** Test code start ***/
//type myKbStub struct{}
//
//func (s myKbStub) Get(key string, q *consulapi.QueryOptions) (*consulapi.KVPair, *consulapi.QueryMeta, error) {
//	log.Println("Trying to get value from KV")
//	return nil, nil, nil
//}
//
//func TestMyTest(t *testing.T) {
//	ch := make(chan string)
//	subscribeToChanges("", ch, &myKbStub{})
//}
