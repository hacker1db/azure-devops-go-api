﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package pipelines

import (
    "time"
)

type ConfigurationType string

type CreatePipelineConfigurationParameters struct {
    Type *ConfigurationType `json:"type,omitempty"`
}

type CreatePipelineParameters struct {
    Configuration *CreatePipelineConfigurationParameters `json:"configuration,omitempty"`
    Folder *string `json:"folder,omitempty"`
    Name *string `json:"name,omitempty"`
}

// $expand options for GetLog and ListLogs.
type GetLogExpandOptions string

type Log struct {
    // The date and time the log was created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // The ID of the log.
    Id *int `json:"id,omitempty"`
    // The date and time the log was last changed.
    LastChangedOn *time.Time `json:"lastChangedOn,omitempty"`
    // The number of lines in the log.
    LineCount *uint64 `json:"lineCount,omitempty"`
    SignedContent *SignedUrl `json:"signedContent,omitempty"`
    Url *string `json:"url,omitempty"`
}

type LogCollection struct {
    Logs *[]Log `json:"logs,omitempty"`
    SignedContent *SignedUrl `json:"signedContent,omitempty"`
    Url *string `json:"url,omitempty"`
}

type Pipeline struct {
    Folder *string `json:"folder,omitempty"`
    Id *int `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Revision *int `json:"revision,omitempty"`
    Links *ReferenceLinks `json:"_links,omitempty"`
    Configuration *PipelineConfiguration `json:"configuration,omitempty"`
    Url *string `json:"url,omitempty"`
}

type PipelineBase struct {
    Folder *string `json:"folder,omitempty"`
    Id *int `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Revision *int `json:"revision,omitempty"`
}

type PipelineConfiguration struct {
    Type *ConfigurationType `json:"type,omitempty"`
}

// A reference to a Pipeline.
type PipelineReference struct {
    Folder *string `json:"folder,omitempty"`
    Id *int `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Revision *int `json:"revision,omitempty"`
    Url *string `json:"url,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

type Repository struct {
    Type *RepositoryType `json:"type,omitempty"`
}

type RepositoryResource struct {
    RefName *string `json:"refName,omitempty"`
    Repository *Repository `json:"repository,omitempty"`
    Version *string `json:"version,omitempty"`
}

type RepositoryResourceParameters struct {
    RefName *string `json:"refName,omitempty"`
    // This is the security token to use when connecting to the repository.
    Token *string `json:"token,omitempty"`
    // Optional. This is the type of the token given. If not provided, a type of "Bearer" is assumed. Note: Use "Basic" for a PAT token.
    TokenType *string `json:"tokenType,omitempty"`
    Version *string `json:"version,omitempty"`
}

type RepositoryType string

type Run struct {
    Id *int `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Links *ReferenceLinks `json:"_links,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    FinishedDate *time.Time `json:"finishedDate,omitempty"`
    Pipeline *PipelineReference `json:"pipeline,omitempty"`
    Resources *RunResources `json:"resources,omitempty"`
    Result *RunResult `json:"result,omitempty"`
    State *RunState `json:"state,omitempty"`
    Url *string `json:"url,omitempty"`
    Variables *map[string]Variable `json:"variables,omitempty"`
}

type RunPipelineParameters struct {
    // Deprecated: Use Context instead
    Resources *RunResourcesParameters `json:"resources,omitempty"`
    Secrets *map[string]string `json:"secrets,omitempty"`
    Variables *map[string]Variable `json:"variables,omitempty"`
}

type RunReference struct {
    Id *int `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}

type RunResources struct {
    Repositories *map[string]RepositoryResource `json:"repositories,omitempty"`
}

type RunResourcesParameters struct {
    Repositories *map[string]RepositoryResourceParameters `json:"repositories,omitempty"`
}

// This is not a Flags enum because we don't want to set multiple results on a build. However, when adding values, please stick to powers of 2 as if it were a Flags enum. This will make it easier to query multiple results.
type RunResult string

// This is not a Flags enum because we don't want to set multiple states on a build. However, when adding values, please stick to powers of 2 as if it were a Flags enum. This will make it easier to query multiple states.
type RunState string

// A signed url allowing limited-time anonymous access to private resources.
type SignedUrl struct {
    SignatureExpires *time.Time `json:"signatureExpires,omitempty"`
    Url *string `json:"url,omitempty"`
}

type Variable struct {
    IsSecret *bool `json:"isSecret,omitempty"`
    Value *string `json:"value,omitempty"`
}