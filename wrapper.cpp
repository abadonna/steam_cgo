class nullptr_t 
{
  public:
    template<class T>
    inline operator T*() const // convertible to any type of null non-member pointer...
    { return 0; }

    template<class C, class T>
    inline operator T C::*() const   // or any type of null member pointer...
    { return 0; }

  private:
    void operator&() const;  // Can't take address of nullptr

} nullptr = {};

#include "./sdk/public/steam/steam_api.h"

extern "C" {
#include "wrapper.h"
}

enum CallbackType {
    NONE,         //0
    USERSTATS    //1
};

uint32_t Init() {
	if (SteamAPI_Init()) {
		SteamAPI_ManualDispatch_Init();
		return SteamUtils()->GetAppID();
	}
	return 0;
}

void Shutdown() {
	SteamAPI_Shutdown();
}

bool RequestUserStats() {
	return SteamUserStats()->RequestCurrentStats();
}

int Dispatch(uint32_t gameID) {
	HSteamPipe hSteamPipe = SteamAPI_GetHSteamPipe();
	SteamAPI_ManualDispatch_RunFrame( hSteamPipe );
	CallbackMsg_t callback;

	while (SteamAPI_ManualDispatch_GetNextCallback( hSteamPipe, &callback ) )
	{
		// Check for dispatching API call results
		if (callback.m_iCallback == SteamAPICallCompleted_t::k_iCallback )
		{
			SteamAPICallCompleted_t *pCallCompleted = (SteamAPICallCompleted_t *)callback.m_pubParam;
			void *pTmpCallResult = malloc( callback.m_cubParam );
			bool bFailed;
			if ( SteamAPI_ManualDispatch_GetAPICallResult( hSteamPipe, pCallCompleted->m_hAsyncCall, pTmpCallResult, callback.m_cubParam, callback.m_iCallback, &bFailed ) )
			{
				// Dispatch the call result to the registered handler(s) for the
				// call identified by pCallCompleted->m_hAsyncCall
			}
			free( pTmpCallResult );
		}
		else
		{
			if (callback.m_iCallback == UserStatsReceived_t::k_iCallback )
			{
				UserStatsReceived_t *pCallCompleted = (UserStatsReceived_t *)callback.m_pubParam;
				if (gameID == pCallCompleted->m_nGameID)
				{
					SteamAPI_ManualDispatch_FreeLastCallback(hSteamPipe);
					return USERSTATS;
				}
		
			}
		}
		SteamAPI_ManualDispatch_FreeLastCallback( hSteamPipe );
	}

	return NONE;
}

bool GetAchievement(const char * name) {
	bool b;
	SteamUserStats()->GetAchievement(name, &b);
	return b;
}

uint32_t GetNumAchievements() {
	return SteamUserStats()->GetNumAchievements();
}

const char* GetAchievementName(uint32_t index) {
	return SteamUserStats()->GetAchievementName(index);
}

bool UnlockAchievement(const char * name){
	return SteamUserStats()->SetAchievement(name);
}