/*

Copyright 2017 Continusec Pty Ltd

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

package api

import (
	"golang.org/x/net/context"

	"github.com/continusec/verifiabledatastructures/pb"
)

func (s *LocalService) LogFetchEntries(ctx context.Context, req *pb.LogFetchEntriesRequest) (*pb.LogFetchEntriesResponse, error) {
	err := s.verifyAccessForLogOperation(req.Log, operationReadEntry)
	if err != nil {
		return nil, err
	}

	if req.First < 0 || req.Last < 0 {
		return nil, ErrInvalidTreeRange
	}

	var rv *pb.LogFetchEntriesResponse
	err = s.Reader.ExecuteReadOnly(func(kr KeyReader) error {
		head, err := s.lookupLogTreeHead(kr, req.Log)
		if err != nil {
			return err
		}

		last := req.Last

		// Do we have it already?
		if last == 0 {
			last = head.TreeSize
		}

		// Are we asking for something silly?
		if last > head.TreeSize || req.First >= last {
			return ErrInvalidTreeRange
		}

		hashes, err := s.lookupLogEntryHashes(kr, req.Log, req.First, last)
		if err != nil {
			return err
		}

		vals := make([]*pb.LeafData, len(hashes))
		for i, h := range hashes {
			vals[i], err = s.lookupDataByLeafHash(kr, req.Log, h)
			if err != nil {
				return err
			}
		}

		rv = &pb.LogFetchEntriesResponse{
			Values: vals,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return rv, nil
}