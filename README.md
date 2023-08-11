## GOでREST, gRPC, GraphQLのAPIを設計する

- ORM使うかどうか

**技術選定**
| 技術 | 役割 |
| - | - |
| net/http | サーバー (標準ライブラリ) |
| OpenAPI | REST用ライブラリ |
| oapi-codegen | OpenAPIのコード生成 |
| gqlgen | GraphQLライブラリ |
| testify | Goのテストライブラリ |
| godoc | ドキュメント生成 |
| Postman | APIの動作確認 |
| PostgreSQL | DB |
| pgAdmin | DBの監視 |
