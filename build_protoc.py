import os
import sys


def run_protoc():
    protoc_command = (
        f"{sys.executable} -m grpc_tools.protoc "
        "--proto_path=./api "  # proto文件的根目录
        "--python_out=./api "  # 生成Python代码的输出目录
        "--grpc_python_out=./api "  # 生成gRPC Python代码的输出目录
        "./api/hello.proto "
    )

    # 执行protoc命令
    os.system(protoc_command)


if __name__ == "__main__":
    run_protoc()
