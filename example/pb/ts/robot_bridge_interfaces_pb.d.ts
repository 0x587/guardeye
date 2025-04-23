// package: robot_bridge_interfaces
// file: robot_bridge_interfaces.proto

import * as jspb from "google-protobuf";
import * as geometry_msgs_pb from "./geometry_msgs_pb";

export class StartNavToPoseRsp extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  getTaskName(): string;
  setTaskName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartNavToPoseRsp.AsObject;
  static toObject(includeInstance: boolean, msg: StartNavToPoseRsp): StartNavToPoseRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartNavToPoseRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartNavToPoseRsp;
  static deserializeBinaryFromReader(message: StartNavToPoseRsp, reader: jspb.BinaryReader): StartNavToPoseRsp;
}

export namespace StartNavToPoseRsp {
  export type AsObject = {
    success: boolean,
    message: string,
    taskName: string,
  }
}

export class StartTransportReq extends jspb.Message {
  hasPose(): boolean;
  clearPose(): void;
  getPose(): geometry_msgs_pb.Pose | undefined;
  setPose(value?: geometry_msgs_pb.Pose): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartTransportReq.AsObject;
  static toObject(includeInstance: boolean, msg: StartTransportReq): StartTransportReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartTransportReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartTransportReq;
  static deserializeBinaryFromReader(message: StartTransportReq, reader: jspb.BinaryReader): StartTransportReq;
}

export namespace StartTransportReq {
  export type AsObject = {
    pose?: geometry_msgs_pb.Pose.AsObject,
  }
}

export class StartPickAndLoadReq extends jspb.Message {
  getItemName(): string;
  setItemName(value: string): void;

  getItemPositionX(): number;
  setItemPositionX(value: number): void;

  getItemPositionY(): number;
  setItemPositionY(value: number): void;

  getLoadPositionX(): number;
  setLoadPositionX(value: number): void;

  getLoadPositionY(): number;
  setLoadPositionY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartPickAndLoadReq.AsObject;
  static toObject(includeInstance: boolean, msg: StartPickAndLoadReq): StartPickAndLoadReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartPickAndLoadReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartPickAndLoadReq;
  static deserializeBinaryFromReader(message: StartPickAndLoadReq, reader: jspb.BinaryReader): StartPickAndLoadReq;
}

export namespace StartPickAndLoadReq {
  export type AsObject = {
    itemName: string,
    itemPositionX: number,
    itemPositionY: number,
    loadPositionX: number,
    loadPositionY: number,
  }
}

export class StartUnloadAndPlaceReq extends jspb.Message {
  getItemName(): string;
  setItemName(value: string): void;

  getUnloadPositionX(): number;
  setUnloadPositionX(value: number): void;

  getUnloadPositionY(): number;
  setUnloadPositionY(value: number): void;

  getShelfPositionX(): number;
  setShelfPositionX(value: number): void;

  getShelfPositionY(): number;
  setShelfPositionY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartUnloadAndPlaceReq.AsObject;
  static toObject(includeInstance: boolean, msg: StartUnloadAndPlaceReq): StartUnloadAndPlaceReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartUnloadAndPlaceReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartUnloadAndPlaceReq;
  static deserializeBinaryFromReader(message: StartUnloadAndPlaceReq, reader: jspb.BinaryReader): StartUnloadAndPlaceReq;
}

export namespace StartUnloadAndPlaceReq {
  export type AsObject = {
    itemName: string,
    unloadPositionX: number,
    unloadPositionY: number,
    shelfPositionX: number,
    shelfPositionY: number,
  }
}

export class GetStateReq extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStateReq.AsObject;
  static toObject(includeInstance: boolean, msg: GetStateReq): GetStateReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetStateReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStateReq;
  static deserializeBinaryFromReader(message: GetStateReq, reader: jspb.BinaryReader): GetStateReq;
}

export namespace GetStateReq {
  export type AsObject = {
  }
}

export class StartNavToPoseReq extends jspb.Message {
  hasPose(): boolean;
  clearPose(): void;
  getPose(): geometry_msgs_pb.Pose | undefined;
  setPose(value?: geometry_msgs_pb.Pose): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartNavToPoseReq.AsObject;
  static toObject(includeInstance: boolean, msg: StartNavToPoseReq): StartNavToPoseReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartNavToPoseReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartNavToPoseReq;
  static deserializeBinaryFromReader(message: StartNavToPoseReq, reader: jspb.BinaryReader): StartNavToPoseReq;
}

