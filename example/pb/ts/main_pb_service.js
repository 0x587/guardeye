// package: 
// file: main.proto

var main_pb = require("./main_pb");
var rcl_interfaces_pb = require("./rcl_interfaces_pb");
var shawn_define_pb = require("./shawn_define_pb");
var std_msgs_pb = require("./std_msgs_pb");
var pkg_d_pb = require("./pkg_d_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Api = (function () {
  function Api() {}
  Api.serviceName = "Api";
  return Api;
}());

Api.PublishTopicParameter_events = {
  methodName: "PublishTopicParameter_events",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ParameterEvent,
  responseType: rcl_interfaces_pb.ParameterEvent
};

Api.PublishTopicRosout = {
  methodName: "PublishTopicRosout",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.Log,
  responseType: rcl_interfaces_pb.Log
};

Api.PublishTopicChatter = {
  methodName: "PublishTopicChatter",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: std_msgs_pb.String,
  responseType: std_msgs_pb.String
};

Api.CallServiceListenerDescribe_parameters = {
  methodName: "CallServiceListenerDescribe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

Api.CallServiceListenerGet_parameters = {
  methodName: "CallServiceListenerGet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParametersReq,
  responseType: rcl_interfaces_pb.GetParametersRsp
};

Api.CallServiceService_server_02Set_parameters = {
  methodName: "CallServiceService_server_02Set_parameters",
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

Api.CallServiceService_server_02Get_parameter_types = {
  methodName: "CallServiceService_server_02Get_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceFoxglove_bridgeSet_parameters_atomically = {
  methodName: "CallServiceFoxglove_bridgeSet_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceService_server_02Describe_parameters = {
  methodName: "CallServiceService_server_02Describe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

Api.CallServiceListenerList_parameters = {
  methodName: "CallServiceListenerList_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceFoxglove_bridgeGet_parameter_types = {
  methodName: "CallServiceFoxglove_bridgeGet_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceAdd_two_ints_srv = {
  methodName: "CallServiceAdd_two_ints_srv",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: shawn_define_pb.AddReq,
  responseType: shawn_define_pb.AddRsp
};

Api.CallServiceGet_person = {
  methodName: "CallServiceGet_person",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: pkg_d_pb.GetPersonReq,
  responseType: pkg_d_pb.GetPersonRsp
};

Api.CallServiceListenerGet_parameter_types = {
  methodName: "CallServiceListenerGet_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceFoxglove_bridgeList_parameters = {
  methodName: "CallServiceFoxglove_bridgeList_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceListenerSet_parameters_atomically = {
  methodName: "CallServiceListenerSet_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceFoxglove_bridgeSet_parameters = {
  methodName: "CallServiceFoxglove_bridgeSet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersReq,
  responseType: rcl_interfaces_pb.SetParametersRsp
};

Api.CallServiceService_server_02List_parameters = {
  methodName: "CallServiceService_server_02List_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceFoxglove_bridgeDescribe_parameters = {
  methodName: "CallServiceFoxglove_bridgeDescribe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

Api.CallServiceService_server_02Set_parameters_atomically = {
  methodName: "CallServiceService_server_02Set_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceFoxglove_bridgeGet_parameters = {
  methodName: "CallServiceFoxglove_bridgeGet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParametersReq,
  responseType: rcl_interfaces_pb.GetParametersRsp
};

Api.CallServiceListenerSet_parameters = {
  methodName: "CallServiceListenerSet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersReq,
  responseType: rcl_interfaces_pb.SetParametersRsp
};

exports.Api = Api;

function ApiClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

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

ApiClient.prototype.publishTopicChatter = function publishTopicChatter(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.PublishTopicChatter, {
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

ApiClient.prototype.callServiceListenerDescribe_parameters = function callServiceListenerDescribe_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceListenerDescribe_parameters, {
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

ApiClient.prototype.callServiceListenerGet_parameters = function callServiceListenerGet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceListenerGet_parameters, {
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

ApiClient.prototype.callServiceListenerList_parameters = function callServiceListenerList_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceListenerList_parameters, {
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

ApiClient.prototype.callServiceGet_person = function callServiceGet_person(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceGet_person, {
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

ApiClient.prototype.callServiceListenerGet_parameter_types = function callServiceListenerGet_parameter_types(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceListenerGet_parameter_types, {
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

ApiClient.prototype.callServiceListenerSet_parameters_atomically = function callServiceListenerSet_parameters_atomically(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceListenerSet_parameters_atomically, {
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

ApiClient.prototype.callServiceListenerSet_parameters = function callServiceListenerSet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceListenerSet_parameters, {
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

