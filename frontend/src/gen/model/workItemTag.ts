import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

export interface WorkItemTag {
  color: string;
  deletedAt?: Date | null;
  description: string;
  name: string;
  projectID: EntityIDs.ProjectID;
  workItemTagID: EntityIDs.WorkItemTagID;
}
