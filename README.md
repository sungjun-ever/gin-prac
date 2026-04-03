# go gin 프레임워크 공부용

```
.
├── cmd/   
│   └── main.go    # 어플리케이션 진입점
├── internal           # 외부에서 임포트할 수 없는 비공개 로직/
│   ├── controller     # HTTP 요청 처리 및 응답 반환
│   ├── service        # 비즈니스 로직
│   ├── repository     # 데이터베이스 접근
│   ├── model          # 데이터 구조체
│   └── middleware     # gin 커스텀 미들웨어
├── registry           # 서비스 등록
├── database           # db 연결
├── config             # 설정 파일  
├── router             # router
├── test               # 통합 테스트
├── go.mod
└── go.sum
```