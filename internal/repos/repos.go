package repos

import (
	"context"

	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
)

// Boards limited to one per project. All teams in a project share the same board.
// IMPORTANT: If a board does not exist for a project, it will not be possible
// to create teams, activities, etc. The first time we must go through the project
// creation steps and create everything at once.
// Later on everything can be updated in the project settings panel, and new elements created.
// Once a board is created it cannot be deleted.
type ProjectBoardCreateParams struct {
	ProjectID int `json:"projectID"`
	// TeamIDs   []int `json:"teamIDs"` // completely useless. the only check needed is to ensure at least one team
	// exacts associated to projectID, else prompt the user to create at least 1 team before creating a board.
	Activities    []db.ActivityCreateParams     `json:"activities"`
	KanbanSteps   []db.KanbanStepCreateParams   `json:"kanbanSteps"`
	Teams         []db.TeamCreateParams         `json:"teams"`
	WorkItemTypes []db.WorkItemTypeCreateParams `json:"workItemTypes"`
	WorkItemTags  []db.WorkItemTagCreateParams  `json:"workItemTags"`
}

// ProjectBoard defines the datastore/repository handling persisting ProjectBoard records.
type ProjectBoard interface {
	/*
		Create corresponds to initial info to be filled in once a project table has been manually
		 created, before it can be used:
		 - kanban columns and their info (order, name, can log time, etc.)
		 - types of workitems (by all teams)
		 - initial teams associated (at least 1 id initially)

		  If we manually added a new project record
		 in migrations (insert into projects) and created its specific work_items_<new project> table
			at this point everything is empty and project.is_setup is False

			!project.is_setup --> dashboard shows single centered [+ Initialize project] button
			and we let the project admins create teams, work item types, etc. all at once.
			if form submitted successfully, is_setup=False and won't show again. Else, the whole thing
			is rolled back at once (should try to save this form state with zustand's persist just in case but not important)


	*/
	Create(ctx context.Context, d db.DBTX, params ProjectBoardCreateParams) (*models.ProjectBoard, error)
	ByID(ctx context.Context, d db.DBTX, projectID int) (*models.ProjectBoard, error)
}

type DemoProjectWorkItemUpdateParams struct {
	DemoProject *db.DemoProjectWorkItemUpdateParams
	Base        *db.WorkItemUpdateParams
}

type DemoProjectWorkItemCreateParams struct {
	DemoProject db.DemoProjectWorkItemCreateParams
	Base        db.WorkItemCreateParams
}

// DemoProjectWorkItem defines the datastore/repository handling persisting DemoProjectWorkItem records.
type DemoProjectWorkItem interface {
	ByID(ctx context.Context, d db.DBTX, id int64, opts ...db.DemoProjectWorkItemSelectConfigOption) (*db.DemoProjectWorkItem, error)
	// params for dedicated workItem require workItemID (FK-as-PK)
	Create(ctx context.Context, d db.DBTX, params DemoProjectWorkItemCreateParams) (*db.DemoProjectWorkItem, error)
	Delete(ctx context.Context, d db.DBTX, workItemID int64) (*db.DemoProjectWorkItem, error)
	// repo has Update only, then service has Close() (Update with closed=True), Move() (Update with kanban step change), ...)
	// params for dedicated workItem require workItemID (FK-as-PK)
	Update(ctx context.Context, d db.DBTX, params DemoProjectWorkItemUpdateParams) (*db.DemoProjectWorkItem, error)
	// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)
}

// WorkItem defines the datastore/repository handling persisting WorkItem records.
/**
 * TODO:
 * instead pass `, project models.Project` and do appropiate joins in workitem.xo.go depending on it.
 * in case we need specific indexes from {project}workitem.xo.go we can do e.g. DemoProjectWorkItemsByRefLine with no workitem join
 * and then just loop over those and filter by ID taking the performance hit. it will be rare to use those filters anyway.
 * If not we can always go back to sqlc or use jet + reuse the xo model
 *
 */
type WorkItem interface {
	ByID(ctx context.Context, d db.DBTX, id int64, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error)
	ByTeam(ctx context.Context, d db.DBTX, teamID int, closed bool, deleted bool, opts ...db.WorkItemSelectConfigOption) ([]*db.WorkItem, error)
	// params for dedicated workItem require workItemID (FK-as-PK)
	Create(ctx context.Context, d db.DBTX, params db.WorkItemCreateParams) (*db.WorkItem, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error)
}

