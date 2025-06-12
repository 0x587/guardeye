// package: 
// file: main.proto

var main_pb = require("./main_pb");
var rcl_interfaces_pb = require("./rcl_interfaces_pb");
var shawn_define_pb = require("./shawn_define_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Api = (function () {
  function Api() {}
  Api.serviceName = "Api";
  return Api;
}());

Api.PublishTopicRosout = {
  methodName: "PublishTopicRosout",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.Log,
  responseType: rcl_interfaces_pb.Log
};

Api.SubscribeTopicRosout = {
  methodName: "SubscribeTopicRosout",
  service: Api,
  requestStream: false,
  responseStream: true,
  requestType: rcl_interfaces_pb.Log,
  responseType: rcl_interfaces_pb.Log
};

Api.PublishTopicV2_publisher = {
  methodName: "PublishTopicV2_publisher",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: shawn_define_pb.V2,
  responseType: shawn_define_pb.V2
};

Api.SubscribeTopicV2_publisher = {
  methodName: "SubscribeTopicV2_publisher",
  service: Api,
  requestStream: false,
  responseStream: true,
  requestType: shawn_define_pb.V2,
  responseType: shawn_define_pb.V2
};

Api.PublishTopicParameter_events = {
  methodName: "PublishTopicParameter_events",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ParameterEvent,
  responseType: rcl_interfaces_pb.ParameterEvent
};

Api.SubscribeTopicParameter_events = {
  methodName: "SubscribeTopicParameter_events",
  service: Api,
  requestStream: false,
  responseStream: true,
  requestType: rcl_interfaces_pb.ParameterEvent,
  responseType: rcl_interfaces_pb.ParameterEvent
};

Api.CallServiceAdd_two_ints_srv = {
  methodName: "CallServiceAdd_two_ints_srv",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: shawn_define_pb.AddReq,
  responseType: shawn_define_pb.AddRsp
};

Api.CallServiceFoxglove_bridgeDescribe_parameters = {
  methodName: "CallServiceFoxglove_bridgeDescribe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

Api.CallServiceFoxglove_bridgeGet_parameters = {
  methodName: "CallServiceFoxglove_bridgeGet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParametersReq,
  responseType: rcl_interfaces_pb.GetParametersRsp
};

Api.CallServiceFoxglove_bridgeSet_parameters_atomically = {
  methodName: "CallServiceFoxglove_bridgeSet_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceFoxglove_bridgeList_parameters = {
  methodName: "CallServiceFoxglove_bridgeList_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceService_server_02List_parameters = {
  methodName: "CallServiceService_server_02List_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceService_server_02Set_parameters = {
  methodName: "CallServiceService_server_02Set_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersReq,
  responseType: rcl_interfaces_pb.SetParametersRsp
};

Api.CallServiceService_server_02Get_parameter_types = {
  methodName: "CallServiceService_server_02Get_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceFoxglove_bridgeSet_parameters = {
  methodName: "CallServiceFoxglove_bridgeSet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersReq,
  responseType: rcl_interfaces_pb.SetParametersRsp
};

Api.CallServiceService_server_02Get_parameters = {
  methodName: "CallServiceService_server_02Get_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParametersReq,
  responseType: rcl_interfaces_pb.GetParametersRsp
};

Api.CallServiceService_server_02Set_parameters_atomically = {
  methodName: "CallServiceService_server_02Set_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceFoxglove_bridgeGet_parameter_types = {
  methodName: "CallServiceFoxglove_bridgeGet_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceService_server_02Describe_parameters = {
  methodName: "CallServiceService_server_02Describe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

exports.Api = Api;

function ApiClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ApiClient.prototype.publishTopicRosout = function publishTopicRosout(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.PublishTopicRosout, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.subscribeTopicRosout = function subscribeTopicRosout(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Api.SubscribeTopicRosout, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

ApiClient.prototype.publishTopicV2_publisher = function publishTopicV2_publisher(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.PublishTopicV2_publisher, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.subscribeTopicV2_publisher = function subscribeTopicV2_publisher(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Api.SubscribeTopicV2_publisher, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

ApiClient.prototype.publishTopicParameter_events = function publishTopicParameter_events(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.PublishTopicParameter_events, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.subscribeTopicParameter_events = function subscribeTopicParameter_events(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Api.SubscribeTopicParameter_events, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceAdd_two_ints_srv = function callServiceAdd_two_ints_srv(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceAdd_two_ints_srv, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceFoxglove_bridgeDescribe_parameters = function callServiceFoxglove_bridgeDescribe_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceFoxglove_bridgeDescribe_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceFoxglove_bridgeGet_parameters = function callServiceFoxglove_bridgeGet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceFoxglove_bridgeGet_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceFoxglove_bridgeSet_parameters_atomically = function callServiceFoxglove_bridgeSet_parameters_atomically(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceFoxglove_bridgeSet_parameters_atomically, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceFoxglove_bridgeList_parameters = function callServiceFoxglove_bridgeList_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceFoxglove_bridgeList_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceService_server_02List_parameters = function callServiceService_server_02List_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceService_server_02List_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceService_server_02Set_parameters = function callServiceService_server_02Set_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceService_server_02Set_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceService_server_02Get_parameter_types = function callServiceService_server_02Get_parameter_types(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceService_server_02Get_parameter_types, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceFoxglove_bridgeSet_parameters = function callServiceFoxglove_bridgeSet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceFoxglove_bridgeSet_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceService_server_02Get_parameters = function callServiceService_server_02Get_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceService_server_02Get_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceService_server_02Set_parameters_atomically = function callServiceService_server_02Set_parameters_atomically(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceService_server_02Set_parameters_atomically, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceFoxglove_bridgeGet_parameter_types = function callServiceFoxglove_bridgeGet_parameter_types(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceFoxglove_bridgeGet_parameter_types, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ApiClient.prototype.callServiceService_server_02Describe_parameters = function callServiceService_server_02Describe_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceService_server_02Describe_parameters, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ApiClient = ApiClient;

