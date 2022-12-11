/* eslint-disable */
/* tslint:disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export type DbProjectPublic = {
  createdAt: string
  description: string
  initialized: boolean
  name: string
  projectID: number
  updatedAt: string
} & DbProjectPublic1
export type DbProjectPublic1 = {
  createdAt: string
  description: string
  initialized: boolean
  name: string
  projectID: number
  updatedAt: string
} | null
export type DbUserAPIKeyPublic = {
  apiKey: string
  expiresOn: string
  userID: UuidUUID
} & DbUserAPIKeyPublic1
export type UuidUUID = string
export type DbUserAPIKeyPublic1 = {
  apiKey: string
  expiresOn: string
  userID: UuidUUID
} | null
export type Role = 'guest' | 'user' | 'advancedUser' | 'manager' | 'admin' | 'superAdmin'
export type Scope =
  | 'test-scope'
  | 'users:read'
  | 'users:write'
  | 'scopes:write'
  | 'team-settings:write'
  | 'project-settings:write'
  | 'work-item:review'
export type Scopes = Scope[]
export type Location = string[]
export type Message = string
export type ErrorType = string
export type Detail = ValidationError[]
/**
 * string identifiers for SSE event listeners.
 */
export type Topics =
  | 'UserNotifications'
  | 'ManagerNotifications'
  | 'AdminNotifications'
  | 'WorkItemMoved'
  | 'WorkItemClosed'
/**
 * User notification type.
 */
export type NotificationType = 'personal' | 'global'
/**
 * Role in work item for a member.
 */
export type WorkItemRole = 'preparer' | 'reviewer'
export type ModelsRole = string

export interface DemoProjectWorkItemsResponse {
  baseWorkItem: DbWorkItemPublic
  lastMessageAt: string
  line: string
  ref: string
  reopened: boolean
  workItemID: number
}
export interface DbWorkItemPublic {
  closed: string | null
  createdAt: string
  deletedAt: string | null
  description: string
  kanbanStepID: number
  metadata: PgtypeJSONB
  targetDate: string
  teamID: number
  title: string
  updatedAt: string
  workItemID: number
  workItemTypeID: number
}
export interface PgtypeJSONB {}
export interface InitializeProjectRequest {
  activities?: ReposActivityCreateParams[] | null
  kanbanSteps?: ReposKanbanStepCreateParams[] | null
  projectID?: number
  teams?: ReposTeamCreateParams[] | null
  workItemTags?: ReposWorkItemTagCreateParams[] | null
  workItemTypes?: ReposWorkItemTypeCreateParams[] | null
}
export interface ReposActivityCreateParams {
  description?: string
  isProductive?: boolean
  name?: string
  projectID?: number
}
export interface ReposKanbanStepCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
  stepOrder?: number
  timeTrackable?: boolean
}
export interface ReposTeamCreateParams {
  description?: string
  name?: string
  projectID?: number
}
export interface ReposWorkItemTagCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
}
export interface ReposWorkItemTypeCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
}
export interface ProjectBoardResponse {
  activities?: DbActivityPublic[] | null
  kanbanSteps?: DbKanbanStepPublic[] | null
  project?: DbProjectPublic
  teams?: DbTeamPublic[] | null
  workItemTags?: DbWorkItemTagPublic[] | null
  workItemTypes?: DbWorkItemTypePublic[] | null
}
export interface DbActivityPublic {
  activityID: number
  description: string
  isProductive: boolean
  name: string
  projectID: number
}
export interface DbKanbanStepPublic {
  color: string
  description: string
  kanbanStepID: number
  name: string
  projectID: number
  stepOrder: number | null
  timeTrackable: boolean
}
export interface DbTeamPublic {
  createdAt: string
  description: string
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
export interface DbWorkItemTagPublic {
  color: string
  description: string
  name: string
  projectID: number
  workItemTagID: number
}
export interface DbWorkItemTypePublic {
  color: string
  description: string
  name: string
  projectID: number
  workItemTypeID: number
}
export interface UserResponse {
  apiKey?: DbUserAPIKeyPublic
  createdAt: string
  deletedAt: string | null
  email: string
  firstName: string | null
  fullName: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName: string | null
  role: Role
  scopes: Scopes
  teams?: DbTeamPublic[] | null
  userID: UuidUUID
  username: string
}
export interface HTTPValidationError {
  detail?: Detail
}
export interface ValidationError {
  loc: Location
  msg: Message
  type: ErrorType
}
/**
 * represents User data to update
 */
export interface UpdateUserRequest {
  /**
   * originally from auth server but updatable
   */
  first_name?: string
  /**
   * originally from auth server but updatable
   */
  last_name?: string
}
/**
 * represents User authorization data to update
 */
export interface UpdateUserAuthRequest {
  role?: Role
  scopes?: Scopes
}
