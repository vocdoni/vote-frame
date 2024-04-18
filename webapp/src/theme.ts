import { defineStyleConfig, extendTheme } from '@chakra-ui/react'

export const theme = extendTheme({
  fonts: {
    heading: '"Inter", sans-serif',
    body: '"Inter", sans-serif',
  },
  styles: {
    global: {
      body: {
        bg: 'purple.50',
      },
    },
  },
  components: {
    Button: defineStyleConfig({
      defaultProps: {
        colorScheme: 'purple',
      },
    }),
    Link: defineStyleConfig({
      baseStyle: {
        _hover: {
          color: 'purple.500',
          textDecoration: 'none',
        },
      },
      variants: {
        primary: {
          color: 'purple.500',
        },
      },
    }),
    Heading: defineStyleConfig({
      baseStyle: {
        fontWeight: 500,
      },
    }),
    FormLabel: defineStyleConfig({
      baseStyle: {
        fontWeight: 500,
      },
    }),
  },
})
