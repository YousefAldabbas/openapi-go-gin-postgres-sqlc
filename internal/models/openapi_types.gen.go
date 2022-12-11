// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package models

import (
	"time"
)

const (
	Api_keyScopes     = "api_key.Scopes"
	Bearer_authScopes = "bearer_auth.Scopes"
)

// Defines values for NotificationType.
const (
	NotificationTypeGlobal   NotificationType = "global"
	NotificationTypePersonal NotificationType = "personal"
)

// AllNotificationTypeValues returns all possible values for NotificationType.
func AllNotificationTypeValues() []NotificationType {
	return []NotificationType{
		NotificationTypeGlobal,
		NotificationTypePersonal,
	}
}

// Defines values for Role.
const (
	RoleAdmin        Role = "admin"
	RoleAdvancedUser Role = "advancedUser"
	RoleGuest        Role = "guest"
	RoleManager      Role = "manager"
	RoleSuperAdmin   Role = "superAdmin"
	RoleUser         Role = "user"
)

// AllRoleValues returns all possible values for Role.
func AllRoleValues() []Role {
	return []Role{
		RoleAdmin,
		RoleAdvancedUser,
		RoleGuest,
		RoleManager,
		RoleSuperAdmin,
		RoleUser,
	}
}

// Defines values for Scope.
const (
	ScopeProjectSettingsWrite Scope = "project-settings:write"
	ScopeScopesWrite          Scope = "scopes:write"
	ScopeTeamSettingsWrite    Scope = "team-settings:write"
	ScopeTestScope            Scope = "test-scope"
	ScopeUsersRead            Scope = "users:read"
	ScopeUsersWrite           Scope = "users:write"
	ScopeWorkItemReview       Scope = "work-item:review"
)

// AllScopeValues returns all possible values for Scope.
func AllScopeValues() []Scope {
	return []Scope{
		ScopeProjectSettingsWrite,
		ScopeScopesWrite,
		ScopeTeamSettingsWrite,
		ScopeTestScope,
		ScopeUsersRead,
		ScopeUsersWrite,
		ScopeWorkItemReview,
	}
}

// Defines values for Topics.
const (
	TopicsAdminNotifications   Topics = "AdminNotifications"
	TopicsManagerNotifications Topics = "ManagerNotifications"
	TopicsUserNotifications    Topics = "UserNotifications"
	TopicsWorkItemClosed       Topics = "WorkItemClosed"
	TopicsWorkItemMoved        Topics = "WorkItemMoved"
)

// AllTopicsValues returns all possible values for Topics.
func AllTopicsValues() []Topics {
	return []Topics{
		TopicsAdminNotifications,
		TopicsManagerNotifications,
		TopicsUserNotifications,
		TopicsWorkItemClosed,
		TopicsWorkItemMoved,
	}
}

// Defines values for WorkItemRole.
const (
	WorkItemRolePreparer WorkItemRole = "preparer"
	WorkItemRoleReviewer WorkItemRole = "reviewer"
)

// AllWorkItemRoleValues returns all possible values for WorkItemRole.
func AllWorkItemRoleValues() []WorkItemRole {
	return []WorkItemRole{
		WorkItemRolePreparer,
		WorkItemRoleReviewer,
	}
}

// DbActivityPublic defines the model for DbActivityPublic.
type DbActivityPublic struct {
	ActivityID   int    `json:"activityID"`
	Description  string `json:"description"`
	IsProductive bool   `json:"isProductive"`
	Name         string `json:"name"`
	ProjectID    int    `json:"projectID"`
}

// DbKanbanStepPublic defines the model for DbKanbanStepPublic.
type DbKanbanStepPublic struct {
	Color         string `json:"color"`
	Description   string `json:"description"`
	KanbanStepID  int    `json:"kanbanStepID"`
	Name          string `json:"name"`
	ProjectID     int    `json:"projectID"`
	StepOrder     *int   `json:"stepOrder"`
	TimeTrackable bool   `json:"timeTrackable"`
}

