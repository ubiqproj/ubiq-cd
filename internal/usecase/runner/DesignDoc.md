# Design Doc: runner

## Goal, Non goal

### Goal

- パイプライン管理 API の提供
  - 登録、検索、実行、停止、削除
- パイプラインのトランザクション
  - `domain/pipeline`, `domain/service` パッケージの状態遷移を吸収
  - domain 層の処理で冪等性を担保する必要がある
- パイプラインのマニフェスト・ジョブのログを永続化
- 安定稼働したリビジョンを記録し、パイプラインのジョブが失敗または、ジョブ完遂後にサービスの可用性が一定値を下回った際にロールバックする