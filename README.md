#  Steamworks 1.51 wrapper for GO to get\set achievements

Do not forget to copy steam_api library.

## Use example

```
var steam Steamworks

func main() {
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
```
