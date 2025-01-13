# Go언어

<hr>

#### 1. 컴파일러 vs 인터프리터

###### 컴파일러

- 컴파일러는 C언어 처럼 모든 소스 코드를 한번에 읽고, 기계어로 번역
- 일반적으로, 속도가 굉장히 빠름 (C++은 1초에 약 1억번 연산 가능)
- C, C++, Go, Rust, Java 등이 있음

###### 인터프리터

- 인터프리터는 Python처럼 소스 코드를 한줄 한줄 읽고 차례대로 기계어로 번역
- 일반적으로, 컴파일러 보다 속도가 느림(Python은 1초에 약 2000만번 정도 연산 가능)
- 하지만, 파일을 처음 실행할 때 속도는 컴파일러 보다 빠름

      - 모든 소스 코드를 읽을 필요가 없기 때문

- Python, JavaScript, MATLAB, Ruby, R, SQL 등이 있음

###### 자바 컴파일러

- 자바는 프로그래밍언어중 거의 유일하게 컴파일러와 인터프리터를 함께 사용
- 컴파일러를 통해 소스코드를 바이트코드로 변환한 후 바이트코드를 인터프리터로 읽음
- 인터프리터 보다 속도가 빠르지만 컴파일러 만큼 빠르진 않음 (Java는 1초에 약 5500만번 정도 연산 가능)

#### 2. Go언어

- Go언어는 2012년 발표된 Google에서 개발된 프로그래밍 언어

![alt text](../images/gologo.png)

###### 장점

- 문법이 어렵지 않아서 배우기 쉬움
- Python, JavaScript와 다르게 정적 타입/강타입 언어이기 때문에 오류가 적을 수 있음
- 고루틴(GoRoutine)이 존재

      - 고루틴은 쓰레드보다 매우 가볍고 비교적 쉽게 이용 가능

- 기본 모듈이 매우 다양하여 프레임워크의 제한을 받지 않음

###### 단점

- 지원하지 않는 문법이 많음

      - 제너릭, 클래스, 예외 처리, Public/Private 키워드로 접근 제어 하지 않음, this 문법 없음, 자료형이 많지 않음

- 예외 처리가 없기 때문에 if를 많이 사용하여 코드가 깔끔하지 않음
- 중앙 저장소가 없어 Github 의존도가 강함
- 강력한 라이브러리가 따로 존재하지 않음(대표적인 라이브러리가 없음)

#### 3. 패키지 관리

- Go언어는 패키지를 github를 통해서 가져옴

###### go mod init

- go mod init을 통해 go.mod 파일을 생성
- 기본이 되는 모듈 주소를 확인, 설정 할 수 있음

###### go get

- go get을 통해 외부 라이브러리를 가져옴
- 기본적으로 최신 버전을 가져옴
- 특정 버전을 가져오고 싶다면 아래 명령어 실행

```bash
$ go get package@version
```

#### 4. Hello World

- 아래 코드는 Go언어를 통해 Hello World를 출력하는 코드

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}
```

```
출력: Hello World
```

#### 5. struct

- Go언어는 class 문법을 지원하지 않기 때문에, struct를 자주 사용
- 아래 코드는 Go언어를 통해 간단한 Person 객체를 정의하는 코드

```go
package main

import "fmt"

type person struct {
  name string
  age int
}

func main() {
  p := person{}

  p.name = "Lee"
  p.age = 10
  fmt.Println(p)
}
```

```
출력: {Lee 10}
```

#### 6. Go루틴

- Go루틴을 통해 비동기 처리 가능
- 아래 코드는 간단한 Go루틴 호출 코드

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	say("Sync")

	go say("Async1")
	go say("Async2")
	go say("Async3")

	time.Sleep(time.Second * 3)
}
```

```
출력:
Sync *** 0
Sync *** 1
Sync *** 2
Async1 *** 0
Async1 *** 1
Async1 *** 2
Async2 *** 0
Async3 *** 0
Async3 *** 1
Async3 *** 2
Async2 *** 1
Async2 *** 2
```
