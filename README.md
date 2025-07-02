# プロジェクトのアーキテクチャ

## 採用したアーキテクチャ

本プロジェクトは、**クリーンアーキテクチャ**（Clean Architecture）の原則に基づき、以下のレイヤー構成で設計されています：

- **internal/models** — [SQLBoiler](https://github.com/volatiletech/sqlboiler) によって自動生成されたデータモデル層。型安全性とDBスキーマとの同期を保証します。
- **internal/repo** — リポジトリ層（データアクセス層）。ビジネスロジックとデータ保存の詳細を分離するためにRepositoryパターンを採用。テストにはsqlmockを利用。
- **internal/handler** — ハンドラー層。ビジネスロジックや外部インターフェース（gRPC/HTTPなど）とのやり取りを担当します。
- **gen/proto** — gRPC/Protobuf（およびConnect）による自動生成ファイル。サービスとクライアント間の契約（インターフェース）を保証します。
- **cmd/server** — サーバーアプリケーションのエントリーポイント。
- **cmd/client** — サーバーとやり取りするCLIクライアント。
