#!/bin/bash

# To regenerate this script:
# find . | grep _test.go | xargs grep -n Test | grep -iv // | sed s/.*func/go\ test\ -run/ | sed s/\*testing.T\)\ {// | sed s/\(//

go test -run TestGetAdgroupCriterion
go test -run TestSetMaxCpm
go test -run TestAddRemoveAdgroupCriterion
go test -run TestGetAdgroup
go test -run TestAddRemoveAdgroups
go test -run TestGetBulkMutateJob
go test -run TestAddRemoveBulkMutateJob
go test -run TestGetCampaign
go test -run TestDecodingSoapFault
go test -run TestServicedAccountGet