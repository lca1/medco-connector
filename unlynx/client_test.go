package unlynx

import (
	"github.com/lca1/medco-connector/util"
	"github.com/lca1/unlynx/lib"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

// warning: needs medco deployment dev-3nodes-local running

func init() {
	util.SetLogLevel("5")
	util.UnlynxGroupFileIdx = 0

	_, filename, _, _ := runtime.Caller(0)
	util.UnlynxGroupFilePath = filepath.Dir(filename) + "/test_group.toml"
	setupUnlynxClient()
}

func TestGetQueryTermsDDT(t *testing.T) {
	encryptedInt0 := libunlynx.EncryptInt(cothorityRoster.Aggregate, 366).Serialize()
	encryptedInt1 := libunlynx.EncryptInt(cothorityRoster.Aggregate, 2).Serialize()
	t.Log(encryptedInt0, encryptedInt1)

	tags, err := GetQueryTermsDDT("test query " + time.Now().Format(time.RFC3339Nano), []string{
		encryptedInt0,
		encryptedInt1,
	})

	if err != nil {
		t.Fail()
	}
	t.Log(tags)
}
//
//func TestAggregateAndKeySwitchDummyFlags(t *testing.T) {
//	privKey, pubKey := libunlynx.GenKey()
//	pubKeySer, err := libunlynx.SerializePoint(pubKey)
//	if err != nil {
//		t.Fail()
//	}
//
//	value := make(chan int64)
//	for i := 0 ; i < 3 ; i++ {
//
//		//fmt.Println("id:", i)
//		//fmt.Println(cothorityRoster.List[i])
//		iCp := i
//		go func() {
//			util.UnlynxGroupFileIdx = iCp
//			fmt.Println("id:", iCp)
//			fmt.Println(cothorityRoster.List[iCp])
//			setupUnlynxClient()
//
//			agg, err := AggregateAndKeySwitchDummyFlags(
//				"test query " + time.Now().Format(time.RFC3339Nano),
//				[]string{
//					libunlynx.EncryptInt(cothorityRoster.Aggregate, 0).Serialize(),
//					libunlynx.EncryptInt(cothorityRoster.Aggregate, 1).Serialize(),
//				},
//				pubKeySer,
//			)
//
//			if err != nil {
//				t.Fail()
//			}
//
//			aggDes := &libunlynx.CipherText{}
//			err = aggDes.Deserialize(agg)
//			if err != nil {
//				t.Fail()
//			}
//
//
//			value <-libunlynx.DecryptInt(privKey, *aggDes)
//		}()
//
//
//	}
//	t.Log(<-value)
//	t.Log(<-value)
//	t.Log(<-value)
//}
