// package: pkg_b
// file: pkg_b.proto

import * as jspb from "google-protobuf";
import * as pkg_a_pb from "./pkg_a_pb";

export class Pos extends jspb.Message {
  getType(): string;
  setType(value: string): void;

  hasV(): boolean;
  clearV(): void;
  getV(): pkg_a_pb.V3 | undefined;
  setV(value?: pkg_a_pb.V3): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pos.AsObject;
  static toObject(includeInstance: boolean, msg: Pos): Pos.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Pos, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pos;
  static deserializeBinaryFromReader(message: Pos, reader: jspb.BinaryReader): Pos;
}

export namespace Pos {
  export type AsObject = {
    type: string,
    v?: pkg_a_pb.V3.AsObject,
  }
}