export namespace StartNavToPoseReq {
  export type AsObject = {
    pose?: geometry_msgs_pb.Pose.AsObject,
  }
}

export class StartTransportRsp extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  getTaskName(): string;
  setTaskName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartTransportRsp.AsObject;
  static toObject(includeInstance: boolean, msg: StartTransportRsp): StartTransportRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartTransportRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartTransportRsp;
  static deserializeBinaryFromReader(message: StartTransportRsp, reader: jspb.BinaryReader): StartTransportRsp;
}

export namespace StartTransportRsp {
  export type AsObject = {
    success: boolean,
    message: string,
    taskName: string,
  }
}

export class StartPickAndLoadRsp extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  getTaskName(): string;
  setTaskName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartPickAndLoadRsp.AsObject;
  static toObject(includeInstance: boolean, msg: StartPickAndLoadRsp): StartPickAndLoadRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartPickAndLoadRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartPickAndLoadRsp;
  static deserializeBinaryFromReader(message: StartPickAndLoadRsp, reader: jspb.BinaryReader): StartPickAndLoadRsp;
}

export namespace StartPickAndLoadRsp {
  export type AsObject = {
    success: boolean,
    message: string,
    taskName: string,
  }
}

export class StartUnloadAndPlaceRsp extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  getTaskName(): string;
  setTaskName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartUnloadAndPlaceRsp.AsObject;
  static toObject(includeInstance: boolean, msg: StartUnloadAndPlaceRsp): StartUnloadAndPlaceRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartUnloadAndPlaceRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartUnloadAndPlaceRsp;
  static deserializeBinaryFromReader(message: StartUnloadAndPlaceRsp, reader: jspb.BinaryReader): StartUnloadAndPlaceRsp;
}

export namespace StartUnloadAndPlaceRsp {
  export type AsObject = {
    success: boolean,
    message: string,
    taskName: string,
  }
}

export class RobotState extends jspb.Message {
  getRobotName(): string;
  setRobotName(value: string): void;

  getRobotType(): string;
  setRobotType(value: string): void;

  getIsWorking(): boolean;
  setIsWorking(value: boolean): void;

  getTaskName(): string;
  setTaskName(value: string): void;

  getTaskSuccess(): boolean;
  setTaskSuccess(value: boolean): void;

  getTaskMessage(): string;
  setTaskMessage(value: string): void;

  getPositionX(): number;
  setPositionX(value: number): void;

  getPositionY(): number;
  setPositionY(value: number): void;

  getPositionZ(): number;
  setPositionZ(value: number): void;

  getOrientationX(): number;
  setOrientationX(value: number): void;

  getOrientationY(): number;
  setOrientationY(value: number): void;

  getOrientationZ(): number;
  setOrientationZ(value: number): void;

  getOrientationW(): number;
  setOrientationW(value: number): void;

  getMotor1JointAngle(): number;
  setMotor1JointAngle(value: number): void;

  getMotor2JointAngle(): number;
  setMotor2JointAngle(value: number): void;

  getMotor3JointAngle(): number;
  setMotor3JointAngle(value: number): void;

  getEndJointAngle(): number;
  setEndJointAngle(value: number): void;

  getOdomPositionX(): number;
  setOdomPositionX(value: number): void;

  getOdomPositionXUncertainty(): number;
  setOdomPositionXUncertainty(value: number): void;

  getOdomPositionY(): number;
  setOdomPositionY(value: number): void;

  getOdomPositionYUncertainty(): number;
  setOdomPositionYUncertainty(value: number): void;

  getOdomPositionZ(): number;
  setOdomPositionZ(value: number): void;

  getOdomPositionZUncertainty(): number;
  setOdomPositionZUncertainty(value: number): void;

  getOdomOrientationX(): number;
  setOdomOrientationX(value: number): void;

  getOdomOrientationXUncertainty(): number;
  setOdomOrientationXUncertainty(value: number): void;

  getOdomOrientationY(): number;
  setOdomOrientationY(value: number): void;

  getOdomOrientationYUncertainty(): number;
  setOdomOrientationYUncertainty(value: number): void;

  getOdomOrientationZ(): number;
  setOdomOrientationZ(value: number): void;

