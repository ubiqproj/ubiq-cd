# Design Doc: pipeline, service

## Objective

## Goal, Non goal

### Goal

- `domain/pipeline` パッケージ
  - Git リモートリポジトリの更新を確認し、CI/CD パイプラインを実行
  - `domain/service` パッケージを使用し、アプリケーションを管理
  - 抽象
    - Git
      - Git リモートリポジトリの更新確認
      - Git リモートリポジトリの clone
  - 具体
- `domain/service` パッケージ
  - 抽象
  - 具体

### Non goal

## High Level Structure

### Sequence

1. Git リモートリポジトリの更新を検知
2. サービスの更新
   1. ビルド
   2. インストール

## References
