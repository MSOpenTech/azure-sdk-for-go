//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armstoragemover.ClientFactory type.
type ServerFactory struct {
	AgentsServer         AgentsServer
	EndpointsServer      EndpointsServer
	JobDefinitionsServer JobDefinitionsServer
	JobRunsServer        JobRunsServer
	OperationsServer     OperationsServer
	ProjectsServer       ProjectsServer
	StorageMoversServer  StorageMoversServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armstoragemover.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armstoragemover.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                    *ServerFactory
	trMu                   sync.Mutex
	trAgentsServer         *AgentsServerTransport
	trEndpointsServer      *EndpointsServerTransport
	trJobDefinitionsServer *JobDefinitionsServerTransport
	trJobRunsServer        *JobRunsServerTransport
	trOperationsServer     *OperationsServerTransport
	trProjectsServer       *ProjectsServerTransport
	trStorageMoversServer  *StorageMoversServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AgentsClient":
		initServer(s, &s.trAgentsServer, func() *AgentsServerTransport { return NewAgentsServerTransport(&s.srv.AgentsServer) })
		resp, err = s.trAgentsServer.Do(req)
	case "EndpointsClient":
		initServer(s, &s.trEndpointsServer, func() *EndpointsServerTransport { return NewEndpointsServerTransport(&s.srv.EndpointsServer) })
		resp, err = s.trEndpointsServer.Do(req)
	case "JobDefinitionsClient":
		initServer(s, &s.trJobDefinitionsServer, func() *JobDefinitionsServerTransport {
			return NewJobDefinitionsServerTransport(&s.srv.JobDefinitionsServer)
		})
		resp, err = s.trJobDefinitionsServer.Do(req)
	case "JobRunsClient":
		initServer(s, &s.trJobRunsServer, func() *JobRunsServerTransport { return NewJobRunsServerTransport(&s.srv.JobRunsServer) })
		resp, err = s.trJobRunsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ProjectsClient":
		initServer(s, &s.trProjectsServer, func() *ProjectsServerTransport { return NewProjectsServerTransport(&s.srv.ProjectsServer) })
		resp, err = s.trProjectsServer.Do(req)
	case "StorageMoversClient":
		initServer(s, &s.trStorageMoversServer, func() *StorageMoversServerTransport {
			return NewStorageMoversServerTransport(&s.srv.StorageMoversServer)
		})
		resp, err = s.trStorageMoversServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
