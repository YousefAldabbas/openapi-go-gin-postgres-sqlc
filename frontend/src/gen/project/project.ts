import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import {
  useInfiniteQuery,
  useMutation,
  useQuery
} from '@tanstack/react-query'
import type {
  MutationFunction,
  QueryFunction,
  QueryKey,
  UseInfiniteQueryOptions,
  UseInfiniteQueryResult,
  UseMutationOptions,
  UseQueryOptions,
  UseQueryResult
} from '@tanstack/react-query'
import type {
  CreateProjectBoardRequest
} from '../model/createProjectBoardRequest'
import type {
  DbProject
} from '../model/dbProject'
import type {
  GetProjectWorkitems200
} from '../model/getProjectWorkitems200'
import type {
  GetProjectWorkitemsParams
} from '../model/getProjectWorkitemsParams'
import type {
  ProjectBoard
} from '../model/projectBoard'
import type {
  ProjectConfig
} from '../model/projectConfig'
import { customInstance } from '../../api/mutator';


// eslint-disable-next-line
  type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P,
) => any
  ? P
  : never;


/**
 * @summary creates initial data (teams, tags...) for a new project
 */
export const initializeProject = (
    projectName: 'demo' | 'demo_two',
    createProjectBoardRequest: CreateProjectBoardRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/project/${projectName}/initialize`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: createProjectBoardRequest
    },
      options);
    }
  


export const getInitializeProjectMutationOptions = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof initializeProject>>, TError,{projectName: 'demo' | 'demo_two';data: CreateProjectBoardRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof initializeProject>>, TError,{projectName: 'demo' | 'demo_two';data: CreateProjectBoardRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof initializeProject>>, {projectName: 'demo' | 'demo_two';data: CreateProjectBoardRequest}> = (props) => {
          const {projectName,data} = props ?? {};

          return  initializeProject(projectName,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type InitializeProjectMutationResult = NonNullable<Awaited<ReturnType<typeof initializeProject>>>
    export type InitializeProjectMutationBody = CreateProjectBoardRequest
    export type InitializeProjectMutationError = unknown

    /**
 * @summary creates initial data (teams, tags...) for a new project
 */
export const useInitializeProject = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof initializeProject>>, TError,{projectName: 'demo' | 'demo_two';data: CreateProjectBoardRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getInitializeProjectMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary returns board data for a project
 */
export const getProject = (
    projectName: 'demo' | 'demo_two',
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<DbProject>(
      {url: `/project/${projectName}/`, method: 'GET', signal
    },
      options);
    }
  

export const getGetProjectQueryKey = (projectName: 'demo' | 'demo_two',) => {
    return [`/project/${projectName}/`] as const;
    }

    
export const getGetProjectInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(projectName: 'demo' | 'demo_two', options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectQueryKey(projectName);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) => getProject(projectName, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectInfiniteQueryError = unknown

/**
 * @summary returns board data for a project
 */
export const useGetProjectInfinite = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(
 projectName: 'demo' | 'demo_two', options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectInfiniteQueryOptions(projectName,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetProjectQueryOptions = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(projectName: 'demo' | 'demo_two', options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectQueryKey(projectName);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) => getProject(projectName, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectQueryError = unknown

/**
 * @summary returns board data for a project
 */
export const useGetProject = <TData = Awaited<ReturnType<typeof getProject>>, TError = unknown>(
 projectName: 'demo' | 'demo_two', options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectQueryOptions(projectName,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary returns the project configuration
 */
export const getProjectConfig = (
    projectName: 'demo' | 'demo_two',
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<ProjectConfig>(
      {url: `/project/${projectName}/config`, method: 'GET', signal
    },
      options);
    }
  

export const getGetProjectConfigQueryKey = (projectName: 'demo' | 'demo_two',) => {
    return [`/project/${projectName}/config`] as const;
    }

    
export const getGetProjectConfigInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = unknown>(projectName: 'demo' | 'demo_two', options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectConfigQueryKey(projectName);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) => getProjectConfig(projectName, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectConfigInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigInfiniteQueryError = unknown

/**
 * @summary returns the project configuration
 */
export const useGetProjectConfigInfinite = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = unknown>(
 projectName: 'demo' | 'demo_two', options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectConfigInfiniteQueryOptions(projectName,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetProjectConfigQueryOptions = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = unknown>(projectName: 'demo' | 'demo_two', options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectConfigQueryKey(projectName);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) => getProjectConfig(projectName, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectConfigQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigQueryError = unknown

/**
 * @summary returns the project configuration
 */
export const useGetProjectConfig = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = unknown>(
 projectName: 'demo' | 'demo_two', options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectConfigQueryOptions(projectName,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary updates the project configuration
 */
export const updateProjectConfig = (
    projectName: 'demo' | 'demo_two',
    projectConfig: ProjectConfig,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/project/${projectName}/config`, method: 'PUT',
      headers: {'Content-Type': 'application/json', },
      data: projectConfig
    },
      options);
    }
  


