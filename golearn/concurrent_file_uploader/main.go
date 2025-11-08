package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var wg = sync.WaitGroup{}

func checkfileexist(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func createsshconnection(user string, password string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

func copyFileToDestination(destinationpath string, file_path string, sftpClient *sftp.Client) {
	localFile, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Failed to open local file:", err)
		return
	}
	defer localFile.Close()

	remoteFile, err := sftpClient.Create(filepath.Join(destinationpath, filepath.Base(file_path)))
	if err != nil {
		fmt.Println("Failed to create remote file:", err)
		return
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		fmt.Println("Failed to upload file:", err)
		return
	}
}

func scpfiletovm(file_path string) {
	if !checkfileexist(file_path) {
		fmt.Printf("file %s does not exist\n", file_path)
		return
	}
	host := "10.50.2.60"
	port := 22
	user := "devuser"
	password := "CrestLogin@24"
	destinationpath := "/home/devuser"

	config := createsshconnection(user, password)
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		fmt.Println("Failed to connect to SSH server:", err)
		return
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println("Failed to open SFTP session:", err)
		return
	}
	defer sftpClient.Close()

	copyFileToDestination(destinationpath, file_path, sftpClient)
	defer wg.Done()
}

func main() {
	files_paths := os.Args[1:]
	sharedsftpimplementation(files_paths)
	// individualsftpimplementation(files_paths)
}

func sharedsftpimplementation(files_paths []string) {

	// Create a single connection
	host := "10.50.2.60"
	port := 22
	user := "devuser"
	password := "CrestLogin@24"
	destinationpath := "/home/devuser"

	config := createsshconnection(user, password)
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		fmt.Println("Failed to connect to SSH server:", err)
		return
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println("Failed to open SFTP session:", err)
		return
	}
	defer sftpClient.Close()

	start_time := time.Now()

	// Use a mutex to protect concurrent access to the SFTP client
	var mu sync.Mutex

	for _, file_path := range files_paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			if !checkfileexist(path) {
				fmt.Printf("file %s does not exist\n", path)
				return
			}

			// Lock when accessing the shared SFTP client
			mu.Lock()
			copyFileToDestination(destinationpath, path, sftpClient)
			mu.Unlock()
		}(file_path)
	}

	wg.Wait()
	end_time := time.Now()
	fmt.Printf("Total time taken %v\n", end_time.Sub(start_time))
}

func individualsftpimplementation(files_paths []string) {
	start_time := time.Now()
	for _, file_path := range files_paths {
		wg.Add(1)
		go scpfiletovm(file_path)
	}
	end_time := time.Now()
	wg.Wait()
	fmt.Printf("Total time taken %v\n", end_time.Sub(start_time))
}
