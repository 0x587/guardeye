// package: pkg_a
// file: pkg_a.proto

import * as jspb from "google-protobuf";

export class V3 extends jspb.Message {
  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  getZ(): number;
  setZ(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): V3.AsObject;
  static toObject(includeInstance: boolean, msg: V3): V3.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: V3, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): V3;
  static deserializeBinaryFromReader(message: V3, reader: jspb.BinaryReader): V3;
}

export namespace V3 {
  export type AsObject = {
    x: number,
    y: number,
    z: number,
  }
}

