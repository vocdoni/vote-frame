import {
  Box,
  CircularProgress,
  CircularProgressLabel,
  Flex,
  Icon,
  SimpleGrid,
  Stat,
  StatLabel,
  StatNumber,
  useColorModeValue,
} from '@chakra-ui/react'
import { FaHeart, FaRegFaceGrinStars } from 'react-icons/fa6'
import { ImStatsDots } from 'react-icons/im'
import { MdOutlineHowToVote } from 'react-icons/md'
import { SlPencil } from 'react-icons/sl'
import type { Reputation } from './useAuthProvider'

export const ReputationCard = ({ reputation }: { reputation: Reputation }) => {
  const bg = useColorModeValue('whiteAlpha.500', 'whiteAlpha.200')
  const boxShadow = useColorModeValue('0px 4px 6px rgba(0, 0, 0, 0.1)', '0px 4px 6px rgba(0, 0, 0, 0.3)')

  if (!reputation) {
    return null
  }

  return (
    <Box
      p={4}
      bg={bg}
      boxShadow={boxShadow}
      borderRadius='lg'
      bgGradient='linear(to-r, purple.700, purple.400)'
      color='white'
    >
      <Flex justifyContent='space-around' alignItems='center'>
        <Stat>
          <StatLabel fontSize='lg'>
            <Icon as={ImStatsDots} boxSize={4} /> Farcaster.vote reputation
          </StatLabel>
          <StatNumber fontSize='2xl'>{reputation.reputation}</StatNumber>
        </Stat>
        <CircularProgress value={reputation.reputation} max={100} color='purple.600' thickness='12px' mr={3}>
          <CircularProgressLabel>{reputation.reputation}%</CircularProgressLabel>
        </CircularProgress>
      </Flex>
      <SimpleGrid columns={2} spacing={3} mt={4}>
        <Stat>
          <StatLabel fontSize='x-small'>Casted votes</StatLabel>
          <StatNumber fontSize='sm'>
            {reputation.data.castedVotes}
            {` `}
            <Icon as={MdOutlineHowToVote} boxSize={3.5} />
          </StatNumber>
        </Stat>
        <Stat>
          <StatLabel fontSize='x-small'>Created polls</StatLabel>
          <StatNumber fontSize='sm'>
            {reputation.data.electionsCreated}
            {` `}
            <Icon as={SlPencil} boxSize={3} />
          </StatNumber>
        </Stat>
        <Stat>
          <StatLabel fontSize='x-small'>Followers</StatLabel>
          <StatNumber fontSize='sm'>
            {reputation.data.followersCount}
            {` `}
            <Icon as={FaHeart} boxSize={3} />
          </StatNumber>
        </Stat>
        <Stat>
          <StatLabel fontSize='x-small'>Participation in created polls</StatLabel>
          <StatNumber fontSize='sm'>
            {reputation.data.participationAchievement}
            {` `}
            <Icon as={FaRegFaceGrinStars} boxSize={3} />
          </StatNumber>
        </Stat>
      </SimpleGrid>
    </Box>
  )
}
