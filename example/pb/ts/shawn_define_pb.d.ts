// package: shawn_define
// file: shawn_define.proto

import * as jspb from "google-protobuf";

export class AddReq extends jspb.Message {
  hasA(): boolean;
  clearA(): void;
  getA(): V2 | undefined;
  setA(value?: V2): void;

  hasB(): boolean;
  clearB(): void;
  getB(): V2 | undefined;
  setB(value?: V2): void;

  clearCList(): void;
  getCList(): Array<number>;
  setCList(value: Array<number>): void;
  addC(value: number, index?: number): number;

  getD(): number;
  setD(value: number): void;

  clearEList(): void;
  getEList(): Array<number>;
  setEList(value: Array<number>): void;
  addE(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddReq.AsObject;
  static toObject(includeInstance: boolean, msg: AddReq): AddReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddReq;
  static deserializeBinaryFromReader(message: AddReq, reader: jspb.BinaryReader): AddReq;
}

export namespace AddReq {
  export type AsObject = {
    a?: V2.AsObject,
    b?: V2.AsObject,
    cList: Array<number>,
    d: number,
    eList: Array<number>,
  }
}

export class V2 extends jspb.Message {
  getA(): number;
  setA(value: number): void;

  getB(): number;
  setB(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): V2.AsObject;
  static toObject(includeInstance: boolean, msg: V2): V2.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: V2, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): V2;
  static deserializeBinaryFromReader(message: V2, reader: jspb.BinaryReader): V2;
}

export namespace V2 {
  export type AsObject = {
    a: number,
    b: number,
  }
}

export class AddRsp extends jspb.Message {
  hasRes(): boolean;
  clearRes(): void;
  getRes(): V2 | undefined;
  setRes(value?: V2): void;

  clearCList(): void;
  getCList(): Array<number>;
  setCList(value: Array<number>): void;
  addC(value: number, index?: number): number;

  getD(): number;
  setD(value: number): void;

  clearEList(): void;
  getEList(): Array<number>;
  setEList(value: Array<number>): void;
  addE(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddRsp.AsObject;
  static toObject(includeInstance: boolean, msg: AddRsp): AddRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddRsp;
  static deserializeBinaryFromReader(message: AddRsp, reader: jspb.BinaryReader): AddRsp;
}

export namespace AddRsp {
  export type AsObject = {
    res?: V2.AsObject,
    cList: Array<number>,
    d: number,
    eList: Array<number>,
  }
}

