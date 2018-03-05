// +build integration

package dal

import (
	"github.com/vahdet/go-user-store-redis/models"
	"testing"
	dockerclient "github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
	"fmt"
	"io"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"os"
	"github.com/go-redis/redis"
	"strings"
	"strconv"
	"github.com/docker/go-connections/nat"
	"github.com/vahdet/go-user-store-redis/dal/interfaces"
	"time"
)

const (
	redisImageName = "redis"
	redisContainerName = "redisIntegrationTest"
)


//var client *redis.Client
var ctx context.Context
var dalInstance interfaces.UserDal

func TestMain(m *testing.M) {
	setupTest()
	retCode := m.Run()
	endTest()
	os.Exit(retCode)
}

func TestCreateUser(t *testing.T) {
	fmt.Println("Create Test started")

	//var foundUser *models.User

	dalInstance = NewUserDal()

	testUser := models.User{
		Id: 1,
		Name: "Freddie Mercury",
		Email: "freddie@heaven.com",
		Location: "Alpha Centauri",
		Language: "en-GB",
		Created: time.Now(),
		LastChanged: time.Now(),
	}

	if err := dalInstance.Create(&testUser); err != nil {
		t.Errorf("An error occurred while testing create: %v\n",err)
	}
	/*
	foundUser, err := dalInstance.Get(testUser.Id)
	if err != nil {
		t.Errorf("An error occurred while testing getting newly created data: %v\n",err)
	}

	assert.NotEmpty(t, foundUser)
	*/

}

func setupTest() {
	// If left default, it throws an error as "client api version 1.36 is too new" :D
	os.Setenv("DOCKER_API_VERSION", "1.35")
	runDockerRedisContainer()
	initRedisClient()
}

func endTest() {
	closeRedisClient()
	stopDockerRedisContainer()
}

func runDockerRedisContainer() {

	ctx = context.Background()

	dockerCli, err := dockerclient.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{
		All:true,
	})
	if err != nil {
		panic(err)
	}

	for _, cnt := range containers {
		// Check if the cnt already exists
		if contains(cnt.Names, "/" + redisContainerName) {
			cId := cnt.ID
			// Stop Container
			dockerCli.ContainerStop(ctx, cId, nil)
			// Remove Container
			dockerCli.ContainerRemove(ctx, cId, types.ContainerRemoveOptions{})
		}
	}

	out, err := dockerCli.ImagePull(ctx, redisImageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	resp, err := dockerCli.ContainerCreate(ctx, &container.Config{
		Image: redisImageName,
		ExposedPorts: map[nat.Port]struct{}{
			"6379/tcp": {}},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"6379/tcp": []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
					HostPort: "6379",
				},
			},
		},
	}, nil, redisContainerName)
	if err != nil {
		panic(err)
	}

	if err := dockerCli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("the testing redis cnt is up and running")


}

func stopDockerRedisContainer() {

	dockerCli, err := dockerclient.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{
		All:true,
	})
	if err != nil {
		panic(err)
	}

	for _, cnt := range containers {
		// Check if the cnt already exists
		if contains(cnt.Names, "/" + redisContainerName) {
			cId := cnt.ID
			// Stop Container
			dockerCli.ContainerStop(ctx, cId, nil)
			// Remove Container
			dockerCli.ContainerRemove(ctx, cId, types.ContainerRemoveOptions{})
		}
	}
	fmt.Println("the testing redis cnt is stopped and removed")
}

func initRedisClient() {

	var host string
	var port uint16

	dockerCli, err := dockerclient.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{
		All:true,
	})
	if err != nil {
		panic(err)
	}

	for _, cnt := range containers {
		// Check if the cnt already exists
		if contains(cnt.Names, "/" + redisContainerName) {
			fmt.Println("yay! found some container for redis test")
			redisContainerPortStruct := cnt.Ports[0]
			// host = redisContainerPortStruct.IP
			host = redisContainerName
			fmt.Printf("Redis Host: %v\n", host)
			port = redisContainerPortStruct.PublicPort
			fmt.Printf("Redis Port: %v\n", port)
		}
	}


	url := strings.Join([]string{host, strconv.FormatUint(uint64(port), 10)}, ":")

	client = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)

}

func closeRedisClient() {

	client.Close()

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
