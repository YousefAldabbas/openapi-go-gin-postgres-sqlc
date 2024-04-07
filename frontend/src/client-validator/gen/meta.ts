/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import {
  Activity,
  CreateActivityRequest,
  CreateDemoTwoWorkItemRequest,
  CreateDemoWorkItemRequest,
  CreateProjectBoardRequest,
  CreateTeamRequest,
  CreateWorkItemCommentRequest,
  CreateWorkItemTagRequest,
  CreateWorkItemTypeRequest,
  DbDemoTwoWorkItem,
  DbDemoTwoWorkItemCreateParams,
  DbDemoWorkItem,
  DbDemoWorkItemCreateParams,
  DbKanbanStep,
  DbNotification,
  DbProject,
  DbTeam,
  DbTeamCreateParams,
  DbTimeEntry,
  DbUser,
  DbUserAPIKey,
  DbUserID,
  DbUserWIAUWorkItem,
  DbWorkItem,
  DbWorkItemComment,
  DbWorkItemCreateParams,
  DbWorkItemTag,
  DbWorkItemTagCreateParams,
  DbWorkItemType,
  Notification,
  PaginatedNotificationsResponse,
  PaginatedUsersResponse,
  PaginationPage,
  ProjectBoard,
  ServicesMember,
  SharedWorkItemJoins,
  Team,
  UpdateActivityRequest,
  UpdateTeamRequest,
  UpdateWorkItemCommentRequest,
  UpdateWorkItemTagRequest,
  UpdateWorkItemTypeRequest,
  User,
  WorkItemComment,
  WorkItemTag,
  WorkItemType,
  Direction,
  DbActivity,
  ProjectConfig,
  ProjectConfigField,
  HTTPValidationError,
  ErrorCode,
  HTTPError,
  Topics,
  Topic,
  Scope,
  Scopes,
  Role,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  ValidationError,
  UuidUUID,
  WorkItem,
  CreateWorkItemRequest,
  Project,
  DbActivityCreateParams,
  DbWorkItemRole,
  NotificationType,
  DemoTwoWorkItemTypes,
  DemoWorkItemTypes,
  DbWorkItemID,
  DbProjectID,
  DbWorkItemTypeID,
  DbNotificationID,
  DbUserNotification,
  DemoKanbanSteps,
  DemoTwoKanbanSteps,
  DbUserWIAWorkItem,
  DbWorkItemM2MAssigneeWIA,
  CreateTimeEntryRequest,
  TimeEntry,
  UpdateTimeEntryRequest,
  DemoTwoWorkItem,
  DemoWorkItem,
  WorkItemBase,
  PaginationFilterPrimitive,
  PaginationFilterArray,
  PaginationFilter,
  Pagination,
  PaginationItems,
  PaginationCursors,
  PaginationCursor,
  GetPaginatedUsersQueryParameters,
  PaginationFilterModes,
  DbCacheDemoWorkItemJoins,
  DbUserJoins,
  GetCacheDemoWorkItemQueryParameters,
  GetCurrentUserQueryParameters,
} from './models'

