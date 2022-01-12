# DoS

## What's DoS

DoSはDenial of Serviceの略称で、攻撃対象のサーバーに対して大量のリクエストを送ってサーバーのリソースを枯渇させ、サービスを停止させる攻撃手法。

## Sample

PHP + Apacheの構成で動いているサービスに対し、Goのスクリプトから簡易的なDoS攻撃を行う。

### Prepare

#### App Side

```bash
cd app
docker-compose up
```

#### Attacker Side

```bash
cd attack
go run ./main.go
```

### Measure

#### Usual Time

```bash
$ curl localhost:8000
time: 0.53205800056458[s]
result: 4999999950000000
```

#### DoS Time

```bash
$ curl localhost:8000
time: 18.589974164963[s]
result: 4999999950000000
```