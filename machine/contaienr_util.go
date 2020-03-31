package machine

import (
	"context"
	"strings"
)

func checkStatus(id string, target ...string) bool {
	_, ok := db.getItem(id)
	if !ok {
		return false
	}

	inspectInfo, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		logger.WithError(err).Error("failed to get container inspect info")
		return false
	}

	source := inspectInfo.State.Status
	for _, s := range target {
		if s == source {
			return true
		}
	}

	return false
}


func envSliceToMap(s []string) map[string]string {
	env := map[string]string{}
	for _, str := range s {
		t := strings.SplitN(str, "=", 2)
		if t == nil || t[1] == "" {
			continue
		}
		env[t[0]] = t[1]
	}
	return env
}
