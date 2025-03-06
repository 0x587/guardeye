const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

const {parse: foxglove_parse} = require("@foxglove/rosmsg");
const {
    MessageReader,
    MessageWriter,
} = require("@foxglove/rosmsg2-serialization");

const msgpackr = require("@msgpack/msgpack");
var JSONbig = require('json-bigint');
const BigNumber = require("bignumber.js");


// 加载 proto 文件
const packageDefinition = protoLoader.loadSync("proto/foxglove_service.proto", {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const FoxgloveService = protoDescriptor.FoxgloveService;

function parse(v) {
    return foxglove_parse(v, {ros2: true});
}

function read(d, v) {
    const l = v.length;
    const ab = new ArrayBuffer(l);
    const view = new Uint8Array(ab);
    for (let i = 0; i < l; i++) {
        view[i] = v[i];
    }
    const dv = new DataView(ab, 0, l);
    const reader = new MessageReader(d);
    const res = reader.readMessage(dv)
    return res;
}

function write(d, v) {
    const writer = new MessageWriter(d);
    const size = writer.calculateByteSize(v);
    const ab = new ArrayBuffer(size);
    const res = writer.writeMessage(v, ab);
    return res;
}

function convertTypedArrays(obj) {
    if (ArrayBuffer.isView(obj) && !(obj instanceof DataView)) {
        // 将 TypedArray 转换为普通数组
        return Array.from(obj);
    } else if (Array.isArray(obj)) {
        // 递归处理数组
        return obj.map(convertTypedArrays);
    } else if (typeof obj === "object" && obj !== null) {
        // 递归处理对象
        return Object.fromEntries(
            Object.entries(obj).map(([key, value]) => [key, convertTypedArrays(value)])
        );
    } else {
        // 其他类型直接返回
        return obj;
    }
}
function convertBigNumberToBigInt(obj) {
    if (obj && typeof obj === "object" && obj instanceof BigNumber) {
        // 识别 BigNumber 并转换为 BigInt
        return BigInt(obj)
    } else if (Array.isArray(obj)) {
        // 如果是数组，递归处理每个元素
        return obj.map(convertBigNumberToBigInt);
    } else if (typeof obj === "object" && obj !== null) {
        // 如果是对象，递归处理每个属性
        return Object.fromEntries(
            Object.entries(obj).map(([key, value]) => [key, convertBigNumberToBigInt(value)])
        );
    }
    // 其他类型直接返回
    return obj;
}

// 实现服务方法
function CdrRead(call, callback) {
    const d = parse(call.request.ros_schema)
    const v = call.request.buf
    let res = read(d, v)

    res = convertTypedArrays(res)
    const strRes = JSONbig.stringify(res)
    console.log('------------READ----------------')
    console.log(res)
    console.log(strRes)
    callback(null, {trans_data: strRes})
}


function CdrWrite(call, callback) {
    const d = parse(call.request.ros_schema)
    let v = JSONbig.parse(call.request.trans_data)
    v = convertBigNumberToBigInt(v)
    console.log('------------WRITE----------------')
    console.log(call.request.trans_data)
    console.log(v)
    const res = write(d, v)
    callback(null, {buf: res})
}

// 启动 gRPC 服务器
function main() {
    const server = new grpc.Server();
    server.addService(FoxgloveService.FoxgloveService.service,
        {
            CdrRead,
            CdrWrite
        });

    const address = "0.0.0.0:50051";
    server.bindAsync(address, grpc.ServerCredentials.createInsecure(), () => {
        console.log(`gRPC server running at ${address}`);
    });
}

main();
