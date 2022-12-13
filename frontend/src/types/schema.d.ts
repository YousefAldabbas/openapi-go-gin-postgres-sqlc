/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable */
// @ts-nocheck
export type schemas = components['schemas']

/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


/** Type helpers */
type Without<T, U> = { [P in Exclude<keyof T, keyof U>]?: never };
type XOR<T, U> = (T | U) extends object ? (Without<T, U> & U) | (Without<U, T> & T) : T | U;
type OneOf<T extends any[]> = T extends [infer Only] ? Only : T extends [infer A, infer B, ...infer Rest] ? OneOf<[XOR<A, B>, ...Rest]> : never;

export interface paths {
  "/auth/myprovider/callback": {
    get: operations["MyProviderCallback"];
  };
  "/auth/myprovider/login": {
    get: operations["MyProviderLogin"];
  };
  "/events": {
    get: operations["Events"];
  };
  "/ping": {
    /** Ping pongs */
    get: operations["Ping"];
  };
  "/openapi.yaml": {
    /** Returns this very OpenAPI spec. */
    get: operations["OpenapiYamlGet"];
  };
  "/admin/ping": {
    /** Ping pongs */
    get: operations["AdminPing"];
  };
  "/user/me": {
    /** returns the logged in user */
    get: operations["GetCurrentUser"];
  };
  "/user/{id}/authorization": {
    /** updates user role and scopes by id */
    patch: operations["UpdateUserAuthorization"];
  };
  "/user/{id}": {
    /** deletes the user by id */
    delete: operations["DeleteUser"];
    /** updates the user by id */
    patch: operations["UpdateUser"];
  };
  "/project/{id}/initialize": {
    /** creates initial data (teams, work item types, tags...) for a new project */
    post: operations["InitializeProject"];
  };
  "/project/{id}/board": {
    /** returns board data for a project */
    get: operations["GetProjectBoard"];
  };
  "/project/{id}/workitems": {
    /** returns workitems for a project */
    get: operations["GetProjectWorkitems"];
  };
}

export interface components {
  schemas: {
    DemoProjectWorkItemsResponse: {
      baseWorkItem: components["schemas"]["DbWorkItemPublic"];
      /** Format: date-time */
      lastMessageAt: string;
      line: string;
      ref: string;
      reopened: boolean;
      workItemID: number;
    };
    InitializeProjectRequest: {
      activities?: (components["schemas"]["ReposActivityCreateParams"])[] | null;
      kanbanSteps?: (components["schemas"]["ReposKanbanStepCreateParams"])[] | null;
      projectID?: number;
      teams?: (components["schemas"]["ReposTeamCreateParams"])[] | null;
      workItemTags?: (components["schemas"]["ReposWorkItemTagCreateParams"])[] | null;
      workItemTypes?: (components["schemas"]["ReposWorkItemTypeCreateParams"])[] | null;
    };
    ProjectBoardResponse: {
      activities?: (components["schemas"]["DbActivityPublic"])[] | null;
      kanbanSteps?: (components["schemas"]["DbKanbanStepPublic"])[] | null;
      project?: components["schemas"]["DbProjectPublic"];
      teams?: (components["schemas"]["DbTeamPublic"])[] | null;
      workItemTags?: (components["schemas"]["DbWorkItemTagPublic"])[] | null;
      workItemTypes?: (components["schemas"]["DbWorkItemTypePublic"])[] | null;
    };
    UserResponse: {
      apiKey?: components["schemas"]["DbUserAPIKeyPublic"];
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt: string | null;
      email: string;
      firstName: string | null;
      fullName: string | null;
      hasGlobalNotifications: boolean;
      hasPersonalNotifications: boolean;
      lastName: string | null;
      role: components["schemas"]["Role"];
      scopes: components["schemas"]["Scopes"];
      teams?: (components["schemas"]["DbTeamPublic"])[] | null;
      userID: components["schemas"]["UuidUUID"];
      username: string;
    };
    /** HTTPValidationError */
    HTTPValidationError: {
      /** Detail */
      detail?: (components["schemas"]["ValidationError"])[];
    };
    /**
     * @description string identifiers for SSE event listeners. 
     * @enum {string}
     */
    Topics: "UserNotifications" | "ManagerNotifications" | "AdminNotifications" | "WorkItemMoved" | "WorkItemClosed";
    /** @enum {string} */
    Scope: "test-scope" | "users:read" | "users:write" | "scopes:write" | "team-settings:write" | "project-settings:write" | "work-item:review";
    Scopes: (components["schemas"]["Scope"])[];
    /** @enum {string} */
    Role: "guest" | "user" | "advancedUser" | "manager" | "admin" | "superAdmin";
    /**
     * Notification type 
     * @description User notification type. 
     * @enum {string}
     */
    NotificationType: "personal" | "global";
    /**
     * WorkItem role 
     * @description Role in work item for a member. 
     * @enum {string}
     */
    WorkItemRole: "preparer" | "reviewer";
    /**
     * @description represents User data to update 
     * @example {
     *   "first_name": "Jane",
     *   "last_name": "Doe"
     * }
     */
    UpdateUserRequest: {
      /** @description originally from auth server but updatable */
      first_name?: string;
      /** @description originally from auth server but updatable */
      last_name?: string;
    };
    /**
     * @description represents User authorization data to update 
     * @example {
     *   "role": "manager",
     *   "scopes": [
     *     "test-scope"
     *   ]
     * }
     */
    UpdateUserAuthRequest: {
      role?: components["schemas"]["Role"];
      scopes?: components["schemas"]["Scopes"];
    };
    /** ValidationError */
    ValidationError: {
      /** Location */
      loc: (string)[];
      /** Message */
      msg: string;
      /** Error Type */
      type: string;
    };
    DbTeamPublic: {
      /** Format: date-time */
      createdAt: string;
      description: string;
      name: string;
      projectID: number;
      teamID: number;
      /** Format: date-time */
      updatedAt: string;
    };
    DbUserAPIKeyPublic: {
      apiKey: string;
      /** Format: date-time */
      expiresOn: string;
      userID: components["schemas"]["UuidUUID"];
    } | null;
    DbActivityPublic: {
      activityID: number;
      description: string;
      isProductive: boolean;
      name: string;
      projectID: number;
    };
    DbKanbanStepPublic: {
      color: string;
      description: string;
      kanbanStepID: number;
      name: string;
      projectID: number;
      stepOrder: number | null;
      timeTrackable: boolean;
    };
    DbProjectPublic: {
      /** Format: date-time */
      createdAt: string;
      description: string;
      initialized: boolean;
      name: string;
      projectID: number;
      /** Format: date-time */
      updatedAt: string;
    } | null;
    DbWorkItemTagPublic: {
      color: string;
      description: string;
      name: string;
      projectID: number;
      workItemTagID: number;
    };
    DbWorkItemTypePublic: {
      color: string;
      description: string;
      name: string;
      projectID: number;
      workItemTypeID: number;
    };
    ReposActivityCreateParams: {
      description?: string;
      isProductive?: boolean;
      name?: string;
      projectID?: number;
    };
    ReposKanbanStepCreateParams: {
      color?: string;
      description?: string;
      name?: string;
      projectID?: number;
      stepOrder?: number;
      timeTrackable?: boolean;
    };
    ReposTeamCreateParams: {
      description?: string;
      name?: string;
      projectID?: number;
    };
    ReposWorkItemTagCreateParams: {
      color?: string;
      description?: string;
      name?: string;
      projectID?: number;
    };
    ReposWorkItemTypeCreateParams: {
      color?: string;
      description?: string;
      name?: string;
      projectID?: number;
    };
    ModelsRole: string;
    UuidUUID: string;
    DbWorkItemPublic: {
      /** Format: date-time */
      closed: string | null;
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt: string | null;
      description: string;
      kanbanStepID: number;
      metadata: components["schemas"]["PgtypeJSONB"];
      /** Format: date-time */
      targetDate: string;
      teamID: number;
      title: string;
      /** Format: date-time */
      updatedAt: string;
      workItemID: number;
      workItemTypeID: number;
    };
    PgtypeJSONB: Record<string, never>;
  };
  responses: never;
  parameters: {
    /**
     * @description UUID identifier of entity that needs to be updated 
     * @example 123e4567-e89b-12d3-a456-426614174000
     */
    uuid: string;
    /**
     * @description integer identifier that needs to be updated 
     * @example 41131
     */
    PathSerial: number;
  };
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type external = Record<string, never>;

export interface operations {

