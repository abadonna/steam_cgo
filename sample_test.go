package steam_cgo

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var steam Steamworks

func TestMain(m *testing.M) {
	steam = Steamworks{}
	if steam.Init() {
		defer steam.Shutdown()
		steam.RequestUserStats(onUserStatsReceived)
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func onUserStatsReceived() {
	fmt.Println("stats recieved!")

	for name, value := range steam.GetAllAchievements() {
		fmt.Println(name, value)
	}

	fmt.Println("Test ACH_TRAVEL_FAR_ACCUM", steam.GetAchievement("ACH_TRAVEL_FAR_ACCUM"))
	steam.UnlockAchievement("ACH_TRAVEL_FAR_ACCUM")
	fmt.Println("Is ACH_TRAVEL_FAR_ACCUM unlocked now?", steam.GetAchievement("ACH_TRAVEL_FAR_ACCUM"))
}
