/* Generated by openapi-typescript v6.4.3 */
/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable */
// @ts-nocheck
export type schemas = components['schemas']

/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


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
  "/notifications/user/page": {
    /** Get paginated user notifications */
    get: operations["GetPaginatedNotifications"];
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
  "/project/{projectName}/team/": {
    /** create team. */
    post: operations["CreateTeam"];
  };
  "/project/{projectName}/team/{id}/": {
    /** get team. */
    get: operations["GetTeam"];
    /** delete team. */
    delete: operations["DeleteTeam"];
    /** update team. */
    patch: operations["UpdateTeam"];
  };
  "/project/{projectName}/workItemTag/": {
    /** create workitemtag. */
    post: operations["CreateWorkItemTag"];
  };
  "/project/{projectName}/workItemTag/{id}/": {
    /** get workitemtag. */
    get: operations["GetWorkItemTag"];
    /** delete workitemtag. */
    delete: operations["DeleteWorkItemTag"];
    /** update workitemtag. */
    patch: operations["UpdateWorkItemTag"];
  };
  "/project/{projectName}/workItemType/": {
    /** create workitemtype. */
    post: operations["CreateWorkItemType"];
  };
  "/project/{projectName}/workItemType/{id}/": {
    /** get workitemtype. */
    get: operations["GetWorkItemType"];
    /** delete workitemtype. */
    delete: operations["DeleteWorkItemType"];
    /** update workitemtype. */
    patch: operations["UpdateWorkItemType"];
  };
  "/user/me": {
    /** returns the logged in user */
    get: operations["GetCurrentUser"];
  };
  [path: `/user/${string}/authorization`]: {
    /** updates user role and scopes by id */
    patch: operations["UpdateUserAuthorization"];
  };
  [path: `/user/${string}`]: {
    /** deletes the user by id */
    delete: operations["DeleteUser"];
    /** updates the user by id */
    patch: operations["UpdateUser"];
  };
  "/project/{projectName}/initialize": {
    /** creates initial data (teams, tags...) for a new project */
    post: operations["InitializeProject"];
  };
  "/project/{projectName}/": {
    /** returns board data for a project */
    get: operations["GetProject"];
  };
  "/project/{projectName}/config": {
    /** returns the project configuration */
    get: operations["GetProjectConfig"];
    /** updates the project configuration */
    put: operations["UpdateProjectConfig"];
  };
  "/project/{projectName}/board": {
    /** returns board data for a project */
    get: operations["GetProjectBoard"];
  };
  "/project/{projectName}/workitems": {
    /** returns workitems for a project */
    get: operations["GetProjectWorkitems"];
  };
  "/workitem/": {
    /** create workitem */
    post: operations["CreateWorkitem"];
  };
  "/workitem/{id}/": {
    /** get workitem */
    get: operations["GetWorkItem"];
    /** delete workitem */
    delete: operations["DeleteWorkitem"];
    /** update workitem */
    patch: operations["UpdateWorkitem"];
  };
  "/workitem/{id}/comments/": {
    /** create workitem comment */
    post: operations["CreateWorkitemComment"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    WorkItemTagCreateRequest: {
      color: string;
      description: string;
      name: string;
    };
    WorkItemTagUpdateRequest: {
      color?: string;
      description?: string;
      name?: string;
    };
    WorkItemTag: {
      color: string;
      description: string;
      name: string;
      projectID: number;
      workItemTagID: number;
    };
    WorkItemTypeCreateRequest: {
      color: string;
      description: string;
      name: string;
    };
    WorkItemTypeUpdateRequest: {
      color?: string;
      description?: string;
      name?: string;
    };
    WorkItemType: {
      color: string;
      description: string;
      name: string;
      projectID: number;
      workItemTypeID: number;
    };
    TeamCreateRequest: {
      description: string;
      name: string;
      projectID: number;
    };
    TeamUpdateRequest: {
      description?: string;
      name?: string;
      projectID?: number;
    };
    Team: {
      /** Format: date-time */
      createdAt: string;
      description: string;
      name: string;
      projectID: number;
      teamID: number;
      /** Format: date-time */
      updatedAt: string;
    };
    /** @enum {string} */
    Direction: "asc" | "desc";
    PaginatedNotificationsResponse: {
      items: components["schemas"]["RestNotification"][] | null;
      page: components["schemas"]["RestPaginationPage"];
    };
    DbActivity: {
      activityID: number;
      description: string;
      isProductive: boolean;
      name: string;
      projectID: number;
    };
    DbKanbanStep: {
      color: string;
      description: string;
      kanbanStepID: number;
      name: string;
      projectID: number;
      stepOrder: number;
      timeTrackable: boolean;
    };
    DbProject: {
      boardConfig: components["schemas"]["ProjectConfig"];
      /** Format: date-time */
      createdAt: string;
      description: string;
      name: components["schemas"]["Project"];
      projectID: number;
      /** Format: date-time */
      updatedAt: string;
    };
    DbTeam: {
      /** Format: date-time */
      createdAt: string;
      description: string;
      name: string;
      projectID: number;
      teamID: number;
      /** Format: date-time */
      updatedAt: string;
    };
    DbWorkItemTag: {
      color: string;
      description: string;
      name: string;
      projectID: number;
      workItemTagID: number;
    };
    DbWorkItemType: {
      color: string;
      description: string;
      name: string;
      projectID: number;
      workItemTypeID: number;
    };
    DbDemoWorkItem: {
      /** Format: date-time */
      lastMessageAt: string;
      line: string;
      ref: string;
      reopened: boolean;
      workItemID: number;
    };
    DbUserAPIKey: {
      apiKey: string;
      /** Format: date-time */
      expiresOn: string;
      userID: components["schemas"]["DbUserID"];
    };
    DbUser: {
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt?: string | null;
      email: string;
      firstName?: string | null;
      fullName?: string | null;
      hasGlobalNotifications: boolean;
      hasPersonalNotifications: boolean;
      lastName?: string | null;
      scopes: components["schemas"]["Scopes"];
      userID: components["schemas"]["DbUserID"];
      username: string;
    };
    DbTimeEntry: {
      activityID: number;
      comment: string;
      durationMinutes?: number | null;
      /** Format: date-time */
      start: string;
      teamID?: number | null;
      timeEntryID: number;
      userID: components["schemas"]["DbUserID"];
      workItemID?: number | null;
    };
    DbWorkItemComment: {
      /** Format: date-time */
      createdAt: string;
      message: string;
      /** Format: date-time */
      updatedAt: string;
      userID: components["schemas"]["DbUserID"];
      workItemCommentID: number;
      workItemID: number;
    };
    ProjectConfig: {
      fields: components["schemas"]["ProjectConfigField"][];
      header: string[];
      visualization?: Record<string, never>;
    };
    ProjectConfigField: {
      isEditable: boolean;
      isVisible: boolean;
      name: string;
      path: string;
      showCollapsed: boolean;
    };
    DemoWorkItems: {
      /** Format: date-time */
      closedAt?: string | null;
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt?: string | null;
      demoWorkItem: components["schemas"]["DbDemoWorkItem"];
      description: string;
      kanbanStepID: number;
      members?: components["schemas"]["DbUser"][] | null;
      metadata: {
        [key: string]: unknown;
      };
      /** Format: date-time */
      targetDate: string;
      teamID: number;
      timeEntries?: components["schemas"]["DbTimeEntry"][] | null;
      title: string;
      /** Format: date-time */
      updatedAt: string;
      workItemComments?: components["schemas"]["DbWorkItemComment"][] | null;
      workItemID: number;
      workItemTags?: components["schemas"]["DbWorkItemTag"][] | null;
      workItemType?: components["schemas"]["DbWorkItemType"];
      workItemTypeID: number;
    };
    DemoTwoWorkItems: {
      /** Format: date-time */
      closedAt?: string | null;
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt?: string | null;
      demoTwoWorkItem: components["schemas"]["DbDemoTwoWorkItem"];
      description: string;
      kanbanStepID: number;
      members?: components["schemas"]["DbUser"][] | null;
      metadata: {
        [key: string]: unknown;
      };
      /** Format: date-time */
      targetDate: string;
      teamID: number;
      timeEntries?: components["schemas"]["DbTimeEntry"][] | null;
      title: string;
      /** Format: date-time */
      updatedAt: string;
      workItemComments?: components["schemas"]["DbWorkItemComment"][] | null;
      workItemID: number;
      workItemTags?: components["schemas"]["DbWorkItemTag"][] | null;
      workItemType?: components["schemas"]["DbWorkItemType"];
      workItemTypeID: number;
    };
    InitializeProjectRequest: {
      tags?: components["schemas"]["DbWorkItemTagCreateParams"][] | null;
      teams?: components["schemas"]["DbTeamCreateParams"][] | null;
    };
    ProjectBoard: {
      projectName: components["schemas"]["Project"];
    };
    User: {
      apiKey?: components["schemas"]["DbUserAPIKey"];
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt?: string | null;
      email: string;
      firstName?: string | null;
      fullName?: string | null;
      hasGlobalNotifications: boolean;
      hasPersonalNotifications: boolean;
      lastName?: string | null;
      projects?: components["schemas"]["DbProject"][] | null;
      role: components["schemas"]["Role"];
      scopes: components["schemas"]["Scopes"];
      teams?: components["schemas"]["DbTeam"][] | null;
      userID: components["schemas"]["DbUserID"];
      username: string;
    };
    /** HTTPValidationError */
    HTTPValidationError: {
      /**
       * Detail
       * @description Additional details for validation errors
       */
      detail?: components["schemas"]["ValidationError"][];
      /**
       * Messages
       * @description Descriptive error messages to show in a callout
       */
      messages: string[];
    };
    /**
     * @description Represents standardized HTTP error types.
     * Notes:
     * - 'Private' marks an error to be hidden in response.
     *
     * @enum {string}
     */
    ErrorCode: "Unknown" | "Private" | "NotFound" | "InvalidArgument" | "AlreadyExists" | "Unauthorized" | "Unauthenticated" | "RequestValidation" | "ResponseValidation" | "OIDC" | "InvalidRole" | "InvalidScope" | "InvalidUUID";
    /** @description represents an error message response. */
    HTTPError: {
      title: string;
      detail: string;
      status: number;
      error: string;
      /**
       * Location
       * @description location in body path, if any
       */
      loc?: string[];
      type: components["schemas"]["ErrorCode"];
      validationError?: components["schemas"]["HTTPValidationError"];
    };
    /**
     * @description string identifiers for SSE event listeners.
     * @enum {string}
     */
    Topics: "GlobalAlerts";
    /** @enum {string} */
    Scope: "users:read" | "users:write" | "users:delete" | "scopes:write" | "team-settings:write" | "project-settings:write" | "work-item-tag:create" | "work-item-tag:edit" | "work-item-tag:delete" | "work-item:review";
    Scopes: components["schemas"]["Scope"][];
    /** @enum {string} */
    Role: "guest" | "user" | "advancedUser" | "manager" | "admin" | "superAdmin";
    /**
     * WorkItem role
     * @description represents a database 'work_item_role'
     * @enum {string}
     */
    WorkItemRole: "preparer" | "reviewer";
    /**
     * @description represents User data to update
     * @example {
     *   "firstName": "Jane",
     *   "lastName": "Doe"
     * }
     */
    UpdateUserRequest: {
      /** @description originally from auth server but updatable */
      firstName?: string;
      /** @description originally from auth server but updatable */
      lastName?: string;
    };
    /**
     * @description represents User authorization data to update
     * @example {
     *   "role": "manager",
     *   "scopes": [
     *     "users:read"
     *   ]
     * }
     */
    UpdateUserAuthRequest: {
      role?: components["schemas"]["Role"];
      scopes?: components["schemas"]["Scopes"];
    };
    /** ValidationError */
    ValidationError: {
      /**
       * Location
       * @description location in body path, if any
       */
      loc: string[];
      /**
       * Message
       * @description should always be shown to the user
       */
      msg: string;
      /**
       * Error details
       * @description verbose details of the error
       */
      detail: {
        schema: Record<string, never>;
        value: string;
      };
      /** Contextual information */
      ctx?: Record<string, never>;
    };
    UuidUUID: string;
    WorkItemCreateRequest: components["schemas"]["DemoWorkItemCreateRequest"] | components["schemas"]["DemoTwoWorkItemCreateRequest"];
    DbWorkItem: {
      /** Format: date-time */
      closedAt?: string | null;
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt?: string | null;
      description: string;
      kanbanStepID: number;
      metadata: {
        [key: string]: unknown;
      };
      /** Format: date-time */
      targetDate: string;
      teamID: number;
      title: string;
      /** Format: date-time */
      updatedAt: string;
      workItemID: number;
      workItemTypeID: number;
    };
    DemoTwoWorkItemCreateRequest: {
      base: components["schemas"]["DbWorkItemCreateParams"];
      demoTwoProject: components["schemas"]["DbDemoTwoWorkItemCreateParams"];
      members: components["schemas"]["ServicesMember"][];
      projectName: components["schemas"]["Project"];
      tagIDs: number[];
    };
    DemoWorkItemCreateRequest: {
      base: components["schemas"]["DbWorkItemCreateParams"];
      demoProject: components["schemas"]["DbDemoWorkItemCreateParams"];
      members: components["schemas"]["ServicesMember"][];
      projectName: components["schemas"]["Project"];
      tagIDs: number[];
    };
    WorkItemCommentCreateRequest: {
      message: string;
      userID: components["schemas"]["DbUserID"];
      workItemID: number;
    };
    /** @enum {string} */
    Project: "demo" | "demo_two";
    DbActivityCreateParams: {
      description: string;
      isProductive: boolean;
      name: string;
      projectID?: number;
    };
    DbTeamCreateParams: {
      description: string;
      name: string;
      projectID: number;
    };
    DbWorkItemTagCreateParams: {
      color: string;
      description: string;
      name: string;
    };
    DbWorkItemRole: string;
    /**
     * @description represents a database 'notification_type'
     * @enum {string}
     */
    NotificationType: "personal" | "global";
    /** @enum {string} */
    DemoProjectKanbanSteps: "Disabled" | "Received" | "Under review" | "Work in progress";
    /** @enum {string} */
    DemoProject2KanbanSteps: "Received";
    /** @enum {string} */
    Demo2WorkItemTypes: "Type 1" | "Type 2" | "Another type";
    /** @enum {string} */
    DemoKanbanSteps: "Disabled" | "Received" | "Under review" | "Work in progress";
    /** @enum {string} */
    DemoTwoKanbanSteps: "Received";
    /** @enum {string} */
    DemoTwoWorkItemTypes: "Type 1" | "Type 2" | "Another type";
    /** @enum {string} */
    DemoWorkItemTypes: "Type 1";
    DbDemoWorkItemCreateParams: {
      /** Format: date-time */
      lastMessageAt: string;
      line: string;
      ref: string;
      reopened: boolean;
    };
    DbWorkItemCreateParams: {
      /** Format: date-time */
      closedAt?: string | null;
      description: string;
      kanbanStepID: number;
      metadata: {
        [key: string]: unknown;
      };
      /** Format: date-time */
      targetDate: string;
      teamID: number;
      title: string;
      workItemTypeID: number;
    };
    ServicesMember: {
      role: components["schemas"]["WorkItemRole"];
      userID: components["schemas"]["DbUserID"];
    };
    DbDemoTwoWorkItem: {
      /** Format: date-time */
      customDateForProject2?: string | null;
      workItemID: components["schemas"]["DbWorkItemID"];
    };
    DbDemoTwoWorkItemCreateParams: {
      /** Format: date-time */
      customDateForProject2?: string | null;
    };
    DbWorkItemID: unknown;
    DbProjectID: unknown;
    DbUserID: string;
    DbWorkItemTypeID: unknown;
    DbNotificationID: unknown;
    DbNotification: {
      body: string;
      /** Format: date-time */
      createdAt: string;
      labels: string[];
      link?: string | null;
      notificationID: components["schemas"]["DbNotificationID"];
      notificationType: components["schemas"]["NotificationType"];
      receiver?: components["schemas"]["DbUserID"];
      sender: components["schemas"]["DbUserID"];
      title: string;
    };
    DbUserNotification: {
      notificationID: number;
      read: boolean;
      userID: components["schemas"]["DbUserID"];
      userNotificationID: number;
    };
    RestPaginationPage: {
      nextCursor?: string;
    };
    RestNotification: {
      notification: components["schemas"]["DbNotification"];
      notificationID: number;
      read: boolean;
      userID: components["schemas"]["DbUserID"];
      userNotificationID: number;
    };
  };
  responses: never;
  parameters: {
    /**
     * @description Project name
     * @example demo
     */
    ProjectName: components["schemas"]["Project"];
    /**
     * @description UUID identifier
     * @example 123e4567-e89b-12d3-a456-426614174000
     */
    UUID: string;
    /**
     * @description integer identifier
     * @example 41131
     */
    SerialID: number;
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
    parameters: {
      query: {
        projectName: components["schemas"]["Project"];
      };
    };
    responses: {
      /** @description events */
      200: {
        content: {
          "text/event-stream": string;
        };
      };
    };
  };
  /** Get paginated user notifications */
  GetPaginatedNotifications: {
    parameters: {
      query: {
        limit: number;
        direction: components["schemas"]["Direction"];
        cursor: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          "application/json": components["schemas"]["PaginatedNotificationsResponse"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** Ping pongs */
  Ping: {
    responses: {
      /** @description OK */
      200: {
        content: {
          "text/plain": string;
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** Returns this very OpenAPI spec. */
  OpenapiYamlGet: {
    responses: {
      /** @description OpenAPI YAML file. */
      200: {
        content: {
          "application/x-yaml": string;
        };
      };
    };
  };
  /** Ping pongs */
  AdminPing: {
    responses: {
      /** @description OK */
      200: {
        content: {
          "text/plain": string;
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** create team. */
  CreateTeam: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["TeamCreateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      201: {
        content: {
          "application/json": components["schemas"]["Team"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** get team. */
  GetTeam: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["Team"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** delete team. */
  DeleteTeam: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      204: never;
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** update team. */
  UpdateTeam: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["TeamUpdateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["Team"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** create workitemtag. */
  CreateWorkItemTag: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["WorkItemTagCreateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      201: {
        content: {
          "application/json": components["schemas"]["WorkItemTag"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** get workitemtag. */
  GetWorkItemTag: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["WorkItemTag"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** delete workitemtag. */
  DeleteWorkItemTag: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      204: never;
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** update workitemtag. */
  UpdateWorkItemTag: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["WorkItemTagUpdateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["WorkItemTag"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** create workitemtype. */
  CreateWorkItemType: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["WorkItemTypeCreateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      201: {
        content: {
          "application/json": components["schemas"]["WorkItemType"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** get workitemtype. */
  GetWorkItemType: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["WorkItemType"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** delete workitemtype. */
  DeleteWorkItemType: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      204: never;
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** update workitemtype. */
  UpdateWorkItemType: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
        id: components["parameters"]["SerialID"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["WorkItemTypeUpdateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["WorkItemType"];
        };
      };
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** returns the logged in user */
  GetCurrentUser: {
    responses: {
      /** @description ok */
      200: {
        content: {
          "application/json": components["schemas"]["User"];
        };
      };
    };
  };
  /** updates user role and scopes by id */
  UpdateUserAuthorization: {
    parameters: {
      path: {
        /**
         * @description UUID identifier
         * @example 123e4567-e89b-12d3-a456-426614174000
         */
        id: string;
      };
    };
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
  /** deletes the user by id */
  DeleteUser: {
    parameters: {
      path: {
        /**
         * @description UUID identifier
         * @example 123e4567-e89b-12d3-a456-426614174000
         */
        id: string;
      };
    };
    responses: {
      /** @description User deleted successfully */
      204: never;
      /** @description Unauthenticated */
      401: never;
      /** @description Unauthorized */
      403: never;
      /** @description Error response */
      "4XX": {
        content: {
          "application/json": components["schemas"]["HTTPError"];
        };
      };
    };
  };
  /** updates the user by id */
  UpdateUser: {
    parameters: {
      path: {
        /**
         * @description UUID identifier
         * @example 123e4567-e89b-12d3-a456-426614174000
         */
        id: string;
      };
    };
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
          "application/json": components["schemas"]["User"];
        };
      };
    };
  };
  /** creates initial data (teams, tags...) for a new project */
  InitializeProject: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["InitializeProjectRequest"];
      };
    };
    responses: {
      /** @description Success. */
      204: never;
    };
  };
  /** returns board data for a project */
  GetProject: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    responses: {
      /** @description Project. */
      200: {
        content: {
          "application/json": components["schemas"]["DbProject"];
        };
      };
    };
  };
  /** returns the project configuration */
  GetProjectConfig: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    responses: {
      /** @description Project config. */
      200: {
        content: {
          "application/json": components["schemas"]["ProjectConfig"];
        };
      };
    };
  };
  /** updates the project configuration */
  UpdateProjectConfig: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    requestBody?: {
      content: {
        "application/json": components["schemas"]["ProjectConfig"];
      };
    };
    responses: {
      /** @description Config updated successfully. */
      204: never;
    };
  };
  /** returns board data for a project */
  GetProjectBoard: {
    parameters: {
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["ProjectBoard"];
        };
      };
    };
  };
  /** returns workitems for a project */
  GetProjectWorkitems: {
    parameters: {
      query?: {
        open?: boolean;
        deleted?: boolean;
      };
      path: {
        projectName: components["parameters"]["ProjectName"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["DemoWorkItems"] | components["schemas"]["DemoTwoWorkItems"];
        };
      };
    };
  };
  /** create workitem */
  CreateWorkitem: {
    requestBody: {
      content: {
        "application/json": components["schemas"]["WorkItemCreateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      201: {
        content: {
          "application/json": components["schemas"]["DbWorkItem"];
        };
      };
    };
  };
  /** get workitem */
  GetWorkItem: {
    parameters: {
      path: {
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["DbWorkItem"];
        };
      };
    };
  };
  /** delete workitem */
  DeleteWorkitem: {
    parameters: {
      path: {
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      204: never;
    };
  };
  /** update workitem */
  UpdateWorkitem: {
    parameters: {
      path: {
        id: components["parameters"]["SerialID"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["DbWorkItem"];
        };
      };
    };
  };
  /** create workitem comment */
  CreateWorkitemComment: {
    parameters: {
      path: {
        id: components["parameters"]["SerialID"];
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["WorkItemCommentCreateRequest"];
      };
    };
    responses: {
      /** @description Success. */
      200: {
        content: {
          "application/json": components["schemas"]["DbWorkItemComment"];
        };
      };
    };
  };
}
