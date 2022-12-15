/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { rest } from 'msw'
import { faker } from '@faker-js/faker'

export const getGetProjectMock = () => ({
  createdAt: (() => faker.date.past())(),
  description: faker.random.word(),
  initialized: faker.datatype.boolean(),
  name: faker.random.word(),
  projectID: faker.datatype.number({ min: undefined, max: undefined }),
  updatedAt: (() => faker.date.past())(),
})

export const getGetProjectBoardMock = () => ({
  activities: faker.helpers.arrayElement([
    Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      activityID: faker.datatype.number({ min: undefined, max: undefined }),
      description: faker.random.word(),
      isProductive: faker.datatype.boolean(),
      name: faker.random.word(),
      projectID: faker.datatype.number({ min: undefined, max: undefined }),
    })),
    undefined,
  ]),
  kanbanSteps: faker.helpers.arrayElement([
    Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      color: faker.random.word(),
      description: faker.random.word(),
      kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
      name: faker.random.word(),
      projectID: faker.datatype.number({ min: undefined, max: undefined }),
      stepOrder: faker.helpers.arrayElement([faker.datatype.number({ min: undefined, max: undefined }), null]),
      timeTrackable: faker.datatype.boolean(),
    })),
    undefined,
  ]),
  project: faker.helpers.arrayElement([
    {
      createdAt: (() => faker.date.past())(),
      description: faker.random.word(),
      initialized: faker.datatype.boolean(),
      name: faker.random.word(),
      projectID: faker.datatype.number({ min: undefined, max: undefined }),
      updatedAt: (() => faker.date.past())(),
    },
    undefined,
  ]),
  teams: faker.helpers.arrayElement([
    Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      createdAt: (() => faker.date.past())(),
      description: faker.random.word(),
      name: faker.random.word(),
      projectID: faker.datatype.number({ min: undefined, max: undefined }),
      teamID: faker.datatype.number({ min: undefined, max: undefined }),
      updatedAt: (() => faker.date.past())(),
    })),
    undefined,
  ]),
  workItemTags: faker.helpers.arrayElement([
    Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      color: faker.random.word(),
      description: faker.random.word(),
      name: faker.random.word(),
      projectID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTagID: faker.datatype.number({ min: undefined, max: undefined }),
    })),
    undefined,
  ]),
  workItemTypes: faker.helpers.arrayElement([
    Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      color: faker.random.word(),
      description: faker.random.word(),
      name: faker.random.word(),
      projectID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
    })),
    undefined,
  ]),
})

export const getGetProjectWorkitemsMock = () =>
  faker.helpers.arrayElement([
    {
      baseWorkItem: {
        closed: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        createdAt: (() => faker.date.past())(),
        deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        description: faker.random.word(),
        kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
        metadata: {},
        targetDate: (() => faker.date.past())(),
        teamID: faker.datatype.number({ min: undefined, max: undefined }),
        title: faker.random.word(),
        updatedAt: (() => faker.date.past())(),
        workItemID: faker.datatype.number({ min: undefined, max: undefined }),
        workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
      },
      lastMessageAt: (() => faker.date.past())(),
      line: faker.random.word(),
      ref: faker.random.word(),
      reopened: faker.datatype.boolean(),
      workItemID: faker.datatype.number({ min: undefined, max: undefined }),
    },
  ])

export const getProjectMSW = () => [
  rest.post('*/project/:id/initialize', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'))
  }),
  rest.get('*/project/:id/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetProjectMock()))
  }),
  rest.get('*/project/:id/board', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetProjectBoardMock()))
  }),
  rest.get('*/project/:id/workitems', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetProjectWorkitemsMock()))
  }),
]
