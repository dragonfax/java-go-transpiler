package steam

import "github.com/dragonfax/delver_converted/com/badlogic/gdx/utils"

type SteamApiInterface interface {
    Init() bool
    GetWorkshopFolders() utils.Array[String]
    RunCallbacks()
    Achieve(achievementName string)
    Achieve(achievementName string, numProgress int, maxProgress int)
    Dispose()
    UploadToWorkshop(workshopId int64, modImagePath string, modTitle string, modFolderPath string)
    IsAvailable() bool
}