export const getUpdateProjectConfigMutationOptions = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateProjectConfig>>, TError,{projectName: 'demo' | 'demo_two';data: ProjectConfig}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateProjectConfig>>, TError,{projectName: 'demo' | 'demo_two';data: ProjectConfig}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateProjectConfig>>, {projectName: 'demo' | 'demo_two';data: ProjectConfig}> = (props) => {
          const {projectName,data} = props ?? {};

          return  updateProjectConfig(projectName,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type UpdateProjectConfigMutationResult = NonNullable<Awaited<ReturnType<typeof updateProjectConfig>>>
    export type UpdateProjectConfigMutationBody = ProjectConfig
    export type UpdateProjectConfigMutationError = unknown

    /**
 * @summary updates the project configuration
 */
export const useUpdateProjectConfig = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateProjectConfig>>, TError,{projectName: 'demo' | 'demo_two';data: ProjectConfig}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getUpdateProjectConfigMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary returns board data for a project
 */
export const getProjectBoard = (
    projectName: 'demo' | 'demo_two',
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<ProjectBoard>(
      {url: `/project/${projectName}/board`, method: 'GET', signal
    },
      options);
    }
  

export const getGetProjectBoardQueryKey = (projectName: 'demo' | 'demo_two',) => {
    return [`/project/${projectName}/board`] as const;
    }

    
export const getGetProjectBoardInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = unknown>(projectName: 'demo' | 'demo_two', options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectBoardQueryKey(projectName);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) => getProjectBoard(projectName, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectBoardInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardInfiniteQueryError = unknown

/**
 * @summary returns board data for a project
 */
export const useGetProjectBoardInfinite = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = unknown>(
 projectName: 'demo' | 'demo_two', options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectBoardInfiniteQueryOptions(projectName,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetProjectBoardQueryOptions = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = unknown>(projectName: 'demo' | 'demo_two', options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectBoardQueryKey(projectName);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) => getProjectBoard(projectName, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectBoardQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardQueryError = unknown

/**
 * @summary returns board data for a project
 */
export const useGetProjectBoard = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = unknown>(
 projectName: 'demo' | 'demo_two', options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectBoardQueryOptions(projectName,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary returns workitems for a project
 */
export const getProjectWorkitems = (
    projectName: 'demo' | 'demo_two',
    params?: GetProjectWorkitemsParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<GetProjectWorkitems200>(
      {url: `/project/${projectName}/workitems`, method: 'GET',
        params, signal
    },
      options);
    }
  

export const getGetProjectWorkitemsQueryKey = (projectName: 'demo' | 'demo_two',
    params?: GetProjectWorkitemsParams,) => {
    return [`/project/${projectName}/workitems`, ...(params ? [params]: [])] as const;
    }

    
export const getGetProjectWorkitemsInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getProjectWorkitems>>, TError = unknown>(projectName: 'demo' | 'demo_two',
    params?: GetProjectWorkitemsParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(projectName,params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) => getProjectWorkitems(projectName,params, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectWorkitemsInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsInfiniteQueryError = unknown

/**
 * @summary returns workitems for a project
 */
export const useGetProjectWorkitemsInfinite = <TData = Awaited<ReturnType<typeof getProjectWorkitems>>, TError = unknown>(
 projectName: 'demo' | 'demo_two',
    params?: GetProjectWorkitemsParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectWorkitemsInfiniteQueryOptions(projectName,params,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetProjectWorkitemsQueryOptions = <TData = Awaited<ReturnType<typeof getProjectWorkitems>>, TError = unknown>(projectName: 'demo' | 'demo_two',
    params?: GetProjectWorkitemsParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(projectName,params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) => getProjectWorkitems(projectName,params, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(projectName),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData> & { queryKey: QueryKey }
}

export type GetProjectWorkitemsQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsQueryError = unknown

/**
 * @summary returns workitems for a project
 */
export const useGetProjectWorkitems = <TData = Awaited<ReturnType<typeof getProjectWorkitems>>, TError = unknown>(
 projectName: 'demo' | 'demo_two',
    params?: GetProjectWorkitemsParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetProjectWorkitemsQueryOptions(projectName,params,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



