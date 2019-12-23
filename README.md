# 카페인 프로젝트 알림 서비스

FCM과 연동하여 주문자, 매장용 프론트엔드에 알림을 보내기 위한 애플리케이션입니다.

## 설정

* `golang` >= 1.13을 권장합니다.
* 실행을 위해 Redis가 필요합니다.

**main.go**:

```go
repository := NewRedisRepository("localhost:6379", "", 0) // Redis 연결에 필요한 값
apiWrapper := NewFcmApiWrapper("key=AAAA...KD") // Push notification 요청 시 Authorization 헤더의 값
```

## 빌드 & 실행

```bash
go build
./caffeine-notification # 윈도우인 경우 caffeine-notification.exe
```
