# google-grpc
This is a sample grpc service

This can be tested using postman. Ref: https://learning.postman.com/docs/sending-requests/grpc/first-grpc-request/

Commands:
1. make // This generates a folder named demo and creates contents of protobufs in it.
2. go run main.go // Listens on port 9000 for requests.
Optional:
3. sh build.sh // creates a binary of this sample service
4. docker tag <name> <image-name>
5. docker push <image-name>
