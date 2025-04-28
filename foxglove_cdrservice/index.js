const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

const {parse: foxglove_parse} = require("@foxglove/rosmsg");
const {
    MessageReader,
    MessageWriter,
} = require("@foxglove/rosmsg2-serialization");

var JSONbig = require('json-bigint');
const BigNumber = require("bignumber.js");
const Long = require("long");
const tempfs = require("temp-fs");
const fs = require("fs");
const protobuf = require("protobufjs");


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

function convertBigIntToLong(obj) {
    if (obj && typeof obj === "bigint") {
        // 识别 BigInt 并转换为 Long
        return Long.fromBigInt(obj)
    } else if (Array.isArray(obj)) {
        // 如果是数组，递归处理每个元素
        return obj.map(convertBigIntToLong);
    } else if (typeof obj === "object" && obj !== null) {
        // 如果是对象，递归处理每个属性
        return Object.fromEntries(
            Object.entries(obj).map(([key, value]) => [key, convertBigIntToLong(value)])
        );
    }
    // 其他类型直接返回
    return obj;
}

function convertBigNumberAndLongToBigInt(obj) {
    if (obj && typeof obj === "object" && (obj instanceof BigNumber || obj instanceof Long)) {
        // 识别 BigNumber 并转换为 BigInt
        return BigInt(obj)
    } else if (Array.isArray(obj)) {
        // 如果是数组，递归处理每个元素
        return obj.map(convertBigNumberAndLongToBigInt);
    } else if (typeof obj === "object" && obj !== null) {
        // 如果是对象，递归处理每个属性
        return Object.fromEntries(
            Object.entries(obj).map(([key, value]) => [key, convertBigNumberAndLongToBigInt(value)])
        );
    }
    // 其他类型直接返回
    return obj;
}

function convertKeysToCamelCase(obj) {
    function toCamelCase(str) {
        return str.replace(/_([a-z])/g, (_, char) => char.toUpperCase());
    }
    if (Array.isArray(obj)) {
        return obj.map(convertKeysToCamelCase);
    } else if (obj !== null && typeof obj === "object") {
        return Object.entries(obj).reduce((acc, [key, value]) => {
            const camelKey = toCamelCase(key);
            acc[camelKey] = convertKeysToCamelCase(value);
            return acc;
        }, {});
    }
    return obj;
}

function camelToSnake(obj) {
    const newObj = {};
    for (const key in obj) {
        if (obj.hasOwnProperty(key)) {
            const snakeKey = key.replace(/([A-Z])/g, '_$1').toLowerCase();
            newObj[snakeKey] = obj[key];
        }
    }
    return newObj;
}

// 实现服务方法
function CdrRead(call, callback) {
    const {
        ros_schema,
        pb_schema,
        pb_type_name,
        cdr_data
    } = call.request;
    tempfs.open(function (err, file) {
        if (err)
            throw err
        fs.writeFileSync(file.path, pb_schema);
        protobuf.load(file.path, function (err, root) {
            if (err) {
                throw err;
            }
            const Msg = root.lookupType(pb_type_name);
            const d = parse(ros_schema)
            const v = cdr_data
            let res = read(d, v)
            res = convertTypedArrays(res)
            const jsonRes = JSONbig.stringify(res)
            res = convertBigIntToLong(res)
            const pbRes = Msg.encode(convertKeysToCamelCase(res)).finish()
            console.log('------------READ----------------')
            console.log(res)
            console.log(jsonRes)
            console.log(pbRes)
            callback(null, {trans_data: pbRes, json_data: jsonRes})
            file.unlink();
        });
    });
}

function CdrWrite(call, callback) {
    const {
        ros_schema,
        pb_schema,
        pb_type_name,
        trans_data
    } = call.request;
    console.log(pb_type_name)
    console.log(pb_schema)
    tempfs.open(function (err, file) {
        if (err)
            throw err
        fs.writeFileSync(file.path, pb_schema);
        protobuf.load(file.path, function (err, root) {
            if (err) {
                throw err;
            }
            var Msg = root.lookupType(pb_type_name);
            const d = parse(ros_schema)
            let v = Msg.decode(trans_data)
            v = convertBigNumberAndLongToBigInt(v)
            console.log('------------WRITE----------------')
            console.log(trans_data)
            console.log(v)
            const res = write(d, camelToSnake(v))
            const jsonRes = JSONbig.stringify(v)
            callback(null, {cdr_data: res, json_data: jsonRes})
            file.unlink();
        });
    });
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
