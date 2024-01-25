# task-management

# After adding package use these commands

go mod vendor
go mod tidy

# Run using
docker-compose -f docker-compose.dev.yml up -d

This will build your development image and start the containers.

With these files, you have a production Dockerfile (Dockerfile) and a development Dockerfile (Dockerfile.dev). You can use docker-compose up -d for production and docker-compose -f docker-compose.dev.yml up -d for development. The development setup allows you to make changes to your code, and the changes will be immediately reflected in the running container.


# Test API

## Create Task

`localhost:8080/api/tasks/create`

```
{
    "taskName":"5th",
    "description":"5th Task",
    "status":"NEW"
}
```

## List Tasks

`localhost:8080/api/tasks/list`

```
{"id":"65b15c3bb925e036f6560d64","taskName":"3rd","description":"3rd Task","status":"NEW"},
{"id":"65b15d9793167a708db2fdfa","taskName":"4th","description":"4th Task","status":"NEW"},
{"id":"65b15db293167a708db2fdfb","taskName":"5th","description":"5th Task","status":"NEW"}
```

## Edit Tasks

`localhost:8080/api/tasks/update?id=65b15c3bb925e036f6560d64`

```
{
    "taskName":"3rd",
    "description":"3rd Task-Updated",
    "status":"NEW"
}
```

## Delete Tasks

`localhost:8080/api/tasks/delete?id=65b15a47a501cb0afaa6c573`

## Update Task Status

`localhost:8080/api/tasks/update-status?id=65b15aa3a501cb0afaa6c574`
```
{"status":"PENDING"}
```