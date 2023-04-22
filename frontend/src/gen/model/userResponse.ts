/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbUserAPIKey } from './dbUserAPIKey'
import type { DbProject } from './dbProject'
import type { Role } from './role'
import type { Scopes } from './scopes'
import type { DbTeam } from './dbTeam'
import type { DbUser } from './dbUser'

export interface UserResponse {
  apiKey?: DbUserAPIKey
  projects?: DbProject[] | null
  role: Role
  scopes: Scopes
  teams?: DbTeam[] | null
  user: DbUser
}
