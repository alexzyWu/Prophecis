/*
 * Copyright 2017-2018 IBM Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

// JobStatusClient is a client interface for updating the status of training jobs
type JobStatusClient interface {
	UpdateJobStatus()
}

// jobStatusClientRPC implements job status updates by calling the gRPC methods of the storage service
type jobStatusClientRPC struct {
}

// NewStorageClient ...
func NewStorageClient() JobStatusClient {
	return &jobStatusClientRPC{}
}

func (*jobStatusClientRPC) UpdateJobStatus() {
	// TODO
}