  MyProviderCallback: {
    responses: {
      /** @description callback for MyProvider auth server */
      200: never;
    };
  };
  MyProviderLogin: {
    responses: {
      /** @description redirect to MyProvider auth server login */
      302: never;
    };
  };
  Events: {
    responses: {
      /** @description events */
      200: {
        content: {
          "text/event-stream": string;
        };
      };
    };
  };
  Ping: {
    /** Ping pongs */
    responses: {
      /** @description OK */
      200: {
        content: {
          "text/plain": string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          "application/json": components["schemas"]["HTTPValidationError"];
        };
      };
    };
  };
  OpenapiYamlGet: {
    /** Returns this very OpenAPI spec. */
    responses: {
      /** @description OpenAPI YAML file. */
      200: {
        content: {
          "application/x-yaml": string;
        };
      };
    };
  };
  AdminPing: {
    /** Ping pongs */
    responses: {
      /** @description OK */
      200: {
        content: {
          "text/plain": string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          "application/json": components["schemas"]["HTTPValidationError"];
        };
      };
    };
  };
  GetCurrentUser: {
    /** returns the logged in user */
    responses: {
      /** @description ok */
      200: {
        content: {
          "application/json": components["schemas"]["UserResponse"];
        };
      };
    };
  };
  UpdateUserAuthorization: {
    /** updates user role and scopes by id */
    /** @description Updated user object */
    requestBody: {
      content: {
        "application/json": components["schemas"]["UpdateUserAuthRequest"];
      };
    };
    responses: {
      /** @description User auth updated successfully. */
      204: never;
    };
  };
  DeleteUser: {
    /** deletes the user by id */
    responses: {
      /** @description User not found */
      404: never;
    };
  };
  UpdateUser: {
    /** updates the user by id */
    /** @description Updated user object */
    requestBody: {
      content: {
        "application/json": components["schemas"]["UpdateUserRequest"];
      };
    };
    responses: {
      /** @description ok */
      200: {
        content: {
          "application/json": components["schemas"]["UserResponse"];
        };
      };
    };
  };
  InitializeProject: {
    /** creates initial data (teams, work item types, tags...) for a new project */
    requestBody: {
      content: {
        "application/json": components["schemas"]["InitializeProjectRequest"];
      };
    };
    responses: {
      /** @description Project successfully initialized. */
      204: never;
    };
  };
  GetProjectBoard: {
    /** returns board data for a project */
    responses: {
      /** @description Project successfully initialized. */
      200: {
        content: {
          "application/json": components["schemas"]["ProjectBoardResponse"];
        };
      };
    };
  };
  GetProjectWorkitems: {
    /** returns workitems for a project */
    parameters?: {
      query?: {
        open?: boolean;
        deleted?: boolean;
      };
    };
    responses: {
      /** @description Project successfully initialized. */
      200: {
        content: {
          "application/json": components["schemas"]["DemoProjectWorkItemsResponse"];
        };
      };
    };
  };
}
