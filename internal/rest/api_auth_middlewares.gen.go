// Code generated by pregen. DO NOT EDIT.

package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) authMiddlewares(opID OperationID) []gin.HandlerFunc {
	switch opID {
	case AdminPing:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
				}),
		}
	case CreateActivity:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("manager"),
					RequiredScopes: models.Scopes{
						models.Scope("activity:create"),
					},
				}),
		}
	case CreateTeam:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case CreateTimeEntry:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case CreateWorkItemComment:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					RequiredScopes: models.Scopes{
						models.Scope("work-item-comment:create"),
					},
				}),
		}
	case CreateWorkItemTag:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("manager"),
					RequiredScopes: models.Scopes{
						models.Scope("work-item-tag:create"),
					},
				}),
		}
	case CreateWorkItemType:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case CreateWorkitem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case DeleteActivity:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("manager"),
					RequiredScopes: models.Scopes{
						models.Scope("activity:delete"),
					},
				}),
		}
	case DeleteTeam:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case DeleteTimeEntry:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case DeleteUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
					RequiredScopes: models.Scopes{
						models.Scope("users:delete"),
					},
				}),
		}
	case DeleteWorkItemComment:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					RequiredScopes: models.Scopes{
						models.Scope("work-item-comment:delete"),
					},
				}),
		}
	case DeleteWorkItemTag:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case DeleteWorkItemType:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case DeleteWorkitem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case Events:
		return []gin.HandlerFunc{}
	case GetActivity:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetCurrentUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetPaginatedNotifications:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetPaginatedUsers:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetPaginatedWorkItem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProject:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProjectBoard:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProjectConfig:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
				}),
		}
	case GetProjectWorkitems:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetTeam:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetTimeEntry:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetWorkItem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetWorkItemComment:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetWorkItemTag:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetWorkItemType:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case InitializeProject:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
				}),
		}
	case MyProviderCallback:
		return []gin.HandlerFunc{}
	case MyProviderLogin:
		return []gin.HandlerFunc{}
	case OpenapiYamlGet:
		return []gin.HandlerFunc{}
	case Ping:
		return []gin.HandlerFunc{}
	case UpdateActivity:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("manager"),
					RequiredScopes: models.Scopes{
						models.Scope("activity:edit"),
					},
				}),
		}
	case UpdateProjectConfig:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
				}),
		}
	case UpdateTeam:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case UpdateTimeEntry:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case UpdateUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("user"),
				}),
		}
	case UpdateUserAuthorization:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
					RequiredScopes: models.Scopes{
						models.Scope("scopes:write"),
					},
				}),
		}
	case UpdateWorkItemComment:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					RequiredScopes: models.Scopes{
						models.Scope("work-item-comment:edit"),
					},
				}),
		}
	case UpdateWorkItemTag:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case UpdateWorkItemType:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case UpdateWorkitem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	default:
		return []gin.HandlerFunc{}
	}
}
