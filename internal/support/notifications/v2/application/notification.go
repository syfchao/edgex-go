//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package application

import (
	"context"

	"github.com/edgexfoundry/edgex-go/internal/pkg/correlation"
	v2NotificationsContainer "github.com/edgexfoundry/edgex-go/internal/support/notifications/v2/bootstrap/container"

	"github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// The AddNotification function accepts the new Notification model from the controller function
// and then invokes AddNotification function of infrastructure layer to add new Notification
func AddNotification(d models.Notification, ctx context.Context, dic *di.Container) (id string, edgeXerr errors.EdgeX) {
	dbClient := v2NotificationsContainer.DBClientFrom(dic.Get)
	lc := container.LoggingClientFrom(dic.Get)

	addedNotification, edgeXerr := dbClient.AddNotification(d)
	if edgeXerr != nil {
		return "", errors.NewCommonEdgeXWrapper(edgeXerr)
	}

	lc.Debugf("Notification created on DB successfully. Notification ID: %s, Correlation-ID: %s ",
		addedNotification.Id,
		correlation.FromContext(ctx))

	// TODO: distribute notification

	return addedNotification.Id, nil
}

// NotificationsByCategory queries notifications with offset, limit, and category
func NotificationsByCategory(offset, limit int, category string, dic *di.Container) (notifications []dtos.Notification, err errors.EdgeX) {
	if category == "" {
		return notifications, errors.NewCommonEdgeX(errors.KindContractInvalid, "category is empty", nil)
	}
	dbClient := v2NotificationsContainer.DBClientFrom(dic.Get)
	notificationModels, err := dbClient.NotificationsByCategory(offset, limit, category)
	if err != nil {
		return notifications, errors.NewCommonEdgeXWrapper(err)
	}
	notifications = make([]dtos.Notification, len(notificationModels))
	for i, n := range notificationModels {
		notifications[i] = dtos.FromNotificationModelToDTO(n)
	}
	return notifications, nil
}
