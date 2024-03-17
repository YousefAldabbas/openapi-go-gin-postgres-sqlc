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
  CreateTimeEntryRequest
} from '../model/createTimeEntryRequest'
import type {
  HTTPError
} from '../model/hTTPError'
import type {
  TimeEntry
} from '../model/timeEntry'
import type {
  UpdateTimeEntryRequest
} from '../model/updateTimeEntryRequest'
import { customInstance } from '../../api/mutator';
import type { ErrorType } from '../../api/mutator';


// eslint-disable-next-line
  type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P,
) => any
  ? P
  : never;


/**
 * @summary create time entry.
 */
export const createTimeEntry = (
    createTimeEntryRequest: CreateTimeEntryRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<TimeEntry>(
      {url: `/time-entry/`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: createTimeEntryRequest
    },
      options);
    }
  


export const getCreateTimeEntryMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createTimeEntry>>, TError,{data: CreateTimeEntryRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createTimeEntry>>, TError,{data: CreateTimeEntryRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createTimeEntry>>, {data: CreateTimeEntryRequest}> = (props) => {
          const {data} = props ?? {};

          return  createTimeEntry(data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type CreateTimeEntryMutationResult = NonNullable<Awaited<ReturnType<typeof createTimeEntry>>>
    export type CreateTimeEntryMutationBody = CreateTimeEntryRequest
    export type CreateTimeEntryMutationError = ErrorType<void | HTTPError>

    /**
 * @summary create time entry.
 */
export const useCreateTimeEntry = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createTimeEntry>>, TError,{data: CreateTimeEntryRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getCreateTimeEntryMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary get time-entry.
 */
export const getTimeEntry = (
    timeEntryID: EntityIDs.TimeEntryID,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<TimeEntry>(
      {url: `/time-entry/${timeEntryID}`, method: 'GET', signal
    },
      options);
    }
  

export const getGetTimeEntryQueryKey = (timeEntryID: EntityIDs.TimeEntryID,) => {
    return [`/time-entry/${timeEntryID}`] as const;
    }

    
export const getGetTimeEntryInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getTimeEntry>>, TError = ErrorType<void | HTTPError>>(timeEntryID: EntityIDs.TimeEntryID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getTimeEntry>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetTimeEntryQueryKey(timeEntryID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getTimeEntry>>> = ({ signal }) => getTimeEntry(timeEntryID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(timeEntryID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getTimeEntry>>, TError, TData> & { queryKey: QueryKey }
}

export type GetTimeEntryInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getTimeEntry>>>
export type GetTimeEntryInfiniteQueryError = ErrorType<void | HTTPError>

/**
 * @summary get time-entry.
 */
export const useGetTimeEntryInfinite = <TData = Awaited<ReturnType<typeof getTimeEntry>>, TError = ErrorType<void | HTTPError>>(
 timeEntryID: EntityIDs.TimeEntryID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getTimeEntry>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetTimeEntryInfiniteQueryOptions(timeEntryID,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetTimeEntryQueryOptions = <TData = Awaited<ReturnType<typeof getTimeEntry>>, TError = ErrorType<void | HTTPError>>(timeEntryID: EntityIDs.TimeEntryID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getTimeEntry>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetTimeEntryQueryKey(timeEntryID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getTimeEntry>>> = ({ signal }) => getTimeEntry(timeEntryID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(timeEntryID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getTimeEntry>>, TError, TData> & { queryKey: QueryKey }
}

export type GetTimeEntryQueryResult = NonNullable<Awaited<ReturnType<typeof getTimeEntry>>>
export type GetTimeEntryQueryError = ErrorType<void | HTTPError>

/**
 * @summary get time-entry.
 */
export const useGetTimeEntry = <TData = Awaited<ReturnType<typeof getTimeEntry>>, TError = ErrorType<void | HTTPError>>(
 timeEntryID: EntityIDs.TimeEntryID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getTimeEntry>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetTimeEntryQueryOptions(timeEntryID,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary update time-entry.
 */
export const updateTimeEntry = (
    timeEntryID: EntityIDs.TimeEntryID,
    updateTimeEntryRequest: UpdateTimeEntryRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<TimeEntry>(
      {url: `/time-entry/${timeEntryID}`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: updateTimeEntryRequest
    },
      options);
    }
  


export const getUpdateTimeEntryMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateTimeEntry>>, TError,{timeEntryID: EntityIDs.TimeEntryID;data: UpdateTimeEntryRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateTimeEntry>>, TError,{timeEntryID: EntityIDs.TimeEntryID;data: UpdateTimeEntryRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateTimeEntry>>, {timeEntryID: EntityIDs.TimeEntryID;data: UpdateTimeEntryRequest}> = (props) => {
          const {timeEntryID,data} = props ?? {};

          return  updateTimeEntry(timeEntryID,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type UpdateTimeEntryMutationResult = NonNullable<Awaited<ReturnType<typeof updateTimeEntry>>>
    export type UpdateTimeEntryMutationBody = UpdateTimeEntryRequest
    export type UpdateTimeEntryMutationError = ErrorType<void | HTTPError>

    /**
 * @summary update time-entry.
 */
export const useUpdateTimeEntry = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateTimeEntry>>, TError,{timeEntryID: EntityIDs.TimeEntryID;data: UpdateTimeEntryRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getUpdateTimeEntryMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary delete time-entry.
 */
export const deleteTimeEntry = (
    timeEntryID: EntityIDs.TimeEntryID,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/time-entry/${timeEntryID}`, method: 'DELETE'
    },
      options);
    }
  


export const getDeleteTimeEntryMutationOptions = <TError = ErrorType<HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteTimeEntry>>, TError,{timeEntryID: EntityIDs.TimeEntryID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteTimeEntry>>, TError,{timeEntryID: EntityIDs.TimeEntryID}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteTimeEntry>>, {timeEntryID: EntityIDs.TimeEntryID}> = (props) => {
          const {timeEntryID} = props ?? {};

          return  deleteTimeEntry(timeEntryID,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type DeleteTimeEntryMutationResult = NonNullable<Awaited<ReturnType<typeof deleteTimeEntry>>>
    
    export type DeleteTimeEntryMutationError = ErrorType<HTTPError>

    /**
 * @summary delete time-entry.
 */
export const useDeleteTimeEntry = <TError = ErrorType<HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteTimeEntry>>, TError,{timeEntryID: EntityIDs.TimeEntryID}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getDeleteTimeEntryMutationOptions(options);

      return useMutation(mutationOptions);
    }
    