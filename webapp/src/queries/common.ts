import { appUrl } from '~constants'

export const fetchShortURL = (bfetch: FetchFunction) => async (url: string) => {
  const response = await bfetch(`${appUrl}/short?url=${url}`)
  const { result } = (await response.json()) as { result: string }
  return result
}
