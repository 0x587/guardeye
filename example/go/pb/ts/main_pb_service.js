// package: 
// file: main.proto

var main_pb = require("./main_pb");
var rcl_interfaces_pb = require("./rcl_interfaces_pb");
var robot_bridge_interfaces_pb = require("./robot_bridge_interfaces_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Api = (function () {
  function Api() {}
  Api.serviceName = "Api";
  return Api;
}());

Api.PublishTopicArebot_transportRobot_bridgeRobot_state = {
  methodName: "PublishTopicArebot_transportRobot_bridgeRobot_state",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.RobotState,
  responseType: robot_bridge_interfaces_pb.RobotState
};

Api.PublishTopicArebot_loadRobot_bridgeRobot_state = {
  methodName: "PublishTopicArebot_loadRobot_bridgeRobot_state",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.RobotState,
  responseType: robot_bridge_interfaces_pb.RobotState
};

Api.PublishTopicArebot_unloadRobot_bridgeRobot_state = {
  methodName: "PublishTopicArebot_unloadRobot_bridgeRobot_state",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.RobotState,
  responseType: robot_bridge_interfaces_pb.RobotState
};

Api.CallServiceArebot_unloadRobot_bridgeGet_parameters = {
  methodName: "CallServiceArebot_unloadRobot_bridgeGet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParametersReq,
  responseType: rcl_interfaces_pb.GetParametersRsp
};

Api.CallServiceArebot_transportRobot_bridgeGet_parameter_types = {
  methodName: "CallServiceArebot_transportRobot_bridgeGet_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceArebot_unloadRobot_bridgeSet_parameters_atomically = {
  methodName: "CallServiceArebot_unloadRobot_bridgeSet_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceArebot_transportRobot_bridgeSet_parameters_atomically = {
  methodName: "CallServiceArebot_transportRobot_bridgeSet_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceArebot_loadRobot_bridgeGet_parameter_types = {
  methodName: "CallServiceArebot_loadRobot_bridgeGet_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceArebot_loadRobot_bridgeGet_parameters = {
  methodName: "CallServiceArebot_loadRobot_bridgeGet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParametersReq,
  responseType: rcl_interfaces_pb.GetParametersRsp
};

