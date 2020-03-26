package image

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/yanlingqiankun/Executor/conf"
	"io"
	"strings"
	"time"
)

//func importDocekrImage(ctx context.Context, name, tag, file string) error {
//	fileReader, err := os.Open(file)
//	if err != nil {
//		return err
//	}
//	imageImportSource := types.ImageImportSource{
//		Source:     fileReader,
//		SourceName: name,
//	}
//	out, err := cli.ImageImport(ctx, imageImportSource, "", types.ImageImportOptions{Tag:tag})
//	if err != nil {
//		return err
//	}
//	return out.Close()
//}
//
//func SaveDockerImage(ctx context.Context, name, file string) error {
//	fileReader, err := os.Open(file)
//	if err != nil {
//		return err
//	}
//	resp, err := cli.ImageLoad(ctx, fileReader, true)
//	if err != nil {
//		return err
//	}
//	fmt.Println(resp.JSON)
//	buf := bytes.Buffer{}
//	defer resp.Body.Close()
//	io.Copy(&buf, resp.Body)
//	logger.Infoln("docker save infomation : \n", buf.String())
//	return nil
//}

func PullDockerImage(ctx context.Context, name string) error {
	if _, ok := db[name]; ok {
		logger.Debug("The docker image has been in repository : ", name)
		return nil
	}
	buf := bytes.Buffer{}
	storeNode := conf.GetString("StoreNode")
	registryPort := conf.GetString("DockerRegistryPort")
	refStr := storeNode + ":" + registryPort + "/" + name
	resp, err := cli.ImagePull(ctx, refStr, types.ImagePullOptions{All: false})
	if err != nil {
		return  err
	}
	defer resp.Close()
	io.Copy(&buf, resp)
	logger.Infoln("docker pull infomation : \n", buf.String())
	if err := cli.ImageTag(ctx, refStr, name); err != nil {
		return err
	}
	imageId, err := getDockerImageID(name)
	if err != nil {
		return err
	}
	db[name] = &ImageEntry{
		Name:          name,
		ID:            imageId,
		CreateTime:    time.Now(),
		Type:          "docker_pull",
		IsDockerImage: true,
		Counter:       0,
	}
	return db.save()
}

func getDockerImageID (name string) (string, error) {
	inspect, _, err := cli.ImageInspectWithRaw(context.Background(), name)
	if err != nil {
		return "", err
	}
	return strings.TrimLeft(inspect.ID, "sha256:"), nil
}
