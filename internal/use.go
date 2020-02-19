package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const goGetURL = "golang.org/dl/"

func Use(ctx context.Context, version string) error {

	if !versionExists(version) {
		getcmd := exec.CommandContext(ctx, "go", "get", goGetURL+version)
		log.Println("start go get...")
		log.Printf("go get %s%s\n", goGetURL, version)

		getcmd.Stdout = log.Writer()
		getcmd.Stderr = log.Writer()

		if err := getcmd.Run(); err != nil {
			return fmt.Errorf("go get failed: %w", err)
		}
		log.Println("go get finished. start download")
	} else {
		log.Printf("already exists version: %s, start download\n", version)
	}

	dlcmd := exec.CommandContext(ctx, version, "download")

	dlcmd.Stdout = log.Writer()
	dlcmd.Stderr = log.Writer()

	if err := dlcmd.Run(); err != nil {
		return fmt.Errorf("%s download failed: %w", version, err)
	}

	log.Println("download finished.")
	log.Println("switch go version")

	gobin := gobinPath()

	if info, err := os.Lstat(gobin + "/go"); err == nil {
		if info.Mode()&os.ModeSymlink != os.ModeSymlink {
			return fmt.Errorf("go command is not a symlink but %s, goswitch use symlink. please delete or make a symlink", info.Mode())
		}

		if err := os.Remove(gobin + "/go"); err != nil {
			return fmt.Errorf("remove go link failed: %w", err)
		}
	}

	if err := os.Symlink(gobin+"/"+version, gobin+"/go"); err != nil {
		return fmt.Errorf("create symlink failed: %w", err)
	}

	return nil
}

func gobinPath() string {
	gobin := os.Getenv("GOBIN")

	if gobin == "" {
		gp := os.Getenv("GOPATH")
		gobin = gp + "/bin"
	}

	return gobin
}

func versionExists(version string) bool {
	gobin := gobinPath()
	_, err := os.Stat(gobin + "/" + version)

	return err == nil
}
