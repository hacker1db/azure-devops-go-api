﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package fileContainer

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API] Creates the specified items in in the referenced container.
// ctx
// items (required)
// containerId (required)
// scope (optional): A guid representing the scope of the container. This is often the project id.
func (client Client) CreateItems(ctx context.Context, items *azureDevops.VssJsonCollectionWrapper, containerId *int, scope *uuid.UUID) (*[]FileContainerItem, error) {
    if items == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "items"}
    }
    routeValues := make(map[string]string)
    if containerId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = strconv.Itoa(*containerId)

    queryParams := url.Values{}
    if scope != nil {
        queryParams.Add("scope", (*scope).String())
    }
    body, marshalErr := json.Marshal(*items)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.4", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []FileContainerItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes the specified items in a container.
// ctx
// containerId (required): Container Id.
// itemPath (required): Path to delete.
// scope (optional): A guid representing the scope of the container. This is often the project id.
func (client Client) DeleteItem(ctx context.Context, containerId *uint64, itemPath *string, scope *uuid.UUID) error {
    routeValues := make(map[string]string)
    if containerId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = strconv.FormatUint(*containerId, 10)

    queryParams := url.Values{}
    if itemPath == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "itemPath"}
    }
    queryParams.Add("itemPath", *itemPath)
    if scope != nil {
        queryParams.Add("scope", (*scope).String())
    }
    locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.4", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Gets containers filtered by a comma separated list of artifact uris within the same scope, if not specified returns all containers
// ctx
// scope (optional): A guid representing the scope of the container. This is often the project id.
// artifactUris (optional)
func (client Client) GetContainers(ctx context.Context, scope *uuid.UUID, artifactUris *string) (*[]FileContainer, error) {
    queryParams := url.Values{}
    if scope != nil {
        queryParams.Add("scope", (*scope).String())
    }
    if artifactUris != nil {
        queryParams.Add("artifactUris", *artifactUris)
    }
    locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.4", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []FileContainer
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// containerId (required)
// scope (optional)
// itemPath (optional)
// metadata (optional)
// format (optional)
// downloadFileName (optional)
// includeDownloadTickets (optional)
// isShallow (optional)
func (client Client) GetItems(ctx context.Context, containerId *uint64, scope *uuid.UUID, itemPath *string, metadata *bool, format *string, downloadFileName *string, includeDownloadTickets *bool, isShallow *bool) (*[]FileContainerItem, error) {
    routeValues := make(map[string]string)
    if containerId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "containerId"} 
    }
    routeValues["containerId"] = strconv.FormatUint(*containerId, 10)

    queryParams := url.Values{}
    if scope != nil {
        queryParams.Add("scope", (*scope).String())
    }
    if itemPath != nil {
        queryParams.Add("itemPath", *itemPath)
    }
    if metadata != nil {
        queryParams.Add("metadata", strconv.FormatBool(*metadata))
    }
    if format != nil {
        queryParams.Add("$format", *format)
    }
    if downloadFileName != nil {
        queryParams.Add("downloadFileName", *downloadFileName)
    }
    if includeDownloadTickets != nil {
        queryParams.Add("includeDownloadTickets", strconv.FormatBool(*includeDownloadTickets))
    }
    if isShallow != nil {
        queryParams.Add("isShallow", strconv.FormatBool(*isShallow))
    }
    locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.4", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []FileContainerItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

