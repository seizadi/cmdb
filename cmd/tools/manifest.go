package main

import (
	"errors"
	"flag"
	"fmt"
	"go.uber.org/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	
	"github.com/seizadi/cmdb/helm"
)

func getManifestConfig(filenames []string) (string, error) {

	var v map[interface{}]interface{}
	var sources []config.YAMLOption

	sources, err := ReadFiles(filenames...)
	if err != nil {
		return "", err
	}
	//fmt.Printf("Sources  #%v \n", sources)

	provider, err := config.NewYAML(sources...)
	if err != nil {
		return "", err
	}
	//fmt.Printf("provider  #%v \n", provider)

	err = provider.Get(config.Root).Populate(&v)
	//fmt.Printf("V  #%v \n", v)

	// Create the Yaml File
	c, err := yaml.Marshal(&v)
	if err != nil {
		return "", err
	}
	//fmt.Printf("c  #%v \n", len(c))

	// Originally I was using helm to resolve the values, it was taking about 370ms
	// I wrote it to sue go template engine directly which reduced the time around 70ms
	values := helm.Values{Values: v}
	r := helm.Renderable{Tpl: string(c), Vals: values}
	config, err := helm.RenderWithReferences(r)
	if err != nil {
		return "", err
	}

	return config, nil
}

func ReadFiles(filenames ...string) ([]config.YAMLOption, error) {

	var sources []config.YAMLOption

	if len(filenames) <= 0 {
		//fmt.Printf("No files to read\n")
		return sources, errors.New("You must provide at least one filename for reading Values")

	}
	for _, filename := range filenames {
		source, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		sources = append(sources, config.Source(strings.NewReader(string(source))))
	}
	return sources, nil
}

func generate_files(output_file string) {

	// open the out file for writing
	outfile, err := os.Create(output_file)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	//cmd.Stdout = outfile
	//err = cmd.Run()
	if err != nil {
		panic(err)
	}
	//fmt.Println("Done")
}

func find_files(filename string, app_name string, env string, lifecycle string) []string {
	//var basePath = "/Users/cpadhan/cicd/hack/personal-repo/dc-hack" //get_base_path()
	var basePath = get_base_path()
	var filenames []string
	var paths []string
	paths = append(paths, fmt.Sprintf("%s/envs/%s", basePath, filename))
	paths = append(paths, fmt.Sprintf("%s/envs/%s/%s", basePath, lifecycle, filename))
	paths = append(paths, fmt.Sprintf("%s/envs/%s/%s/%s", basePath, lifecycle, env, filename))
	paths = append(paths, fmt.Sprintf("%s/envs/%s-%s", basePath, app_name, filename))
	paths = append(paths, fmt.Sprintf("%s/envs/%s/%s-%s", basePath, lifecycle, app_name, filename))
	paths = append(paths, fmt.Sprintf("%s/envs/%s/%s/%s-%s", basePath, lifecycle, env, app_name, filename))
	for _, path := range paths {
		//fmt.Printf("Path %v:\n", path)
		if file_stat, err := os.Stat(path); err == nil {
			if file_stat.Size() != 0 {
				filenames = append(filenames, path)
			}
		}
	}
	return filenames
}

func get_base_path() string {
	pwd, err := os.Getwd()
	if err != nil {
		os.Exit(0)
	}
	return pwd + "/tmp/repo"
}

func main() {

	LifeCyclePtr := flag.String("lifecycle", "integration", "the lifecycle the env is part of (dev, integration, preprod, prod)")
	EnvPtr := flag.String("env", "env-1", "name of the environment the app is being deployed to")
	AppNamePtr := flag.String("app_name", "acme", "name of the app being deployed")
	OutputFilePtr := flag.String("output_file", "/tmp", "output file, place to generate file")
	flag.Parse()

	var filename = "values.yaml"
	var filenames []string
	filenames = find_files(filename, *AppNamePtr, *EnvPtr, *LifeCyclePtr)
	//fmt.Printf("Filenames: \n", filenames)
	data, err := getManifestConfig(filenames)

	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v \n", err)
	}

	//fmt.Printf("%v", data)
	// open the out file for writing
	f, err := os.Create(*OutputFilePtr)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		panic(err)
	}
	f.Sync()
}