  getOdomOrientationZUncertainty(): number;
  setOdomOrientationZUncertainty(value: number): void;

  getOdomOrientationW(): number;
  setOdomOrientationW(value: number): void;

  getOdomOrientationWUncertainty(): number;
  setOdomOrientationWUncertainty(value: number): void;

  getLinearVelocityX(): number;
  setLinearVelocityX(value: number): void;

  getLinearVelocityXUncertainty(): number;
  setLinearVelocityXUncertainty(value: number): void;

  getLinearVelocityY(): number;
  setLinearVelocityY(value: number): void;

  getLinearVelocityYUncertainty(): number;
  setLinearVelocityYUncertainty(value: number): void;

  getLinearVelocityZ(): number;
  setLinearVelocityZ(value: number): void;

  getLinearVelocityZUncertainty(): number;
  setLinearVelocityZUncertainty(value: number): void;

  getAngularVelocityRoll(): number;
  setAngularVelocityRoll(value: number): void;

  getAngularVelocityRollUncertainty(): number;
  setAngularVelocityRollUncertainty(value: number): void;

  getAngularVelocityPitch(): number;
  setAngularVelocityPitch(value: number): void;

  getAngularVelocityPitchUncertainty(): number;
  setAngularVelocityPitchUncertainty(value: number): void;

  getAngularVelocityYaw(): number;
  setAngularVelocityYaw(value: number): void;

  getAngularVelocityYawUncertainty(): number;
  setAngularVelocityYawUncertainty(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RobotState.AsObject;
  static toObject(includeInstance: boolean, msg: RobotState): RobotState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RobotState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RobotState;
  static deserializeBinaryFromReader(message: RobotState, reader: jspb.BinaryReader): RobotState;
}

export namespace RobotState {
  export type AsObject = {
    robotName: string,
    robotType: string,
    isWorking: boolean,
    taskName: string,
    taskSuccess: boolean,
    taskMessage: string,
    positionX: number,
    positionY: number,
    positionZ: number,
    orientationX: number,
    orientationY: number,
    orientationZ: number,
    orientationW: number,
    motor1JointAngle: number,
    motor2JointAngle: number,
    motor3JointAngle: number,
    endJointAngle: number,
    odomPositionX: number,
    odomPositionXUncertainty: number,
    odomPositionY: number,
    odomPositionYUncertainty: number,
    odomPositionZ: number,
    odomPositionZUncertainty: number,
    odomOrientationX: number,
    odomOrientationXUncertainty: number,
    odomOrientationY: number,
    odomOrientationYUncertainty: number,
    odomOrientationZ: number,
    odomOrientationZUncertainty: number,
    odomOrientationW: number,
    odomOrientationWUncertainty: number,
    linearVelocityX: number,
    linearVelocityXUncertainty: number,
    linearVelocityY: number,
    linearVelocityYUncertainty: number,
    linearVelocityZ: number,
    linearVelocityZUncertainty: number,
    angularVelocityRoll: number,
    angularVelocityRollUncertainty: number,
    angularVelocityPitch: number,
    angularVelocityPitchUncertainty: number,
    angularVelocityYaw: number,
    angularVelocityYawUncertainty: number,
  }
}

export class GetStateRsp extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getMessage(): string;
  setMessage(value: string): void;

  getRobotName(): string;
  setRobotName(value: string): void;

  getRobotType(): string;
  setRobotType(value: string): void;

  getIsWorking(): boolean;
  setIsWorking(value: boolean): void;

  getTaskName(): string;
  setTaskName(value: string): void;

  getTaskSuccess(): boolean;
  setTaskSuccess(value: boolean): void;

  getTaskMessage(): string;
  setTaskMessage(value: string): void;

  getPositionX(): number;
  setPositionX(value: number): void;

  getPositionY(): number;
  setPositionY(value: number): void;

  getPositionZ(): number;
  setPositionZ(value: number): void;

  getOrientationX(): number;
  setOrientationX(value: number): void;

  getOrientationY(): number;
  setOrientationY(value: number): void;

  getOrientationZ(): number;
  setOrientationZ(value: number): void;

  getOrientationW(): number;
  setOrientationW(value: number): void;

  getMotor1JointAngle(): number;
  setMotor1JointAngle(value: number): void;

  getMotor2JointAngle(): number;
  setMotor2JointAngle(value: number): void;

  getMotor3JointAngle(): number;
  setMotor3JointAngle(value: number): void;