// Notification defines the datastore/repository handling persisting Notification records.
type Notification interface {
	LatestUserNotifications(ctx context.Context, d db.DBTX, params db.GetUserNotificationsParams) ([]db.GetUserNotificationsRow, error)
	Create(ctx context.Context, d db.DBTX, params db.NotificationCreateParams) (*db.Notification, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Notification, error)
}

// User defines the datastore/repository handling persisting User records.
type User interface {
	ByID(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error)
	ByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error)
	ByUsername(ctx context.Context, d db.DBTX, username string) (*db.User, error)
	ByExternalID(ctx context.Context, d db.DBTX, extID string) (*db.User, error)
	ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error)
	Create(ctx context.Context, d db.DBTX, params db.UserCreateParams) (*db.User, error)
	Update(ctx context.Context, d db.DBTX, id uuid.UUID, params db.UserUpdateParams) (*db.User, error)
	Delete(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error)
	// CreateAPIKey requires an existing user.
	CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error)
}

// Project defines the datastore/repository handling persisting Project records.
// Projects are manually created on demand.
type Project interface {
	ByName(ctx context.Context, d db.DBTX, name internalmodels.Project) (*db.Project, error)
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Project, error)
}

// Team defines the datastore/repository handling persisting Team records.
type Team interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Team, error)
	Create(ctx context.Context, d db.DBTX, params db.TeamCreateParams) (*db.Team, error)
	Update(ctx context.Context, d db.DBTX, id int, params db.TeamUpdateParams) (*db.Team, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Team, error)
}

// KanbanStep defines the datastore/repository handling persisting KanbanStep records.
type KanbanStep interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error)
	ByProject(ctx context.Context, d db.DBTX, projectID int) ([]db.KanbanStep, error)
	Create(ctx context.Context, d db.DBTX, params db.KanbanStepCreateParams) (*db.KanbanStep, error)
	Update(ctx context.Context, d db.DBTX, id int, params db.KanbanStepUpdateParams) (*db.KanbanStep, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error)
}

// WorkItemType defines the datastore/repository handling persisting WorkItemType records.
type WorkItemType interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.WorkItemType, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemType, error)
	Create(ctx context.Context, d db.DBTX, params db.WorkItemTypeCreateParams) (*db.WorkItemType, error)
	Update(ctx context.Context, d db.DBTX, id int, params db.WorkItemTypeUpdateParams) (*db.WorkItemType, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error)
}

// WorkItemComment defines the datastore/repository handling persisting WorkItemComment records.
type WorkItemComment interface {
	ByID(ctx context.Context, d db.DBTX, id int64) (*db.WorkItemComment, error)
	Create(ctx context.Context, d db.DBTX, params db.WorkItemCommentCreateParams) (*db.WorkItemComment, error)
	Update(ctx context.Context, d db.DBTX, id int64, params db.WorkItemCommentUpdateParams) (*db.WorkItemComment, error)
	Delete(ctx context.Context, d db.DBTX, id int64) (*db.WorkItemComment, error)
}

// WorkItemTag defines the datastore/repository handling persisting WorkItemTag records.
type WorkItemTag interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.WorkItemTag, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemTag, error)
	Create(ctx context.Context, d db.DBTX, params db.WorkItemTagCreateParams) (*db.WorkItemTag, error)
	Update(ctx context.Context, d db.DBTX, id int, params db.WorkItemTagUpdateParams) (*db.WorkItemTag, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error)
}

// Activity defines the datastore/repository handling persisting Activity records.
type Activity interface {
	ByID(ctx context.Context, d db.DBTX, id int) (*db.Activity, error)
	// TODO ByProjectID(ctx context.Context, d db.DBTX, id int) ([]*db.Activity, error)
	ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Activity, error)
	Create(ctx context.Context, d db.DBTX, params db.ActivityCreateParams) (*db.Activity, error)
	Update(ctx context.Context, d db.DBTX, id int, params db.ActivityUpdateParams) (*db.Activity, error)
	Delete(ctx context.Context, d db.DBTX, id int) (*db.Activity, error)
}
