# server-scaffold

This is an example repository for showing the standard module design used in Re:Earth's Go server application.

- **cmd**: Entrypoint
- **internal**
  - **boot**: Read the config, initialize DB drivers, etc., and finally initialize usecases and run the appropriate transport layer.
  - **infra**: Implements repos, policies, and gateways.
    - **gcp**
    - **mongo**
    - ...
  - **transport**: UI layer
    - **cli**
    - **echo**
    - **gql**
    - ...
  - **usecase**: Define gateways I/F, transaction I/F, gateways container, repo container, and policy container
    - **xxxuc** Actual usecase implementation
    - **gateway**: Gateways interfaces
    - **transaction.go**: Transaciton interfaces
    - **usecase.go**: Usecase container
    - ...
- **pkg**: Domain models, repo interfaces, policy interfaces

## Dependency Flows

```mermaid
graph TD
  %% CMD calls InitializeEcho and InitializeCLI
  CMD -->|calls| InitializeEcho
  CMD -->|calls| InitializeCLI

  %% InitializeEcho dependencies
  subgraph Echo Server
    InitializeEcho -->|Config & DB| boot.LoadConfig
    InitializeEcho -->|Config & DB| boot.InitMongo
    InitializeEcho -->|Usecases| usecase.NewUsecases
    InitializeEcho -->|Echo Server| echo.NewEchoConfig
    InitializeEcho -->|Echo Server| echo.New
  end

  %% InitializeCLI dependencies
  subgraph CLI
    InitializeCLI -->|Config & DB| boot.LoadConfig
    InitializeCLI -->|Config & DB| boot.InitMongo
    InitializeCLI -->|CLI Setup| cli.NewCLIConfig
    InitializeCLI -->|CLI Setup| cli.NewCLI
  end

  %% Usecases initialization
  subgraph Usecases
    usecase.NewUsecases --> assetuc.New
    usecase.NewUsecases --> projectuc.New
    usecase.NewUsecases --> workspaceuc.New
    usecase.NewUsecases --> useruc.New
  end

  %% Asset Usecase dependencies
  subgraph Asset Usecase
    assetuc.New --> assetuc.NewFindByIDsUsecase
    assetuc.New --> assetuc.NewFindByProjectUsecase
    assetuc.New --> assetuc.NewCreateUsecase
    assetuc.New --> assetuc.NewUpdateUsecase
  end

  %% Infrastructure dependencies
  subgraph Infrastructure
    mongo.NewAsset -->|binds| asset.Repo
    mongo.NewWorkspace -->|binds| workspace.Repo
    mongo.NewUser -->|binds| user.Repo
    mongo.NewProject -->|binds| project.Repo
    gcp.NewStorage -->|binds| gateway.Storage
  end
```