// DbProjectPublic defines the model for DbProjectPublic.
type DbProjectPublic struct {
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	Initialized bool      `json:"initialized"`
	Name        string    `json:"name"`
	ProjectID   int       `json:"projectID"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// DbTeamPublic defines the model for DbTeamPublic.
type DbTeamPublic struct {
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	ProjectID   int       `json:"projectID"`
	TeamID      int       `json:"teamID"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// DbUserAPIKeyPublic defines the model for DbUserAPIKeyPublic.
type DbUserAPIKeyPublic struct {
	ApiKey    string    `json:"apiKey"`
	ExpiresOn time.Time `json:"expiresOn"`
	UserID    UuidUUID  `json:"userID"`
}

// DbWorkItemPublic defines the model for DbWorkItemPublic.
type DbWorkItemPublic struct {
	Closed         *time.Time  `json:"closed"`
	CreatedAt      time.Time   `json:"createdAt"`
	DeletedAt      *time.Time  `json:"deletedAt"`
	Description    string      `json:"description"`
	KanbanStepID   int         `json:"kanbanStepID"`
	Metadata       PgtypeJSONB `json:"metadata"`
	TargetDate     time.Time   `json:"targetDate"`
	TeamID         int         `json:"teamID"`
	Title          string      `json:"title"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	WorkItemID     int         `json:"workItemID"`
	WorkItemTypeID int         `json:"workItemTypeID"`
}

// DbWorkItemTagPublic defines the model for DbWorkItemTagPublic.
type DbWorkItemTagPublic struct {
	Color         string `json:"color"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	ProjectID     int    `json:"projectID"`
	WorkItemTagID int    `json:"workItemTagID"`
}

// DbWorkItemTypePublic defines the model for DbWorkItemTypePublic.
type DbWorkItemTypePublic struct {
	Color          string `json:"color"`
	Description    string `json:"description"`
	Name           string `json:"name"`
	ProjectID      int    `json:"projectID"`
	WorkItemTypeID int    `json:"workItemTypeID"`
}

// DemoProjectWorkItemsResponse defines the model for DemoProjectWorkItemsResponse.
type DemoProjectWorkItemsResponse struct {
	BaseWorkItem  DbWorkItemPublic `json:"baseWorkItem"`
	LastMessageAt time.Time        `json:"lastMessageAt"`
	Line          string           `json:"line"`
	Ref           string           `json:"ref"`
	Reopened      bool             `json:"reopened"`
	WorkItemID    int              `json:"workItemID"`
}

// HTTPValidationError defines the model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// InitializeProjectRequest defines the model for InitializeProjectRequest.
type InitializeProjectRequest struct {
	Activities    *[]ReposActivityCreateParams     `json:"activities"`
	KanbanSteps   *[]ReposKanbanStepCreateParams   `json:"kanbanSteps"`
	ProjectID     *int                             `json:"projectID,omitempty"`
	Teams         *[]ReposTeamCreateParams         `json:"teams"`
	WorkItemTags  *[]ReposWorkItemTagCreateParams  `json:"workItemTags"`
	WorkItemTypes *[]ReposWorkItemTypeCreateParams `json:"workItemTypes"`
}

// ModelsRole defines the model for ModelsRole.
type ModelsRole = string

// NotificationType User notification type.
type NotificationType string

// PgtypeJSONB defines the model for PgtypeJSONB.
type PgtypeJSONB = map[string]interface{}

// ProjectBoardResponse defines the model for ProjectBoardResponse.
type ProjectBoardResponse struct {
	Activities    *[]DbActivityPublic     `json:"activities"`
	KanbanSteps   *[]DbKanbanStepPublic   `json:"kanbanSteps"`
	Project       *DbProjectPublic        `json:"project"`
	Teams         *[]DbTeamPublic         `json:"teams"`
	WorkItemTags  *[]DbWorkItemTagPublic  `json:"workItemTags"`
	WorkItemTypes *[]DbWorkItemTypePublic `json:"workItemTypes"`
}

// ReposActivityCreateParams defines the model for ReposActivityCreateParams.
type ReposActivityCreateParams struct {
	Description  *string `json:"description,omitempty"`
	IsProductive *bool   `json:"isProductive,omitempty"`
	Name         *string `json:"name,omitempty"`
	ProjectID    *int    `json:"projectID,omitempty"`
}

// ReposKanbanStepCreateParams defines the model for ReposKanbanStepCreateParams.
type ReposKanbanStepCreateParams struct {
	Color         *string `json:"color,omitempty"`
	Description   *string `json:"description,omitempty"`
	Name          *string `json:"name,omitempty"`
	ProjectID     *int    `json:"projectID,omitempty"`
	StepOrder     *int    `json:"stepOrder,omitempty"`
	TimeTrackable *bool   `json:"timeTrackable,omitempty"`
}

// ReposTeamCreateParams defines the model for ReposTeamCreateParams.
type ReposTeamCreateParams struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProjectID   *int    `json:"projectID,omitempty"`
}

