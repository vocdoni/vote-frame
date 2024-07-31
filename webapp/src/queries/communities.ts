import { useQuery } from '@tanstack/react-query'
import { useAuth } from '~components/Auth/useAuth'
import { appUrl, paginationItemsPerPage } from '~constants'

export const fetchCommunities =
  (bfetch: FetchFunction, { limit = paginationItemsPerPage, offset = 0 }) =>
  async () => {
    const response = await bfetch(`${appUrl}/communities?limit=${limit}&offset=${offset}`)
    const { communities, pagination } = (await response.json()) as { communities: Community[]; pagination: Pagination }

    return { communities, pagination }
  }

export const fetchFeatured =
  (bfetch: FetchFunction, { limit = paginationItemsPerPage, offset = 0 }) =>
  async () => {
    const response = await bfetch(`${appUrl}/communities?featured=true&limit=${limit}&offset=${offset}`)
    const { communities, pagination } = (await response.json()) as { communities: Community[]; pagination: Pagination }

    return { communities, pagination }
  }

export const fetchCommunitiesByAdmin =
  (bfetch: FetchFunction, profile: Profile, { limit = paginationItemsPerPage, offset = 0 }) =>
  async () => {
    const response = await bfetch(`${appUrl}/communities?byAdminFID=${profile.fid}&limit=${limit}&offset=${offset}`)
    const { communities, pagination } = (await response.json()) as { communities: Community[]; pagination: Pagination }

    return { communities, pagination }
  }

export const fetchCommunity = (bfetch: FetchFunction, id: string) => async () => {
  const response = await bfetch(`${appUrl}/communities/${id}`)
  const community = (await response.json()) as Community

  return community
}

export const useCommunity = (id?: string) => {
  const { bfetch } = useAuth()

  return useQuery({
    queryFn: fetchCommunity(bfetch, id!),
    queryKey: ['community', id],
    enabled: !!id,
  })
}

export const updateCommunity = async (bfetch: FetchFunction, community: Community) => {
  const response = await bfetch(`${appUrl}/communities/${community.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(community),
  })
  return response.json()
}

const fetchDelegations = (bfetch: FetchFunction, community: Community) => async () => {
  const response = await bfetch(`${appUrl}/communities/${community.id}/delegations`)
  try {
    const delegates = (await response.json()) as Delegation[]
    return delegates
  } catch (e) {
    return null
  }
}

export const useDelegations = (community: Community) => {
  const { bfetch, isAuthenticated } = useAuth()

  return useQuery({
    queryFn: fetchDelegations(bfetch, community),
    queryKey: ['delegations', community.id],
    enabled: isAuthenticated && !!community.id,
    retry: false,
  })
}