export const schemaDefinitions = {
  Activity: info<Activity>('Activity', '#/definitions/Activity'),
  CreateActivityRequest: info<CreateActivityRequest>('CreateActivityRequest', '#/definitions/CreateActivityRequest'),
  CreateDemoTwoWorkItemRequest: info<CreateDemoTwoWorkItemRequest>(
    'CreateDemoTwoWorkItemRequest',
    '#/definitions/CreateDemoTwoWorkItemRequest',
  ),
  CreateDemoWorkItemRequest: info<CreateDemoWorkItemRequest>(
    'CreateDemoWorkItemRequest',
    '#/definitions/CreateDemoWorkItemRequest',
  ),
  CreateProjectBoardRequest: info<CreateProjectBoardRequest>(
    'CreateProjectBoardRequest',
    '#/definitions/CreateProjectBoardRequest',
  ),
  CreateTeamRequest: info<CreateTeamRequest>('CreateTeamRequest', '#/definitions/CreateTeamRequest'),
  CreateWorkItemCommentRequest: info<CreateWorkItemCommentRequest>(
    'CreateWorkItemCommentRequest',
    '#/definitions/CreateWorkItemCommentRequest',
  ),
  CreateWorkItemTagRequest: info<CreateWorkItemTagRequest>(
    'CreateWorkItemTagRequest',
    '#/definitions/CreateWorkItemTagRequest',
  ),
  CreateWorkItemTypeRequest: info<CreateWorkItemTypeRequest>(
    'CreateWorkItemTypeRequest',
    '#/definitions/CreateWorkItemTypeRequest',
  ),
  DbDemoTwoWorkItem: info<DbDemoTwoWorkItem>('DbDemoTwoWorkItem', '#/definitions/DbDemoTwoWorkItem'),
  DbDemoTwoWorkItemCreateParams: info<DbDemoTwoWorkItemCreateParams>(
    'DbDemoTwoWorkItemCreateParams',
    '#/definitions/DbDemoTwoWorkItemCreateParams',
  ),
  DbDemoWorkItem: info<DbDemoWorkItem>('DbDemoWorkItem', '#/definitions/DbDemoWorkItem'),
  DbDemoWorkItemCreateParams: info<DbDemoWorkItemCreateParams>(
    'DbDemoWorkItemCreateParams',
    '#/definitions/DbDemoWorkItemCreateParams',
  ),
  DbKanbanStep: info<DbKanbanStep>('DbKanbanStep', '#/definitions/DbKanbanStep'),
  DbNotification: info<DbNotification>('DbNotification', '#/definitions/DbNotification'),
  DbProject: info<DbProject>('DbProject', '#/definitions/DbProject'),
  DbTeam: info<DbTeam>('DbTeam', '#/definitions/DbTeam'),
  DbTeamCreateParams: info<DbTeamCreateParams>('DbTeamCreateParams', '#/definitions/DbTeamCreateParams'),
  DbTimeEntry: info<DbTimeEntry>('DbTimeEntry', '#/definitions/DbTimeEntry'),
  DbUser: info<DbUser>('DbUser', '#/definitions/DbUser'),
  DbUserAPIKey: info<DbUserAPIKey>('DbUserAPIKey', '#/definitions/DbUserAPIKey'),
  DbUserID: info<DbUserID>('DbUserID', '#/definitions/DbUserID'),
  DbUserWIAUWorkItem: info<DbUserWIAUWorkItem>('DbUserWIAUWorkItem', '#/definitions/DbUserWIAUWorkItem'),
  DbWorkItem: info<DbWorkItem>('DbWorkItem', '#/definitions/DbWorkItem'),
  DbWorkItemComment: info<DbWorkItemComment>('DbWorkItemComment', '#/definitions/DbWorkItemComment'),
  DbWorkItemCreateParams: info<DbWorkItemCreateParams>(
    'DbWorkItemCreateParams',
    '#/definitions/DbWorkItemCreateParams',
  ),
  DbWorkItemTag: info<DbWorkItemTag>('DbWorkItemTag', '#/definitions/DbWorkItemTag'),
  DbWorkItemTagCreateParams: info<DbWorkItemTagCreateParams>(
    'DbWorkItemTagCreateParams',
    '#/definitions/DbWorkItemTagCreateParams',
  ),
  DbWorkItemType: info<DbWorkItemType>('DbWorkItemType', '#/definitions/DbWorkItemType'),
  Notification: info<Notification>('Notification', '#/definitions/Notification'),
  PaginatedNotificationsResponse: info<PaginatedNotificationsResponse>(
    'PaginatedNotificationsResponse',
    '#/definitions/PaginatedNotificationsResponse',
  ),
  PaginatedUsersResponse: info<PaginatedUsersResponse>(
    'PaginatedUsersResponse',
    '#/definitions/PaginatedUsersResponse',
  ),
  PaginationPage: info<PaginationPage>('PaginationPage', '#/definitions/PaginationPage'),
  ProjectBoard: info<ProjectBoard>('ProjectBoard', '#/definitions/ProjectBoard'),
  ServicesMember: info<ServicesMember>('ServicesMember', '#/definitions/ServicesMember'),
  SharedWorkItemJoins: info<SharedWorkItemJoins>('SharedWorkItemJoins', '#/definitions/SharedWorkItemJoins'),
  Team: info<Team>('Team', '#/definitions/Team'),
  UpdateActivityRequest: info<UpdateActivityRequest>('UpdateActivityRequest', '#/definitions/UpdateActivityRequest'),
  UpdateTeamRequest: info<UpdateTeamRequest>('UpdateTeamRequest', '#/definitions/UpdateTeamRequest'),
  UpdateWorkItemCommentRequest: info<UpdateWorkItemCommentRequest>(
    'UpdateWorkItemCommentRequest',
    '#/definitions/UpdateWorkItemCommentRequest',
  ),
  UpdateWorkItemTagRequest: info<UpdateWorkItemTagRequest>(
    'UpdateWorkItemTagRequest',
    '#/definitions/UpdateWorkItemTagRequest',
  ),
  UpdateWorkItemTypeRequest: info<UpdateWorkItemTypeRequest>(
    'UpdateWorkItemTypeRequest',
    '#/definitions/UpdateWorkItemTypeRequest',
  ),
  User: info<User>('User', '#/definitions/User'),
  WorkItemComment: info<WorkItemComment>('WorkItemComment', '#/definitions/WorkItemComment'),
  WorkItemTag: info<WorkItemTag>('WorkItemTag', '#/definitions/WorkItemTag'),
  WorkItemType: info<WorkItemType>('WorkItemType', '#/definitions/WorkItemType'),
  Direction: info<Direction>('Direction', '#/definitions/Direction'),
  DbActivity: info<DbActivity>('DbActivity', '#/definitions/DbActivity'),
  ProjectConfig: info<ProjectConfig>('ProjectConfig', '#/definitions/ProjectConfig'),
  ProjectConfigField: info<ProjectConfigField>('ProjectConfigField', '#/definitions/ProjectConfigField'),
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  ErrorCode: info<ErrorCode>('ErrorCode', '#/definitions/ErrorCode'),
  HTTPError: info<HTTPError>('HTTPError', '#/definitions/HTTPError'),
  Topics: info<Topics>('Topics', '#/definitions/Topics'),
  Topic: info<Topic>('Topic', '#/definitions/Topic'),
  Scope: info<Scope>('Scope', '#/definitions/Scope'),
  Scopes: info<Scopes>('Scopes', '#/definitions/Scopes'),
  Role: info<Role>('Role', '#/definitions/Role'),
  WorkItemRole: info<WorkItemRole>('WorkItemRole', '#/definitions/WorkItemRole'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  UpdateUserAuthRequest: info<UpdateUserAuthRequest>('UpdateUserAuthRequest', '#/definitions/UpdateUserAuthRequest'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  WorkItem: info<WorkItem>('WorkItem', '#/definitions/WorkItem'),
  CreateWorkItemRequest: info<CreateWorkItemRequest>('CreateWorkItemRequest', '#/definitions/CreateWorkItemRequest'),
  Project: info<Project>('Project', '#/definitions/Project'),
  DbActivityCreateParams: info<DbActivityCreateParams>(
    'DbActivityCreateParams',
    '#/definitions/DbActivityCreateParams',
  ),
  DbWorkItemRole: info<DbWorkItemRole>('DbWorkItemRole', '#/definitions/DbWorkItemRole'),
  NotificationType: info<NotificationType>('NotificationType', '#/definitions/NotificationType'),
  DemoTwoWorkItemTypes: info<DemoTwoWorkItemTypes>('DemoTwoWorkItemTypes', '#/definitions/DemoTwoWorkItemTypes'),
  DemoWorkItemTypes: info<DemoWorkItemTypes>('DemoWorkItemTypes', '#/definitions/DemoWorkItemTypes'),
  DbWorkItemID: info<DbWorkItemID>('DbWorkItemID', '#/definitions/DbWorkItemID'),
  DbProjectID: info<DbProjectID>('DbProjectID', '#/definitions/DbProjectID'),
  DbWorkItemTypeID: info<DbWorkItemTypeID>('DbWorkItemTypeID', '#/definitions/DbWorkItemTypeID'),
  DbNotificationID: info<DbNotificationID>('DbNotificationID', '#/definitions/DbNotificationID'),
  DbUserNotification: info<DbUserNotification>('DbUserNotification', '#/definitions/DbUserNotification'),
  DemoKanbanSteps: info<DemoKanbanSteps>('DemoKanbanSteps', '#/definitions/DemoKanbanSteps'),
  DemoTwoKanbanSteps: info<DemoTwoKanbanSteps>('DemoTwoKanbanSteps', '#/definitions/DemoTwoKanbanSteps'),
  DbUserWIAWorkItem: info<DbUserWIAWorkItem>('DbUserWIAWorkItem', '#/definitions/DbUserWIAWorkItem'),
  DbWorkItemM2MAssigneeWIA: info<DbWorkItemM2MAssigneeWIA>(
    'DbWorkItemM2MAssigneeWIA',
    '#/definitions/DbWorkItemM2MAssigneeWIA',
  ),
  CreateTimeEntryRequest: info<CreateTimeEntryRequest>(
    'CreateTimeEntryRequest',
    '#/definitions/CreateTimeEntryRequest',
  ),
  TimeEntry: info<TimeEntry>('TimeEntry', '#/definitions/TimeEntry'),
  UpdateTimeEntryRequest: info<UpdateTimeEntryRequest>(
    'UpdateTimeEntryRequest',
    '#/definitions/UpdateTimeEntryRequest',
  ),
  DemoTwoWorkItem: info<DemoTwoWorkItem>('DemoTwoWorkItem', '#/definitions/DemoTwoWorkItem'),
  DemoWorkItem: info<DemoWorkItem>('DemoWorkItem', '#/definitions/DemoWorkItem'),
  WorkItemBase: info<WorkItemBase>('WorkItemBase', '#/definitions/WorkItemBase'),
  PaginationFilterPrimitive: info<PaginationFilterPrimitive>(
    'PaginationFilterPrimitive',
    '#/definitions/PaginationFilterPrimitive',
  ),
  PaginationFilterArray: info<PaginationFilterArray>('PaginationFilterArray', '#/definitions/PaginationFilterArray'),
  PaginationFilter: info<PaginationFilter>('PaginationFilter', '#/definitions/PaginationFilter'),
  Pagination: info<Pagination>('Pagination', '#/definitions/Pagination'),
  PaginationItems: info<PaginationItems>('PaginationItems', '#/definitions/PaginationItems'),
  PaginationCursors: info<PaginationCursors>('PaginationCursors', '#/definitions/PaginationCursors'),
  PaginationCursor: info<PaginationCursor>('PaginationCursor', '#/definitions/PaginationCursor'),
  GetPaginatedUsersQueryParameters: info<GetPaginatedUsersQueryParameters>(
    'GetPaginatedUsersQueryParameters',
    '#/definitions/GetPaginatedUsersQueryParameters',
  ),
  PaginationFilterModes: info<PaginationFilterModes>('PaginationFilterModes', '#/definitions/PaginationFilterModes'),
  DbCacheDemoWorkItemJoins: info<DbCacheDemoWorkItemJoins>(
    'DbCacheDemoWorkItemJoins',
    '#/definitions/DbCacheDemoWorkItemJoins',
  ),
  DbUserJoins: info<DbUserJoins>('DbUserJoins', '#/definitions/DbUserJoins'),
  GetCacheDemoWorkItemQueryParameters: info<GetCacheDemoWorkItemQueryParameters>(
    'GetCacheDemoWorkItemQueryParameters',
    '#/definitions/GetCacheDemoWorkItemQueryParameters',
  ),
  GetCurrentUserQueryParameters: info<GetCurrentUserQueryParameters>(
    'GetCurrentUserQueryParameters',
    '#/definitions/GetCurrentUserQueryParameters',
  ),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
