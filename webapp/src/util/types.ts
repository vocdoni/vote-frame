import {CensusType} from "../components/CensusTypeSelector.tsx";

export type FetchFunction = (input: RequestInfo, init?: RequestInit) => Promise<Response>

export type Address = {
  address: string
  blockchain: string
}

export type Profile = {
  fid: number
  username: string
  displayName: string
  bio: string
  pfpUrl: string
  custody: string
  verifications: string[]
}

export type Poll = {
  electionId: string
  title: string
  createdByUsername: string
  createdByDisplayname: string
  voteCount: number
  createdTime: Date
  lastVoteTime: Date
}

export interface HTTPErrorResponse {
  response?: {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data?: any
  }
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const isErrorWithHTTPResponse = (error: any): error is HTTPErrorResponse =>
  error && typeof error === 'object' && 'response' in error && 'data' in error.response

// This enum comes from the contract repo
enum ContractCensusType {
  FC,         /// All Farcaster users.
  CHANNEL,    /// Users in a specific channel.
  FOLLOWERS,  /// Users following a specific account.
  CSV,        /// Users in a CSV file.
  ERC20,      /// Users holding a specific ERC20 token.
  NFT         /// Users holding a specific NFT.
}

export const censusTypeToEnum = (census: CensusType) => {
  switch (census) {
    case 'channel':
      return ContractCensusType.CHANNEL
    case 'followers':
      return ContractCensusType.FOLLOWERS
    case 'custom':
      return ContractCensusType.CSV
    case 'erc20':
      return ContractCensusType.ERC20
    case 'nft':
      return ContractCensusType.NFT
    case 'farcaster':
    default:
      return ContractCensusType.FC
  }
}
