// package: rcl_interfaces
// file: rcl_interfaces.proto

import * as jspb from "google-protobuf";
import * as builtin_interfaces_pb from "./builtin_interfaces_pb";

export class Log extends jspb.Message {
  hasStamp(): boolean;
  clearStamp(): void;
  getStamp(): builtin_interfaces_pb.Time | undefined;
  setStamp(value?: builtin_interfaces_pb.Time): void;

  getLevel(): number;
  setLevel(value: number): void;

  getName(): string;
  setName(value: string): void;

  getMsg(): string;
  setMsg(value: string): void;

  getFile(): string;
  setFile(value: string): void;

  getFunction(): string;
  setFunction(value: string): void;

  getLine(): number;
  setLine(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Log.AsObject;
  static toObject(includeInstance: boolean, msg: Log): Log.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Log, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Log;
  static deserializeBinaryFromReader(message: Log, reader: jspb.BinaryReader): Log;
}

export namespace Log {
  export type AsObject = {
    stamp?: builtin_interfaces_pb.Time.AsObject,
    level: number,
    name: string,
    msg: string,
    file: string,
    pb_function: string,
    line: number,
  }
}

export class DescribeParametersReq extends jspb.Message {
  clearNamesList(): void;
  getNamesList(): Array<string>;
  setNamesList(value: Array<string>): void;
  addNames(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DescribeParametersReq.AsObject;
  static toObject(includeInstance: boolean, msg: DescribeParametersReq): DescribeParametersReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DescribeParametersReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DescribeParametersReq;
  static deserializeBinaryFromReader(message: DescribeParametersReq, reader: jspb.BinaryReader): DescribeParametersReq;
}

export namespace DescribeParametersReq {
  export type AsObject = {
    namesList: Array<string>,
  }
}

export class GetParametersReq extends jspb.Message {
  clearNamesList(): void;
  getNamesList(): Array<string>;
  setNamesList(value: Array<string>): void;
  addNames(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetParametersReq.AsObject;
  static toObject(includeInstance: boolean, msg: GetParametersReq): GetParametersReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetParametersReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetParametersReq;
  static deserializeBinaryFromReader(message: GetParametersReq, reader: jspb.BinaryReader): GetParametersReq;
}

export namespace GetParametersReq {
  export type AsObject = {
    namesList: Array<string>,
  }
}

export class ListParametersRsp extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): ListParametersResult | undefined;
  setResult(value?: ListParametersResult): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListParametersRsp.AsObject;
  static toObject(includeInstance: boolean, msg: ListParametersRsp): ListParametersRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListParametersRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListParametersRsp;
  static deserializeBinaryFromReader(message: ListParametersRsp, reader: jspb.BinaryReader): ListParametersRsp;
}

export namespace ListParametersRsp {
  export type AsObject = {
    result?: ListParametersResult.AsObject,
  }
}

export class ParameterEvent extends jspb.Message {
  hasStamp(): boolean;
  clearStamp(): void;
  getStamp(): builtin_interfaces_pb.Time | undefined;
  setStamp(value?: builtin_interfaces_pb.Time): void;

  getNode(): string;
  setNode(value: string): void;

  clearNewParametersList(): void;
  getNewParametersList(): Array<Parameter>;
  setNewParametersList(value: Array<Parameter>): void;
  addNewParameters(value?: Parameter, index?: number): Parameter;

  clearChangedParametersList(): void;
  getChangedParametersList(): Array<Parameter>;
  setChangedParametersList(value: Array<Parameter>): void;
  addChangedParameters(value?: Parameter, index?: number): Parameter;

  clearDeletedParametersList(): void;
  getDeletedParametersList(): Array<Parameter>;
  setDeletedParametersList(value: Array<Parameter>): void;
  addDeletedParameters(value?: Parameter, index?: number): Parameter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ParameterEvent.AsObject;
  static toObject(includeInstance: boolean, msg: ParameterEvent): ParameterEvent.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ParameterEvent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ParameterEvent;
  static deserializeBinaryFromReader(message: ParameterEvent, reader: jspb.BinaryReader): ParameterEvent;
}

export namespace ParameterEvent {
  export type AsObject = {
    stamp?: builtin_interfaces_pb.Time.AsObject,
    node: string,
    newParametersList: Array<Parameter.AsObject>,
    changedParametersList: Array<Parameter.AsObject>,
    deletedParametersList: Array<Parameter.AsObject>,
  }
}

export class IntegerRange extends jspb.Message {
  getFromValue(): number;
  setFromValue(value: number): void;

