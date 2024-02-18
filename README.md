# GDS-OneCV-SWE-Assignment

### Database Schema
![entity relation diagram](erd.jpg)

While emails are used to identify the teachers and students, auto-incrementing bigint id fields are used as primary keys for the `teachers` and `students` tables for better extensibility (e.g. we may want to allow teachers/students to update their email addresses in the future)

### Local Development
**Prerequisites**
- Docker
- Docker Compose

**Environment variables**\
Set the environmental variables as follow
| Key | Value | Remarks |
|--------------|--------------------------------------|-----------------------------------------------------------------------|
| DB_DIALECT | postgres | |
| DB_HOST | database | `database` is the name of the golang server service in docker-compsoe |
| DB_PORT | 5432 | |
| DB_USER | | e.g. user |
| DB_NAME | | e.g. postgres |
| DB_PASSWORD | | e.g. mysecretpassword |
| SERVER_PORT | 8000 | |
| SERVER_HOST | server | `server` is the name of the golang server service in docker-compsoe |

Run the following to spin up the containers
```
bash run.sh --dev
```
Alternatively, run
```
docker-compose -f docker-compose-dev.yml up -d
```

*The only difference between running the contains in dev/prod environment is that the nginx server is configured for HTTPS instead of HTTP in the latter, as the SSL certificate is associated with the Elastic IP address which is only accessible in the production environment.

After the containers are up, you can access the backend at `http://localhost` (exposed by default on port 80)

**Seeding**\
For seeding, send a POST request with an empty body to `api/seed`

**Docker compose**\
The docker-compose service consists of 3 containers:
- HTTP server in Golang
- PostgreSQL database
- Nginx reverse proxy

For local development, the database container is also exposed on port 5432 on localhost for easy debugging

### Cloud Deployment and CI/CD
The API is deployed on AWS EC2 at https://54.254.110.35

Elastic IP address is also associated with a TLS/SSL certificate obtained from ZeroSSL to enable the instance to handle HTTPS traffic.

Terraform is used to quickly launch and tear down the instance, with secrets stored in Hashicorp Vault Secrets.

Automated unit-testing is triggered via GitHub actions whenever a new commit is made, and when the master branch receives a pull request.

### Project structure
```
    .
    ├── ent
    │   ├── ...
    │   │   
    │   ├── schema
    │   │   ├── student.go
    │   │   └── teacher.go
    │   │
    │   ├── ... (generated files by ent)
    ├── pkg
    │   ├── api
    │   │   ├── api.go
    │   │   ├── api_test.go
    │   │   ├── errors.go
    │   │   └── errors_test.go
    │   ├── database
    │   │   ├── database.go
    │   │   ├── students.go
    │   │   └── teachers.go
    │   ├── handlers
    │   │   ├── commonstudents.go
    │   │   ├── commonstudents_test.go
    │   │   ├── register.go
    │   │   ├── register_test.go
    │   │   ├── retrievefornotification.go
    │   │   ├── retrievefornotification_test.go
    │   │   ├── suspend.go
    │   │   └── suspend_test.go
    │   ├── middleware
    │   ├── router 
    │   ├── seed
    │   ├── server
    │   └── testutil
    ├── terraform
    ├── certificate.crt
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── (nginx configs...)
    └── (Dockerfiles, docker-compose.yml files...)
```

### Commit messages

This project follows [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) guideline for commit messages. See the table below for the list of commit types.

| Type     | Description                                                                                            |
| -------- | ------------------------------------------------------------------------------------------------------ |
| feat     | A new feature                                                                                          |
| fix      | Bug fixes                                                                                              |
| test     | Adding missing tests or correcting existing tests                                                      |
| refactor | Changes to source code that neither add a feature nor fixes a bug                                      |
| build    | Changes to CI or build configuration files (Docker, github actions)                                    |
| chore    | Anything else that doesn't modify any `internal` or `test` files (linters, configs, etc.)              |
| revert   | Reverts a previous commit                                                                              |
| docs     | Documentation only changes                                                                             |
| perf     | A code change that improves performance                                                                |
| style    | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc) |
