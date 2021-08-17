module github.com/mongodb/mongodb-atlas-kubernetes

go 1.15

require (
	github.com/fatih/structtag v1.2.0
	github.com/go-logr/zapr v0.4.0
	github.com/google/go-cmp v0.5.5
	github.com/mongodb-forks/digest v1.0.2
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/pborman/uuid v1.2.1
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/atlas v0.7.2
	go.uber.org/zap v1.18.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	sigs.k8s.io/controller-runtime v0.9.6
)
