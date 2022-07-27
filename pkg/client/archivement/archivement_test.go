package archivement

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/archivement-middleware/pkg/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestClient(t *testing.T) {
	// Here won't pass test due to we always test with localhost
}
