package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v3"
)

func installDockerCli(ctx context.Context, cmd *cli.Command) error {
	say("Installing docker-cli...")
	var err error
	say("Updating apt-get package cache...")
	err = execSudoCommand(ctx, true, "apt-get", "update")
	if err != nil {
		return err
	}
	say("Getting required packages...")
	err = execSudoCommand(ctx, false, "apt-get", "install", "ca-certificates", "curl", "gnupg")
	if err != nil {
		return err
	}

	say("Setting up docker repository...")
	err = execSudoCommand(ctx, false, "install", "-m", "0755", "-d", "/etc/apt/keyrings")
	if err != nil {
		return err
	}
	err = downloadFile("https://download.docker.com/linux/ubuntu/gpg", "/tmp/docker-gpg")
	if err != nil {
		return err
	}
	err = execSudoCommand(ctx, false, "gpg","--dearmor", "-o", "/etc/apt/keyrings/docker.gpg", "/tmp/docker-gpg" )
	if err != nil {
		return err
	}
	err = execSudoCommand(ctx, false, "chmod", "0644", "/etc/apt/keyrings/docker.gpg")
	if err != nil {
		return err
	}
	arch := runtime.GOARCH
	dockerList := fmt.Sprintf(`deb [arch=amd64 signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu %v stable`, arch)
	listFile := "/etc/apt/sources.list.d/docker.list"
	if err = os.WriteFile(listFile, []byte(dockerList), 0644); err != nil {
		return err
	}
	err = execSudoCommand(ctx, true, "apt-get", "update")
	if err != nil {
		return err
	}
	err = execSudoCommand(ctx, true, "apt-get", "install", "--no-install-recommends", "-y", "docker-ce-cli")
	if err != nil {
		return err
	}
	return nil
}