  getToValue(): number;
  setToValue(value: number): void;

  getStep(): number;
  setStep(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IntegerRange.AsObject;
  static toObject(includeInstance: boolean, msg: IntegerRange): IntegerRange.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IntegerRange, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IntegerRange;
  static deserializeBinaryFromReader(message: IntegerRange, reader: jspb.BinaryReader): IntegerRange;
}

export namespace IntegerRange {
  export type AsObject = {
    fromValue: number,
    toValue: number,
    step: number,
  }
}

export class SetParametersRsp extends jspb.Message {
  clearResultsList(): void;
  getResultsList(): Array<SetParametersResult>;
  setResultsList(value: Array<SetParametersResult>): void;
  addResults(value?: SetParametersResult, index?: number): SetParametersResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetParametersRsp.AsObject;
  static toObject(includeInstance: boolean, msg: SetParametersRsp): SetParametersRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetParametersRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetParametersRsp;
  static deserializeBinaryFromReader(message: SetParametersRsp, reader: jspb.BinaryReader): SetParametersRsp;
}

export namespace SetParametersRsp {
  export type AsObject = {
    resultsList: Array<SetParametersResult.AsObject>,
  }
}

export class SetParametersResult extends jspb.Message {
  getSuccessful(): boolean;
  setSuccessful(value: boolean): void;

  getReason(): string;
  setReason(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetParametersResult.AsObject;
  static toObject(includeInstance: boolean, msg: SetParametersResult): SetParametersResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetParametersResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetParametersResult;
  static deserializeBinaryFromReader(message: SetParametersResult, reader: jspb.BinaryReader): SetParametersResult;
}

export namespace SetParametersResult {
  export type AsObject = {
    successful: boolean,
    reason: string,
  }
}

export class SetParametersAtomicallyReq extends jspb.Message {
  clearParametersList(): void;
  getParametersList(): Array<Parameter>;
  setParametersList(value: Array<Parameter>): void;
  addParameters(value?: Parameter, index?: number): Parameter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetParametersAtomicallyReq.AsObject;
  static toObject(includeInstance: boolean, msg: SetParametersAtomicallyReq): SetParametersAtomicallyReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetParametersAtomicallyReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetParametersAtomicallyReq;
  static deserializeBinaryFromReader(message: SetParametersAtomicallyReq, reader: jspb.BinaryReader): SetParametersAtomicallyReq;
}

export namespace SetParametersAtomicallyReq {
  export type AsObject = {
    parametersList: Array<Parameter.AsObject>,
  }
}

export class SetParametersAtomicallyRsp extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): SetParametersResult | undefined;
  setResult(value?: SetParametersResult): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetParametersAtomicallyRsp.AsObject;
  static toObject(includeInstance: boolean, msg: SetParametersAtomicallyRsp): SetParametersAtomicallyRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetParametersAtomicallyRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetParametersAtomicallyRsp;
  static deserializeBinaryFromReader(message: SetParametersAtomicallyRsp, reader: jspb.BinaryReader): SetParametersAtomicallyRsp;
}

export namespace SetParametersAtomicallyRsp {
  export type AsObject = {
    result?: SetParametersResult.AsObject,
  }
}

export class Parameter extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasValue(): boolean;
  clearValue(): void;
  getValue(): ParameterValue | undefined;
  setValue(value?: ParameterValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Parameter.AsObject;
  static toObject(includeInstance: boolean, msg: Parameter): Parameter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Parameter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Parameter;
  static deserializeBinaryFromReader(message: Parameter, reader: jspb.BinaryReader): Parameter;
}

export namespace Parameter {
  export type AsObject = {
    name: string,
    value?: ParameterValue.AsObject,
  }
}

export class ParameterValue extends jspb.Message {
  getType(): number;
  setType(value: number): void;

  getBoolValue(): boolean;
  setBoolValue(value: boolean): void;

  getIntegerValue(): number;
  setIntegerValue(value: number): void;

  getDoubleValue(): number;
  setDoubleValue(value: number): void;

  getStringValue(): string;
  setStringValue(value: string): void;

