import { lazy } from 'react'
import { createHashRouter, redirect, RouterProvider } from 'react-router-dom'
import { Layout } from '~components/Layout'
import { SuspenseLoader } from './SuspenseLoader'

const About = lazy(() => import('~pages/About'))
const Home = lazy(() => import('~pages/Home'))
const AppForm = lazy(() => import('~pages/Form'))
const CommunitiesLayout = lazy(() => import('~pages/communities/layout'))
const CommunitiesNew = lazy(() => import('~pages/communities/new'))
const AllCommunitiesList = lazy(() => import('~pages/communities'))
const MyCommunitiesList = lazy(() => import('~pages/communities/mine'))
const Community = lazy(() => import('~pages/communities/view'))
const CommunityPoll = lazy(() => import('~pages/communities/poll'))
const FarcasterAccountProtectedRoute = lazy(() => import('./FarcasterAccountProtectedRoute'))
const Leaderboards = lazy(() => import('~pages/Leaderboards'))
const Points = lazy(() => import('~pages/points'))
const Poll = lazy(() => import('~pages/Poll'))
const Profile = lazy(() => import('~pages/Profile'))
const ProtectedRoute = lazy(() => import('./ProtectedRoute'))

export const Router = () => {
  const router = createHashRouter([
    {
      path: '/',
      element: <Layout />,
      children: [
        {
          path: '/',
          element: (
            <SuspenseLoader>
              <Home />
            </SuspenseLoader>
          ),
        },
        {
          path: '/form/:id?',
          element: (
            <SuspenseLoader>
              <AppForm />
            </SuspenseLoader>
          ),
        },
        {
          path: '/about',
          element: (
            <SuspenseLoader>
              <About />
            </SuspenseLoader>
          ),
        },
        {
          path: '/leaderboards',
          element: (
            <SuspenseLoader>
              <Leaderboards />
            </SuspenseLoader>
          ),
        },
        {
          path: '/poll/:pid',
          element: (
            <SuspenseLoader>
              <Poll />
            </SuspenseLoader>
          ),
        },
        {
          path: '/communities/:id',
          loader: ({ params: { id } }) => {
            return redirect(`/communities/degen/${id}`)
          },
        },
        {
          path: '/communities/:id/poll/:pid',
          loader: ({ params: { id, pid } }) => {
            return redirect(`/communities/degen/${id}/poll/${pid}`)
          },
        },
        {
          path: '/communities/:chain/:id',
          element: (
            <SuspenseLoader>
              <Community />
            </SuspenseLoader>
          ),
        },
        {
          path: '/communities/:chain/:community/poll/:poll',
          element: (
            <SuspenseLoader>
              <CommunityPoll />
            </SuspenseLoader>
          ),
        },
        {
          path: '/profile/:id',
          element: (
            <SuspenseLoader>
              <Profile />
            </SuspenseLoader>
          ),
        },
        {
          element: (
            <SuspenseLoader>
              <ProtectedRoute />
            </SuspenseLoader>
          ),
          children: [
            {
              path: '/profile',
              element: (
                <SuspenseLoader>
                  <Profile />
                </SuspenseLoader>
              ),
            },
            {
              path: '/points',
              element: (
                <SuspenseLoader>
                  <Points />
                </SuspenseLoader>
              ),
            },
          ],
        },
        {
          element: (
            <SuspenseLoader>
              <FarcasterAccountProtectedRoute />
            </SuspenseLoader>
          ),
          children: [
            {
              path: '/communities/new',
              element: (
                <SuspenseLoader>
                  <CommunitiesNew />
                </SuspenseLoader>
              ),
            },
          ],
        },
        {
          element: (
            <SuspenseLoader>
              <CommunitiesLayout />
            </SuspenseLoader>
          ),
          children: [
            {
              path: '/communities/page?/:page?',
              element: (
                <SuspenseLoader>
                  <AllCommunitiesList />
                </SuspenseLoader>
              ),
            },
            {
              element: (
                <SuspenseLoader>
                  <ProtectedRoute />
                </SuspenseLoader>
              ),
              children: [
                {
                  path: '/communities/mine/:page?',
                  element: (
                    <SuspenseLoader>
                      <MyCommunitiesList />
                    </SuspenseLoader>
                  ),
                },
              ],
            },
          ],
        },
      ],
    },
  ])

  return <RouterProvider router={router} />
}
