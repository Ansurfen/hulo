import time
from concurrent import futures
import grpc
import hulo_pb2_grpc
import hulo_pb2
from typing import Callable
import argparse

parser = argparse.ArgumentParser(
    description='Start a grpc server on the specified port.')
parser.add_argument('-p', '--port', metavar='port', type=int, default=0,
                    help='the port number to start the server on (default: 0)')
args = parser.parse_args()

if args.port == 0:
    raise "invalid port"


class Hulo(hulo_pb2_grpc.HuloInterface):
    funcs: dict[str, Callable[[hulo_pb2.CallRequest],
                              hulo_pb2.CallResponse]]

    def __init__(self) -> None:
        self.funcs = {}

    def Completion(self, request: hulo_pb2.CompletionRequest, context) -> hulo_pb2.CompletionResponse:
        return hulo_pb2.CompletionResponse()

    def Call(self, request: hulo_pb2.CallRequest, context) -> hulo_pb2.CallResponse:
        if request.Func in self.funcs:
            return self.funcs[request.Func](request)
        return hulo_pb2.CallResponse(Ok=False, Buf="unknown")

    def run(self) -> None:
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        hulo_pb2_grpc.add_HuloInterfaceServicer_to_server(
            self, server)
        server.add_insecure_port('[::]:{port}'.format(port=args.port))
        server.start()
        print("start service...")
        try:
            while True:
                time.sleep(60 * 60 * 24)
        except KeyboardInterrupt:
            server.stop(0)

    def registerCallback(self, name, cb):
        self.funcs[name] = cb


HuloBuilder = Hulo()


def Call(fn: str):
    def wrapper(func: Callable[[hulo_pb2.CallRequest], hulo_pb2.CallResponse]):
        HuloBuilder.registerCallback(fn, func)
    return wrapper


def HuloDialCall(port: int, req: hulo_pb2.CallRequest):
    with grpc.insecure_channel('localhost:{port}'.format(port=port)) as channel:
        stub = hulo_pb2_grpc.HuloInterfaceStub(channel)
        res = stub.Callback(req)
        print(res)