Api.CallServiceArebot_unloadRobot_bridgeDescribe_parameters = {
  methodName: "CallServiceArebot_unloadRobot_bridgeDescribe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

Api.CallServiceArebot_transportRobot_bridgeList_parameters = {
  methodName: "CallServiceArebot_transportRobot_bridgeList_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceArebot_loadRobot_bridgeStart_nav_to_pose = {
  methodName: "CallServiceArebot_loadRobot_bridgeStart_nav_to_pose",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.StartNavToPoseReq,
  responseType: robot_bridge_interfaces_pb.StartNavToPoseRsp
};

Api.CallServiceArebot_loadRobot_bridgeSet_parameters_atomically = {
  methodName: "CallServiceArebot_loadRobot_bridgeSet_parameters_atomically",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersAtomicallyReq,
  responseType: rcl_interfaces_pb.SetParametersAtomicallyRsp
};

Api.CallServiceArebot_unloadRobot_bridgeGet_parameter_types = {
  methodName: "CallServiceArebot_unloadRobot_bridgeGet_parameter_types",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParameterTypesReq,
  responseType: rcl_interfaces_pb.GetParameterTypesRsp
};

Api.CallServiceArebot_transportRobot_bridgeGet_state = {
  methodName: "CallServiceArebot_transportRobot_bridgeGet_state",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.GetStateReq,
  responseType: robot_bridge_interfaces_pb.GetStateRsp
};

Api.CallServiceArebot_transportRobot_bridgeSet_parameters = {
  methodName: "CallServiceArebot_transportRobot_bridgeSet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersReq,
  responseType: rcl_interfaces_pb.SetParametersRsp
};

Api.CallServiceArebot_unloadRobot_bridgeStart_nav_to_pose = {
  methodName: "CallServiceArebot_unloadRobot_bridgeStart_nav_to_pose",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.StartNavToPoseReq,
  responseType: robot_bridge_interfaces_pb.StartNavToPoseRsp
};

Api.CallServiceArebot_loadRobot_bridgeGet_state = {
  methodName: "CallServiceArebot_loadRobot_bridgeGet_state",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.GetStateReq,
  responseType: robot_bridge_interfaces_pb.GetStateRsp
};

Api.CallServiceArebot_loadRobot_bridgeList_parameters = {
  methodName: "CallServiceArebot_loadRobot_bridgeList_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceArebot_unloadRobot_bridgeSet_parameters = {
  methodName: "CallServiceArebot_unloadRobot_bridgeSet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.SetParametersReq,
  responseType: rcl_interfaces_pb.SetParametersRsp
};

Api.CallServiceArebot_transportRobot_bridgeStart_nav_to_pose = {
  methodName: "CallServiceArebot_transportRobot_bridgeStart_nav_to_pose",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.StartNavToPoseReq,
  responseType: robot_bridge_interfaces_pb.StartNavToPoseRsp
};

Api.CallServiceArebot_loadRobot_bridgeDescribe_parameters = {
  methodName: "CallServiceArebot_loadRobot_bridgeDescribe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

Api.CallServiceArebot_unloadRobot_bridgeList_parameters = {
  methodName: "CallServiceArebot_unloadRobot_bridgeList_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.ListParametersReq,
  responseType: rcl_interfaces_pb.ListParametersRsp
};

Api.CallServiceArebot_loadRobot_bridgeStart_pick_and_load = {
  methodName: "CallServiceArebot_loadRobot_bridgeStart_pick_and_load",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.StartPickAndLoadReq,
  responseType: robot_bridge_interfaces_pb.StartPickAndLoadRsp
};

Api.CallServiceArebot_unloadRobot_bridgeGet_state = {
  methodName: "CallServiceArebot_unloadRobot_bridgeGet_state",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.GetStateReq,
  responseType: robot_bridge_interfaces_pb.GetStateRsp
};

Api.CallServiceArebot_unloadRobot_bridgeStart_unload_and_place = {
  methodName: "CallServiceArebot_unloadRobot_bridgeStart_unload_and_place",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.StartUnloadAndPlaceReq,
  responseType: robot_bridge_interfaces_pb.StartUnloadAndPlaceRsp
};

Api.CallServiceArebot_transportRobot_bridgeDescribe_parameters = {
  methodName: "CallServiceArebot_transportRobot_bridgeDescribe_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.DescribeParametersReq,
  responseType: rcl_interfaces_pb.DescribeParametersRsp
};

Api.CallServiceArebot_transportRobot_bridgeGet_parameters = {
  methodName: "CallServiceArebot_transportRobot_bridgeGet_parameters",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: rcl_interfaces_pb.GetParametersReq,
  responseType: rcl_interfaces_pb.GetParametersRsp
};

Api.CallServiceArebot_transportRobot_bridgeStart_transport = {
  methodName: "CallServiceArebot_transportRobot_bridgeStart_transport",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: robot_bridge_interfaces_pb.StartTransportReq,
  responseType: robot_bridge_interfaces_pb.StartTransportRsp
};

Api.CallServiceArebot_loadRobot_bridgeSet_parameters = {
  methodName: "CallServiceArebot_loadRobot_bridgeSet_parameters",
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

ApiClient.prototype.publishTopicArebot_transportRobot_bridgeRobot_state = function publishTopicArebot_transportRobot_bridgeRobot_state(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.PublishTopicArebot_transportRobot_bridgeRobot_state, {
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

ApiClient.prototype.publishTopicArebot_loadRobot_bridgeRobot_state = function publishTopicArebot_loadRobot_bridgeRobot_state(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.PublishTopicArebot_loadRobot_bridgeRobot_state, {
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

ApiClient.prototype.publishTopicArebot_unloadRobot_bridgeRobot_state = function publishTopicArebot_unloadRobot_bridgeRobot_state(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.PublishTopicArebot_unloadRobot_bridgeRobot_state, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeGet_parameters = function callServiceArebot_unloadRobot_bridgeGet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeGet_parameters, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeGet_parameter_types = function callServiceArebot_transportRobot_bridgeGet_parameter_types(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeGet_parameter_types, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeSet_parameters_atomically = function callServiceArebot_unloadRobot_bridgeSet_parameters_atomically(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeSet_parameters_atomically, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeSet_parameters_atomically = function callServiceArebot_transportRobot_bridgeSet_parameters_atomically(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeSet_parameters_atomically, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeGet_parameter_types = function callServiceArebot_loadRobot_bridgeGet_parameter_types(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeGet_parameter_types, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeGet_parameters = function callServiceArebot_loadRobot_bridgeGet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeGet_parameters, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeDescribe_parameters = function callServiceArebot_unloadRobot_bridgeDescribe_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeDescribe_parameters, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeList_parameters = function callServiceArebot_transportRobot_bridgeList_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeList_parameters, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeStart_nav_to_pose = function callServiceArebot_loadRobot_bridgeStart_nav_to_pose(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeStart_nav_to_pose, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeSet_parameters_atomically = function callServiceArebot_loadRobot_bridgeSet_parameters_atomically(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeSet_parameters_atomically, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeGet_parameter_types = function callServiceArebot_unloadRobot_bridgeGet_parameter_types(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeGet_parameter_types, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeGet_state = function callServiceArebot_transportRobot_bridgeGet_state(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeGet_state, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeSet_parameters = function callServiceArebot_transportRobot_bridgeSet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeSet_parameters, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeStart_nav_to_pose = function callServiceArebot_unloadRobot_bridgeStart_nav_to_pose(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeStart_nav_to_pose, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeGet_state = function callServiceArebot_loadRobot_bridgeGet_state(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeGet_state, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeList_parameters = function callServiceArebot_loadRobot_bridgeList_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeList_parameters, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeSet_parameters = function callServiceArebot_unloadRobot_bridgeSet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeSet_parameters, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeStart_nav_to_pose = function callServiceArebot_transportRobot_bridgeStart_nav_to_pose(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeStart_nav_to_pose, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeDescribe_parameters = function callServiceArebot_loadRobot_bridgeDescribe_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeDescribe_parameters, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeList_parameters = function callServiceArebot_unloadRobot_bridgeList_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeList_parameters, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeStart_pick_and_load = function callServiceArebot_loadRobot_bridgeStart_pick_and_load(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeStart_pick_and_load, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeGet_state = function callServiceArebot_unloadRobot_bridgeGet_state(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeGet_state, {
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

ApiClient.prototype.callServiceArebot_unloadRobot_bridgeStart_unload_and_place = function callServiceArebot_unloadRobot_bridgeStart_unload_and_place(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_unloadRobot_bridgeStart_unload_and_place, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeDescribe_parameters = function callServiceArebot_transportRobot_bridgeDescribe_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeDescribe_parameters, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeGet_parameters = function callServiceArebot_transportRobot_bridgeGet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeGet_parameters, {
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

ApiClient.prototype.callServiceArebot_transportRobot_bridgeStart_transport = function callServiceArebot_transportRobot_bridgeStart_transport(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_transportRobot_bridgeStart_transport, {
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

ApiClient.prototype.callServiceArebot_loadRobot_bridgeSet_parameters = function callServiceArebot_loadRobot_bridgeSet_parameters(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Api.CallServiceArebot_loadRobot_bridgeSet_parameters, {
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

