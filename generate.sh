#!/bin/bash

protoc gamefeed/gamefeedpb/gamefeed.proto --go_out=plugins=grpc:.