  getEndJointAngle(): number;
  setEndJointAngle(value: number): void;

  getOdomPositionX(): number;
  setOdomPositionX(value: number): void;

  getOdomPositionXUncertainty(): number;
  setOdomPositionXUncertainty(value: number): void;

  getOdomPositionY(): number;
  setOdomPositionY(value: number): void;

  getOdomPositionYUncertainty(): number;
  setOdomPositionYUncertainty(value: number): void;

  getOdomPositionZ(): number;
  setOdomPositionZ(value: number): void;

  getOdomPositionZUncertainty(): number;
  setOdomPositionZUncertainty(value: number): void;

  getOdomOrientationX(): number;
  setOdomOrientationX(value: number): void;

  getOdomOrientationXUncertainty(): number;
  setOdomOrientationXUncertainty(value: number): void;

  getOdomOrientationY(): number;
  setOdomOrientationY(value: number): void;

  getOdomOrientationYUncertainty(): number;
  setOdomOrientationYUncertainty(value: number): void;

  getOdomOrientationZ(): number;
  setOdomOrientationZ(value: number): void;

  getOdomOrientationZUncertainty(): number;
  setOdomOrientationZUncertainty(value: number): void;

  getOdomOrientationW(): number;
  setOdomOrientationW(value: number): void;

  getOdomOrientationWUncertainty(): number;
  setOdomOrientationWUncertainty(value: number): void;

  getLinearVelocityX(): number;
  setLinearVelocityX(value: number): void;

  getLinearVelocityXUncertainty(): number;
  setLinearVelocityXUncertainty(value: number): void;

  getLinearVelocityY(): number;
  setLinearVelocityY(value: number): void;

  getLinearVelocityYUncertainty(): number;
  setLinearVelocityYUncertainty(value: number): void;

  getLinearVelocityZ(): number;
  setLinearVelocityZ(value: number): void;

  getLinearVelocityZUncertainty(): number;
  setLinearVelocityZUncertainty(value: number): void;

  getAngularVelocityRoll(): number;
  setAngularVelocityRoll(value: number): void;

  getAngularVelocityRollUncertainty(): number;
  setAngularVelocityRollUncertainty(value: number): void;

  getAngularVelocityPitch(): number;
  setAngularVelocityPitch(value: number): void;

  getAngularVelocityPitchUncertainty(): number;
  setAngularVelocityPitchUncertainty(value: number): void;

  getAngularVelocityYaw(): number;
  setAngularVelocityYaw(value: number): void;

  getAngularVelocityYawUncertainty(): number;
  setAngularVelocityYawUncertainty(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetStateRsp.AsObject;
  static toObject(includeInstance: boolean, msg: GetStateRsp): GetStateRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetStateRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetStateRsp;
  static deserializeBinaryFromReader(message: GetStateRsp, reader: jspb.BinaryReader): GetStateRsp;
}

export namespace GetStateRsp {
  export type AsObject = {
    success: boolean,
    message: string,
    robotName: string,
    robotType: string,
    isWorking: boolean,
    taskName: string,
    taskSuccess: boolean,
    taskMessage: string,
    positionX: number,
    positionY: number,
    positionZ: number,
    orientationX: number,
    orientationY: number,
    orientationZ: number,
    orientationW: number,
    motor1JointAngle: number,
    motor2JointAngle: number,
    motor3JointAngle: number,
    endJointAngle: number,
    odomPositionX: number,
    odomPositionXUncertainty: number,
    odomPositionY: number,
    odomPositionYUncertainty: number,
    odomPositionZ: number,
    odomPositionZUncertainty: number,
    odomOrientationX: number,
    odomOrientationXUncertainty: number,
    odomOrientationY: number,
    odomOrientationYUncertainty: number,
    odomOrientationZ: number,
    odomOrientationZUncertainty: number,
    odomOrientationW: number,
    odomOrientationWUncertainty: number,
    linearVelocityX: number,
    linearVelocityXUncertainty: number,
    linearVelocityY: number,
    linearVelocityYUncertainty: number,
    linearVelocityZ: number,
    linearVelocityZUncertainty: number,
    angularVelocityRoll: number,
    angularVelocityRollUncertainty: number,
    angularVelocityPitch: number,
    angularVelocityPitchUncertainty: number,
    angularVelocityYaw: number,
    angularVelocityYawUncertainty: number,
  }
}

