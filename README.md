# echodock
## 構成
- Go v1.15
- echo v4

## セットアップ
```shell script
$ git clone https://github.com/shinjiezumi/echodock.git
$ cd echodock
$ docker-compose up -d
$ sh src/scripts/setup.sh
```

## 動作確認
```shell script
$ sh src/scripts/run.sh
```

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

## localstack

1. localstack用のプロファイル作成

```shell
aws configure
```

2. 動作確認

```shell
$ aws sqs create-queue --queue-name test --endpoint-url=http://localhost:4566 --profile=localstack
{
    "QueueUrl": "http://localstack:4566/000000000000/test"
}
```

```shell
$ aws sqs list-queues  --endpoint-url=http://localhost:4566 --profile=localstack
{
    "QueueUrls": [
        "http://localstack:4566/000000000000/test"
    ]
}
```

