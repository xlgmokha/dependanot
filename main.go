package main

import (
  "github.com/sirupsen/logrus"
  "time"
)

func init() {
  logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
  lastYear := time.Now().AddDate(-1, 0, 0)
  logrus.WithTime(lastYear).Info("hello, world!")
}
