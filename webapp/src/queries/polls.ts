import { appUrl } from '~constants'

export const fetchPollInfo = (bfetch: FetchFunction, electionID: string) => async (): Promise<PollInfo> => {
  const response = await bfetch(`${appUrl}/poll/info/${electionID}`)
  const { poll } = (await response.json()) as { poll: PollInfo }
  return poll
}

export const fetchPollsVoters = (bfetch: FetchFunction, electionId: string) => async (): Promise<string[]> => {
  const response = await bfetch(`${appUrl}/votersOf/${electionId}`)
  const { voters } = (await response.json()) as { voters: string[] }
  return voters
}
