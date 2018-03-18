# (git + hub/lab/bucket) * interactivity = gico
gico - interactive & integrated git utility tool

## Features
* multi platform support
* multi git hosting services support
* multi interactive filter support

## gico checkout

## gico checkout --pr
remeteのホスト名からいい感じにPR一覧を取得してgo 

## gico id
選択したlogのidを標準出力に出力します

# TODO
- [x] Tokenを設定ファイルから読むようにする
- [x] GitLab対応
- [ ] BitBucket対応
- [ ] GitBucket対応
- [x] ホスト名とリポジトリの種類の対応を設定ファイルで管理
- [ ] fzf対応
- [x] logがバグってる
- [ ] browseのテスト
- [x] browseでpecoをキャンセルすると例外を吐く
- [ ] テスト書く
- [ ] 各種repoServiceごとのstructを作って、hostTypeごとに生成するクライアントを分ける
- [ ] finderの処理を共通化
- [ ] git logを雑にやると多すぎて帰ってこなくなる


pecoへFilterbleStrnigerの配列を渡す
FilterbleStringerは以下のinterafceを持つ
serialize
deserialize
選択されたFilterbleStringerを返す

