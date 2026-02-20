import grpc
from api.hello_pb2_grpc import GreeterStub
from api.hello_pb2 import HelloRequest

def main():
    channel = grpc.insecure_channel("127.0.0.1:8080")

    stub = GreeterStub(channel)

    # 可携带 metadata（对应 Go 里的 metadata.FromIncomingContext）
    metadata = (
        ("version", "v1"),
    )

    resp = stub.SayHello(
        HelloRequest(name="python-client"),
        metadata=metadata,
        timeout=1.0,
    )

    print("Response:", resp.message)


if __name__ == "__main__":
    main()