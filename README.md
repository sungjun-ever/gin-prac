# go gin 프레임워크 공부용

```
.
├── cmd/
│   └── server/
│       └── main.go    # 어플리케이션 진입점
├── internal           # 외부에서 임포트할 수 없는 비공개 로직/
│   ├── controller     # HTTP 요청 처리 및 응답 반환
│   ├── service        # 비즈니스 로직
│   ├── repository     # 데이터베이스 접근
│   ├── model          # 데이터 구조체
│   ├── middleware     # gin 커스텀 미들웨어
│   └── platform       # 공통 유틸리티 (DB 연결, redis 설정 등)
├── pkg                # 외부 프로젝트에서도 사용할 수 있는 공개 라이브러리
├── configs            # 설정 파일
├── api                # api 정의 파일
├── scripts            # 빌드, 배포, 마이그레이션 스크립트
├── test               # 통합 테스트
├── go.mod
└── go.sum
```