import { appUrl } from '~constants'

export const fetchPollsByVotes = (bfetch: FetchFunction) => async (): Promise<PollRanking[]> => {
  const response = await bfetch(`${appUrl}/rankings/pollsByVotes`)
  const { polls } = (await response.json()) as { polls: PollRanking[] }
  return polls
}

export const fetchTopVoters = (bfetch: FetchFunction) => async (): Promise<UserRanking[]> => {
  const response = await bfetch(`${appUrl}/rankings/usersByCastedVotes`)
  const { users } = (await response.json()) as { users: UserRanking[] }
  return users
}

export const fetchTopCreators = (bfetch: FetchFunction) => async (): Promise<UserRanking[]> => {
  const response = await bfetch(`${appUrl}/rankings/usersByCreatedPolls`)
  const { users } = (await response.json()) as { users: UserRanking[] }
  return users
}

export const latestPolls =
  (bfetch: FetchFunction, { limit = 10 }: Partial<Pagination> = {}) =>
  async (): Promise<PollRanking[]> => {
    const response = await bfetch(`${appUrl}/rankings/latestElections?limit=${limit}`)
    const { polls } = (await response.json()) as { polls: PollRanking[] }
    return polls
  }

export const fetchPollsByCommunity = (bfetch: FetchFunction, community: Community) => async (): Promise<Poll[]> => {
  const response = await bfetch(`${appUrl}/rankings/pollsByCommunity/${community.id}`)
  const { polls } = (await response.json()) as { polls: Poll[] }
  return polls
}
