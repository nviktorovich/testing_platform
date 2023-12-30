# Testing platform

## Service descriptions
The testing platform has a simple name. This service has only one global task - to provide the testing process. Before meeting with a tutor, a student should test the minimum level of his/her knowledge.
The service has a large field for development, but at the first stage it provides the possibility of online testing through the user's browser.
## Development platform and tools
The testing_service core is developed in Golang. Golang is a simple and fast programming language.
Golangci-lint is used as the linter.
PostgreSQL is chosen as the default storage because it is idiomatic to Golang. Migrations are used to update schemas and tables.
To avoid deployment issues on different platforms, Docker with docker-file and with docker compose was used.
Pipelines are used to provide ci/cd process with stages: linking, migrations, testing, build and deployment.
A make-file is used for automation.
These are not only the presented tools and platforms used in this project, this list will expand.

