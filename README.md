# go_chat_project
> go언어기반의 websocket을 활용한 채팅시스템 구축

![Static Badge](https://img.shields.io/badge/Go-%233776AB)
![Static Badge](https://img.shields.io/badge/Gin-%23009688)
![Static Badge](https://img.shields.io/badge/Websocket-%234169E1)

## 해당프로젝트 상세내용
#### 1. go언어기반의 gin프레임워크사용.<br>
#### 2. gin프레임워크의 http연결을 websocket연결로 업그레이드하기위해 gorilla websocket 라이브러리 사용<br>
#### 3. UI를위한 프론트엔드 템플릿 지원<br>

![](../header.png)

## 개발 환경 설정
go와 yarn설치되어 있어야합니다.

## 실행 방법
winodws 기준
1. git clone을한다
2. frontend경로로가서 ```npm install``` 명령어 실행하여 node_module들을 설치해준다.
3. 설치가 완료되면 ```yarn start``` 명령어 실행한다.
4. backend경로로가서 ```go mod tidy``` 명령어 실행하여 go.mod에 명시된 의존성기반으로 필요한 패키지들을 설치한다.
5. ```go run .``` 명령어를 실행하여 채팅서버를 실행시켜준다. </br>


### 로그인 화면<br>
![image](https://github.com/CHOJUNGHO96/go_chat_project/assets/61762674/656b3e30-b96a-4984-a431-a8bc7188d981)
### 채팅방 화면 (one유저와 two유저의 실시간 대화)<br>
![image](https://github.com/CHOJUNGHO96/go_chat_project/assets/61762674/14a41875-5ab6-47b6-b883-7da1d50a5dd8)
![image](https://github.com/CHOJUNGHO96/go_chat_project/assets/61762674/708e8cd4-e5d5-4d39-adf0-9fbbdbe3a56c)



## 정보

조정호 – jo4186@naver.com

[https://github.com/CHOJUNGHO96/github-link](https://github.com/CHOJUNGHO96)

## 기여 방법

1. (<https://github.com/CHOJUNGHO96/fastapi-chat-project/fork>)을 포크합니다.
2. (`git checkout -b feature/fooBar`) 명령어로 새 브랜치를 만드세요.
3. (`git commit -am 'Add some fooBar'`) 명령어로 커밋하세요.
4. (`git push origin feature/fooBar`) 명령어로 브랜치에 푸시하세요. 
5. 풀리퀘스트를 보내주세요.

<!-- Markdown link & img dfn's -->
[npm-image]: https://img.shields.io/npm/v/datadog-metrics.svg?style=flat-square
[npm-url]: https://npmjs.org/package/datadog-metrics
[npm-downloads]: https://img.shields.io/npm/dm/datadog-metrics.svg?style=flat-square
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[wiki]: https://github.com/yourname/yourproject/wiki
