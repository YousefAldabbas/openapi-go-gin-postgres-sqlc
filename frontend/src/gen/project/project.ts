/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { useQuery, useInfiniteQuery, useMutation } from '@tanstack/react-query'
import type {
  UseQueryOptions,
  UseInfiniteQueryOptions,
  UseMutationOptions,
  QueryFunction,
  MutationFunction,
  UseQueryResult,
  UseInfiniteQueryResult,
  QueryKey,
} from '@tanstack/react-query'
import type {
  InitializeProjectRequest,
  DbProject,
  ProjectConfig,
  ProjectBoardResponse,
  GetProjectWorkitems200,
  GetProjectWorkitemsParams,
  DbWorkItemTag,
  WorkItemTagCreateRequest,
} from '.././model'
import { customInstance } from '../../api/mutator'

type AwaitedInput<T> = PromiseLike<T> | T

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary creates initial data (teams, work item types, tags...) for a new project
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

export const getInitializeProjectMutationOptions = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof initializeProject>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof initializeProject>>,
    { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return initializeProject(projectName, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type InitializeProjectMutationResult = NonNullable<Awaited<ReturnType<typeof initializeProject>>>
export type InitializeProjectMutationBody = InitializeProjectRequest
export type InitializeProjectMutationError = unknown

export const useInitializeProject = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getInitializeProjectMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary returns board data for a project
 */
export const getProject = (
  projectName: 'demo' | 'demo_two',
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<DbProject>({ url: `/project/${projectName}/`, method: 'get', signal }, options)
}

export const getGetProjectQueryKey = (projectName: 'demo' | 'demo_two') => [`/project/${projectName}/`] as const

export const getGetProjectInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) =>
    getProject(projectName, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectInfiniteQueryError = unknown

export const useGetProjectInfinite = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectInfiniteQueryOptions(projectName, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectQueryOptions = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) =>
    getProject(projectName, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectQueryError = unknown

export const useGetProject = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectQueryOptions(projectName, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary returns the project configuration
 */
export const getProjectConfig = (
  projectName: 'demo' | 'demo_two',
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<ProjectConfig>({ url: `/project/${projectName}/config`, method: 'get', signal }, options)
}

export const getGetProjectConfigQueryKey = (projectName: 'demo' | 'demo_two') =>
  [`/project/${projectName}/config`] as const

export const getGetProjectConfigInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectConfig>>,
  TError = unknown,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectConfigQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) =>
    getProjectConfig(projectName, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectConfigInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigInfiniteQueryError = unknown

export const useGetProjectConfigInfinite = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectConfigInfiniteQueryOptions(projectName, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectConfigQueryOptions = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectConfigQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) =>
    getProjectConfig(projectName, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectConfigQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigQueryError = unknown

export const useGetProjectConfig = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectConfigQueryOptions(projectName, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
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

export const getUpdateProjectConfigMutationOptions = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateProjectConfig>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: ProjectConfig },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof updateProjectConfig>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: ProjectConfig },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateProjectConfig>>,
    { projectName: 'demo' | 'demo_two'; data: ProjectConfig }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return updateProjectConfig(projectName, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type UpdateProjectConfigMutationResult = NonNullable<Awaited<ReturnType<typeof updateProjectConfig>>>
export type UpdateProjectConfigMutationBody = ProjectConfig
export type UpdateProjectConfigMutationError = unknown

export const useUpdateProjectConfig = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateProjectConfig>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: ProjectConfig },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getUpdateProjectConfigMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary returns board data for a project
 */
export const getProjectBoard = (
  projectName: 'demo' | 'demo_two',
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<ProjectBoardResponse>({ url: `/project/${projectName}/board`, method: 'get', signal }, options)
}

export const getGetProjectBoardQueryKey = (projectName: 'demo' | 'demo_two') =>
  [`/project/${projectName}/board`] as const

export const getGetProjectBoardInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectBoard>>,
  TError = unknown,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(projectName, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectBoardInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardInfiniteQueryError = unknown

export const useGetProjectBoardInfinite = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectBoardInfiniteQueryOptions(projectName, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectBoardQueryOptions = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(projectName, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectBoardQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardQueryError = unknown

export const useGetProjectBoard = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectBoardQueryOptions(projectName, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary returns workitems for a project
 */
export const getProjectWorkitems = (
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<GetProjectWorkitems200>(
    { url: `/project/${projectName}/workitems`, method: 'get', params, signal },
    options,
  )
}

export const getGetProjectWorkitemsQueryKey = (projectName: 'demo' | 'demo_two', params?: GetProjectWorkitemsParams) =>
  [`/project/${projectName}/workitems`, ...(params ? [params] : [])] as const

export const getGetProjectWorkitemsInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = unknown,
>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(projectName, params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) =>
    getProjectWorkitems(projectName, params, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectWorkitemsInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsInfiniteQueryError = unknown

export const useGetProjectWorkitemsInfinite = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = unknown,
>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectWorkitemsInfiniteQueryOptions(projectName, params, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectWorkitemsQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = unknown,
>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(projectName, params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) =>
    getProjectWorkitems(projectName, params, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectWorkitemsQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsQueryError = unknown

export const useGetProjectWorkitems = <TData = Awaited<ReturnType<typeof getProjectWorkitems>>, TError = unknown>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectWorkitemsQueryOptions(projectName, params, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary create workitem tag
 */
export const createWorkitemTag = (
  projectName: 'demo' | 'demo_two',
  workItemTagCreateRequest: WorkItemTagCreateRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<DbWorkItemTag>(
    {
      url: `/project/${projectName}/tag/`,
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: workItemTagCreateRequest,
    },
    options,
  )
}

export const getCreateWorkitemTagMutationOptions = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitemTag>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: WorkItemTagCreateRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof createWorkitemTag>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: WorkItemTagCreateRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof createWorkitemTag>>,
    { projectName: 'demo' | 'demo_two'; data: WorkItemTagCreateRequest }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return createWorkitemTag(projectName, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateWorkitemTagMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkitemTag>>>
export type CreateWorkitemTagMutationBody = WorkItemTagCreateRequest
export type CreateWorkitemTagMutationError = unknown

export const useCreateWorkitemTag = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitemTag>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: WorkItemTagCreateRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getCreateWorkitemTagMutationOptions(options)

  return useMutation(mutationOptions)
}
