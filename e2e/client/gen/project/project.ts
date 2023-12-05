/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type {
  InitializeProjectRequest,
  DbProject,
  ProjectConfig,
  ProjectBoard,
  GetProjectWorkitems200,
  GetProjectWorkitemsParams,
} from '.././model'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary creates initial data (teams, tags...) for a new project
 */
export const initializeProject = (
  projectName: 'demo' | 'demo_two',
  initializeProjectRequest: InitializeProjectRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>(
    {
      url: `/project/${projectName}/initialize`,
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: initializeProjectRequest,
    },
    options,
  )
}
/**
 * @summary returns board data for a project
 */
export const getProject = (projectName: 'demo' | 'demo_two', options?: SecondParameter<typeof customInstance>) => {
  return customInstance<DbProject>({ url: `/project/${projectName}/`, method: 'get' }, options)
}
/**
 * @summary returns the project configuration
 */
export const getProjectConfig = (
  projectName: 'demo' | 'demo_two',
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<ProjectConfig>({ url: `/project/${projectName}/config`, method: 'get' }, options)
}
/**
 * @summary updates the project configuration
 */
export const updateProjectConfig = (
  projectName: 'demo' | 'demo_two',
  projectConfig: ProjectConfig,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>(
    {
      url: `/project/${projectName}/config`,
      method: 'put',
      headers: { 'Content-Type': 'application/json' },
      data: projectConfig,
    },
    options,
  )
}
/**
 * @summary returns board data for a project
 */
export const getProjectBoard = (projectName: 'demo' | 'demo_two', options?: SecondParameter<typeof customInstance>) => {
  return customInstance<ProjectBoard>({ url: `/project/${projectName}/board`, method: 'get' }, options)
}
/**
 * @summary returns workitems for a project
 */
export const getProjectWorkitems = (
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<GetProjectWorkitems200>(
    { url: `/project/${projectName}/workitems`, method: 'get', params },
    options,
  )
}
export type InitializeProjectResult = NonNullable<Awaited<ReturnType<typeof initializeProject>>>
export type GetProjectResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectConfigResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type UpdateProjectConfigResult = NonNullable<Awaited<ReturnType<typeof updateProjectConfig>>>
export type GetProjectBoardResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectWorkitemsResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
