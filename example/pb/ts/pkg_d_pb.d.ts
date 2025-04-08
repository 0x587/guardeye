// package: pkg_d
// file: pkg_d.proto

import * as jspb from "google-protobuf";
import * as pkg_a_pb from "./pkg_a_pb";
import * as pkg_b_pb from "./pkg_b_pb";
import * as pkg_c_pb from "./pkg_c_pb";

export class GetPersonReq extends jspb.Message {
  hasReq(): boolean;
  clearReq(): void;
  getReq(): Person | undefined;
  setReq(value?: Person): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPersonReq.AsObject;
  static toObject(includeInstance: boolean, msg: GetPersonReq): GetPersonReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPersonReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPersonReq;
  static deserializeBinaryFromReader(message: GetPersonReq, reader: jspb.BinaryReader): GetPersonReq;
}

export namespace GetPersonReq {
  export type AsObject = {
    req?: Person.AsObject,
  }
}

export class Person extends jspb.Message {
  hasV(): boolean;
  clearV(): void;
  getV(): pkg_a_pb.V3 | undefined;
  setV(value?: pkg_a_pb.V3): void;

  hasPos(): boolean;
  clearPos(): void;
  getPos(): pkg_b_pb.Pos | undefined;
  setPos(value?: pkg_b_pb.Pos): void;

  hasState(): boolean;
  clearState(): void;
  getState(): pkg_c_pb.State | undefined;
  setState(value?: pkg_c_pb.State): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Person.AsObject;
  static toObject(includeInstance: boolean, msg: Person): Person.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Person, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Person;
  static deserializeBinaryFromReader(message: Person, reader: jspb.BinaryReader): Person;
}

export namespace Person {
  export type AsObject = {
    v?: pkg_a_pb.V3.AsObject,
    pos?: pkg_b_pb.Pos.AsObject,
    state?: pkg_c_pb.State.AsObject,
  }
}

export class GetPersonRsp extends jspb.Message {
  hasRsp(): boolean;
  clearRsp(): void;
  getRsp(): Person | undefined;
  setRsp(value?: Person): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPersonRsp.AsObject;
  static toObject(includeInstance: boolean, msg: GetPersonRsp): GetPersonRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPersonRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPersonRsp;
  static deserializeBinaryFromReader(message: GetPersonRsp, reader: jspb.BinaryReader): GetPersonRsp;
}

export namespace GetPersonRsp {
  export type AsObject = {
    rsp?: Person.AsObject,
  }
}

