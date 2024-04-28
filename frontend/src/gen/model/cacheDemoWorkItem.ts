import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { CacheDemoWorkItemMetadata } from './cacheDemoWorkItemMetadata';

export interface CacheDemoWorkItem {
  closedAt?: Date | null;
  createdAt: Date;
  deletedAt?: Date | null;
  description: string;
  kanbanStepID: EntityIDs.KanbanStepID;
  lastMessageAt: Date;
  line: string;
  metadata: CacheDemoWorkItemMetadata;
  ref: string;
  reopened: boolean;
  targetDate: Date;
  teamID: EntityIDs.TeamID;
  title: string;
  updatedAt: Date;
  workItemID: EntityIDs.WorkItemID;
  workItemTypeID: EntityIDs.WorkItemTypeID;
}
