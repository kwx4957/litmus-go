package environment

import (
	"strconv"

	experimentTypes "github.com/litmuschaos/litmus-go/pkg/load/locust-loadgen/types"
	"github.com/litmuschaos/litmus-go/pkg/types"
)

// GetENV fetches all the env variables from the runner pod
func GetENV(experimentDetails *experimentTypes.ExperimentDetails) {
	experimentDetails.ExperimentName = types.Getenv("EXPERIMENT_NAME", "locust-loadgen")
	experimentDetails.ChaosNamespace = types.Getenv("CHAOS_NAMESPACE", "litmus")
	experimentDetails.EngineName = types.Getenv("CHAOSENGINE", "")
	experimentDetails.ChaosDuration, _ = strconv.Atoi(types.Getenv("TOTAL_CHAOS_DURATION", "30"))
	experimentDetails.ChaosInterval, _ = strconv.Atoi(types.Getenv("CHAOS_INTERVAL", "10"))
	experimentDetails.RampTime, _ = strconv.Atoi(types.Getenv("RAMP_TIME", "0"))
	experimentDetails.Delay, _ = strconv.Atoi(types.Getenv("STATUS_CHECK_DELAY", "2"))
	experimentDetails.Timeout, _ = strconv.Atoi(types.Getenv("STATUS_CHECK_TIMEOUT", "180"))
	experimentDetails.LIBImage = types.Getenv("LIB_IMAGE", "locustio/locust")
	experimentDetails.LIBImagePullPolicy = types.Getenv("LIB_IMAGE_PULL_POLICY", "Always")
	experimentDetails.Host = types.Getenv("HOST", "http://my-nginx.default.svc.cluster.local/")
	experimentDetails.ConfigMapName = types.Getenv("CONFIG_MAP_FILE", "locust-script")
	experimentDetails.Users, _ = strconv.Atoi(types.Getenv("USERS", "40"))
	experimentDetails.SpawnRate, _ = strconv.Atoi(types.Getenv("SPAWN_RATE", "30"))
	experimentDetails.Replica, _ = strconv.Atoi(types.Getenv("REPLICA", "1"))
	experimentDetails.LoadType = types.Getenv("LOAD_TYPE", "load")
	experimentDetails.GrantType = types.Getenv("GRANT_TYPE", "")
	experimentDetails.NodeNames = types.Getenv("NODE_NAMES", "")
}
