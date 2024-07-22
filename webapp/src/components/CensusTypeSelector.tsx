import {
  Alert,
  AlertDescription,
  Avatar,
  Button,
  Flex,
  FormControl,
  FormControlProps,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  Icon,
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
  Link,
  ListItem,
  Radio,
  RadioGroup,
  Select,
  Spinner,
  Stack,
  Text,
  UnorderedList,
} from '@chakra-ui/react'
import { useQuery } from '@tanstack/react-query'
import { chakraComponents, GroupBase, OptionProps, Select as RSelect } from 'chakra-react-select'
import { useEffect, useState } from 'react'
import { Controller, useFieldArray, useFormContext } from 'react-hook-form'
import { BiTrash } from 'react-icons/bi'
import { MdArrowDropDown } from 'react-icons/md'
import { fetchAirstackBlockchains } from '~queries/census'
import { fetchCommunitiesByAdmin } from '~queries/communities'
import { ucfirst } from '~util/strings'
import Airstack from '../assets/airstack.svg?react'
import { useAuth } from './Auth/useAuth'
import ChannelSelector, { ChannelFormValues } from './Census/ChannelSelector'
import { CreateFarcasterCommunityButton } from './Layout/DegenButton'

export type CensusFormValues = ChannelFormValues & {
  censusType: CensusType
  addresses?: Address[]
  community?: Community
  csv?: File | undefined
}

export type CensusTypeSelectorProps = FormControlProps & {
  complete?: boolean
  communityId?: string
  admin?: boolean
  composer?: boolean
}

