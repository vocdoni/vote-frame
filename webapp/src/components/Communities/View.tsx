import {
  Avatar,
  Box,
  Button,
  Flex,
  Grid,
  GridItem,
  Heading,
  Icon,
  Link,
  Text,
  HStack,
  Table,
  Tag,
  TagLeftIcon,
  TagLabel,
  Thead,
  Tr,
  Td,
  Th,
  Tbody,
  VStack
} from '@chakra-ui/react'
import {PropsWithChildren, ReactElement, Fragment} from 'react'
import {TbExternalLink} from "react-icons/tb"
import {SiFarcaster} from "react-icons/si";
import {BsChatDotsFill} from "react-icons/bs";
import {FaRegCircleStop, FaPlay} from 'react-icons/fa6'
import {useQuery} from '@tanstack/react-query'
import {Link as RouterLink, useNavigate} from 'react-router-dom';

import {appUrl, degenContractAddress} from '../../util/constants'
import {Community} from '../../queries/communities'
import {fetchPollsByCommunity} from '../../queries/tops'
import {useAuth} from '../Auth/useAuth'
import {Poll} from '../../util/types';
import {humanDate} from '../../util/strings'
import {MdHowToVote} from "react-icons/md";

export type CommunitiesViewProps = {
  community: Community
}

const WhiteBox = ({children}: PropsWithChildren) => (
  <Flex alignItems='start' gap={4} padding={6} bg='white' boxShadow='sm' borderRadius='md' flexWrap='wrap' h='100%'
        maxW={'100vw'} overflowX={'auto'}>
    {children}
  </Flex>
)

export const CommunitiesView = ({community}: CommunitiesViewProps) => {
  const {bfetch, profile, isAuthenticated} = useAuth()
  const {data: communityPolls, refetch} = useQuery<Poll[], Error>({
    queryKey: ['communityPolls', community?.id],
    queryFn: fetchPollsByCommunity(bfetch, community as Community),
    enabled: !!community,
  })
  const navigate = useNavigate() // Hook to control navigation

  if (!community) return;


  const imAdmin = isAuthenticated && community.admins.some(admin => admin.fid == profile?.fid);

  const disableCommunity = async () => {
    try {
      await bfetch(`${appUrl}/communities/${community.id}`, {method: 'DELETE'}).then(() => refetch())
    } catch (e) {
      console.error('could not unmute user', e)
    } finally {
      navigate('/communities')
    }
  }

  const channelLinks: ReactElement[] = [];
  community.channels.forEach((channel, index) => {
    channelLinks.push(
      <Link key={`link-${channel}`} fontSize="sm" color="gray" isExternal _hover={{textDecoration: 'underline'}}
            href={`https://warpcast.com/~/channel/${channel}`}>
        /{channel}
      </Link>
    );
    // Add the separator if it's not the last item
    if (index !== community.channels.length - 1) {
      channelLinks.push(<Text as="span" fontSize="sm" mx={1} color={'grey'} key={`separator-${index}`}>&amp;</Text>);
    }
  });

  return (
    <Grid
      w='full'
      gap={4}
      gridTemplateAreas={{base: '"profile" "links" "polls"', md: '"profile links" "polls polls"'}}
      gridTemplateColumns={{base: 'full', md: '50%'}}
    >
      <GridItem gridArea='profile'>
        <WhiteBox>
          <Avatar src={community.logoURL}/>
          <Box>
            <Heading size='md'>{community.name}</Heading>
            <Text fontSize='smaller' fontStyle='italic'>
              Managed by <CommunityAdmins community={community}/>
            </Text>
            <Text fontSize='smaller' mt='6'>
              Deployed on <Link isExternal href={`https://explorer.degen.tips/address/${degenContractAddress}`}><Text
              as={'u'}>🎩 DegenChain</Text></Link>
            </Text>
            {!!imAdmin && <Flex mt={4} gap={4}>
              <Button onClick={disableCommunity} colorScheme={'red'}>
                <Text>Disable community</Text>
              </Button>
              <Button onClick={() => navigate('/')} leftIcon={<MdHowToVote/>}>Create vote</Button></Flex>
            }
          </Box>
        </WhiteBox>
      </GridItem>
      <GridItem gridArea='links'>
        <WhiteBox>
          <Box>
            <Heading size={'sm'} mb={2}>Community Engagement</Heading>
            <HStack spacing={2} align='center'>
              <Icon as={SiFarcaster} size={8}/>
              <Text fontWeight={'semibold'} fontSize={'sm'}>Farcaster channels</Text>
            </HStack>
            <Box ml={6} mb={2}>
              {channelLinks}
            </Box>
            <Link isExternal href={community.groupChat}>
              <HStack spacing={2} align='center'>
                <Icon as={BsChatDotsFill}/>
                <Heading size='xs'><Text as='u'>Group chat</Text></Heading>
                <Icon as={TbExternalLink} size={4}/>
              </HStack>
            </Link>
          </Box>
        </WhiteBox>
      </GridItem>
      {!!communityPolls && <GridItem gridArea='polls'>
        <WhiteBox>
          <Heading size={'md'} mb={4}>Community Polls</Heading>
          <Table style={{overflowX: 'auto'}} maxW="100%">
            <Thead>
              <Tr>
                <Th>Question</Th>
                <Th isNumeric>Votes</Th>
                <Th isNumeric>Census size</Th>
                <Th isNumeric>Participation(%)</Th>
                <Th>Last vote</Th>
                <Th>Status</Th>
              </Tr>
            </Thead>
            <Tbody>
              {communityPolls?.map((poll, index) => (
                <Tr key={index}>
                  <Td>
                    <RouterLink to={`poll/${poll.electionId}`}>{poll.question}</RouterLink>
                    <Text as={'p'} fontSize={'xs'} color='gray'>by {poll.createdByDisplayname}</Text>
                  </Td>
                  <Td isNumeric>{poll.voteCount}</Td>
                  <Td isNumeric>{poll.censusParticipantsCount}</Td>
                  <Td isNumeric>{`${(poll.voteCount / poll.censusParticipantsCount * 100).toFixed(1)}%`}</Td>
                  <Td>{poll.voteCount > 0 ? humanDate(poll.lastVoteTime) : '-'}</Td>
                  <Td>
                    <VStack>
                      {poll.finalized ?
                        <Tag>
                          <TagLeftIcon as={FaRegCircleStop}></TagLeftIcon>
                          <TagLabel>Ended</TagLabel>
                        </Tag> :
                        <Tag colorScheme='green'>
                          <TagLeftIcon as={FaPlay}></TagLeftIcon>
                          <TagLabel>Ongoing</TagLabel>
                        </Tag>}
                      {poll.finalized && <Text fontSize={'xs'} color={'gray'}>{humanDate(poll.endTime)}</Text>}
                    </VStack>
                  </Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
        </WhiteBox>
      </GridItem>}
    </Grid>
  )
}

export const CommunityAdmins = ({community}: CommunitiesViewProps) => {
  return community.admins.map((admin, k) => (
    <Fragment key={k}>
      <Link isExternal href={`https://warpcast.com/${admin.username}`}>
        {admin.displayName || admin.username}
      </Link>
      {k === community.admins.length - 2 ? ' & ' : k < community.admins.length - 2 ? ', ' : ''}
    </Fragment>
  ))
}
