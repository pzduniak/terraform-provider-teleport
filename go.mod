module github.com/gravitational/teleport-plugins/terraform

go 1.16

require (
	github.com/fatih/color v1.9.0 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/gravitational/protoc-gen-terraform v0.0.0-20211108170245-3b37ff28d21e // protoc-gen-terraform master (#13)
	github.com/gravitational/teleport/api v0.0.0-20211213214838-0e304c323c49 // tag v8.0.5
	github.com/gravitational/trace v1.1.16-0.20210609220119-4855e69c89fc
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/sirupsen/logrus v1.8.1-0.20210219125412-f104497f2b21
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	google.golang.org/genproto v0.0.0-20210223151946-22b48be4551b // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/gogo/protobuf => github.com/gravitational/protobuf v1.3.2-0.20201123192827-2b9fcfaffcbf
	github.com/julienschmidt/httprouter => github.com/rw-access/httprouter v1.3.1-0.20210321233808-98e93175c124
)
