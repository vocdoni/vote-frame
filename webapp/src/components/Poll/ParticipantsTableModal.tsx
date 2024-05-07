import { Button, useDisclosure } from '@chakra-ui/react'
import { useQuery } from '@tanstack/react-query'
import { FaUserGroup } from 'react-icons/fa6'
import { useAuth } from '~components/Auth/useAuth'
import { fetchCensus } from '~queries/census'
import { UsersTableModal } from './UsersTableModal'

export const ParticipantsTableModal = ({ id }: { id?: string }) => {
  const { isOpen, onOpen, onClose } = useDisclosure()
  const { bfetch } = useAuth()
  const { data, error, isLoading } = useQuery({
    queryKey: ['census', id],
    queryFn: fetchCensus(bfetch, id!),
    enabled: !!id && isOpen,
  })

  if (!id) return

  return (
    <>
      <Button size='sm' onClick={onOpen} isLoading={isLoading} rightIcon={<FaUserGroup />}>
        Census
      </Button>
      <UsersTableModal
        isOpen={isOpen}
        onClose={onClose}
        downloadText='Download full census'
        error={error}
        isLoading={isLoading}
        title='Participants / census'
        filename='participants.csv'
        data={
          data?.participants &&
          Object.keys(data.participants).map((username) => [username, data.participants[username]])
        }
      />
    </>
  )
}