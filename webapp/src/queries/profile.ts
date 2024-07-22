import { useQuery } from '@tanstack/react-query'
import { isAddress } from 'viem'
import { degen, mainnet } from 'viem/chains'
import abi from '~abis/nftdegen.json'
import { useBlockchain } from '~components/Blockchains/BlockchainContext'
import { useHealthcheck } from '~components/Healthcheck/use-healthcheck'
import { appUrl, degenNameResolverContractAddress } from '~constants'

export const fetchUserProfile = (bfetch: FetchFunction, username: string) => async (): Promise<UserProfileResponse> => {
  const response = await bfetch(`${appUrl}/profile/user/${username}`)
  const user = (await response.json()) as UserProfileResponse

  return user
}

export const fetchUserPolls = (bfetch: FetchFunction, username: string) => async (): Promise<Poll[]> => {
  const response = await bfetch(`${appUrl}/profile/user/${username}`)
  const { polls } = (await response.json()) as { polls: Poll[] }
  if (!polls) {
    throw new Error('Received no elections')
  }
  return polls.map((poll) => ({
    ...poll,
    createdByUsername: username,
  }))
}

export const fetchMutedUsers = (bfetch: FetchFunction) => async (): Promise<Profile[]> => {
  const response = await bfetch(`${appUrl}/profile`)
  const data = await response.json()
  return data.mutedUsers
}

export const fetchWarpcastAPIEnabled = (bfetch: FetchFunction) => async (): Promise<boolean> => {
  const response = await bfetch(`${appUrl}/profile`)
  const data = await response.json()
  return data.warpcastApiEnabled
}

export const useFirstDegenOrEnsName = (addresses: string[] = []) => {
  const { degen: connected } = useHealthcheck()
  const { getContract } = useBlockchain(degen)
  const { client } = useBlockchain(mainnet)

  const contract = getContract(degenNameResolverContractAddress, abi)
  // Process the addresses to ensure a consistent react query function key
  const sortedAddresses = addresses.map((v) => v.toLowerCase()).sort((a, b) => a.localeCompare(b))

  return useQuery({
    queryKey: ['firstDegenOrEnsName', ...sortedAddresses],
    retry: connected,
    queryFn: async () => {
      const degenOrEnsNames = await Promise.all(
        sortedAddresses.map(async (addr) => {
          if (!isAddress(addr)) {
            return null
          }

          const degenName = await contract.read.defaultNames([addr])
          if (degenName) {
            return `${degenName}.degen`
          }

          return client.getEnsName({ address: addr })
        })
      )
      const firstValidName = degenOrEnsNames.find((v) => !!v)
      return firstValidName || null
    },
  })
}

export const getProfileAddresses = (p?: UserProfileResponse) => {
  return p?.user.addresses
}

export const useUserDegenOrEnsName = (user?: UserProfileResponse) => {
  return useFirstDegenOrEnsName(getProfileAddresses(user))
}
