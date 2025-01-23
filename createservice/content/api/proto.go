package api

const ProtoContent = `syntax = "proto3";

import "google/protobuf/timestamp.proto";
import "google/protobuf/empty.proto";

package <service-name>;

option go_package = "github.com/DKhorkov/hmtm-<service-name>/api/protobuf/<service-name>;<service-name>";

`
