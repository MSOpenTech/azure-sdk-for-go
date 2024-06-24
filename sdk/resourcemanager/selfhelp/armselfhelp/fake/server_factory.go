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

// ServerFactory is a fake server for instances of the armselfhelp.ClientFactory type.
type ServerFactory struct {
	CheckNameAvailabilityServer CheckNameAvailabilityServer
	DiagnosticsServer           DiagnosticsServer
	DiscoverySolutionServer     DiscoverySolutionServer
	DiscoverySolutionNLPServer  DiscoverySolutionNLPServer
	OperationsServer            OperationsServer
	SimplifiedSolutionsServer   SimplifiedSolutionsServer
	SolutionServer              SolutionServer
	SolutionSelfHelpServer      SolutionSelfHelpServer
	TroubleshootersServer       TroubleshootersServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armselfhelp.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armselfhelp.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                           *ServerFactory
	trMu                          sync.Mutex
	trCheckNameAvailabilityServer *CheckNameAvailabilityServerTransport
	trDiagnosticsServer           *DiagnosticsServerTransport
	trDiscoverySolutionServer     *DiscoverySolutionServerTransport
	trDiscoverySolutionNLPServer  *DiscoverySolutionNLPServerTransport
	trOperationsServer            *OperationsServerTransport
	trSimplifiedSolutionsServer   *SimplifiedSolutionsServerTransport
	trSolutionServer              *SolutionServerTransport
	trSolutionSelfHelpServer      *SolutionSelfHelpServerTransport
	trTroubleshootersServer       *TroubleshootersServerTransport
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
	case "CheckNameAvailabilityClient":
		initServer(s, &s.trCheckNameAvailabilityServer, func() *CheckNameAvailabilityServerTransport {
			return NewCheckNameAvailabilityServerTransport(&s.srv.CheckNameAvailabilityServer)
		})
		resp, err = s.trCheckNameAvailabilityServer.Do(req)
	case "DiagnosticsClient":
		initServer(s, &s.trDiagnosticsServer, func() *DiagnosticsServerTransport { return NewDiagnosticsServerTransport(&s.srv.DiagnosticsServer) })
		resp, err = s.trDiagnosticsServer.Do(req)
	case "DiscoverySolutionClient":
		initServer(s, &s.trDiscoverySolutionServer, func() *DiscoverySolutionServerTransport {
			return NewDiscoverySolutionServerTransport(&s.srv.DiscoverySolutionServer)
		})
		resp, err = s.trDiscoverySolutionServer.Do(req)
	case "DiscoverySolutionNLPClient":
		initServer(s, &s.trDiscoverySolutionNLPServer, func() *DiscoverySolutionNLPServerTransport {
			return NewDiscoverySolutionNLPServerTransport(&s.srv.DiscoverySolutionNLPServer)
		})
		resp, err = s.trDiscoverySolutionNLPServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SimplifiedSolutionsClient":
		initServer(s, &s.trSimplifiedSolutionsServer, func() *SimplifiedSolutionsServerTransport {
			return NewSimplifiedSolutionsServerTransport(&s.srv.SimplifiedSolutionsServer)
		})
		resp, err = s.trSimplifiedSolutionsServer.Do(req)
	case "SolutionClient":
		initServer(s, &s.trSolutionServer, func() *SolutionServerTransport { return NewSolutionServerTransport(&s.srv.SolutionServer) })
		resp, err = s.trSolutionServer.Do(req)
	case "SolutionSelfHelpClient":
		initServer(s, &s.trSolutionSelfHelpServer, func() *SolutionSelfHelpServerTransport {
			return NewSolutionSelfHelpServerTransport(&s.srv.SolutionSelfHelpServer)
		})
		resp, err = s.trSolutionSelfHelpServer.Do(req)
	case "TroubleshootersClient":
		initServer(s, &s.trTroubleshootersServer, func() *TroubleshootersServerTransport {
			return NewTroubleshootersServerTransport(&s.srv.TroubleshootersServer)
		})
		resp, err = s.trTroubleshootersServer.Do(req)
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
