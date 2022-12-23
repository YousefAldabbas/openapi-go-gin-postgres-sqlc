/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import axios from 'axios'
import type { AxiosRequestConfig, AxiosResponse, AxiosError } from 'axios'
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
  DbProjectPublic,
  ProjectConfigResponse,
  ProjectBoardResponse,
  DemoProjectWorkItemsResponse,
  GetProjectWorkitemsParams,
} from '.././model'

/**
 * @summary creates initial data (teams, work item types, tags...) for a new project
 */
export const initializeProject = (
  id: number,
  initializeProjectRequest: InitializeProjectRequest,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<void>> => {
  return axios.post(`/project/${id}/initialize`, initializeProjectRequest, options)
}

export type InitializeProjectMutationResult = NonNullable<Awaited<ReturnType<typeof initializeProject>>>
export type InitializeProjectMutationBody = InitializeProjectRequest
export type InitializeProjectMutationError = AxiosError<unknown>

export const useInitializeProject = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { id: number; data: InitializeProjectRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof initializeProject>>,
    { id: number; data: InitializeProjectRequest }
  > = (props) => {
    const { id, data } = props ?? {}

    return initializeProject(id, data, axiosOptions)
  }

  return useMutation<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { id: number; data: InitializeProjectRequest },
    TContext
  >(mutationFn, mutationOptions)
}
/**
 * @summary returns board data for a project
 */
export const getProject = (id: number, options?: AxiosRequestConfig): Promise<AxiosResponse<DbProjectPublic>> => {
  return axios.get(`/project/${id}/`, options)
}

export const getGetProjectQueryKey = (id: number) => [`/project/${id}/`]

export type GetProjectInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectInfiniteQueryError = AxiosError<unknown>

export const useGetProjectInfinite = <TData = Awaited<ReturnType<typeof getProject>>, TError = AxiosError<unknown>>(
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) =>
    getProject(id, { signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof getProject>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type GetProjectQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectQueryError = AxiosError<unknown>

export const useGetProject = <TData = Awaited<ReturnType<typeof getProject>>, TError = AxiosError<unknown>>(
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) =>
    getProject(id, { signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof getProject>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary returns the project configuration
 */
export const getProjectConfig = (
  id: number,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<ProjectConfigResponse>> => {
  return axios.get(`/project/${id}/config`, options)
}

export const getGetProjectConfigQueryKey = (id: number) => [`/project/${id}/config`]

export type GetProjectConfigInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigInfiniteQueryError = AxiosError<unknown>

export const useGetProjectConfigInfinite = <
  TData = Awaited<ReturnType<typeof getProjectConfig>>,
  TError = AxiosError<unknown>,
>(
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectConfigQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) =>
    getProjectConfig(id, { signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type GetProjectConfigQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigQueryError = AxiosError<unknown>

export const useGetProjectConfig = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = AxiosError<unknown>>(
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectConfigQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) =>
    getProjectConfig(id, { signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary returns board data for a project
 */
export const getProjectBoard = (
  id: number,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<ProjectBoardResponse>> => {
  return axios.get(`/project/${id}/board`, options)
}

export const getGetProjectBoardQueryKey = (id: number) => [`/project/${id}/board`]

export type GetProjectBoardInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardInfiniteQueryError = AxiosError<unknown>

export const useGetProjectBoardInfinite = <
  TData = Awaited<ReturnType<typeof getProjectBoard>>,
  TError = AxiosError<unknown>,
>(
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(id, { signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type GetProjectBoardQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardQueryError = AxiosError<unknown>

export const useGetProjectBoard = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = AxiosError<unknown>>(
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(id, { signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary returns workitems for a project
 */
export const getProjectWorkitems = (
  id: number,
  params?: GetProjectWorkitemsParams,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<DemoProjectWorkItemsResponse>> => {
  return axios.get(`/project/${id}/workitems`, {
    ...options,
    params: { ...params, ...options?.params },
  })
}

export const getGetProjectWorkitemsQueryKey = (id: number, params?: GetProjectWorkitemsParams) => [
  `/project/${id}/workitems`,
  ...(params ? [params] : []),
]

export type GetProjectWorkitemsInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsInfiniteQueryError = AxiosError<unknown>

export const useGetProjectWorkitemsInfinite = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = AxiosError<unknown>,
>(
  id: number,
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(id, params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) =>
    getProjectWorkitems(id, params, { signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type GetProjectWorkitemsQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsQueryError = AxiosError<unknown>

export const useGetProjectWorkitems = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = AxiosError<unknown>,
>(
  id: number,
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(id, params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) =>
    getProjectWorkitems(id, params, { signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}
