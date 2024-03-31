import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import {
  useInfiniteQuery,
  useQuery
} from '@tanstack/react-query'
import type {
  QueryFunction,
  QueryKey,
  UseInfiniteQueryOptions,
  UseInfiniteQueryResult,
  UseQueryOptions,
  UseQueryResult
} from '@tanstack/react-query'
import type {
  MyProviderLoginParams
} from '.././model'
import { customInstance } from '../../api/mutator';
import type { ErrorType } from '../../api/mutator';


type SecondParameter<T extends (...args: any) => any> = Parameters<T>[1];


export const myProviderCallback = (
    
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<unknown>(
      {url: `/auth/myprovider/callback`, method: 'GET', signal
    },
      options);
    }
  

export const getMyProviderCallbackQueryKey = () => {
    return [`/auth/myprovider/callback`] as const;
    }

    
export const getMyProviderCallbackInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof myProviderCallback>>, TError = ErrorType<void>>( options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getMyProviderCallbackQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderCallback>>> = ({ signal }) => myProviderCallback(requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData> & { queryKey: QueryKey }
}

export type MyProviderCallbackInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderCallback>>>
export type MyProviderCallbackInfiniteQueryError = ErrorType<void>

export const useMyProviderCallbackInfinite = <TData = Awaited<ReturnType<typeof myProviderCallback>>, TError = ErrorType<void>>(
  options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getMyProviderCallbackInfiniteQueryOptions(options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getMyProviderCallbackQueryOptions = <TData = Awaited<ReturnType<typeof myProviderCallback>>, TError = ErrorType<void>>( options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getMyProviderCallbackQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderCallback>>> = ({ signal }) => myProviderCallback(requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData> & { queryKey: QueryKey }
}

export type MyProviderCallbackQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderCallback>>>
export type MyProviderCallbackQueryError = ErrorType<void>

export const useMyProviderCallback = <TData = Awaited<ReturnType<typeof myProviderCallback>>, TError = ErrorType<void>>(
  options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getMyProviderCallbackQueryOptions(options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const myProviderLogin = (
    params: MyProviderLoginParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<unknown>(
      {url: `/auth/myprovider/login`, method: 'GET',
        params, signal
    },
      options);
    }
  

export const getMyProviderLoginQueryKey = (params: MyProviderLoginParams,) => {
    return [`/auth/myprovider/login`, ...(params ? [params]: [])] as const;
    }

    
export const getMyProviderLoginInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof myProviderLogin>>, TError = ErrorType<void>>(params: MyProviderLoginParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getMyProviderLoginQueryKey(params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderLogin>>> = ({ signal, pageParam }) => myProviderLogin({...params, cursor: pageParam || params?.['cursor']}, requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData> & { queryKey: QueryKey }
}

export type MyProviderLoginInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderLogin>>>
export type MyProviderLoginInfiniteQueryError = ErrorType<void>

export const useMyProviderLoginInfinite = <TData = Awaited<ReturnType<typeof myProviderLogin>>, TError = ErrorType<void>>(
 params: MyProviderLoginParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getMyProviderLoginInfiniteQueryOptions(params,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getMyProviderLoginQueryOptions = <TData = Awaited<ReturnType<typeof myProviderLogin>>, TError = ErrorType<void>>(params: MyProviderLoginParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getMyProviderLoginQueryKey(params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderLogin>>> = ({ signal }) => myProviderLogin(params, requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData> & { queryKey: QueryKey }
}

export type MyProviderLoginQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderLogin>>>
export type MyProviderLoginQueryError = ErrorType<void>

export const useMyProviderLogin = <TData = Awaited<ReturnType<typeof myProviderLogin>>, TError = ErrorType<void>>(
 params: MyProviderLoginParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getMyProviderLoginQueryOptions(params,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