const CensusTypeSelector = ({ complete, communityId, admin, composer, ...props }: CensusTypeSelectorProps) => {
  const { bfetch, profile, isAuthenticated } = useAuth()
  const {
    control,
    formState: { errors },
    register,
    setValue,
    watch,
  } = useFormContext<CensusFormValues>()
  const {
    fields: addressFields,
    append: appendAddress,
    remove: removeAddress,
  } = useFieldArray({
    control,
    name: 'addresses',
  })
  const { data: blockchains, isLoading: bloading } = useQuery({
    queryKey: ['blockchains'],
    queryFn: fetchAirstackBlockchains(bfetch),
  })
  const {
    data: communities,
    isLoading: cloading,
    isSuccess,
  } = useQuery({
    queryKey: ['communities', 'byAdmin', profile?.fid],
    queryFn: fetchCommunitiesByAdmin(bfetch, profile!, { offset: 0, limit: 20 }),
    enabled: isAuthenticated && !!complete,
  })
  const [initCommunity, setInitCommunity] = useState<Community | undefined>(undefined)

  const censusType = watch('censusType')
  const addresses = watch('addresses')

  // reset address fields when censusType changes
  useEffect(() => {
    if ((censusType === 'erc20' || censusType === 'nft') && addresses && !addresses.length) {
      // Remove all fields initially
      setValue('addresses', [])
      // Add one field by default
      for (let i = 0; i < 1; i++) {
        appendAddress({ address: '', blockchain: 'base' })
      }
    }
  }, [censusType, addresses])

  // set community id if received (1st step)
  useEffect(() => {
    if (!communityId || cloading || !isSuccess) return

    setInitCommunity(communities?.communities.find((c) => c.id === communityId))
  }, [communityId, cloading, communities, isSuccess])
  // yeah we need to do it in two steps, or use a timeout which would have been a worse solution
  useEffect(() => {
    if (!initCommunity) return

    setValue('censusType', 'community')
    setValue('community', initCommunity)
  }, [initCommunity])

  const required = {
    value: true,
    message: 'This field is required',
  }

  return (
    <>
      <FormControl {...props} isRequired>
        <FormLabel>Census/voters</FormLabel>
        <RadioGroup onChange={(val: CensusType) => setValue('censusType', val)} value={censusType} id='census-type'>
          <Stack direction='column' flexWrap='wrap'>
            {complete && <Radio value='farcaster'>🌐 All farcaster users</Radio>}
            {complete && <Radio value='community'>🏘️ Community based</Radio>}
            <Radio value='channel'>⛩ Farcaster channel gated</Radio>
            <Radio value='followers'>❤️ My Farcaster followers and me</Radio>
            {complete && <Radio value='alfafrens'>💙 My alfafrens channel subscribers</Radio>}
            {complete && !composer && <Radio value='custom'>🦄 Token based via CSV</Radio>}
            <Radio value='nft'>
              <Icon as={Airstack} /> NFT based via airstack
            </Radio>
            <Radio value='erc20'>
              <Icon as={Airstack} /> ERC20 based via airstack
            </Radio>
          </Stack>
        </RadioGroup>
      </FormControl>
      {censusType === 'community' &&
        (communities && communities?.communities.length ? (
          <FormControl isRequired>
            <FormLabel>Select a community</FormLabel>
            <Controller
              name='community'
              control={control}
              render={({ field }) => (
                <RSelect
                  placeholder='Choose a community'
                  isLoading={cloading}
                  options={communities?.communities.filter((c) => !c.disabled) || []}
                  getOptionLabel={(option: Community) => option.name}
                  getOptionValue={(option: Community) => option.id.toString()}
                  components={communitySelector}
                  {...field}
                />
              )}
            />
          </FormControl>
        ) : (
          <Flex alignItems='center' direction='column' w='full'>
            <Text>You don't have a community yet, want to create one?</Text>
            <CreateFarcasterCommunityButton />
          </Flex>
        ))}
      {['erc20', 'nft'].includes(censusType) &&
        addressFields.map((field, index) => (
          <FormControl key={field.id} {...props}>
            <FormLabel>
              {censusType.toUpperCase()} address {index + 1}
            </FormLabel>
            <Flex>
              <Select
                {...register(`addresses.${index}.blockchain`, { required })}
                defaultValue='ethereum'
                w='auto'
                icon={bloading ? <Spinner /> : <MdArrowDropDown />}
              >
                {blockchains &&
                  blockchains.map((blockchain, key) => (
                    <option value={blockchain} key={key}>
                      {ucfirst(blockchain)}
                    </option>
                  ))}
              </Select>
              <InputGroup>
                <Input placeholder='Smart contract address' {...register(`addresses.${index}.address`, { required })} />
                {addressFields.length > 1 && (
                  <InputRightElement>
                    <IconButton
                      aria-label='Remove address'
                      icon={<BiTrash />}
                      onClick={() => removeAddress(index)}
                      size='sm'
                    />
                  </InputRightElement>
                )}
              </InputGroup>
            </Flex>
          </FormControl>
        ))}
      {censusType === 'nft' && addressFields.length < 3 && (
        <Button variant='ghost' onClick={() => appendAddress({ address: '', blockchain: 'ethereum' })}>
          Add address
        </Button>
      )}
      {censusType === 'channel' && (
        <FormControl isRequired isInvalid={!!errors.channel} {...props}>
          <FormLabel htmlFor='channel'>Farcaster channel</FormLabel>
          <Controller name='channel' render={({ field }) => <ChannelSelector admin={admin} {...field} />} />
          <FormErrorMessage>{errors.channel?.message?.toString()}</FormErrorMessage>
        </FormControl>
      )}
      {censusType === 'custom' && (
        <FormControl isRequired {...props}>
          <FormLabel htmlFor='csv'>CSV files</FormLabel>
          <Input
            id='csv'
            placeholder='Upload CSV'
            type='file'
            multiple
            accept='text/csv,application/csv,.csv'
            {...register('csv', {
              required: {
                value: true,
                message: 'This field is required',
              },
            })}
          />
          {errors.csv ? (
            <FormErrorMessage>{errors.csv?.message?.toString()}</FormErrorMessage>
          ) : (
            <FormHelperText>
              <Alert status='info'>
                <AlertDescription>
                  The CSV files <strong>must include Ethereum addresses and their balances</strong> from any network.
                  You can build your own at:
                  <UnorderedList>
                    <ListItem>
                      <Link isExternal href='https://holders.at' variant='primary'>
                        holders.at
                      </Link>{' '}
                      for NFTs
                    </ListItem>
                    <ListItem>
                      <Link isExternal href='https://collectors.poap.xyz' variant='primary'>
                        collectors.poap.xyz
                      </Link>{' '}
                      for POAPs
                    </ListItem>
                  </UnorderedList>
                  <strong>If an address appears multiple times, its balances will be aggregated.</strong>
                </AlertDescription>
              </Alert>
            </FormHelperText>
          )}
        </FormControl>
      )}
    </>
  )
}

export default CensusTypeSelector

const communitySelector = {
  Option: ({ children, ...props }: OptionProps<any, false, GroupBase<any>>) => (
    <chakraComponents.Option {...props}>
      <Avatar size={'sm'} src={(props.data as Community).logoURL} mr={2} /> {children}
    </chakraComponents.Option>
  ),
}
