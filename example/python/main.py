import grpc

import main_pb2_grpc
import pkg_a_pb2, pkg_d_pb2
import guardeye_python_sdk as sdk

channel = sdk.GuardeyeChannel('10.0.4.112:8080', "106bc8ee-6048-4024-8afb-c294ec8fd559")
cli = main_pb2_grpc.ApiStub(channel)
request = pkg_d_pb2.GetPersonReq(req=pkg_d_pb2.Person(v=pkg_a_pb2.V3(x=1)))
response = cli.CallServiceGet_person(request)

print(response)