// package: 
// file: main.proto

import * as main_pb from "./main_pb";
import * as rcl_interfaces_pb from "./rcl_interfaces_pb";
import * as shawn_define_pb from "./shawn_define_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ApiPublishTopicRosout = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.Log;
  readonly responseType: typeof rcl_interfaces_pb.Log;
};

type ApiSubscribeTopicRosout = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof rcl_interfaces_pb.Log;
  readonly responseType: typeof rcl_interfaces_pb.Log;
};

type ApiPublishTopicV2_publisher = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shawn_define_pb.V2;
  readonly responseType: typeof shawn_define_pb.V2;
};

type ApiSubscribeTopicV2_publisher = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof shawn_define_pb.V2;
  readonly responseType: typeof shawn_define_pb.V2;
};

type ApiPublishTopicParameter_events = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ParameterEvent;
  readonly responseType: typeof rcl_interfaces_pb.ParameterEvent;
};

type ApiSubscribeTopicParameter_events = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof rcl_interfaces_pb.ParameterEvent;
  readonly responseType: typeof rcl_interfaces_pb.ParameterEvent;
};

type ApiCallServiceAdd_two_ints_srv = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shawn_define_pb.AddReq;
  readonly responseType: typeof shawn_define_pb.AddRsp;
};

type ApiCallServiceFoxglove_bridgeDescribe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

type ApiCallServiceFoxglove_bridgeGet_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParametersRsp;
};

type ApiCallServiceFoxglove_bridgeSet_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceFoxglove_bridgeList_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceService_server_02List_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.ListParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.ListParametersRsp;
};

type ApiCallServiceService_server_02Set_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersRsp;
};

type ApiCallServiceService_server_02Get_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceFoxglove_bridgeSet_parameters = {
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

type ApiCallServiceService_server_02Set_parameters_atomically = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.SetParametersAtomicallyReq;
  readonly responseType: typeof rcl_interfaces_pb.SetParametersAtomicallyRsp;
};

type ApiCallServiceFoxglove_bridgeGet_parameter_types = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.GetParameterTypesReq;
  readonly responseType: typeof rcl_interfaces_pb.GetParameterTypesRsp;
};

type ApiCallServiceService_server_02Describe_parameters = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof rcl_interfaces_pb.DescribeParametersReq;
  readonly responseType: typeof rcl_interfaces_pb.DescribeParametersRsp;
};

export class Api {
  static readonly serviceName: string;
  static readonly PublishTopicRosout: ApiPublishTopicRosout;
  static readonly SubscribeTopicRosout: ApiSubscribeTopicRosout;
  static readonly PublishTopicV2_publisher: ApiPublishTopicV2_publisher;
  static readonly SubscribeTopicV2_publisher: ApiSubscribeTopicV2_publisher;
  static readonly PublishTopicParameter_events: ApiPublishTopicParameter_events;
  static readonly SubscribeTopicParameter_events: ApiSubscribeTopicParameter_events;
  static readonly CallServiceAdd_two_ints_srv: ApiCallServiceAdd_two_ints_srv;
  static readonly CallServiceFoxglove_bridgeDescribe_parameters: ApiCallServiceFoxglove_bridgeDescribe_parameters;
  static readonly CallServiceFoxglove_bridgeGet_parameters: ApiCallServiceFoxglove_bridgeGet_parameters;
  static readonly CallServiceFoxglove_bridgeSet_parameters_atomically: ApiCallServiceFoxglove_bridgeSet_parameters_atomically;
  static readonly CallServiceFoxglove_bridgeList_parameters: ApiCallServiceFoxglove_bridgeList_parameters;
  static readonly CallServiceService_server_02List_parameters: ApiCallServiceService_server_02List_parameters;
  static readonly CallServiceService_server_02Set_parameters: ApiCallServiceService_server_02Set_parameters;
  static readonly CallServiceService_server_02Get_parameter_types: ApiCallServiceService_server_02Get_parameter_types;
  static readonly CallServiceFoxglove_bridgeSet_parameters: ApiCallServiceFoxglove_bridgeSet_parameters;
  static readonly CallServiceService_server_02Get_parameters: ApiCallServiceService_server_02Get_parameters;
  static readonly CallServiceService_server_02Set_parameters_atomically: ApiCallServiceService_server_02Set_parameters_atomically;
  static readonly CallServiceFoxglove_bridgeGet_parameter_types: ApiCallServiceFoxglove_bridgeGet_parameter_types;
  static readonly CallServiceService_server_02Describe_parameters: ApiCallServiceService_server_02Describe_parameters;
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
  publishTopicRosout(
    requestMessage: rcl_interfaces_pb.Log,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.Log|null) => void
  ): UnaryResponse;
  publishTopicRosout(
    requestMessage: rcl_interfaces_pb.Log,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.Log|null) => void
  ): UnaryResponse;
  subscribeTopicRosout(requestMessage: rcl_interfaces_pb.Log, metadata?: grpc.Metadata): ResponseStream<rcl_interfaces_pb.Log>;
  publishTopicV2_publisher(
    requestMessage: shawn_define_pb.V2,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: shawn_define_pb.V2|null) => void
  ): UnaryResponse;
  publishTopicV2_publisher(
    requestMessage: shawn_define_pb.V2,
    callback: (error: ServiceError|null, responseMessage: shawn_define_pb.V2|null) => void
  ): UnaryResponse;
  subscribeTopicV2_publisher(requestMessage: shawn_define_pb.V2, metadata?: grpc.Metadata): ResponseStream<shawn_define_pb.V2>;
  publishTopicParameter_events(
    requestMessage: rcl_interfaces_pb.ParameterEvent,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ParameterEvent|null) => void
  ): UnaryResponse;
  publishTopicParameter_events(
    requestMessage: rcl_interfaces_pb.ParameterEvent,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ParameterEvent|null) => void
  ): UnaryResponse;
  subscribeTopicParameter_events(requestMessage: rcl_interfaces_pb.ParameterEvent, metadata?: grpc.Metadata): ResponseStream<rcl_interfaces_pb.ParameterEvent>;
  callServiceAdd_two_ints_srv(
    requestMessage: shawn_define_pb.AddReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: shawn_define_pb.AddRsp|null) => void
  ): UnaryResponse;
  callServiceAdd_two_ints_srv(
    requestMessage: shawn_define_pb.AddReq,
    callback: (error: ServiceError|null, responseMessage: shawn_define_pb.AddRsp|null) => void
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
  callServiceFoxglove_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeGet_parameters(
    requestMessage: rcl_interfaces_pb.GetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.GetParametersRsp|null) => void
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
  callServiceFoxglove_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeList_parameters(
    requestMessage: rcl_interfaces_pb.ListParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.ListParametersRsp|null) => void
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
  callServiceService_server_02Set_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Set_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
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
  callServiceFoxglove_bridgeSet_parameters(
    requestMessage: rcl_interfaces_pb.SetParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersRsp|null) => void
  ): UnaryResponse;
  callServiceFoxglove_bridgeSet_parameters(
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
  callServiceService_server_02Set_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Set_parameters_atomically(
    requestMessage: rcl_interfaces_pb.SetParametersAtomicallyReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.SetParametersAtomicallyRsp|null) => void
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
  callServiceService_server_02Describe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
  callServiceService_server_02Describe_parameters(
    requestMessage: rcl_interfaces_pb.DescribeParametersReq,
    callback: (error: ServiceError|null, responseMessage: rcl_interfaces_pb.DescribeParametersRsp|null) => void
  ): UnaryResponse;
}

