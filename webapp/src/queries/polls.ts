import { useQuery } from '@tanstack/react-query'
import { ethers } from 'ethers'
import { useAuth } from '~components/Auth/useAuth'
import { appUrl, degenChainRpc, degenContractAddress } from '~constants'
import { CommunityHub__factory } from '~typechain'
import { toArrayBuffer } from '~util/hex'

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

export const fetchShortURL = (bfetch: FetchFunction) => async (url: string) => {
  const response = await bfetch(`${appUrl}/short?url=${url}`)
  const { result } = (await response.json()) as { result: string }
  return result
}

export const useApiPollInfo = (electionId?: string) => {
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

export const useContractPollInfo = (communityId?: string, electionId?: string) => {
  const provider = new ethers.JsonRpcProvider(degenChainRpc)
  const contract = CommunityHub__factory.connect(degenContractAddress, provider)
  return useQuery({
    queryKey: ['contractPollInfo', communityId, electionId],
    queryFn: async () => {
      const contractData = await contract.getResult(communityId!, toArrayBuffer(electionId!))
      return contractData
    },
    enabled: !!communityId && !!electionId,
  })
}
