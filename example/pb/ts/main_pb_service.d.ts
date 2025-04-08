// package: 
// file: main.proto

import * as main_pb from "./main_pb";
import * as rcl_interfaces_pb from "./rcl_interfaces_pb";
import * as shawn_define_pb from "./shawn_define_pb";
import * as std_msgs_pb from "./std_msgs_pb";
import * as pkg_d_pb from "./pkg_d_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ApiPublishTopicParameter_events = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ParameterEvent;
  readonly responseType: typeof rcl_interfaces_pb.ParameterEvent;
};

type ApiPublishTopicRosout = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.Log;
  readonly responseType: typeof rcl_interfaces_pb.Log;
};

type ApiPublishTopicChatter = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof std_msgs_pb.String;
  readonly responseType: typeof std_msgs_pb.String;
};

type ApiCallServiceListenerDescribe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

type ApiCallServiceListenerGet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParametersRsp;
};

type ApiCallServiceService_server_02Set_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersRsp;
};

type ApiCallServiceService_server_02Get_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParametersRsp;
};

type ApiCallServiceService_server_02Get_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceFoxglove_bridgeSet_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceService_server_02Describe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

type ApiCallServiceListenerList_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceFoxglove_bridgeGet_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceAdd_two_ints_srv = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shawn_define_pb.AddReq;
  readonly responseType: typeof shawn_define_pb.AddRsp;
};

type ApiCallServiceGet_person = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof pkg_d_pb.GetPersonReq;
  readonly responseType: typeof pkg_d_pb.GetPersonRsp;
};

type ApiCallServiceListenerGet_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceFoxglove_bridgeList_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceListenerSet_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceFoxglove_bridgeSet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersRsp;
};

type ApiCallServiceService_server_02List_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceFoxglove_bridgeDescribe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

type ApiCallServiceService_server_02Set_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceFoxglove_bridgeGet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParametersRsp;
};

type ApiCallServiceListenerSet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersRsp;
};

export class Api {
  static readonly serviceName: string;
  static readonly PublishTopicParameter_events: ApiPublishTopicParameter_events;
  static readonly PublishTopicRosout: ApiPublishTopicRosout;
  static readonly PublishTopicChatter: ApiPublishTopicChatter;
  static readonly CallServiceListenerDescribe_parameters: ApiCallServiceListenerDescribe_parameters;
  static readonly CallServiceListenerGet_parameters: ApiCallServiceListenerGet_parameters;
  static readonly CallServiceService_server_02Set_parameters: ApiCallServiceService_server_02Set_parameters;
  static readonly CallServiceService_server_02Get_parameters: ApiCallServiceService_server_02Get_parameters;
  static readonly CallServiceService_server_02Get_parameter_types: ApiCallServiceService_server_02Get_parameter_types;
  static readonly CallServiceFoxglove_bridgeSet_parameters_atomically: ApiCallServiceFoxglove_bridgeSet_parameters_atomically;
  static readonly CallServiceService_server_02Describe_parameters: ApiCallServiceService_server_02Describe_parameters;
  static readonly CallServiceListenerList_parameters: ApiCallServiceListenerList_parameters;
  static readonly CallServiceFoxglove_bridgeGet_parameter_types: ApiCallServiceFoxglove_bridgeGet_parameter_types;
  static readonly CallServiceAdd_two_ints_srv: ApiCallServiceAdd_two_ints_srv;
  static readonly CallServiceGet_person: ApiCallServiceGet_person;
  static readonly CallServiceListenerGet_parameter_types: ApiCallServiceListenerGet_parameter_types;
  static readonly CallServiceFoxglove_bridgeList_parameters: ApiCallServiceFoxglove_bridgeList_parameters;
  static readonly CallServiceListenerSet_parameters_atomically: ApiCallServiceListenerSet_parameters_atomically;
  static readonly CallServiceFoxglove_bridgeSet_parameters: ApiCallServiceFoxglove_bridgeSet_parameters;
  static readonly CallServiceService_server_02List_parameters: ApiCallServiceService_server_02List_parameters;
  static readonly CallServiceFoxglove_bridgeDescribe_parameters: ApiCallServiceFoxglove_bridgeDescribe_parameters;
  static readonly CallServiceService_server_02Set_parameters_atomically: ApiCallServiceService_server_02Set_parameters_atomically;
  static readonly CallServiceFoxglove_bridgeGet_parameters: ApiCallServiceFoxglove_bridgeGet_parameters;
  static readonly CallServiceListenerSet_parameters: ApiCallServiceListenerSet_parameters;
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
  publishTopicParameter_events(
    requestMessage: rcl_interfaces_pb.ParameterEvent,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ParameterEvent|null) => void
  ): UnaryResponse;
  publishTopicParameter_events(
    requestMessage: rcl_interfaces_pb.ParameterEvent,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ParameterEvent|null) => void
  ): UnaryResponse;
  publishTopicRosout(
    requestMessage: rcl_interfaces_pb.Log,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.Log|null) => void
  ): UnaryResponse;
  publishTopicRosout(
    requestMessage: rcl_interfaces_pb.Log,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.Log|null) => void
  ): UnaryResponse;
  publishTopicChatter(
    requestMessage: std_msgs_pb.String,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: std_msgs_pb.String|null) => void
  ): UnaryResponse;
  publishTopicChatter(
    requestMessage: std_msgs_pb.String,
    callback: (error: ServiceError|null, responseMessage: std_msgs_pb.String|null) => void
  ): UnaryResponse;
  callServiceListenerDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Set_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Set_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Get_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Get_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Get_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Get_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Describe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Describe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceAdd_two_ints_srv(
    requestMessage: shawn_define_pb.AddReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: shawn_define_pb.AddRsp|null) => void
  ): UnaryResponse;
  callServiceAdd_two_ints_srv(
    requestMessage: shawn_define_pb.AddReq,
    callback: (error: ServiceError|null, responseMessage: shawn_define_pb.AddRsp|null) => void
  ): UnaryResponse;
  callServiceGet_person(
    requestMessage: pkg_d_pb.GetPersonReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: pkg_d_pb.GetPersonRsp|null) => void
  ): UnaryResponse;
  callServiceGet_person(
    requestMessage: pkg_d_pb.GetPersonReq,
    callback: (error: ServiceError|null, responseMessage: pkg_d_pb.GetPersonRsp|null) => void
  ): UnaryResponse;
  callServiceListenerGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceListenerGet_parameter_types(
    requestMessage: rcl_interfaces_pb.GetParameterTypesReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParameterTypesRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceListenerSet_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02List_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02List_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeDescribe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Set_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Set_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceListenerSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
}

