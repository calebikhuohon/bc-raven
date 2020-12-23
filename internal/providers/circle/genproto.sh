#To give this file the necessary permissions on Mac or Linux, change terminal directory to providers/kudabank and run `chmod +x genproto.sh`
#next, ensure the protobuf-compiler is installed on your PC. Installation can be confirmed by running the command
# `protoc --version`
#Finally, to generate the probuf definitions, run this bash script with in the terminal

PATH=$PATH:$GOPATH/bin
protodir=proto

protoc --go_out=plugins=grpc:genproto -I $protodir $protodir/req.proto
