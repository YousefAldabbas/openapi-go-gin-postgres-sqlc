/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { GetPaginatedNotificationsParams } from '../model/getPaginatedNotificationsParams'
import type { PaginatedNotificationsResponse } from '../model/paginatedNotificationsResponse'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary Get paginated user notifications
 */
export const getPaginatedNotifications = (
  params: GetPaginatedNotificationsParams,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<PaginatedNotificationsResponse>(
    { url: `/notifications/user/page`, method: 'GET', params },
    options,
  )
}
export type GetPaginatedNotificationsResult = NonNullable<Awaited<ReturnType<typeof getPaginatedNotifications>>>
