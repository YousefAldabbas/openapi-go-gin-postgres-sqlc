import { shallowEqual } from 'react-redux'
import { capitalize } from 'lodash'
import { useMemo } from 'react'
import { COLOR_BLIND_PALETTE, generateColor } from 'src/utils/colors'
import roles from '@roles'
import scopes from '@scopes'
import type { UserResponse } from 'src/gen/model'

export const useAuthenticatedUser = () => {
  // TODO for app_env dev, remove Authorization header and comes from backend via x-api-key header
  // or have fallthorugh if authentication failed instead - would need multierror
  const user: UserResponse = {
    hasGlobalNotifications: true,
    hasPersonalNotifications: true,
    role: 'admin',
    userID: crypto.randomUUID(),
    email: 'admin@mail.com',
    firstName: 'John',
    lastName: 'Doe',
    fullName: 'John Doe',
    username: 'john.doe',
    scopes: ['users:read', 'test-scope', 'scopes:write'],
    createdAt: new Date(),
    deletedAt: null,
  }

  const logUserOut = () => {
    null
  }

  return {
    user,
    logUserOut,
  }
}
