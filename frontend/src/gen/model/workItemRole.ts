/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

/**
 * is generated from database enum 'work_item_role'.
 */
export type WorkItemRole = typeof WorkItemRole[keyof typeof WorkItemRole]

// eslint-disable-next-line @typescript-eslint/no-redeclare
export const WorkItemRole = {
  preparer: 'preparer',
  reviewer: 'reviewer',
} as const
