// +build cgo

package steam_cgo

/*
#cgo darwin LDFLAGS: -L${SRCDIR} -lsteam_api

#include "./wrapper.h"
*/
import "C"
import (
	"fmt"
	"time"
)

//Callback - listener func
type Callback func()

//Steamworks client
type Steamworks struct {
	appID             uint
	userStatsListener Callback
}

const (
	userStatsRecieved = 1
)

//Init - call before use
func (client *Steamworks) Init() bool {
	client.appID = uint(C.Init())
	return client.appID > 0
}

//Shutdown - should be called during process shutdown if possible.
func (client *Steamworks) Shutdown() {
	fmt.Println("shutting down")
	C.Shutdown()
}

//IsActive - true if client is ready.
func (client *Steamworks) IsActive() bool {
	return client.appID > 0
}

//RequestUserStats - call and wait for callback
func (client *Steamworks) RequestUserStats(callback Callback) bool {
	client.userStatsListener = callback

	go func() {
		for range time.Tick(time.Second) {
			i := C.Dispatch(C.uint(client.appID))
			if i == userStatsRecieved {
				client.userStatsListener()
				return
			}
		}
	}()

	return bool(C.RequestUserStats())
}

//GetAllAchievements -
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

//GetAchievement -
func (client *Steamworks) GetAchievement(name string) bool {
	value := C.GetAchievement(C.CString(name))
	return bool(value)
}

//UnlockAchievement -
func (client *Steamworks) UnlockAchievement(name string) bool {
	value := C.UnlockAchievement(C.CString(name))
	return bool(value)
}
