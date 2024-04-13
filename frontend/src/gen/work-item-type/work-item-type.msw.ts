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
import type {
  WorkItemType
} from '.././model'

export const getCreateWorkItemTypeResponseMock = (overrideResponse: any = {}): WorkItemType => ({color: faker.word.sample(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID, ...overrideResponse})

export const getGetWorkItemTypeResponseMock = (overrideResponse: any = {}): WorkItemType => ({color: faker.word.sample(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID, ...overrideResponse})

export const getUpdateWorkItemTypeResponseMock = (overrideResponse: any = {}): WorkItemType => ({color: faker.word.sample(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID, ...overrideResponse})


export const getCreateWorkItemTypeMockHandler = (overrideResponse?: WorkItemType) => {
  return http.post('*/project/:projectName/work-item-type/', async () => {
    await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getCreateWorkItemTypeResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetWorkItemTypeMockHandler = (overrideResponse?: WorkItemType) => {
  return http.get('*/work-item-type/:workItemTypeID', async () => {
    await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getGetWorkItemTypeResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateWorkItemTypeMockHandler = (overrideResponse?: WorkItemType) => {
  return http.patch('*/work-item-type/:workItemTypeID', async () => {
    await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getUpdateWorkItemTypeResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteWorkItemTypeMockHandler = () => {
  return http.delete('*/work-item-type/:workItemTypeID', async () => {
    await delay(200);
    return new HttpResponse(null,
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}
export const getWorkItemTypeMock = () => [
  getCreateWorkItemTypeMockHandler(),
  getGetWorkItemTypeMockHandler(),
  getUpdateWorkItemTypeMockHandler(),
  getDeleteWorkItemTypeMockHandler()
]
