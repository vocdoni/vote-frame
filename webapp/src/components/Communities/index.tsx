import { Box, Button, Flex, Heading, HStack, Link, SimpleGrid, Text, VStack } from '@chakra-ui/react'
import { useQuery } from '@tanstack/react-query'
import { useCallback, useState } from 'react'
import { FaRegStar, FaUsers } from 'react-icons/fa'
import { Link as RouterLink } from 'react-router-dom'
import { useAuth } from '~components/Auth/useAuth'
import { Check } from '~components/Check'
import { fetchCommunities, fetchCommunitiesByAdmin } from '~queries/communities'
import { CommunityCard } from './Card'

export const CommunitiesList = () => {
  const { bfetch, profile, isAuthenticated } = useAuth()

  // state to show only the communities the user is part of
  const [showMyCommunities, setShowMyCommunities] = useState(false)

  // callback to toggle showMyCommunities
  const toggleMyCommunities = useCallback(() => {
    setShowMyCommunities(!showMyCommunities)
  }, [showMyCommunities])

  const {
    data: allCommunities,
    error: allCommunitiesError,
    isLoading: isAllCommunitiesLoading,
  } = useQuery({
    queryKey: ['communities'],
    queryFn: fetchCommunities(bfetch),
  })
  const {
    data: myCommunities,
    error: myCommunitiesError,
    isLoading: isMyCommunitiesLoading,
  } = useQuery({
    queryKey: ['communities', 'byAdmin'],
    queryFn: () => fetchCommunitiesByAdmin(bfetch, profile!),
    enabled: profile != null,
  })
  // Filter by community admins fid in case showMyCommunities is true
  const filteredData = showMyCommunities ? myCommunities : allCommunities

  return (
    <VStack spacing={4} w='full' alignItems='start'>
      <Flex my={4} w='full' justifyContent='space-between' alignItems='center' wrap={'wrap'}>
        <Heading size='md' m={4}>
          Communities
        </Heading>

        {isAuthenticated && (
          <ToggleStateComponent
            state={showMyCommunities}
            toggleState={toggleMyCommunities}
            state1text={'All communities'}
            state2text={'My communities'}
          />
        )}
      </Flex>

      <SimpleGrid gap={4} w='full' alignItems='start' columns={{ base: 1, md: 2, lg: 3, xl: 4 }}>
        {filteredData &&
          filteredData.map((community: Community, k: number) => (
            <CommunityCard
              name={community.name}
              slug={community.id.toString()}
              key={k}
              pfpUrl={community.logoURL}
              admins={community.admins}
              disabled={community.disabled}
            />
          ))}
      </SimpleGrid>
      <Check
        error={allCommunitiesError || myCommunitiesError}
        isLoading={isAllCommunitiesLoading || isMyCommunitiesLoading}
      />
      <Box
        w='full'
        boxShadow='sm'
        borderRadius='lg'
        minHeight={300}
        display='flex'
        flexDir='column'
        alignItems='center'
        justifyContent='center'
        bg='white'
        p={10}
        textAlign='center'
        gap={4}
      >
        <Text fontSize='larger' fontWeight='500'>
          Create your own community and start managing its governance
        </Text>
        <Link as={RouterLink} to='/communities/new'>
          <Button>Create a community on 🎩 Degenchain</Button>
        </Link>
      </Box>
    </VStack>
  )
}

interface IToggleStateComponentProps {
  state: boolean
  toggleState: () => void
  state1text: string
  state2text: string
}

export const ToggleStateComponent = ({ state, toggleState, state1text, state2text }: IToggleStateComponentProps) => {
  return (
    <HStack m={4} align='center' gap={4}>
      <Button
        size={'sm'}
        onClick={() => {
          if (!state) toggleState()
        }}
        leftIcon={<FaRegStar />}
        variant={state ? 'solid' : 'ghost'}
      >
        {state2text}
      </Button>
      <Button
        size={'sm'}
        onClick={() => {
          if (state) toggleState()
        }}
        leftIcon={<FaUsers />}
        variant={!state ? 'solid' : 'ghost'}
      >
        {state1text}
      </Button>
    </HStack>
  )
}
