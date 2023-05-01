from hulo import *


@Call(fn="SayHello")
def sayHello(request: hulo_pb2.CallRequest):
    print("recv: ", request)
    return hulo_pb2.CallResponse(Ok=True, Buf="I'm Python")


HuloBuilder.run()
