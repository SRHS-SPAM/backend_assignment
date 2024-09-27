# 백엔드 2차 과제

> 저번 과제에서 이어지는 과제에요. 실패하더라도 PR을 올려 선배들께 코드리뷰를 받아보아요.

이번 과제는 JWT를 사용하여 유저 서비스를 만들어볼거에요.

## 구현사항
### 사장님
- 식당 등록, 수정, 삭제 기능은 그대로 유지합니다.
- 회원가입: 사장님이 계정을 생성할 수 있어야 합니다.
- 로그인: 사장님이 자신의 계정으로 로그인할 수 있어야 합니다.
- JWT 발급: 로그인 시 JWT(Json Web Token)를 발급받아야 합니다.

### 사용자
- 식당 리스트 및 음식 조회 기능은 그대로 유지합니다.
- 회원가입: 사용자가 계정을 생성할 수 있어야 합니다.
- 로그인: 사용자가 자신의 계정으로 로그인할 수 있어야 합니다.
- JWT 발급: 로그인 시 JWT를 발급받아야 합니다.
- 인증된 사용자만 음식 추천/좋아요: 로그인된 사용자만 음식 추천 및 좋아요 기능을 사용할 수 있어야 합니다.

### shop
- 가게 이름으로 가게를 검색할 수 있어야해요.
- 가게 ID로 가게를 검색할 수 있어야해요.
-  onwer_name대신 owner_id를 사용해서 사장님의 uuid를 저장해야되요.

## Endpoint 및 파라미터 설정법
Endpoint 및 파라미터 설정법
### shop/auth
- 사장님 회원가입
- 사장님 로그인
- 사장님 로그아웃

### food/auth
- 사용자 회원가입
- 사용자 로그인
- 사용자 로그아웃


## 가져야하는 body 구조
```json
{
  "shop" : {
    "uuid" : "string", // uuid
    "name" : "string", // 식당이름
    "onwer_id" : "string", // 사장이름 (이번 구현에서는 직접 인력받아야해요)
    "category" : "string" // 한식, 양식, 중식, 일식, 분식 등
  }
}
```
> 음식은 변경된 사항이 없어요.
```json
{
  "owner" : {
     "ID" : "string", // 사장님 계정 아이디
     "password" : "string", // 비밀번호
     "shop_id" : "string" // 가게 아이디
  }
}
```
```json
{
  "user" : {
     "ID" : "string", // 유저 계정 아이디
     "password" : "string", // 비밀번호
     "liked" : []"string" // 좋아요 누른 음식 uuid를 배열로 저장해야되요.
  }
}
```

## 파일 구성
1번에서 진행했던 자신의 이름 폴더와 postgres-data폴더를 모두 복사 붙여넣기를 해서 작업해야해요.

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