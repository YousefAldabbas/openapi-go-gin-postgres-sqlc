import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbDemoTwoWorkItem } from './dbDemoTwoWorkItem';
import type { DbUserWIAUWorkItem } from './dbUserWIAUWorkItem';
import type { DemoTwoWorkItemsMetadata } from './demoTwoWorkItemsMetadata';
import type { DbTimeEntry } from './dbTimeEntry';
import type { DbWorkItemComment } from './dbWorkItemComment';
import type { DbWorkItemTag } from './dbWorkItemTag';
import type { DbWorkItemType } from './dbWorkItemType';

export interface DemoTwoWorkItems {
  closedAt?: Date | null;
  createdAt: Date;
  deletedAt?: Date | null;
  demoTwoWorkItem: DbDemoTwoWorkItem;
  description: string;
  kanbanStepID: EntityIDs.KanbanStepID;
  members?: DbUserWIAUWorkItem[] | null;
  metadata: DemoTwoWorkItemsMetadata;
  targetDate: Date;
  teamID: EntityIDs.TeamID | null;
  timeEntries?: DbTimeEntry[] | null;
  title: string;
  updatedAt: Date;
  workItemComments?: DbWorkItemComment[] | null;
  workItemID: EntityIDs.WorkItemID;
  workItemTags?: DbWorkItemTag[] | null;
  workItemType?: DbWorkItemType;
  workItemTypeID: EntityIDs.WorkItemTypeID;
}
