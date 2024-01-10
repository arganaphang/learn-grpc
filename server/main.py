from concurrent import futures

import grpc
from user import user_pb2_grpc
from user import user_pb2 as user__pb2


class UserServiceServicer(user_pb2_grpc.UserServiceServicer):
    def __init__(self):
        pass

    def GetUsers(self, request, context):
        return user__pb2.GetUsersResponse(
            message="Get User List From Server",
            data=[
                # user__pb2.User(id=1, name="John Doe"),
                # user__pb2.User(id=2, name="Emilie"),
                # user__pb2.User(id=3, name="Kevin"),
                user__pb2.User(id=1, name="John Doe", email="johndoe@mail.com"),
                user__pb2.User(id=2, name="Emilie", email="emilie@mail.com"),
                user__pb2.User(id=3, name="Kevin", email="kevin@mail.com"),
            ],
        )


def run():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    user_pb2_grpc.add_UserServiceServicer_to_server(UserServiceServicer(), server)
    server.add_insecure_port("[::]:8001")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    run()
