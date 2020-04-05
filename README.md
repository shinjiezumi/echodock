# echodock
## 構成
- Go v1.13.7
- echo v3.3.10

## セットアップ
```
$ git clone https://github.com/shinjiezumi/echodock.git
$ cd echodock
$ docker-compose up -d
```

## 動作確認
`http://localhost:8080`を開く

## リモートデバッグ
1. realize.yaml切り替え
    ```
    $ mv .realize.yaml .realize.yaml.bk
    $ mv .realize.yaml.debug .realize.yaml
    ```
2. docker-compose.yamlの`command`をコメントアウト
    ``` yaml
    #    command: realize start --run
    ```
3. コンテナ再起動+realize実行
    ``` 
    $ docker-compose up -d
    $ docker-compose exec go realize start --run
    ```
4. Golandでデバッグ実行