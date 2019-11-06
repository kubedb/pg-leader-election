module kubedb.dev/pg-leader-election

go 1.12

require (
	cloud.google.com/go v0.41.0 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.4.12 // indirect
	github.com/Azure/azure-pipeline-go v0.1.9 // indirect
	github.com/Azure/azure-storage-blob-go v0.6.0 // indirect
	github.com/appscode/go v0.0.0-20191016085057-e186b6c94a3b
	github.com/go-ini/ini v1.25.4 // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/martian v2.1.1-0.20190517191504-25dcb96d9e51+incompatible // indirect
	github.com/jackc/pgx v3.3.0+incompatible // indirect
	github.com/lib/pq v1.1.0 // indirect
	github.com/prometheus/client_golang v0.9.4 // indirect
	github.com/spf13/cobra v0.0.5
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
	google.golang.org/api v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20190508193815-b515fa19cec8 // indirect
	k8s.io/api v0.0.0-20190503110853-61630f889b3c
	k8s.io/apimachinery v0.0.0-20190508063446-a3da69d3723c
	k8s.io/client-go v11.0.0+incompatible
	kmodules.xyz/client-go v0.0.0-20191023042933-b12d1ccfaf57
	kmodules.xyz/custom-resources v0.0.0-20190927035424-65fe358bb045
	kubedb.dev/apimachinery v0.13.0-rc.2
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.34.0
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.2+incompatible
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190811223248-5a95b2df4348
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190314001948-2899ed30580f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190314000054-4a91899592f4
	k8s.io/klog => k8s.io/klog v0.3.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190314000639-da8327669ac5
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190314001731-1bd6a4002213
	k8s.io/utils => k8s.io/utils v0.0.0-20190514214443-0a167cbac756
	sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v0.0.0-20190302045857-e85c7b244fd2
)
