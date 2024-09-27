# 백엔드 3차 과제

> 저번 과제에서 이어지는 과제에요. 실패하더라도 PR을 올려 선배들께 코드리뷰를 받아보아요.

이번과제는 주문 관리 기능을 구현해볼거에요.

## 구현사항
### 사장님
- 주문 목록 조회: 사장님은 자신이 운영하는 식당에 대한 주문 목록을 조회할 수 있어야 합니다.
- 주문 상태 변경: 사장님이 주문의 상태(예: 준비중, 배달중, 완료 등)를 변경할 수 있어야 합니다.

### 사용자
- 주문하기: 사용자는 특정 음식과 수량을 선택하여 주문할 수 있어야 합니다.
- 주문 내역 조회: 사용자는 자신의 주문 내역을 조회할 수 있어야 합니다.
- 주문 취소: 사용자가 주문을 취소할 수 있는 기능이 필요합니다.


## Endpoint 및 파라미터 설정법
Endpoint 및 파라미터 설정법
### orders
1. /orders: 주문 생성
2. /orders/user/{userId}: 특정 사용자의 주문 내역 조회
3. /orders/shop/{shopId}: 특정 식당의 주문 목록 조회
4. /orders/{orderId}: 주문 상태 변경 
5. /orders/{orderId}: 주문 취소

## 가져야하는 body 구조
```json
{
  "order": {
    "userId": "string", // 사용자 ID
    "foodId": "string", // 음식 ID
    "quantity": int, // 수량
    "totalPrice": int // 총 가격
    "status" : "string" // 주문상태
  }
}
```
status의 구성요소는
- 준비중
- 배달중
- 완료
이렇게 3가지만 가능해요.

## 파일 구성
2번에서 진행했던 자신의 이름 폴더와 postgres-data폴더를 모두 복사 붙여넣기를 해서 작업해야해요.

## DataBase
### 데이터베이스 사용법
데이터베이스는 postgresql을 사용할 것 이에요. 그렇기에 ORM을 사용해야해요.<br/> <br>
Node 환경(express, Nest JS 등)에서는 TypeORM 혹은 Prisma를 사용해요. <br> <br>
Java 환경(Java & Kotlin Spring)에서는 JPA를 사용해요. <br> <br>
설치없이 바로 사용할 수 있도록 도커환경을 사용할꺼에요. docker-compose.yaml파일을 건들지 말고, 밑에 있는 명령어만 실행시키면, DB가 실행되요.<br/>
> 도커가 설치 되어있지 않다면 아래 아티클을 보고 도커를 설치해요
> https://mz-moonzoo.tistory.com/40

```shell
docker-compose up
```
하나의 터미널에서 실행시킨 후 새로운 터미널을 열어서 작업해요.
```
USER: postgres
PASSWORD: postgres
DATABASE: baemin
```
postgresql://localhost:5432/baemin

### Tips
owner_auth, user_auth로 사장님과 사용자를 다른 테이블에서 관리해야요.

## API 명세서
Postman 혹은 Swegger를 사용해서 작업한 API를 정리해서 PR에 첨부해야해요.


## ✉️ 커밋 메세지
```
################
# <타입> : <제목> 의 형식으로 제목을 아래 공백줄에 작성
# 제목은 50자 이내 / 변경사항이 "무엇"인지 명확히 작성 / 끝에 마침표 금지
# 예) feat : 로그인 기능 추가
################
# feat : 새로운 기능 추가
# fix : 버그 수정
# docs : 문서 수정
# test : 테스트 코드 추가
# refact : 코드 리팩토링
# style : 코드 의미에 영향을 주지 않는 변경사항
# chore : 빌드 부분 혹은 패키지 매니저 수정사항
################
```

## ✉️ PR 템플릿
```
## 🔘Part

- [x] FE

  <br/>

## 🔎 작업 내용

- 기능에서 어떤 부분이

- 구현되었는지 설명해주세요

  <br/>

## 이미지 첨부

<img src="파일주소" width="50%" height="50%"/>

<br/>

## 🔧 앞으로의 과제

- 고치고 싶은 부분을 적어요.

  <br/>

## ➕ API 명세서

- 명세서 링크

<br/>
```