  getByteArrayValue(): Uint8Array | string;
  getByteArrayValue_asU8(): Uint8Array;
  getByteArrayValue_asB64(): string;
  setByteArrayValue(value: Uint8Array | string): void;

  clearBoolArrayValueList(): void;
  getBoolArrayValueList(): Array<boolean>;
  setBoolArrayValueList(value: Array<boolean>): void;
  addBoolArrayValue(value: boolean, index?: number): boolean;

  clearIntegerArrayValueList(): void;
  getIntegerArrayValueList(): Array<number>;
  setIntegerArrayValueList(value: Array<number>): void;
  addIntegerArrayValue(value: number, index?: number): number;

  clearDoubleArrayValueList(): void;
  getDoubleArrayValueList(): Array<number>;
  setDoubleArrayValueList(value: Array<number>): void;
  addDoubleArrayValue(value: number, index?: number): number;

  clearStringArrayValueList(): void;
  getStringArrayValueList(): Array<string>;
  setStringArrayValueList(value: Array<string>): void;
  addStringArrayValue(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ParameterValue.AsObject;
  static toObject(includeInstance: boolean, msg: ParameterValue): ParameterValue.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ParameterValue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ParameterValue;
  static deserializeBinaryFromReader(message: ParameterValue, reader: jspb.BinaryReader): ParameterValue;
}

export namespace ParameterValue {
  export type AsObject = {
    type: number,
    boolValue: boolean,
    integerValue: number,
    doubleValue: number,
    stringValue: string,
    byteArrayValue: Uint8Array | string,
    boolArrayValueList: Array<boolean>,
    integerArrayValueList: Array<number>,
    doubleArrayValueList: Array<number>,
    stringArrayValueList: Array<string>,
  }
}

export class SetParametersReq extends jspb.Message {
  clearParametersList(): void;
  getParametersList(): Array<Parameter>;
  setParametersList(value: Array<Parameter>): void;
  addParameters(value?: Parameter, index?: number): Parameter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetParametersReq.AsObject;
  static toObject(includeInstance: boolean, msg: SetParametersReq): SetParametersReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetParametersReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetParametersReq;
  static deserializeBinaryFromReader(message: SetParametersReq, reader: jspb.BinaryReader): SetParametersReq;
}

export namespace SetParametersReq {
  export type AsObject = {
    parametersList: Array<Parameter.AsObject>,
  }
}

export class GetParameterTypesRsp extends jspb.Message {
  clearTypesList(): void;
  getTypesList(): Array<number>;
  setTypesList(value: Array<number>): void;
  addTypes(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetParameterTypesRsp.AsObject;
  static toObject(includeInstance: boolean, msg: GetParameterTypesRsp): GetParameterTypesRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetParameterTypesRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetParameterTypesRsp;
  static deserializeBinaryFromReader(message: GetParameterTypesRsp, reader: jspb.BinaryReader): GetParameterTypesRsp;
}

export namespace GetParameterTypesRsp {
  export type AsObject = {
    typesList: Array<number>,
  }
}

export class ListParametersResult extends jspb.Message {
  clearNamesList(): void;
  getNamesList(): Array<string>;
  setNamesList(value: Array<string>): void;
  addNames(value: string, index?: number): string;

  clearPrefixesList(): void;
  getPrefixesList(): Array<string>;
  setPrefixesList(value: Array<string>): void;
  addPrefixes(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListParametersResult.AsObject;
  static toObject(includeInstance: boolean, msg: ListParametersResult): ListParametersResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListParametersResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListParametersResult;
  static deserializeBinaryFromReader(message: ListParametersResult, reader: jspb.BinaryReader): ListParametersResult;
}

export namespace ListParametersResult {
  export type AsObject = {
    namesList: Array<string>,
    prefixesList: Array<string>,
  }
}

export class DescribeParametersRsp extends jspb.Message {
  clearDescriptorsList(): void;
  getDescriptorsList(): Array<ParameterDescriptor>;
  setDescriptorsList(value: Array<ParameterDescriptor>): void;
  addDescriptors(value?: ParameterDescriptor, index?: number): ParameterDescriptor;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DescribeParametersRsp.AsObject;
  static toObject(includeInstance: boolean, msg: DescribeParametersRsp): DescribeParametersRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DescribeParametersRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DescribeParametersRsp;
  static deserializeBinaryFromReader(message: DescribeParametersRsp, reader: jspb.BinaryReader): DescribeParametersRsp;
}

export namespace DescribeParametersRsp {
  export type AsObject = {
    descriptorsList: Array<ParameterDescriptor.AsObject>,
  }
}

export class ParameterDescriptor extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getType(): number;
  setType(value: number): void;

