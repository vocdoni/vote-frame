import { Alert, AlertDescription, Box, Heading, Image, Flex, Link, Skeleton, Progress, VStack, Text, AlertTitle, Button, HStack } from '@chakra-ui/react'
import { useEffect, useMemo, useState } from 'react'
import { useParams } from 'react-router-dom'
import { FaDownload } from 'react-icons/fa6'
import { ethers } from 'ethers'

import { useAuth } from '../components/Auth/useAuth'
import { toArrayBuffer } from '../util/hex'
import { PollResult } from '../util/types'
import { CsvGenerator } from '../generator'
import { CommunityHub__factory } from '../typechain'
import { appUrl, degenChainRpc, degenContractAddress, electionResultsContract } from '../util/constants'
import { fetchPollInfo } from '../queries/polls'

const mockedResults: PollResult = {
  censusRoot: 'a989f2e94f9f7954c96ba2cef784525c5ce5c3cba90f0b3da14349a93f3e7dde',
  censusURI: 'https://census.com',
  createdTime: new Date("2024-04-20T14:28:51.228+00:00"),
  options: ['Option 1', 'Option 2'],
  participants: [237855, 308972, 10080],
  question: 'Whats your favorite love movie?',
  tally: [[1, 2], [], [], []],
  voteCount: 3,
  turnout: 100,
}


const Poll = () => {
  const { bfetch } = useAuth()
  const { pid: electionID, id: communityID } = useParams()
  const [voters, setVoters] = useState([])
  const [loaded, setLoaded] = useState<boolean>(false)
  const [loading, setLoading] = useState<boolean>(false)
  const [results, setResults] = useState<PollResult | null>(null)

  useEffect(() => {
    if (loaded || loading || !electionID || !communityID) return
      ; (async () => {
        try {
          setLoading(true)
          // get results from the contract
          const provider = new ethers.JsonRpcProvider(degenChainRpc)
          const communityHubContract = CommunityHub__factory.connect(degenContractAddress, provider)
          const contractData = await communityHubContract.getResult(communityID, toArrayBuffer(electionID))
          let results: PollResult
          if (contractData.date !== "") {
            const participants = contractData.participants.map((p) => parseInt(p.toString()))
            const tally = contractData.tally.map((t) => t.map((v) => parseInt(v.toString())))
            results = {
              censusRoot: contractData.censusRoot,
              censusURI: contractData.censusURI,
              createdTime: new Date(contractData.date),
              options: contractData.options,
              participants: participants,
              question: contractData.question,
              tally: tally,
              turnout: parseFloat(contractData.turnout.toString()),
              voteCount: parseInt(contractData.totalVotingPower.toString()),
            }
            console.log("results from contract")
          } else {
            try {
              const apiData = await fetchPollInfo(bfetch)(electionID)
              const tally: number[][] = [[]]
              apiData.tally?.forEach((t) => {
                tally[0].push(parseInt(t))
              })
              results = {
                censusRoot: "",
                censusURI: "",
                createdTime: new Date(apiData.createdTime),
                options: apiData.options,
                participants: apiData.participants,
                question: apiData.question,
                tally: tally,
                turnout: apiData.turnout,
                voteCount: apiData.voteCount,
              }
              console.log("results from api")
            } catch (e) {
              console.error(e)
              results = mockedResults
              console.log("mocked results")
            }
            setResults(results)
          }
          // get the voters
          const response = await fetch(`${import.meta.env.APP_URL}/votersOf/${electionID}`)
          const data = await response.json()
          setVoters(data.voters)
        } catch (e) {
          console.error(e)
        } finally {
          setLoaded(true)
          setLoading(false)
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
    }else console.error('clipboard API not available');
  };

  return (
    <Box
      gap={4}
      display='flex'
      flexDir={['column', 'column', 'row']}
      alignItems='start'>
      <Box flex={1} bg='white' p={6} pb={12} boxShadow='md' borderRadius='md'>
        <VStack spacing={8} alignItems='left'>
          <VStack spacing={4} alignItems='left'>
            <Image src={`${import.meta.env.APP_URL}/preview/${electionID}`} fallback={<Skeleton height={200} />} />
            <Link fontSize={'sm'} color={'gray'} onClick={() => copyToClipboard(`${appUrl}/${electionID}`)}>Copy link to the frame</Link>
          </VStack>
          <VStack spacing={4} alignItems='left'>
            <Heading size='md'>Results</Heading>
            <Skeleton isLoaded={!loading}>
              <VStack px={4} alignItems='left'>
                <Text fontSize={'xs'} color={'gray'}>Results</Text>
                <Heading size='sm' fontWeight={'semibold'}>{results?.question}</Heading>
                <Alert status='success' variant='left-accent' rounded={4}>
                  <Box>
                    <AlertTitle fontSize={'sm'}>Results verifiable on Degenchain</AlertTitle>
                    <AlertDescription fontSize={'sm'}>
                      This poll has ended. The results are definitive and have been settled on the 🎩 Degenchain.
                    </AlertDescription>
                  </Box>
                </Alert>
                <Link fontSize={'xs'} color='gray' textAlign={'right'} isExternal href={`https://explorer.degen.tips/address/${electionResultsContract}`}>View contract</Link>
                <VStack spacing={6} alignItems='left'>
                  {results?.options.map((option, index) => (
                    <Box key={index} w='full'>
                      <Flex justifyContent='space-between' w='full'>
                        <Text>{option}</Text>
                        <Text>{results?.tally[0][index]} votes</Text>
                      </Flex>
                      <Progress size='sm' rounded={50} value={results?.tally[0][index] / results?.voteCount * 100} />
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
                This poll started on {`${results?.createdTime}`}. Check the Vocdoni blockchain explorer for <Link textDecoration={'underline'} isExternal href={`https://stg.explorer.vote/processes/show/#/${electionID}`}>more information</Link>.
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
            <Heading pb={4} size='sm'>Votes turnout</Heading>
            <Flex alignItems={'end'} gap={2}>
              <Text fontSize={'xx-large'} lineHeight={1} fontWeight={'semibold'}>{results?.tally[0].reduce((a,b)=>a+b, 0)}</Text>
              <Text>/{results?.voteCount}</Text>
              <Text fontSize={'xl'}>({results?.turnout}%)</Text>
            </Flex>
          </Box>
          <Box flex={1}></Box>
        </Flex>
      </Flex>
    </Box>
  )
}

export default Poll
