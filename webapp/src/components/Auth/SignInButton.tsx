import {
  Box,
  Button,
  ButtonProps,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Spinner,
  Text,
  useDisclosure,
} from '@chakra-ui/react'
import { QRCode } from '@farcaster/auth-kit'
import { useEffect, useState } from 'react'
import { IoPhonePortraitOutline } from 'react-icons/io5'
import { appUrl } from '~constants'
import { FarcasterLogo } from '../FarcasterLogo'
import { useAuth } from './useAuth'

import '@farcaster/auth-kit/styles.css'

export const SignInButton = (props: ButtonProps) => {
  const { login, isAuthenticated } = useAuth()
  const { isOpen, onOpen, onClose } = useDisclosure()
  const [url, setUrl] = useState<string | null>(null)
  const [id, setId] = useState<string | null>(null)

  // retrieve the QR url
  useEffect(() => {
    if (!isOpen || url !== null) return
    ;(async () => {
      try {
        const res = await fetch(`${appUrl}/auth`)
        const { id, url } = await res.json()
        setUrl(url)
        setId(id)
      } catch (e) {
        console.error('error fetching auth url:', e)
      }
    })()
  }, [isOpen, url])

  const bearerCheck = async (id: string) => {
    const res = await fetch(`${appUrl}/auth/${id}`)

    if (res.status === 200) {
      const rjson = await res.json()
      login({
        bearer: rjson.authToken,
        profile: rjson.profile,
      })
      return true
    }
    if (res.status !== 204) {
      throw await res.text()
    }

    return false
  }

  // as soon as we have the QR, start polling for the bearer
  useEffect(() => {
    if (isAuthenticated || !id) return

    const interval = setInterval(async () => {
      if (await bearerCheck(id)) {
        clearInterval(interval)
        onClose()
      }
    }, 2000)

    // clear interval if the modal gets closed
    if (!isOpen) clearInterval(interval)

    return () => clearInterval(interval)
  }, [url, isAuthenticated, isOpen])

  return (
    <>
      <Button colorScheme='purple' leftIcon={<FarcasterLogo height={20} fill='white' />} {...props} onClick={onOpen}>
        Sign in
      </Button>
      <Modal isOpen={isOpen} onClose={onClose} isCentered size='xs'>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Sign in with Farcaster</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text>Scan with your phone's camera to continue.</Text>
            <Box justifyContent='center' display='flex'>
              {url ? <QRCode uri={url} size={264} logoSize={22} logoMargin={12} /> : <Spinner />}
            </Box>
          </ModalBody>
          <ModalFooter justifyContent='center'>
            {url && (
              <Button
                variant='text'
                fontWeight='normal'
                colorScheme='purple'
                onClick={() => (window.location.href = url)}
                leftIcon={<IoPhonePortraitOutline />}
              >
                I'm using my phone →
              </Button>
            )}
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  )
}