  getDescription(): string;
  setDescription(value: string): void;

  getAdditionalConstraints(): string;
  setAdditionalConstraints(value: string): void;

  getReadOnly(): boolean;
  setReadOnly(value: boolean): void;

  getDynamicTyping(): boolean;
  setDynamicTyping(value: boolean): void;

  clearFloatingPointRangeList(): void;
  getFloatingPointRangeList(): Array<FloatingPointRange>;
  setFloatingPointRangeList(value: Array<FloatingPointRange>): void;
  addFloatingPointRange(value?: FloatingPointRange, index?: number): FloatingPointRange;

  clearIntegerRangeList(): void;
  getIntegerRangeList(): Array<IntegerRange>;
  setIntegerRangeList(value: Array<IntegerRange>): void;
  addIntegerRange(value?: IntegerRange, index?: number): IntegerRange;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ParameterDescriptor.AsObject;
  static toObject(includeInstance: boolean, msg: ParameterDescriptor): ParameterDescriptor.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ParameterDescriptor, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ParameterDescriptor;
  static deserializeBinaryFromReader(message: ParameterDescriptor, reader: jspb.BinaryReader): ParameterDescriptor;
}

export namespace ParameterDescriptor {
  export type AsObject = {
    name: string,
    type: number,
    description: string,
    additionalConstraints: string,
    readOnly: boolean,
    dynamicTyping: boolean,
    floatingPointRangeList: Array<FloatingPointRange.AsObject>,
    integerRangeList: Array<IntegerRange.AsObject>,
  }
}

export class FloatingPointRange extends jspb.Message {
  getFromValue(): number;
  setFromValue(value: number): void;

  getToValue(): number;
  setToValue(value: number): void;

  getStep(): number;
  setStep(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FloatingPointRange.AsObject;
  static toObject(includeInstance: boolean, msg: FloatingPointRange): FloatingPointRange.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FloatingPointRange, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FloatingPointRange;
  static deserializeBinaryFromReader(message: FloatingPointRange, reader: jspb.BinaryReader): FloatingPointRange;
}

export namespace FloatingPointRange {
  export type AsObject = {
    fromValue: number,
    toValue: number,
    step: number,
  }
}

export class GetParametersRsp extends jspb.Message {
  clearValuesList(): void;
  getValuesList(): Array<ParameterValue>;
  setValuesList(value: Array<ParameterValue>): void;
  addValues(value?: ParameterValue, index?: number): ParameterValue;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetParametersRsp.AsObject;
  static toObject(includeInstance: boolean, msg: GetParametersRsp): GetParametersRsp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetParametersRsp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetParametersRsp;
  static deserializeBinaryFromReader(message: GetParametersRsp, reader: jspb.BinaryReader): GetParametersRsp;
}

export namespace GetParametersRsp {
  export type AsObject = {
    valuesList: Array<ParameterValue.AsObject>,
  }
}

export class GetParameterTypesReq extends jspb.Message {
  clearNamesList(): void;
  getNamesList(): Array<string>;
  setNamesList(value: Array<string>): void;
  addNames(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetParameterTypesReq.AsObject;
  static toObject(includeInstance: boolean, msg: GetParameterTypesReq): GetParameterTypesReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetParameterTypesReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetParameterTypesReq;
  static deserializeBinaryFromReader(message: GetParameterTypesReq, reader: jspb.BinaryReader): GetParameterTypesReq;
}

export namespace GetParameterTypesReq {
  export type AsObject = {
    namesList: Array<string>,
  }
}

export class ListParametersReq extends jspb.Message {
  clearPrefixesList(): void;
  getPrefixesList(): Array<string>;
  setPrefixesList(value: Array<string>): void;
  addPrefixes(value: string, index?: number): string;

  getDepth(): number;
  setDepth(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListParametersReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListParametersReq): ListParametersReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListParametersReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListParametersReq;
  static deserializeBinaryFromReader(message: ListParametersReq, reader: jspb.BinaryReader): ListParametersReq;
}

export namespace ListParametersReq {
  export type AsObject = {
    prefixesList: Array<string>,
    depth: number,
  }
}

