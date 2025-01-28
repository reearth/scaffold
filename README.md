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
    - **usecase.go**: Usecase container
    - ...
  - **usecase**: Define gateways I/F, transaction I/F, gateways container, repo container, and policy container
    - **xxxuc** Actual usecase implementation
    - **gateway.go**: Gateways interfaces and container
    - **policy.go**: Policy container
    - **repo.go**: Repo container
    - **transaction.go**: Transaciton interfaces
    - ...
- **pkg**: Domain models, repo interfaces, policy interfaces

## Dependency Flows

```mermaid
flowchart
  subgraph boot
    config
  end
  subgraph infra
    infra2[mongo, gcp, auth0, cerbos...]
  end
  subgraph transport
    transport2[cli, echo, gql..]
  end
  subgraph usecase
    uc[xxxuc]
    gatewayIF[Gateway I/F]
    transaction[Transaction I/F]
  end
  subgraph domain
    model[Domain Models]
    repoIF[Repo I/F]
    policyIF[Policy I/F]
  end

  cmd --> boot
  cmd --> transport
  boot -- initialize --> uc
  boot -- initialize --> infra
  transport --> uc
  usecase --> domain
  uc --> gatewayIF
  uc --> transaction
  infra -- impl --> repoIF
  infra -- impl --> policyIF
  infra -- impl --> gatewayIF
  infra -- impl --> transaction
  infra --> domain
```
