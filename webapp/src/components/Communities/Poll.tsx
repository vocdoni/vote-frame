import {
  Alert,
  AlertDescription,
  AlertTitle,
  Box,
  Button,
  Flex,
  Heading,
  Image,
  Link,
  Skeleton,
  Progress,
  VStack,
  Tag,
  TagLeftIcon,
  TagLabel,
  Text,
} from '@chakra-ui/react'
import { useEffect, useMemo, useState } from 'react'
import { FaDownload, FaArrowUp, FaRegCircleStop, FaPlay } from 'react-icons/fa6'

import { useAuth } from '../Auth/useAuth'
import { fetchShortURL } from '../../queries/common'
import type { PollResult } from '../../util/types'
import { humanDate } from '../../util/strings'
import { CsvGenerator } from '../../generator'
import { appUrl, electionResultsContract } from '../../util/constants'


export type CommunitiyPollViewProps = {
  electionId: string | undefined,
  communityId: string | undefined,
  poll: PollResult | null,
  loading: boolean | false,
  loaded: boolean | false,
}

export const CommunityPollView = ({poll, electionId, loading, loaded}: CommunitiyPollViewProps) => {
  const { bfetch } = useAuth()
  const [voters, setVoters] = useState([])
  const [electionURL, setElectionURL] = useState<string>(`${appUrl}/${electionId}`)

  useEffect(() => {
    if (loaded || loading || !electionId || !poll) return
      ; (async () => {
        // get the short url
        try {
          const url = await fetchShortURL(bfetch)(electionURL)
          setElectionURL(url)
        } catch (e) {
          console.log("error getting short url, using default", e)
        }
        // get the voters if there are any
        if (poll.voteCount > 0) {
          try {
            const response = await fetch(`${import.meta.env.APP_URL}/votersOf/${electionId}`)
            const data = await response.json()
            setVoters(data.voters)
          } catch (e) {
            console.error("error geting election voters", e)
          }
        }
      })()
  }, [])

  const usersfile = useMemo(() => {
    if (!voters.length) return { url: '', filename: '' }
    return new CsvGenerator(
      ['Username'],
      voters.map((username) => [username])
    )
  }, [voters])

  const copyToClipboard = (input: string) => {
    if (navigator && navigator.clipboard && navigator.clipboard.writeText) {
      navigator.clipboard.writeText(input).catch(console.error);
    } else console.error('clipboard API not available');
  };

  const participationPercentage = useMemo(() => {
    if (!poll) return 0
    return (poll.voteCount / poll.censusParticipantsCount * 100).toFixed(1)
  }, [poll])

  return (
    <Box
      gap={4}
      display='flex'
      flexDir={['column', 'column', 'row']}
      alignItems='start'>
      <Box flex={1} bg='white' p={6} pb={12} boxShadow='md' borderRadius='md'>
        <VStack spacing={8} alignItems='left'>
          <VStack spacing={4} alignItems='left'>
            <Skeleton isLoaded={!loading}>
              <Flex gap={4}>
                {poll?.finalized ?
                  <Tag>
                    <TagLeftIcon as={FaRegCircleStop}></TagLeftIcon>
                    <TagLabel>Ended</TagLabel>
                  </Tag> :
                  <Tag colorScheme='green'>
                    <TagLeftIcon as={FaPlay}></TagLeftIcon>
                    <TagLabel>Ongoing</TagLabel>
                  </Tag>
                }
                {poll?.finalized && <Tag colorScheme='cyan'>
                  <TagLeftIcon as={FaArrowUp}></TagLeftIcon>
                  <TagLabel>Live</TagLabel>
                </Tag>}
              </Flex>
            </Skeleton>
            <Image src={`${import.meta.env.APP_URL}/preview/${electionId}`} fallback={<Skeleton height={200} />} />
            <Link fontSize={'sm'} color={'gray'} onClick={() => copyToClipboard(electionURL)}>Copy link to the frame</Link>
          </VStack>
          <VStack spacing={4} alignItems='left'>
            <Heading size='md'>Results</Heading>
            <Skeleton isLoaded={!loading}>
              <VStack px={4} alignItems='left'>
                <Heading size='sm' fontWeight={'semibold'}>{poll?.question}</Heading>
                {poll?.finalized && <Alert status='success' variant='left-accent' rounded={4}>
                  <Box>
                    <AlertTitle fontSize={'sm'}>Results verifiable on Degenchain</AlertTitle>
                    <AlertDescription fontSize={'sm'}>
                      This poll has ended. The results are definitive and have been settled on the 🎩 Degenchain.
                    </AlertDescription>
                  </Box>
                </Alert>}
                <Link fontSize={'xs'} color='gray' textAlign={'right'} isExternal href={`https://explorer.degen.tips/address/${electionResultsContract}`}>View contract</Link>
                <VStack spacing={6} alignItems='left'>
                  {poll?.options.map((option, index) => (
                    <Box key={index} w='full'>
                      <Flex justifyContent='space-between' w='full'>
                        <Text>{option}</Text>
                        <Text>{poll?.tally[0][index]} votes</Text>
                      </Flex>
                      <Progress size='sm' rounded={50} value={poll?.tally[0][index] / poll?.voteCount * 100} />
                    </Box>
                  ))}
                </VStack>
              </VStack>
            </Skeleton>
          </VStack>
        </VStack>
      </Box>
      <Flex flex={1} direction={'column'} gap={4}>
        <Box bg='white' p={6} boxShadow='md' borderRadius='md'>
          <Heading size='sm'>Information</Heading>
          <Skeleton isLoaded={!loading}>
            <VStack spacing={6} alignItems='left' fontSize={'sm'}>
              <Text>
                This poll {poll?.finalized ? 'has ended' : 'ends'} on {`${humanDate(poll?.endTime)}`}. Check the Vocdoni blockchain explorer for <Link textDecoration={'underline'} isExternal href={`https://stg.explorer.vote/processes/show/#/${electionId}`}>more information</Link>.
              </Text>
              {voters.length > 0 && <>
                <Text>You can download the list of users who casted their votes.</Text>
                <Link href={usersfile.url} download={'voters-list.csv'}>
                  <Button colorScheme='blue' size='sm' rightIcon={<FaDownload />}>Download voters</Button>
                </Link>
              </>}
            </VStack>
          </Skeleton>
        </Box>
        <Flex gap={6}>
          <Box flex={1} bg='white' p={6} boxShadow='md' borderRadius='md'>
            <Skeleton isLoaded={!loading}>
              <Heading pb={4} size='sm'>Participation</Heading>
              <Flex alignItems={'end'} gap={2}>
                <Text fontSize={'xx-large'} lineHeight={1} fontWeight={'semibold'}>{poll?.voteCount}</Text>
                <Text>/{poll?.censusParticipantsCount}</Text>
                <Text fontSize={'xl'}>{participationPercentage}%</Text>
              </Flex>
            </Skeleton>
          </Box>
          <Box flex={1} bg='white' p={6} boxShadow='md' borderRadius='md'>
            <Skeleton isLoaded={!loading}>
              <Heading pb={4} size='sm'>Turnout</Heading>
              <Flex alignItems={'end'} gap={2}>
                <Text fontSize={'xx-large'} lineHeight={1} fontWeight={'semibold'}>{poll?.turnout}</Text>
                <Text>%</Text>
              </Flex>
            </Skeleton>
          </Box>
        </Flex>
      </Flex>
    </Box>
  )
}
