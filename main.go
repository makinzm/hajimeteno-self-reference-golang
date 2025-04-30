package main

import (
  "go.uber.org/zap"
)

type some struct {
    parent *st
}

type st struct {
    some *some
}

func newST() *st {
    st := &st{}
    st.some = &some{
        parent: st,
    }
    return st
}

var logger *zap.Logger

func init() {
  var err error
  logger, err = zap.NewProduction()
  if err != nil {
    panic(err)
  }
}

func main() {
  st := newST()

  logger.Info(
    "Is some.parent == st?",
    zap.Bool("result", st.some.parent == st),
  )
  logger.Info(
    "Is some.parent.some == st.some?",
    zap.Bool("result", st.some.parent.some == st.some),
  )
  logger.Info(
    "Is some.parent.some.parent == st?",
    zap.Bool("result", st.some.parent.some.parent == st),
  )

  st2 := newST()
  logger.Info(
    "Is st == st2?",
    zap.Bool("result", st == st2),
  )
}

