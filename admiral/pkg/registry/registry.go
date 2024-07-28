package registry

import (
	"encoding/json"
	"os"

	"github.com/istio-ecosystem/admiral/admiral/pkg/controller/common"
	log "github.com/sirupsen/logrus"
	networkingV1Alpha3 "istio.io/api/networking/v1alpha3"
	coreV1 "k8s.io/api/core/v1"
<<<<<<< HEAD
=======
	"os"
	"strings"
>>>>>>> 508caceb (MESH-5069: Operator Shards (#749))
)

// IdentityConfiguration is an interface to fetch configuration from a registry
// backend. The backend can provide an API to give configurations per identity,
// or if given a cluster name, it will provide the configurations for all
// the identities present in that cluster.
type IdentityConfiguration interface {
	GetIdentityConfigByIdentityName(identityAlias string, ctxLogger *log.Entry) (IdentityConfig, error)
	GetIdentityConfigByClusterName(clusterName string, ctxLogger *log.Entry) ([]IdentityConfig, error)
}

type registryClient struct {
	registryEndpoint string
}

func NewRegistryClient(options ...func(client *registryClient)) *registryClient {
	registryClient := &registryClient{}
	for _, o := range options {
		o(registryClient)
	}
	return registryClient
}

func WithRegistryEndpoint(registryEndpoint string) func(*registryClient) {
	return func(c *registryClient) {
		c.registryEndpoint = registryEndpoint
	}
}

type IdentityConfig struct {
	IdentityName string                  `json:"identityName"`
	Clusters     []IdentityConfigCluster `json:"clusters"`
	ClientAssets []map[string]string     `json:"clientAssets"`
}

type IdentityConfigCluster struct {
	Name            string                      `json:"name"`
	Locality        string                      `json:"locality"`
	IngressEndpoint string                      `json:"ingressEndpoint"`
	IngressPort     string                      `json:"ingressPort"`
	IngressPortName string                      `json:"ingressPortName"`
	Environment     []IdentityConfigEnvironment `json:"environment"`
}

type IdentityConfigEnvironment struct {
	Name          string                           `json:"name"`
	Namespace     string                           `json:"namespace"`
	ServiceName   string                           `json:"serviceName"`
	Type          string                           `json:"type"`
	Selectors     map[string]string                `json:"selectors"`
	Ports         []coreV1.ServicePort             `json:"ports"`
	TrafficPolicy networkingV1Alpha3.TrafficPolicy `json:"trafficPolicy"`
}

// GetIdentityConfigByIdentityName calls the registry API to fetch the IdentityConfig for
// the given identityAlias
func (c *registryClient) GetIdentityConfigByIdentityName(identityAlias string, ctxLogger *log.Entry) (IdentityConfig, error) {
	//TODO: Use real result from registry and remove string splitting to match test file names
	byteValue, err := readIdentityConfigFromFile(strings.Split(identityAlias, "."))
	if err != nil {
		ctxLogger.Infof(common.CtxLogFormat, "GetByIdentityName", identityAlias, "", "", err)
	}
	var identityConfigUnmarshalResult IdentityConfig
	err = json.Unmarshal(byteValue, &identityConfigUnmarshalResult)
	if err != nil {
		ctxLogger.Infof(common.CtxLogFormat, "GetByIdentityName", identityAlias, "", "", err)
	}
	return identityConfigUnmarshalResult, err
}

func readIdentityConfigFromFile(shortAlias []string) ([]byte, error) {
	pathName := "testdata/" + shortAlias[len(shortAlias)-1] + "IdentityConfiguration.json"
	if common.GetSecretFilterTags() == "admiral/syncrtay" {
		pathName = "/etc/serviceregistry/config/" + shortAlias[len(shortAlias)-1] + "IdentityConfiguration.json"
	}
	return os.ReadFile(pathName)
}

// GetIdentityConfigByClusterName calls the registry API to fetch the IdentityConfigs for
// every identity on the cluster.
func (c *registryClient) GetIdentityConfigByClusterName(clusterName string, ctxLogger *log.Entry) ([]IdentityConfig, error) {
	//TODO: need to call this function once during startup time to warm the cache
	//jsonResult = os.request(/cluster/{cluster_id}/configurations
	ctxLogger.Infof(common.CtxLogFormat, "GetByClusterName", "", "", clusterName, "")
	//identities := getIdentitiesForCluster(clusterName) - either queries shard CRD or shard CRD controller calls this func with those as parameters
	identities := []string{clusterName}
	identityConfigs := []IdentityConfig{}
	var err error
	for _, identity := range identities {
		identityConfig, identityErr := c.GetIdentityConfigByIdentityName(identity, ctxLogger)
		if identityErr != nil {
			err = identityErr
			ctxLogger.Infof(common.CtxLogFormat, "GetByClusterName", "", "", clusterName, identityErr)
		}
		identityConfigs = append(identityConfigs, identityConfig)
	}
	return identityConfigs, err
}