// ReposWorkItemTagCreateParams defines the model for ReposWorkItemTagCreateParams.
type ReposWorkItemTagCreateParams struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProjectID   *int    `json:"projectID,omitempty"`
}

// ReposWorkItemTypeCreateParams defines the model for ReposWorkItemTypeCreateParams.
type ReposWorkItemTypeCreateParams struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProjectID   *int    `json:"projectID,omitempty"`
}

// Role defines the model for Role.
type Role string

// Scope defines the model for Scope.
type Scope string

// Scopes defines the model for Scopes.
type Scopes = []Scope

// Topics string identifiers for SSE event listeners.
type Topics string

// UpdateUserAuthRequest represents User authorization data to update
type UpdateUserAuthRequest struct {
	Role   *Role   `json:"role,omitempty"`
	Scopes *Scopes `json:"scopes,omitempty"`
}

// UpdateUserRequest represents User data to update
type UpdateUserRequest struct {
	// FirstName originally from auth server but updatable
	FirstName *string `json:"first_name,omitempty"`

	// LastName originally from auth server but updatable
	LastName *string `json:"last_name,omitempty"`
}

// UserResponse defines the model for UserResponse.
type UserResponse struct {
	ApiKey                   *DbUserAPIKeyPublic `json:"apiKey"`
	CreatedAt                time.Time           `json:"createdAt"`
	DeletedAt                *time.Time          `json:"deletedAt"`
	Email                    string              `json:"email"`
	FirstName                *string             `json:"firstName"`
	FullName                 *string             `json:"fullName"`
	HasGlobalNotifications   bool                `json:"hasGlobalNotifications"`
	HasPersonalNotifications bool                `json:"hasPersonalNotifications"`
	LastName                 *string             `json:"lastName"`
	Role                     Role                `json:"role"`
	Scopes                   Scopes              `json:"scopes"`
	Teams                    *[]DbTeamPublic     `json:"teams"`
	UserID                   UuidUUID            `json:"userID"`
	Username                 string              `json:"username"`
}

// UuidUUID defines the model for UuidUUID.
type UuidUUID = string

// ValidationError defines the model for ValidationError.
type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// WorkItemRole Role in work item for a member.
type WorkItemRole string

// PathSerial defines the model for PathSerial.
type PathSerial = int

// Uuid defines the model for uuid.
type Uuid = string

// GetProjectWorkitemsParams defines parameters for GetProjectWorkitems.
type GetProjectWorkitemsParams struct {
	Open *bool `form:"open,omitempty" json:"open,omitempty"`
}

// InitializeProjectJSONRequestBody defines body for InitializeProject for application/json ContentType.
type InitializeProjectJSONRequestBody = InitializeProjectRequest

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UpdateUserRequest

// UpdateUserAuthorizationJSONRequestBody defines body for UpdateUserAuthorization for application/json ContentType.
type UpdateUserAuthorizationJSONRequestBody = UpdateUserAuthRequest
