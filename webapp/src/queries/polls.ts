import { useQuery } from '@tanstack/react-query'
import { useReadContract } from 'wagmi'
import { useAuth } from '~components/Auth/useAuth'
import { useHealthcheck } from '~components/Healthcheck/use-healthcheck'
import { appUrl } from '~constants'
import { communityHubAbi } from '~src/bindings'
import { getChain, getContractForChain } from '~util/chain'
import { config } from '~util/rainbow'

export const fetchPollInfo = (bfetch: FetchFunction, electionID: string) => async (): Promise<PollResponse> => {
  const response = await bfetch(`${appUrl}/poll/info/${electionID}`)
  const { poll } = (await response.json()) as { poll: PollResponse }
  return poll
}

export const fetchPollsVoters = (bfetch: FetchFunction, electionId: string) => async (): Promise<string[]> => {
  const response = await bfetch(`${appUrl}/votersOf/${electionId}`)
  const { usernames } = (await response.json()) as { usernames: string[] }
  return usernames
}

export const fetchPollsRemainingVoters = (bfetch: FetchFunction, electionId: string) => async (): Promise<string[]> => {
  const response = await bfetch(`${appUrl}/remainingVotersOf/${electionId}`)
  const { usernames } = (await response.json()) as { usernames: string[] }
  return usernames
}

export const fetchPollsReminders = (bfetch: FetchFunction, electionId: string) => async (): Promise<PollReminders> => {
  const response = await bfetch(`${appUrl}/poll/${electionId}/reminders`)
  const data = await response.json()
  const remindableVoters: Profile[] = []
  for (const fid in data.remindableVoters) {
    remindableVoters.push({
      fid: parseInt(fid),
      username: data.remindableVoters[fid],
    } as Profile)
  }

  const votersWeight: { [key: string]: string } = {}
  for (const fid in data.votersWeight) {
    votersWeight[data.remindableVoters[fid]] = data.votersWeight[fid]
  }
  return {
    remindableVoters,
    alreadySent: data.alreadySent,
    maxReminders: data.maxReminders,
    votersWeight: votersWeight,
  } as PollReminders
}

export const fetchShortURL = (bfetch: FetchFunction) => async (url: string) => {
  const response = await bfetch(`${appUrl}/short?url=${url}`)
  const { result } = (await response.json()) as { result: string }
  return result
}

export const useApiPollInfo = (electionId: string) => {
  const { bfetch } = useAuth()

  return useQuery<PollResponse, Error, PollInfo>({
    queryKey: ['apiPollInfo', electionId],
    queryFn: fetchPollInfo(bfetch, electionId!),
    enabled: !!electionId,
    select: (data) => ({
      ...data,
      totalWeight: Number(data.totalWeight),
      createdTime: new Date(data.createdTime),
      endTime: new Date(data.endTime),
      lastVoteTime: new Date(data.lastVoteTime),
      tally: [data.tally.map((t) => Number(t))],
    }),
  })
}

export const useContractPollInfo = (chainAlias: ChainKey, communityId: number, electionId: string) => {
  const health = useHealthcheck()
  return useReadContract({
    abi: communityHubAbi,
    chainId: getChain(chainAlias).id,
    config: {
      ...config,
      chains: [getChain(chainAlias)],
    },
    address: getContractForChain(chainAlias),
    functionName: 'getResult',
    args: [BigInt(communityId!), `0x${electionId}`],
    query: {
      retry: (failureCount, error) => {
        const retry = (health[chainAlias] as boolean) && failureCount < 2
        if (retry) {
          console.warn('Retrying contract call', chainAlias, communityId, electionId, failureCount, error)
        }
        return retry
      },
      enabled: !!chainAlias && !!communityId && !!electionId,
    },
  })
}
