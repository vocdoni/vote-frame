import { HStack, Link, Text, VStack } from '@chakra-ui/react'
import { vocdoniExplorer } from '~constants'
import { humanDate } from '~util/strings'
import { ParticipantsTableModal } from './ParticipantsTableModal'
import { RemainingVotersTableModal } from './RemainingVotersTableModal'
import { VotersTableModal } from './VotersTableModal'

export const Information = ({ poll }: { poll?: PollInfo }) => {
  if (!poll) return

  return (
    <VStack spacing={6} alignItems='left' fontSize={'sm'}>
      <Text>
        This poll {poll?.finalized ? 'has ended' : 'ends'} on {`${humanDate(poll?.endTime)}`}.{` `}
        <Link variant='primary' isExternal href={`${vocdoniExplorer}/processes/show/#/${poll.electionId}`}>
          Check the Vocdoni blockchain explorer
        </Link>
        {` `}for more information.
      </Text>
      {!!poll.censusParticipantsCount && (
        <>
          <Text>
            Download the list of members who have already cast their votes, the list of remaining members who still need
            to vote, and the total census of eligible voters.
          </Text>
          <HStack spacing={2} flexWrap='wrap'>
            <VotersTableModal id={poll.electionId} />
            <RemainingVotersTableModal id={poll.electionId} />
            <ParticipantsTableModal id={poll.electionId} />
          </HStack>
        </>
      )}
    </VStack>
  )
}