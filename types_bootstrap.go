// Copyright 2019 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

package mach

/*
#include <bootstrap.h>
*/
import "C"

type BoolArrayT = C.bool_array_t

type BootstrapPropertyArrayT = C.bootstrap_property_array_t

type BootstrapPropertyT = C.bootstrap_property_t

type BootstrapStatusArrayT = C.bootstrap_status_array_t

type BootstrapStatusT = C.bootstrap_status_t

type CmdT = C.cmd_t

type NameArrayT = C.name_array_t

type NameT = C.name_t

type BootstrapPort = C.mach_port_t

const BOOTSTRAP_BAD_COUNT = C.BOOTSTRAP_BAD_COUNT

const BOOTSTRAP_MAX_CMD_LEN = C.BOOTSTRAP_MAX_CMD_LEN

const BOOTSTRAP_MAX_LOOKUP_COUNT = C.BOOTSTRAP_MAX_LOOKUP_COUNT

const BOOTSTRAP_MAX_NAME_LEN = C.BOOTSTRAP_MAX_NAME_LEN

const BOOTSTRAP_NAME_IN_USE = C.BOOTSTRAP_NAME_IN_USE

const BOOTSTRAP_NOT_PRIVILEGED = C.BOOTSTRAP_NOT_PRIVILEGED

const BOOTSTRAP_NO_CHILDREN = C.BOOTSTRAP_NO_CHILDREN

const BOOTSTRAP_NO_MEMORY = C.BOOTSTRAP_NO_MEMORY

const BOOTSTRAP_SERVICE_ACTIVE = C.BOOTSTRAP_SERVICE_ACTIVE

const BOOTSTRAP_STATUS_ACTIVE = C.BOOTSTRAP_STATUS_ACTIVE

const BOOTSTRAP_STATUS_INACTIVE = C.BOOTSTRAP_STATUS_INACTIVE

const BOOTSTRAP_STATUS_ON_DEMAND = C.BOOTSTRAP_STATUS_ON_DEMAND

const BOOTSTRAP_SUCCESS = C.BOOTSTRAP_SUCCESS

const BOOTSTRAP_UNKNOWN_SERVICE = C.BOOTSTRAP_UNKNOWN_SERVICE
