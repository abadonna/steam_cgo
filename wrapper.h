
#include "stdio.h"
#include "stdint.h"
#include <stdlib.h>
#include <stdbool.h>

uint32_t Init();
void Shutdown();
bool RequestUserStats();
int Dispatch(uint32_t gameID);
bool GetAchievement(const char * name);
uint32_t GetNumAchievements();
const char* GetAchievementName(uint32_t index);
bool UnlockAchievement(const char * name);
bool StoreStats();
