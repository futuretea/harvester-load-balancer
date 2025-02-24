package util

const (
	prefix                         = "harvesterhci.io"
	RemovedPVCsAnnotationKey       = prefix + "/removedPersistentVolumeClaims"
	AdditionalCASecretName         = "harvester-additional-ca"
	AdditionalCAFileName           = "additional-ca.pem"
	AnnotationMigrationTarget      = prefix + "/migrationTargetNodeName"
	AnnotationMigrationUID         = prefix + "/migrationUID"
	AnnotationMigrationState       = prefix + "/migrationState"
	AnnotationTimestamp            = prefix + "/timestamp"
	AnnotationVolumeClaimTemplates = prefix + "/volumeClaimTemplates"
	AnnotationImageID              = prefix + "/imageId"
	AnnotationReservedMemory       = prefix + "/reservedMemory"
	AnnotationHash                 = prefix + "/hash"
	AnnotationRunStrategy          = prefix + "/vmRunStrategy"
	LabelImageDisplayName          = prefix + "/imageDisplayName"
	LabelVMName                    = prefix + "/vmName"

	AnnotationStorageClassName          = prefix + "/storageClassName"
	AnnotationStorageProvisioner        = prefix + "/storageProvisioner"
	AnnotationIsDefaultStorageClassName = "storageclass.kubernetes.io/is-default-class"

	// AnnotationMigratingPrefix is used to store the migrating vm in the annotation of ResourceQuota
	// eg: harvesterhci.io/migrating-vm1: jsonOfResourceList, harvesterhci.io/migrating-vm2: jsonOfResourceList
	AnnotationMigratingPrefix = prefix + "/migrating-"

	AnnotationDefaultUserdataSecret = prefix + "/default-userdata-secret"

	ContainerdRegistrySecretName = "harvester-containerd-registry"
	ContainerdRegistryFileName   = "registries.yaml"

	BackupTargetSecretName            = "harvester-backup-target-secret"
	InternalTLSSecretName             = "tls-rancher-internal"
	Rke2IngressNginxAppName           = "rke2-ingress-nginx"
	CattleSystemNamespaceName         = "cattle-system"
	LonghornSystemNamespaceName       = "longhorn-system"
	LonghornDefaultManagerURL         = "http://longhorn-backend.longhorn-system:9500/v1"
	KubeSystemNamespace               = "kube-system"
	FleetLocalNamespaceName           = "fleet-local"
	LocalClusterName                  = "local"
	HarvesterSystemNamespaceName      = "harvester-system"
	RancherLoggingName                = "rancher-logging"
	CattleLoggingSystemNamespaceName  = "cattle-logging-system"
	HarvesterUpgradeImageRepository   = "rancher/harvester-upgrade"
	HarvesterLonghornStorageClassName = "harvester-longhorn"

	HTTPProxyEnv  = "HTTP_PROXY"
	HTTPSProxyEnv = "HTTPS_PROXY"
	NoProxyEnv    = "NO_PROXY"

	LonghornOptionBackingImageName = "backingImage"
	LonghornOptionMigratable       = "migratable"
	AddonValuesAnnotation          = "harvesterhci.io/addon-defaults"

	LabelUpgradeReadMessage          = prefix + "/read-message"
	LabelUpgradeState                = prefix + "/upgradeState"
	UpgradeStateLoggingInfraPrepared = "LoggingInfraPrepared"

	AnnotationArchiveName         = prefix + "/archiveName"
	LabelUpgradeLog               = prefix + "/upgradeLog"
	LabelUpgradeLogComponent      = prefix + "/upgradeLogComponent"
	UpgradeLogInfraComponent      = "infra"
	UpgradeLogShipperComponent    = "shipper"
	UpgradeLogAggregatorComponent = "aggregator"
	UpgradeLogDownloaderComponent = "downloader"
	UpgradeLogFlowComponent       = "clusterflow"
	UpgradeLogArchiveComponent    = "log-archive"
	UpgradeLogOperatorComponent   = "operator"
	UpgradeLogOutputComponent     = "clusteroutput"
	UpgradeLogPackagerComponent   = "packager"

	UpgradeNodeDrainTaintKey   = "kubevirt.io/drain"
	UpgradeNodeDrainTaintValue = "draining"

	FieldCattlePrefix             = "field.cattle.io"
	CattleAnnotationResourceQuota = FieldCattlePrefix + "/resourceQuota"

	ManagementCattlePrefix              = "management.cattle.io"
	LabelManagementDefaultResourceQuota = "resourcequota." + ManagementCattlePrefix + "/default-resource-quota"
)
