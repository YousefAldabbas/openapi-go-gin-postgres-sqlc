package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UserNotification represents a row from 'public.user_notifications'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type UserNotification struct {
	UserNotificationID int64     `json:"userNotificationID" db:"user_notification_id" required:"true"` // user_notification_id
	NotificationID     int       `json:"notificationID" db:"notification_id" required:"true"`          // notification_id
	Read               bool      `json:"read" db:"read" required:"true"`                               // read
	UserID             uuid.UUID `json:"userID" db:"user_id" required:"true"`                          // user_id

	NotificationJoin *Notification `json:"-" db:"notification" openapi-go:"ignore"` // O2O
	UserJoin         *User         `json:"-" db:"user" openapi-go:"ignore"`         // O2O

}

// UserNotificationCreateParams represents insert params for 'public.user_notifications'
type UserNotificationCreateParams struct {
	NotificationID int       `json:"notificationID"` // notification_id
	Read           bool      `json:"read"`           // read
	UserID         uuid.UUID `json:"userID"`         // user_id
}

func (un *UserNotification) SetCreateParams(params *UserNotificationCreateParams) {
	un.NotificationID = params.NotificationID
	un.Read = params.Read
	un.UserID = params.UserID
}

// UserNotificationUpdateParams represents update params for 'public.user_notifications'
type UserNotificationUpdateParams struct {
	NotificationID *int       `json:"notificationID"` // notification_id
	Read           *bool      `json:"read"`           // read
	UserID         *uuid.UUID `json:"userID"`         // user_id
}

func (un *UserNotification) SetUpdateParams(params *UserNotificationUpdateParams) {

	if params.NotificationID != nil {
		un.NotificationID = *params.NotificationID
	}

	if params.Read != nil {
		un.Read = *params.Read
	}

	if params.UserID != nil {
		un.UserID = *params.UserID
	}
}

type UserNotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   UserNotificationJoins
}
type UserNotificationSelectConfigOption func(*UserNotificationSelectConfig)

// WithUserNotificationLimit limits row selection.
func WithUserNotificationLimit(limit int) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type UserNotificationOrderBy = string

const ()

type UserNotificationJoins struct {
	Notification bool
	User         bool
}

// WithUserNotificationJoin joins with the given tables.
func WithUserNotificationJoin(joins UserNotificationJoins) UserNotificationSelectConfigOption {
	return func(s *UserNotificationSelectConfig) {
		s.joins = joins
	}
}

// Insert inserts the UserNotification to the database.
func (un *UserNotification) Insert(ctx context.Context, db DB) (*UserNotification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.user_notifications (` +
		`notification_id, read, user_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, un.NotificationID, un.Read, un.UserID)

	rows, err := db.Query(ctx, sqlstr, un.NotificationID, un.Read, un.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Insert/db.Query: %w", err))
	}
	newun, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Insert/pgx.CollectOneRow: %w", err))
	}

	*un = newun

	return un, nil
}

// Update updates a UserNotification in the database.
func (un *UserNotification) Update(ctx context.Context, db DB) (*UserNotification, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.user_notifications SET ` +
		`notification_id = $1, read = $2, user_id = $3 ` +
		`WHERE user_notification_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, un.NotificationID, un.Read, un.UserID, un.UserNotificationID)

	rows, err := db.Query(ctx, sqlstr, un.NotificationID, un.Read, un.UserID, un.UserNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Update/db.Query: %w", err))
	}
	newun, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserNotification/Update/pgx.CollectOneRow: %w", err))
	}
	*un = newun

	return un, nil
}

// Upsert performs an upsert for UserNotification.
func (un *UserNotification) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.user_notifications (` +
		`user_notification_id, notification_id, read, user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (user_notification_id) DO ` +
		`UPDATE SET ` +
		`notification_id = EXCLUDED.notification_id, read = EXCLUDED.read, user_id = EXCLUDED.user_id ` +
		` RETURNING * `
	// run
	logf(sqlstr, un.UserNotificationID, un.NotificationID, un.Read, un.UserID)
	if _, err := db.Exec(ctx, sqlstr, un.UserNotificationID, un.NotificationID, un.Read, un.UserID); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
}

// Delete deletes the UserNotification from the database.
func (un *UserNotification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.user_notifications ` +
		`WHERE user_notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, un.UserNotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// UserNotificationByNotificationIDUserID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_notification_id_user_id_key'.
func UserNotificationByNotificationIDUserID(ctx context.Context, db DB, notificationID int, userID uuid.UUID, opts ...UserNotificationSelectConfigOption) (*UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and notifications.notification_id is not null then row(notifications.*) end) as notification,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from O2M|M2O)"
left join notifications on notifications.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from O2M|M2O)"
left join users on users.user_id = user_notifications.user_id` +
		` WHERE user_notifications.notification_id = $3 AND user_notifications.user_id = $4 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, notificationID, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, notificationID, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/db.Query: %w", err))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByNotificationIDUserID/pgx.CollectOneRow: %w", err))
	}

	return &un, nil
}

// UserNotificationsByNotificationID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_notification_id_user_id_key'.
func UserNotificationsByNotificationID(ctx context.Context, db DB, notificationID int, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and notifications.notification_id is not null then row(notifications.*) end) as notification,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from O2M|M2O)"
left join notifications on notifications.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from O2M|M2O)"
left join users on users.user_id = user_notifications.user_id` +
		` WHERE user_notifications.notification_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, notificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, notificationID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserNotificationByUserNotificationID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_pkey'.
func UserNotificationByUserNotificationID(ctx context.Context, db DB, userNotificationID int64, opts ...UserNotificationSelectConfigOption) (*UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and notifications.notification_id is not null then row(notifications.*) end) as notification,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from O2M|M2O)"
left join notifications on notifications.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from O2M|M2O)"
left join users on users.user_id = user_notifications.user_id` +
		` WHERE user_notifications.user_notification_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userNotificationID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, userNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/db.Query: %w", err))
	}
	un, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_notifications/UserNotificationByUserNotificationID/pgx.CollectOneRow: %w", err))
	}

	return &un, nil
}

// UserNotificationsByUserID retrieves a row from 'public.user_notifications' as a UserNotification.
//
// Generated from index 'user_notifications_user_id_idx'.
func UserNotificationsByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserNotificationSelectConfigOption) ([]UserNotification, error) {
	c := &UserNotificationSelectConfig{joins: UserNotificationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_notifications.user_notification_id,
user_notifications.notification_id,
user_notifications.read,
user_notifications.user_id,
(case when $1::boolean = true and notifications.notification_id is not null then row(notifications.*) end) as notification,
(case when $2::boolean = true and users.user_id is not null then row(users.*) end) as user ` +
		`FROM public.user_notifications ` +
		`-- O2O join generated from "user_notifications_notification_id_fkey (Generated from O2M|M2O)"
left join notifications on notifications.notification_id = user_notifications.notification_id
-- O2O join generated from "user_notifications_user_id_fkey (Generated from O2M|M2O)"
left join users on users.user_id = user_notifications.user_id` +
		` WHERE user_notifications.user_id = $3 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Notification, c.joins.User, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}
