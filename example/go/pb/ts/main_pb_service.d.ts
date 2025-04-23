// package: 
// file: main.proto

import * as main_pb from "./main_pb";
import * as rcl_interfaces_pb from "./rcl_interfaces_pb";
import * as robot_bridge_interfaces_pb from "./robot_bridge_interfaces_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ApiPublishTopicArebot_transportRobot_bridgeRobot_state = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.RobotState;
  readonly responseType: typeof robot_bridge_interfaces_pb.RobotState;
};

type ApiPublishTopicArebot_loadRobot_bridgeRobot_state = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.RobotState;
  readonly responseType: typeof robot_bridge_interfaces_pb.RobotState;
};

type ApiPublishTopicArebot_unloadRobot_bridgeRobot_state = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.RobotState;
  readonly responseType: typeof robot_bridge_interfaces_pb.RobotState;
};

type ApiCallServiceArebot_unloadRobot_bridgeGet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParametersRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeGet_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeSet_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeSet_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeGet_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeGet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParametersRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeDescribe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeList_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeStart_nav_to_pose = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.StartNavToPoseReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.StartNavToPoseRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeSet_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeGet_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeGet_state = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.GetStateReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.GetStateRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeSet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeStart_nav_to_pose = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.StartNavToPoseReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.StartNavToPoseRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeGet_state = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.GetStateReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.GetStateRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeList_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeSet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeStart_nav_to_pose = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.StartNavToPoseReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.StartNavToPoseRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeDescribe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeList_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeStart_pick_and_load = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.StartPickAndLoadReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.StartPickAndLoadRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeGet_state = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.GetStateReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.GetStateRsp;
};

type ApiCallServiceArebot_unloadRobot_bridgeStart_unload_and_place = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.StartUnloadAndPlaceReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.StartUnloadAndPlaceRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeDescribe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeGet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParametersRsp;
};

type ApiCallServiceArebot_transportRobot_bridgeStart_transport = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof robot_bridge_interfaces_pb.StartTransportReq;
  readonly responseType: typeof robot_bridge_interfaces_pb.StartTransportRsp;
};

type ApiCallServiceArebot_loadRobot_bridgeSet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersRsp;
};

