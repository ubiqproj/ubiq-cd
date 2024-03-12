# Design Doc: ubiq-cd

## Objective

Pull 型の Continuous Deployment パイプラインはホスト上のリソース管理やアプリケーションの監視を環境上で動作するエージェントが行うためエンジニアが意識する必要がなく Push 型の Continuous Deployment パイプライン (GitHub Actions, GitLab CI/CD, Jenkins などで構築される一般的な CD パイプライン) と比べて優位である。
しかし、2023 年現在 Pull 型の Continuous Deployment パイプラインを構築できるツール (ArgoCD など) はすべて Kubernetes 上で使用することを前提としており、コンテナイメージの管理のための CI/CD パイプラインも原則として別途必要であるため構築・運用にコストがかかる。
そこで、Pull 型のパイプラインを IaC から構築でき、コンテナイメージの管理が不要な CI/CD ツールを作成しこれらの課題を解決する。

## Goal, Non goal

### Goal

- Pull 型のパイプラインを低コストで構築
  - Git 操作のみでアプリケーションをホストにデプロイ・削除
- アプリケーション管理機能の実装
  - 動作監視
  - 障害時のロールバック
  - 上記機能内のメトリクス公開 (for Prometheus)

### Non goal

- CI/CD ツールの範疇を超えた機能の実装
  - ubiq-cd を導入したホスト同士のクラスタ化、単一障害点の排除など
  - 外部への webhook 機能 (メトリクスを公開し監視及び通知は Prometheus で行う)

## High Level Structure

- cmd
  - ubiqcd
  - ubiqctl
    - cmd
- internal
  - infrastructure
    - webapi
    - datasource
    - externalapi
      - git
      - docker
      - systemd
  - interface-adapter
  - usecase
    - runner
    - gitops
    - metrics
  - domain
    - service
    - pipeline

## References

- [CI/CD とは - RedHat](https://www.redhat.com/ja/topics/devops/what-is-ci-cd)
