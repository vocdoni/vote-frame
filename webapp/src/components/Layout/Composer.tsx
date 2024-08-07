import { Box, ChakraProvider, ColorModeScript } from '@chakra-ui/react'
import { useEffect } from 'react'
import { Outlet, useLocation } from 'react-router-dom'
import { useAuth } from '~components/Auth/useAuth'
import { composer } from '~src/themes/composer'
import { MaintenanceAlert } from './MaintenanceAlert'

export const ComposerLayout = () => {
  const { search } = useLocation()
  const { searchParamsTokenLogin } = useAuth()

  // login via token (needs to be handled here because the auth provider is outside of the router context)
  useEffect(() => {
    searchParamsTokenLogin(search)
  }, [search])

  return (
    <ChakraProvider theme={composer}>
      <ColorModeScript />
      <Box p={4}>
        <MaintenanceAlert />
        <Outlet />
      </Box>
    </ChakraProvider>
  )
}
