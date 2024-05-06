import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsUserID } from './modelsUserID';

export interface UpdateTimeEntryRequest {
  activityID?: EntityIDs.ActivityID;
  comment?: string;
  durationMinutes?: number | null;
  start?: Date;
  teamID?: EntityIDs.TeamID | null;
  userID?: EntityIDs.UserID;
  workItemID?: EntityIDs.WorkItemID | null;
}
