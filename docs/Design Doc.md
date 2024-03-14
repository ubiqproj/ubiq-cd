# Design Doc: ubiq-cd

## Objective

Pull-based continuous deployment pipelines have advantages in managing hosts over push-based one ( common CD pipelines built with GitHub Actions, GitLab CI/CD, Jenkins, etc.).
This is because you can leave resource management and application monitoring to the agent on the host.

However, as of 2023, all tools that can build pull-based continuous deployment pipelines ( e.g., ArgoCD ) require use on Kubernetes.
This makes them expensive to build and operate because they require Kubernetes cluster and CI/CD pipeline for managing container images.

So, we solve these issues by creating CI/CD tools.
It will build and manage pull-based pipelines, each with a single IaC.

## Goal, Non goal

### Goal

- Simple pull-based pipelines
  - You can deploy and remove applications to a host with just Git operations.
- Manage applications
  - Monitoring
  - Rollback on fail
  - Metrics of them (for Prometheus)

### Non goal

- Features outside of the scope of CI/CD tools
  - Clustering of ubiq-cd hosts between each other, removal of single point of failure, etc.
  - External webhook capability (monitoring and notification will operate with Prometheus)

## High Level Structure

[Matrix](https://www.figma.com/file/JX4i8dXPMGibet47KFIUlH/Architecture-matrix?type=whiteboar&version-id=5343802330&node-id=1-2d&node-id=1%3A2)

- cmd
  - ubiqcd
  - ubiqctl
    - cmd
- internal
  - infrastructure
    - interface
      - http
    - cmd
      - ubiqcd
      - ubiqctl
    - externalapi
      - docker
      - git
      - systemd
  - interface-adapter
    - agent
    - client
    - interface
      - connectrpc
      - handler
  - usecase
    - runner
    - gitops
  - domain
    - service
    - pipeline

## References

- [What is CI/CD? - RedHat](https://www.redhat.com/en/topics/devops/what-is-ci-cd)
