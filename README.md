# Default Credentials Finder

An interactive CLI tool written in Go to retrieve default administrator credentials for popular services and applications. This tool is designed for cybersecurity professionals, penetration testers, and system administrators who need quick access to default login credentials for various systems.

## Key Features

- **Interactive CLI**: Select the service interactively from a terminal menu to quickly get default admin credentials.
- **Wide Service Support**: Retrieve default credentials for services like GitLab, Gitea, MySQL, PostgreSQL, MongoDB, Jenkins, WordPress, and more.
- **Extensible and Customizable**: Easily add new services or modify existing credentials.
- **User-Friendly Interface**: A simple, intuitive menu-driven interface for seamless user experience.
- **Modular Design**: Credentials are handled separately for each service, making it easy to maintain and expand.

## Available Services

- GitLab
- Gitea
- MySQL
- PostgreSQL
- MongoDB
- Redis
- Jenkins
- Docker
- Tomcat
- WordPress
- Drupal
- Apache
- phpMyAdmin
- DokuWiki
- Elasticsearch
- And more...

## Installation

### Prerequisites

Make sure you have [Go](https://golang.org/dl/) installed on your machine. If you don’t have it installed, follow the installation instructions on the Go website.

### Clone the repository

Clone this repository to your local machine:

```bash
git clone https://github.com/yourusername/default-credentials-finder.git
cd default-credentials-finder
```

### Install Dependencies

Run the following command to install the required Go dependencies:

```bash
go get github.com/manifoldco/promptui
```

## Usage

1. **Run the Program**

   To start the tool, run the following command:

   ```bash
   go run main.go
   ```

2. **Interactive Menu**

   After running the program, you will be prompted with an interactive menu to choose a service. Select the service, and the corresponding default admin credentials will be displayed in your terminal.

3. **Example Output**

   When you select a service, the credentials for that service will be shown like this:

   ```bash
   Select a service to retrieve default credentials
    1) gitlab
    2) gitea
    3) mysql
    4) postgresql
    5) mongodb
    6) redis
    7) jenkins
    8) docker
    9) tomcat
   10) wordpress
   11) drupal
   12) apache
   13) phpmyadmin
   14) dokuwiki
   15) elasticsearch
   16) other
   > 1
   GitLab Admin Credentials:
   Username: root
   Password: 5iveL!fe
   ```

## Contributing

Feel free to fork the repository, create branches, and contribute improvements or bug fixes.

### How to contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add new service'`)
5. Push to your forked repository (`git push origin feature-branch`)
6. Create a pull request
