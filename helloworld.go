package hello

import "github.com/sirupsen/logrus"

func sayHello() string {
	logrus.Info("saying hello...")
	return "Hello, world."
}
