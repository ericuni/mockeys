# mockeys
当使用 [mockey](https://github.com/bytedance/mockey) 但是又想直接使用testing.T 来分组, 而不是使用 PatchConvey 时, 需要使用下面的方式.
```go
import "github.com/bytedance/mockey"

func Foo(x int) int {
	return x + 1
}

func TestMockey(t *testing.T) {
  assert := assert.New(t)

  defer mockey.Mock(Foo).Return(2).Build().UnPatch()

  assert.Equal(2, Foo(0))
}
```

有如下几个痛点

1. 当要替换的结构比较复杂时, gopls 可能会提示不出来 Build 或者 UnPatch, 需要手动输入
2. 当需要多个mock 时, 每个都加上 defer, Build, UnPatch 比较繁琐
3. 多个单测公用的mock 不好封装

mockeys 就是为了解决上面的几个问题
```go
import (
  "github.com/bytedance/mockey"
  "github.com/ericuni/mockeys"
)

func TestMockeys(t *testing.T) {
  patches := mockeys.NewPatches()
  defer patches.Reset()

  // common mock
  mock := mockey.Mock(Foo).Return(2)
  patches.Apply(mock)

  t.Run("case 1", func() {
    patches := mockeys.NewPatches()
    defer patches.Reset()

    mock := mockey.Mock(Bar).To(BarTarget1)
    patches.Apply(mock)

    mockVar := mockey.MockValue(&Value).To(ValueTarget1)
    patches.ApplyVar(mockVar)

    // do xxx
  })

  t.Run("case 2", func() {
    patches := mockeys.NewPatches()
    defer patches.Reset()

    mock := mockey.Mock(Bar).To(BarTarget2)
    patches.Apply(mock)

    mockVar := mockey.MockValue(&Value).To(ValueTarget2)
    patches.ApplyVar(mockVar)

    // do xxx
  })
}
```

多个单测公用的 mock
```go
import (
  "github.com/bytedance/mockey"
  "github.com/ericuni/mockeys"
)

type Repos struct {
  mysqlRepo *mock_dbops.Repo
  redisRepo *mock_redisops.Repo
  // ...
}

func mockDal(t *testing.T) (Repos, *mockeys.Patches) {
  ctrl := gomock.NewController(t)

  repos := Repos{}
  patches := mockeys.NewPatches()

  mysqlRepo := mock_dbops.NewMockRepo(ctrl)
  mock := mockey.Mock(dbops.GetRepo).Return(mysqlRepo)
  patches.Apply(mock)
  repos.mysqlRepo = mysqlRepo

  redisRepo := mock_redisops.NewMockRepo(ctrl)
  mock = mockey.Mock(redisops.GetRepo).Return(redisRepo)
  patches.Apply(mock)
  repos.redisRepo = redisRepo

  return repos, patches
}

func TestFoo(t *testing.T) {
  repos, patches := mockDal(t)
  defer patches.Reset()

  // do xxx
}

func TestBar(t *testing.T) {
  repos, patches := mockDal(t)
  defer patches.Reset()

  // do xxx
}
```

