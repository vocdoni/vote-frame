import { useCallback, useEffect, useMemo, useState } from 'react'
import { appUrl } from '~constants'
import { userToProfile } from '~util/mappings'

export type AuthState = ReturnType<typeof useAuthProvider>

const baseRep = {
  activity: {
    followersCount: 0,
    electionsCreated: 0,
    castedVotes: 0,
    participationAchievement: 0,
    communitiesCount: 0,
  },
  activityCounts: {
    followersCount: 0,
    electionsCreated: 0,
    castedVotes: 0,
    participationAchievement: 0,
    communitiesCount: 0,
  },
  boosters: {
    hasVotecasterNFTPass: false,
    hasVotecasterLaunchNFT: false,
    isVotecasterAlphafrensFollower: false,
    isVotecasterFarcasterFollower: false,
    isVocdoniFarcasterFollower: false,
    votecasterAnnouncementRecasted: false,
    hasKIWI: false,
    hasDegenDAONFT: false,
    hasHaberdasheryNFT: false,
    has10kDegenAtLeast: false,
    hasTokyoDAONFT: false,
    has5ProxyAtLeast: false,
    hasProxyStudioNFT: false,
    hasNameDegen: false,
    hasFarcasterOGNFT: false,
  },
  points: {
    ownerPoints: 0,
    voterPoints: 0,
    totalPoints: 0,
  },
  totalReputation: 0,
  activityInfo: {
    maxCastedReputation: 0,
    maxCommunityReputation: 0,
    maxElectionsReputation: 0,
    maxFollowersReputation: 0,
    maxReputation: 0,
    maxVotesReputation: 0,
  },
  boostersInfo: {
    degenAtLeast10kPuntuaction: 0,
    degenDAONFTPuntuaction: 0,
    farcasterOGNFTPuntuaction: 0,
    haberdasheryNFTPuntuaction: 0,
    kiwiPuntuaction: 0,
    nameDegenPuntuaction: 0,
    proxyAtLeast5Puntuaction: 0,
    proxyStudioNFTPuntuaction: 0,
    tokyoDAONFTPuntuaction: 0,
    vocdoniFarcasterFollowerPuntuaction: 0,
    votecasterAlphafrensFollowerPuntuaction: 0,
    votecasterAnnouncementRecastedPuntuaction: 0,
    votecasterFarcasterFollowerPuntuaction: 0,
    votecasterLaunchNFTPuntuaction: 0,
    votecasterNFTPassPuntuaction: 0,
  },
}

export type Reputation = typeof baseRep

type LoginParams = {
  profile: Profile
  bearer: string
  reputation: Reputation
}

export const useAuthProvider = () => {
  const [bearer, setBearer] = useState<string | null>(localStorage.getItem('bearer'))
  const [profile, setProfile] = useState<Profile | null>(JSON.parse(localStorage.getItem('profile') || 'null'))
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const [reputation, setReputation] = useState<Reputation | undefined>(
    JSON.parse(localStorage.getItem('reputation') || '{}')
  )

  const isAuthenticated = useMemo(() => !!bearer && !!profile && !!reputation, [bearer, profile, reputation])

  const bearedFetch = useCallback(
    async (input: RequestInfo, init: RequestInit = {}) => {
      const headers = new Headers(init.headers || {})
      if (bearer) {
        headers.append('Authorization', `Bearer ${bearer}`)
      }
      const updatedInit = { ...init, headers }
      return fetch(input, updatedInit).then(async (response) => {
        if (!response.ok) {
          const text = await response.text()
          const sanitized = text.replace('\n', '')
          throw new Error(sanitized.length ? sanitized : response.statusText)
        }

        return response
      })
    },
    [bearer]
  )

  const storeReputation = ({ reputation }: { reputation: Reputation }) => {
    setReputation(reputation)
    localStorage.setItem('reputation', JSON.stringify(reputation))

    return rep
  }

  const tokenLogin = useCallback((token: string) => {
    setError(null)
    setLoading(true)
    return bearedFetch(`${appUrl}/profile`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((resp) => resp.json())
      .then(({ user, reputation, reputationData }: UserProfileResponse) =>
        login({
          profile: userToProfile(user),
          bearer: token,
          reputation: {
            reputation,
            data: {
              ...reputationData,
            },
          },
        })
      )
      .catch((err) => {
        setError(err.message)
      })
      .finally(() => setLoading(false))
  }, [])

  const searchParamsTokenLogin = useCallback(
    (search: string) => {
      const params = new URLSearchParams(search.replace(/^\?/, ''))
      const token = params.get('token')

      if (!token || isAuthenticated) return

      tokenLogin(token)
    },
    [isAuthenticated]
  )

  // if no bearer but profile, logout
  useEffect(() => {
    if (!bearer && !!profile) {
      logout()
    }
  }, [bearer, profile])

  // check if the token is still valid
  useEffect(() => {
    if (!bearer) return

    bearedFetch(`${appUrl}/auth/check`)
      // Invalid token or expired, so logout
      .then(async (response) => {
        if (response.status !== 200) {
          logout()
        }

        return response.json()
      })
      // update reputation
      .then(storeReputation)
      // network errors or other issues
      .catch(() => {
        logout()
      })
  }, [])

  const login = useCallback(({ profile, bearer, reputation }: LoginParams) => {
    localStorage.setItem('bearer', bearer)
    localStorage.setItem('profile', JSON.stringify(profile))
    localStorage.setItem('reputation', JSON.stringify(reputation))
    setBearer(bearer)
    setProfile(profile)
    setReputation(reputation)
  }, [])

  const logout = useCallback(() => {
    localStorage.removeItem('bearer')
    localStorage.removeItem('profile')
    localStorage.removeItem('reputation')
    setBearer(null)
    setProfile(null)
    setReputation(undefined)
  }, [])

  return {
    bearer,
    bfetch: bearedFetch,
    error,
    isAuthenticated,
    loading,
    login,
    logout,
    profile,
    reputation,
    searchParamsTokenLogin,
    tokenLogin,
  }
}
