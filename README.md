# kispi/goexamples

#### go 사용법에 대한 사소한 예시들

### Compiles and Run
```
go run main.go
```

go는 현재 gopath모드가 아닌 module 모드가 default이며, 이 상황에서는 local package를 import하는게 다소 직관적이지 않다.<br><br>어떻게 하는지 도저히 찾지 못해 꽤 많은 시간을 삽질하다가 겨우 이해하게 되어 참고용으로 남긴다.

<b>localpackage</b>: go에서 local package를 만드는 방법을 설명.

1. 디렉토리를 하나 만든다. (ex: ```mkdir goexamples```)
1. 그 안에 들어가서 ```go mod init goexamples```를 입력한다. (보통은 github.com으로 시작하도록 지음)<br>
=> "goexamples"라는 프로젝트(패키지)를 생성하는 것이다. go.mod 파일이 생성되며, 해당 파일에는 "module goexamples"라는 문장으로 현재 패키지의 이름을 정하는 코드가 생성된다. (nodejs 기반 package.json의 "name"필드로 생각할 수도 있긴 한데, 그와는 달리 실제 "module goexamples"는 컴파일시 영향을 준다.)
2. 로컬 패키지를 만들어본다. (ex: localpackage)
3. 내부의 파일들의 서두에 선언되는 package "패키지명"은 전부 동일해야 한다.
4. 보통은 디렉토리명과 동일하게 사용(ex: localpackage)하나, 달라도 작동은 된다.
5. main.go에서 사용시에는 "메인패키지명/상대패키지명"으로 import하여 사용 가능하다.<br>(ex: ``` import "goexamples/localpackage"```)

이 레포를 클론받고 값들을 바꿔보며 어떨 때 VSCode가 에러를 내는지 확인해보면 쉽게 이해할 수 있다.