// package: pkg_c
// file: pkg_c.proto

import * as jspb from "google-protobuf";
import * as pkg_a_pb from "./pkg_a_pb";

export class State extends jspb.Message {
  getCount(): number;
  setCount(value: number): void;

  hasV(): boolean;
  clearV(): void;
  getV(): pkg_a_pb.V3 | undefined;
  setV(value?: pkg_a_pb.V3): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): State.AsObject;
  static toObject(includeInstance: boolean, msg: State): State.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: State, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): State;
  static deserializeBinaryFromReader(message: State, reader: jspb.BinaryReader): State;
}

export namespace State {
  export type AsObject = {
    count: number,
    v?: pkg_a_pb.V3.AsObject,
  }
}

