package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func getGitLabAdminCredentials() string {
	return "GitLab Admin Credentials:\nUsername: root\nPassword: 5iveL!fe"
}

func getGiteaAdminCredentials() string {
	return "Gitea Admin Credentials:\nUsername: admin\nPassword: password"
}

func getMySQLAdminCredentials() string {
	return "MySQL Admin Credentials:\nUsername: root\nPassword: root"
}

func getPostgreSQLAdminCredentials() string {
	return "PostgreSQL Admin Credentials:\nUsername: postgres\nPassword: postgres"
}

func getMongoDBAdminCredentials() string {
	return "MongoDB Admin Credentials:\nUsername: admin\nPassword: admin"
}

func getRedisAdminCredentials() string {
	return "Redis Admin Credentials:\nUsername: default\nPassword: default"
}

func getJenkinsAdminCredentials() string {
	return "Jenkins Admin Credentials:\nUsername: admin\nPassword: admin"
}

func getDockerAdminCredentials() string {
	return "Docker Admin Credentials:\nUsername: root\nPassword: docker"
}

func getTomcatAdminCredentials() string {
	return "Tomcat Admin Credentials:\nUsername: admin\nPassword: admin"
}

func getWordPressAdminCredentials() string {
	return "WordPress Admin Credentials:\nUsername: admin\nPassword: admin"
}

func getDrupalAdminCredentials() string {
	return "Drupal Admin Credentials:\nUsername: admin\nPassword: admin"
}

func getApacheAdminCredentials() string {
	return "Apache Admin Credentials:\nUsername: admin\nPassword: admin"
}

func getphpMyAdminAdminCredentials() string {
	return "phpMyAdmin Admin Credentials:\nUsername: root\nPassword: root"
}

func getDokuWikiAdminCredentials() string {
	return "DokuWiki Admin Credentials:\nUsername: admin\nPassword: admin"
}

func getElasticsearchAdminCredentials() string {
	return "Elasticsearch Admin Credentials:\nUsername: elastic\nPassword: changeme"
}

func getOtherCredentials() string {
	return "Other Service Admin Credentials:\nUsername: admin\nPassword: admin123"
}

func listAvailableServices() []string {
	return []string{
		"gitlab", "gitea", "mysql", "postgresql", "mongodb", "redis",
		"jenkins", "docker", "tomcat", "wordpress", "drupal", "apache",
		"phpmyadmin", "dokuwiki", "elasticsearch", "other",
	}
}

func main() {
	services := listAvailableServices()

	prompt := promptui.Select{
		Label: "Select a service to retrieve default credentials",
		Items: services,
	}

	_, serviceName, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	serviceName = strings.ToLower(serviceName)

	switch serviceName {
	case "gitlab":
		fmt.Println(getGitLabAdminCredentials())
	case "gitea":
		fmt.Println(getGiteaAdminCredentials())
	case "mysql":
		fmt.Println(getMySQLAdminCredentials())
	case "postgresql":
		fmt.Println(getPostgreSQLAdminCredentials())
	case "mongodb":
		fmt.Println(getMongoDBAdminCredentials())
	case "redis":
		fmt.Println(getRedisAdminCredentials())
	case "jenkins":
		fmt.Println(getJenkinsAdminCredentials())
	case "docker":
		fmt.Println(getDockerAdminCredentials())
	case "tomcat":
		fmt.Println(getTomcatAdminCredentials())
	case "wordpress":
		fmt.Println(getWordPressAdminCredentials())
	case "drupal":
		fmt.Println(getDrupalAdminCredentials())
	case "apache":
		fmt.Println(getApacheAdminCredentials())
	case "phpmyadmin":
		fmt.Println(getphpMyAdminAdminCredentials())
	case "dokuwiki":
		fmt.Println(getDokuWikiAdminCredentials())
	case "elasticsearch":
		fmt.Println(getElasticsearchAdminCredentials())
	case "other":
		fmt.Println(getOtherCredentials())
	default:
		fmt.Println("Service not recognized.")
	}
}
