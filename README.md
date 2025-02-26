# scaffold

A scaffolding repository to speed up the launch of new application development.

- [server](server)
- [web](web)

# How to launch a new application

## 1. Init a new repo on GitHub

## 2. Set variables and secrets to the GitHub repo

Variable names and example values:

- `REPO`: `reearth/scaffold`
- `IMAGE_SERVER`: `ghcr.io/reearth/scaffold/scaffold-web`
- `IMAGE_WEB`: `ghcr.io/reearth/scaffold/scaffold-api`
- `REGION`: `asia-northeast1`
- `SERVICE_NAME_API`: `reearth-scaffold-api`
- `SERVICE_NAME_WEB`: `reearth-scaffold-web`

Secrets (organization secrets are also OK):

- `GCP_PROJECT_ID`
- `GCP_SA_EMAIL`
- `GCP_WORKLOAD_IDENTITY_PROVIDER`
- `GH_APP_ID`
- `GH_APP_PRIVATE_KEY`
- `GH_APP_USER`
- `DOCKERHUB_USERNAME` (if needed)
- `DOCKERHUB_TOKEN` (if needed)

## 3. Copy files in this repository to a new repo

Do not fork this repository. Init a new repository on GitHub and just copy files in this repository to it.

## 4. Edit files

- `server/go.mod`

  Rename this to your new module name.

  ```
  module github.com/reearth/scaffold/server
  ```

- `web/index.html`

  ```HTML
  <title>Vite + React + TS</title>
  ```

- `web/package.json`

  ```json
  {
    "name": "scaffold" // EDIT
    // ...
  }
  ```

- `terraform/reearth_scaffold_{api,web}`

  Rename these directories.

- `terraform/reearth_scaffold_{api,web}/locals.tf`

  Edit a local: `service_name`.

- `**/README.md`

- `**/*.go`

  Replace import statements in each go files to your new module name. It's good to run `go mod tidy` after renaming.

  ```go
  import (
    "github.com/reearth/scaffold/server/foobar" // -> "<YOUR MODULE NAME>/server/foobar"
  )
  ```

## 5. Trigger CI

Ensure Docker images have been saved at a registry. Deployment will fail, but don't worry, as long as CI is successful.

## 6. Deploy infrastructure

(WIP)

## 7. Trigger CI/CD again

Make sure the deployment workflows are successful.
