import type { Branded } from 'src/types/utils'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbNotification } from './dbNotification'
import type { DbUserID } from './dbUserID'

export interface Notification {
  notification: DbNotification
  notificationID: Branded<number, 'notificationID'>
  read: boolean
  userID: Branded<string, 'userID'>
  userNotificationID: Branded<number, 'userNotificationID'>
}
