PWD 	:= $(shell pwd)

OS_NAME := $(shell uname -s | tr A-Z a-z)

check:
	@[ "${module}" ] && echo "Selected module is ${module}" || (echo "module list: "; ls $(PWD)/proto ; echo "ex) make go module=core" ; exit 1)

go_pre:
	@which protoc protoc-gen-go protoc-gen-go-grpc protoc-gen-grpc-gateway protoc-gen-validate 1>/dev/null || (env bash $(PWD)/scripts/proto_install_go.sh ${OS_NAME})

swift_pre:
	@which protoc protoc-gen-swift protoc-gen-grpc-swift > /dev/null 2>&1 || (env bash $(PWD)/scripts/proto_install_swift.sh ${OS_NAME})

kotlin_pre:
	@which protoc protoc-gen-grpc-kotlin.jar > /dev/null 2>&1 || (env bash $(PWD)/scripts/proto_install_kotlin.sh ${OS_NAME})

go: check go_pre
	@(env bash $(PWD)/scripts/proto_gen_go.sh)

swift: check swift_pre
	@(env bash $(PWD)/scripts/proto_gen_swift.sh)

kotlin: check kotlin_pre
	@(env bash $(PWD)/scripts/proto_gen_kotlin.sh)

openapi:
	@(env bash $(PWD)/scripts/proto_gen_openapi.sh)

postman:
	@(env bash $(PWD)/scripts/proto_gen_postman.sh)

clean:
	@echo "Cleaning up all the generated files"