export class Api {
  static readonly serviceName: string;
  static readonly PublishTopicArebot_transportRobot_bridgeRobot_state: ApiPublishTopicArebot_transportRobot_bridgeRobot_state;
  static readonly PublishTopicArebot_loadRobot_bridgeRobot_state: ApiPublishTopicArebot_loadRobot_bridgeRobot_state;
  static readonly PublishTopicArebot_unloadRobot_bridgeRobot_state: ApiPublishTopicArebot_unloadRobot_bridgeRobot_state;
  static readonly CallServiceArebot_unloadRobot_bridgeGet_parameters: ApiCallServiceArebot_unloadRobot_bridgeGet_parameters;
  static readonly CallServiceArebot_transportRobot_bridgeGet_parameter_types: ApiCallServiceArebot_transportRobot_bridgeGet_parameter_types;
  static readonly CallServiceArebot_unloadRobot_bridgeSet_parameters_atomically: ApiCallServiceArebot_unloadRobot_bridgeSet_parameters_atomically;
  static readonly CallServiceArebot_transportRobot_bridgeSet_parameters_atomically: ApiCallServiceArebot_transportRobot_bridgeSet_parameters_atomically;
  static readonly CallServiceArebot_loadRobot_bridgeGet_parameter_types: ApiCallServiceArebot_loadRobot_bridgeGet_parameter_types;
  static readonly CallServiceArebot_loadRobot_bridgeGet_parameters: ApiCallServiceArebot_loadRobot_bridgeGet_parameters;
  static readonly CallServiceArebot_unloadRobot_bridgeDescribe_parameters: ApiCallServiceArebot_unloadRobot_bridgeDescribe_parameters;
  static readonly CallServiceArebot_transportRobot_bridgeList_parameters: ApiCallServiceArebot_transportRobot_bridgeList_parameters;
  static readonly CallServiceArebot_loadRobot_bridgeStart_nav_to_pose: ApiCallServiceArebot_loadRobot_bridgeStart_nav_to_pose;
  static readonly CallServiceArebot_loadRobot_bridgeSet_parameters_atomically: ApiCallServiceArebot_loadRobot_bridgeSet_parameters_atomically;
  static readonly CallServiceArebot_unloadRobot_bridgeGet_parameter_types: ApiCallServiceArebot_unloadRobot_bridgeGet_parameter_types;
  static readonly CallServiceArebot_transportRobot_bridgeGet_state: ApiCallServiceArebot_transportRobot_bridgeGet_state;
  static readonly CallServiceArebot_transportRobot_bridgeSet_parameters: ApiCallServiceArebot_transportRobot_bridgeSet_parameters;
  static readonly CallServiceArebot_unloadRobot_bridgeStart_nav_to_pose: ApiCallServiceArebot_unloadRobot_bridgeStart_nav_to_pose;
  static readonly CallServiceArebot_loadRobot_bridgeGet_state: ApiCallServiceArebot_loadRobot_bridgeGet_state;
  static readonly CallServiceArebot_loadRobot_bridgeList_parameters: ApiCallServiceArebot_loadRobot_bridgeList_parameters;
  static readonly CallServiceArebot_unloadRobot_bridgeSet_parameters: ApiCallServiceArebot_unloadRobot_bridgeSet_parameters;
  static readonly CallServiceArebot_transportRobot_bridgeStart_nav_to_pose: ApiCallServiceArebot_transportRobot_bridgeStart_nav_to_pose;
  static readonly CallServiceArebot_loadRobot_bridgeDescribe_parameters: ApiCallServiceArebot_loadRobot_bridgeDescribe_parameters;
  static readonly CallServiceArebot_unloadRobot_bridgeList_parameters: ApiCallServiceArebot_unloadRobot_bridgeList_parameters;
  static readonly CallServiceArebot_loadRobot_bridgeStart_pick_and_load: ApiCallServiceArebot_loadRobot_bridgeStart_pick_and_load;
  static readonly CallServiceArebot_unloadRobot_bridgeGet_state: ApiCallServiceArebot_unloadRobot_bridgeGet_state;
  static readonly CallServiceArebot_unloadRobot_bridgeStart_unload_and_place: ApiCallServiceArebot_unloadRobot_bridgeStart_unload_and_place;
  static readonly CallServiceArebot_transportRobot_bridgeDescribe_parameters: ApiCallServiceArebot_transportRobot_bridgeDescribe_parameters;
  static readonly CallServiceArebot_transportRobot_bridgeGet_parameters: ApiCallServiceArebot_transportRobot_bridgeGet_parameters;
  static readonly CallServiceArebot_transportRobot_bridgeStart_transport: ApiCallServiceArebot_transportRobot_bridgeStart_transport;
  static readonly CallServiceArebot_loadRobot_bridgeSet_parameters: ApiCallServiceArebot_loadRobot_bridgeSet_parameters;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ApiClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  publishTopicArebot_transportRobot_bridgeRobot_state(
    requestMessage: robot_bridge_interfaces_pb.RobotState,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.RobotState|null) => void
  ): UnaryResponse;
  publishTopicArebot_transportRobot_bridgeRobot_state(
    requestMessage: robot_bridge_interfaces_pb.RobotState,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.RobotState|null) => void
  ): UnaryResponse;
  publishTopicArebot_loadRobot_bridgeRobot_state(
    requestMessage: robot_bridge_interfaces_pb.RobotState,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.RobotState|null) => void
  ): UnaryResponse;
  publishTopicArebot_loadRobot_bridgeRobot_state(
    requestMessage: robot_bridge_interfaces_pb.RobotState,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.RobotState|null) => void
  ): UnaryResponse;
  publishTopicArebot_unloadRobot_bridgeRobot_state(
    requestMessage: robot_bridge_interfaces_pb.RobotState,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.RobotState|null) => void
  ): UnaryResponse;
  publishTopicArebot_unloadRobot_bridgeRobot_state(
    requestMessage: robot_bridge_interfaces_pb.RobotState,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.RobotState|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeStart_nav_to_pose(
    requestMessage: robot_bridge_interfaces_pb.StartNavToPoseReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartNavToPoseRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeStart_nav_to_pose(
    requestMessage: robot_bridge_interfaces_pb.StartNavToPoseReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartNavToPoseRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeGet_state(
    requestMessage: robot_bridge_interfaces_pb.GetStateReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.GetStateRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeGet_state(
    requestMessage: robot_bridge_interfaces_pb.GetStateReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.GetStateRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeStart_nav_to_pose(
    requestMessage: robot_bridge_interfaces_pb.StartNavToPoseReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartNavToPoseRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeStart_nav_to_pose(
    requestMessage: robot_bridge_interfaces_pb.StartNavToPoseReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartNavToPoseRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeGet_state(
    requestMessage: robot_bridge_interfaces_pb.GetStateReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.GetStateRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeGet_state(
    requestMessage: robot_bridge_interfaces_pb.GetStateReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.GetStateRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeStart_nav_to_pose(
    requestMessage: robot_bridge_interfaces_pb.StartNavToPoseReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartNavToPoseRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeStart_nav_to_pose(
    requestMessage: robot_bridge_interfaces_pb.StartNavToPoseReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartNavToPoseRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeStart_pick_and_load(
    requestMessage: robot_bridge_interfaces_pb.StartPickAndLoadReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartPickAndLoadRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeStart_pick_and_load(
    requestMessage: robot_bridge_interfaces_pb.StartPickAndLoadReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartPickAndLoadRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeGet_state(
    requestMessage: robot_bridge_interfaces_pb.GetStateReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.GetStateRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeGet_state(
    requestMessage: robot_bridge_interfaces_pb.GetStateReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.GetStateRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeStart_unload_and_place(
    requestMessage: robot_bridge_interfaces_pb.StartUnloadAndPlaceReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartUnloadAndPlaceRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_unloadRobot_bridgeStart_unload_and_place(
    requestMessage: robot_bridge_interfaces_pb.StartUnloadAndPlaceReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartUnloadAndPlaceRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeStart_transport(
    requestMessage: robot_bridge_interfaces_pb.StartTransportReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartTransportRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_transportRobot_bridgeStart_transport(
    requestMessage: robot_bridge_interfaces_pb.StartTransportReq,
    callback: (error: ServiceError|null, responseMessage: robot_bridge_interfaces_pb.StartTransportRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceArebot_loadRobot_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
}

