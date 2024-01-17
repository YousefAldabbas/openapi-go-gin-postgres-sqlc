package servicetestutil

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/require"
)

type CreateNotificationParams struct {
	Receiver     *db.UserID
	ReceiverRole *models.Role
}

// CreatePersonalNotification creates a new notification with the given configuration.
func (ff *FixtureFactory) CreatePersonalNotification(ctx context.Context, params CreateNotificationParams) *db.UserNotification {
	admin := ff.CreateUser(ctx, CreateUserParams{Role: models.RoleAdmin})

	n, err := ff.svc.Notification.CreateNotification(ctx, ff.d, &services.NotificationCreateParams{
		NotificationCreateParams: db.NotificationCreateParams{
			Body:             testutil.RandomString(10),
			Labels:           []string{"label " + fmt.Sprint(testutil.RandomInt(1, 9999))},
			Link:             pointers.New("https://somelink"),
			Title:            testutil.RandomNameIdentifier(0, "-"),
			Sender:           admin.User.UserID,
			NotificationType: db.NotificationTypePersonal,
			Receiver:         params.Receiver,
		},
	})
	require.NoError(ff.t, err)

	return n
}

// CreateGlobalNotification creates a new global notification with the given configuration.
// Returns a single user notification from the fan out.
func (ff *FixtureFactory) CreateGlobalNotification(ctx context.Context, params CreateNotificationParams) *db.UserNotification {
	admin := ff.CreateUser(ctx, CreateUserParams{Role: models.RoleAdmin})

	n, err := ff.svc.Notification.CreateNotification(ctx, ff.d, &services.NotificationCreateParams{
		NotificationCreateParams: db.NotificationCreateParams{
			Body:             testutil.RandomString(10),
			Labels:           []string{"label " + fmt.Sprint(testutil.RandomInt(1, 9999))},
			Link:             pointers.New("https://somelink"),
			Title:            testutil.RandomNameIdentifier(0, "-"),
			Sender:           admin.User.UserID,
			NotificationType: db.NotificationTypeGlobal,
		},
		ReceiverRole: params.ReceiverRole,
	})
	require.NoError(ff.t, err)

	return n
}
