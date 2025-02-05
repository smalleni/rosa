/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/statusboard/v1

import (
	"io"
	"net/http"
)

func readServiceDependencyDeleteRequest(request *ServiceDependencyDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeServiceDependencyDeleteRequest(request *ServiceDependencyDeleteRequest, writer io.Writer) error {
	return nil
}
func readServiceDependencyDeleteResponse(response *ServiceDependencyDeleteResponse, reader io.Reader) error {
	return nil
}
func writeServiceDependencyDeleteResponse(response *ServiceDependencyDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
func readServiceDependencyGetRequest(request *ServiceDependencyGetServerRequest, r *http.Request) error {
	return nil
}
func writeServiceDependencyGetRequest(request *ServiceDependencyGetRequest, writer io.Writer) error {
	return nil
}
func readServiceDependencyGetResponse(response *ServiceDependencyGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalServiceDependency(reader)
	return err
}
func writeServiceDependencyGetResponse(response *ServiceDependencyGetServerResponse, w http.ResponseWriter) error {
	return MarshalServiceDependency(response.body, w)
}
func readServiceDependencyUpdateRequest(request *ServiceDependencyUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalServiceDependency(r.Body)
	return err
}
func writeServiceDependencyUpdateRequest(request *ServiceDependencyUpdateRequest, writer io.Writer) error {
	return MarshalServiceDependency(request.body, writer)
}
func readServiceDependencyUpdateResponse(response *ServiceDependencyUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalServiceDependency(reader)
	return err
}
func writeServiceDependencyUpdateResponse(response *ServiceDependencyUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalServiceDependency(response.body, w)
}
