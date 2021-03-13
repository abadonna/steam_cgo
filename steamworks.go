// +build cgo

package steam_cgo

/*
#cgo CXXFLAGS: -std=c++11
#cgo CPPFLAGS: -isystem ${SRCDIR}/sdk/public
#cgo LDFLAGS: -Wl,-rpath,$ORIGIN

#cgo windows,386 LDFLAGS: -L ${SRCDIR}/sdk/redistributable_bin
#cgo windows,amd64 LDFLAGS: -L ${SRCDIR}/sdk/redistributable_bin/win64

#cgo linux,386 LDFLAGS: -L ${SRCDIR}/sdk/redistributable_bin/linux32
#cgo linux,amd64 LDFLAGS: -L ${SRCDIR}/sdk/redistributable_bin/linux64

#cgo darwin LDFLAGS: -L ${SRCDIR}/sdk/redistributable_bin/osx

#cgo linux windows,386 darwin LDFLAGS: -lsteam_api
#cgo windows,amd64 LDFLAGS: -lsteam_api64

#include "./wrapper.h"
*/
import "C"
import (
	"time"
)

//Callback - listener func
type Callback func()

//Steamworks client
type Steamworks struct {
	appID uint
}

const (
	userStatsRecieved = 1
)

//Init - call before use other methods
func (client *Steamworks) Init() bool {
	client.appID = uint(C.Init())
	return client.appID > 0
}

//Shutdown - should be called during process shutdown if possible.
func (client *Steamworks) Shutdown() {
	C.Shutdown()
}

//IsValid - true if client is ready.
func (client *Steamworks) IsValid() bool {
	return client.appID > 0
}

//RequestUserStats - call and wait for callback
func (client *Steamworks) RequestUserStats(callback Callback) bool {

	go func() {
		for range time.Tick(time.Second) {
			i := C.Dispatch(C.uint(client.appID))
			if i == userStatsRecieved {

				if callback != nil {
					callback()
				}

				return
			}
		}
	}()

	return bool(C.RequestUserStats())
}

//GetAllAchievements - should be called after receiving callback from RequestUserStats
func (client *Steamworks) GetAllAchievements() map[string]bool {
	achievements := make(map[string]bool)
	count := uint(C.GetNumAchievements())

	var i uint = 0
	for ; i < count; i++ {
		name := C.GetAchievementName(C.uint(i))
		value := C.GetAchievement(name)
		achievements[C.GoString(name)] = bool(value)
	}
	return achievements
}

//GetAchievement - should be called after receiving callback from RequestUserStats
func (client *Steamworks) GetAchievement(name string) bool {
	value := C.GetAchievement(C.CString(name))
	return bool(value)
}

//UnlockAchievement - should be called after receiving callback from RequestUserStats
func (client *Steamworks) UnlockAchievement(name string) bool {
	value := C.UnlockAchievement(C.CString(name))

	go func() {
		for range time.Tick(time.Second) {
			//If StoreStats fails then nothing is sent to the server.
			//It's advisable to keep trying until the call is successful.
			if bool(C.StoreStats()) {
				return
			}
		}
	}()

	return bool(value)
}
