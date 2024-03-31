import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import {
  faker
} from '@faker-js/faker'
import {
  HttpResponse,
  delay,
  http
} from 'msw'

export const getAdminPingResponseMock = (): string => (faker.word.sample())


export const getAdminPingMockHandler = () => {
  return http.get('*/admin/ping', async () => {
    await delay(1000);
    return new HttpResponse(getAdminPingResponseMock(),
      {
        status: 200,
        headers: {
          'Content-Type': 'text/plain',
        }
      }
    )
  })
}
export const getAdminMock = () => [
  getAdminPingMockHandler()
]
