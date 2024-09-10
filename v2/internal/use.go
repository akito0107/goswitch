package internal

import (
	"bytes"
	"context"
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const goGetURL = "golang.org/dl/"

func Use(ctx context.Context, version string) error {

	gobin := gobinPath()
	symlinkPath := filepath.Join(gobin, "go")

	if version == "system" {
		log.Println("using system go")
		log.Println("removing symlink...")
		return rmSymlink(symlinkPath)
	}

	if !versionExists(version) {
		v, err := currentGoVersion(ctx)
		if err != nil {
			return fmt.Errorf("get current go version: %w", err)
		}
		var cmd *exec.Cmd
		if v.Major() > 1 || (v.Major() == 1 && v.Minor() >= 17) {
			cmd = exec.CommandContext(ctx, "go", "install", goGetURL+version+"@latest")
		} else {
			cmd = exec.CommandContext(ctx, "go", "get", goGetURL+version)
			cmd.Env = append(os.Environ(), "GO111MODULE=off")
		}
		sub := cmd.Args[1]
		log.Printf("start go %s...", sub)
		log.Print(strings.Join(cmd.Args, " "))

		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("go %s failed: %w", sub, err)
		}
		log.Printf("go %s finished. start download", sub)
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

	if err := rmSymlink(symlinkPath); err != nil {
		return err
	}

	if err := os.Symlink(filepath.Join(gobin, version), symlinkPath); err != nil {
		return fmt.Errorf("create symlink failed: %w", err)
	}

	return nil
}

func rmSymlink(symlinkPath string) error {
	if info, err := os.Lstat(symlinkPath); err == nil {
		if info.Mode()&os.ModeSymlink != os.ModeSymlink {
			return fmt.Errorf("go command is not a symlink but %s, goswitch use symlink. please delete or make a symlink", info.Mode())
		}

		if err := os.Remove(symlinkPath); err != nil {
			return fmt.Errorf("remove go link failed: %w", err)
		}
	}

	return nil
}

func gobinPath() string {
	gobin := os.Getenv("GOBIN")

	if gobin == "" {
		gobin = filepath.Join(build.Default.GOPATH, "bin")
	}

	return gobin
}

func versionExists(version string) bool {
	gobin := gobinPath()
	_, err := os.Stat(filepath.Join(gobin, version))

	return err == nil
}

func currentGoVersion(ctx context.Context) (goversion, error) {
	b, err := exec.CommandContext(ctx, "go", "env", "GOVERSION").Output()
	if err != nil {
		return "", fmt.Errorf("exec go env GOVERSION: %w", err)
	}
	return goversion(bytes.TrimSuffix(b, []byte("\n"))), nil
}
