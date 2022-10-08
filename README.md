# 井草屋Blog
個人用内製ブログです。  
以下を利用して製造してます。
- Go-server  
- Next.js
- Openapi
- MySql

## Doc
- [Github Actionsを利用したCroudRunへのデプロイ](./doc/001_deploy_to_CroudRun.md)

## 要install
```shell
$ go install golang.org/x/tools/cmd/goimports@latest  
$ go install github.com/k0kubun/sqldef/cmd/mysqldef@latest  
```

## TODO
- DB設計
  - articles
  - tags
  - articles_tags
  - comments
  - articles-comments
  - admin
- DB接続
- テスト自動化
- 機能
  - 記事投稿
  - 記事一覧
  - 記事詳細
  - 記事更新
  - 記事削除
  - ログイン

今後の予定
1. 記事投稿spec作成 
   - ok
1. 記事投稿実装  
   - 疎通確認 ok
   - 残：テストコード
1. develop環境整備
   - DB用意
   - DB結合確認
1. front作成
1. 結合
1. develop環境での結合
1. Google認証実装(spec front api)
