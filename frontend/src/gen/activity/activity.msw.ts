import type { Branded } from 'src/types/utils'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { HttpResponse, delay, http } from 'msw'

export const getCreateActivityMock = () => ({
  activityID: faker.number.int({ min: undefined, max: undefined }),
  deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
  description: faker.word.sample(),
  isProductive: faker.datatype.boolean(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
})

export const getGetActivityMock = () => ({
  activityID: faker.number.int({ min: undefined, max: undefined }),
  deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
  description: faker.word.sample(),
  isProductive: faker.datatype.boolean(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
})

export const getUpdateActivityMock = () => ({
  activityID: faker.number.int({ min: undefined, max: undefined }),
  deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
  description: faker.word.sample(),
  isProductive: faker.datatype.boolean(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
})

export const getActivityMock = () => [
  http.post('*/project/:projectName/activity/', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getCreateActivityMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/activity/:activityID', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetActivityMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.patch('*/activity/:activityID', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getUpdateActivityMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.delete('*/activity/:activityID', async () => {
    await delay(1000)
    return new HttpResponse(null, {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
]
