import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
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
import {
  NotificationType
} from '.././model'

export const getGetPaginatedNotificationsMock = () => ({items: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({notification: {body: faker.word.sample(), createdAt: (() => faker.date.past())(), labels: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => (faker.word.sample())), link: faker.helpers.arrayElement([faker.word.sample(), null]), notificationID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.NotificationID, notificationType: faker.helpers.arrayElement(Object.values(NotificationType)), receiver: faker.word.sample(), sender: faker.word.sample(), title: faker.word.sample()}, notificationID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.NotificationID, read: faker.datatype.boolean(), userID: faker.string.uuid() as EntityIDs.UserID, userNotificationID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.UserNotificationID})), page: {nextCursor: faker.word.sample()}})

export const getNotificationsMock = () => [
http.get('*/notifications/user/page', async () => {
        await delay(1000);
        return new HttpResponse(JSON.stringify(getGetPaginatedNotificationsMock()),
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),]
