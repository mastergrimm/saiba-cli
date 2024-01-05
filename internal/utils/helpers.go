package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func GotoDir(dir string) {
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	cleanDir, err := filepath.Abs(filepath.Clean(dir))
	if err != nil {
		return
	}

	if currentDir == cleanDir {
		return
	}

	if _, err := os.Stat(cleanDir); os.IsNotExist(err) {
		return
	}

	os.Chdir(cleanDir)
}

func CloneAndCopySubdir(repo, subdir, destination string) error {
	tempDir, err := os.MkdirTemp("", "repo-")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	cloneCmd := exec.Command("git", "clone", "--quiet", repo, tempDir)
	if err := runCommand(cloneCmd); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	srcPath := filepath.Join(tempDir, subdir)
	if err := copyDir(srcPath, destination); err != nil {
		return fmt.Errorf("failed to copy subdirectory: %w", err)
	}

	return nil
}

func runCommand(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyDir(src, dst string) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
