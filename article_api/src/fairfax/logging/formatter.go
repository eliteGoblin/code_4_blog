package logging

import (
	"github.com/sirupsen/logrus"
	"time"
)

type FairFaxFormatter struct {
	jsonFormatter *logrus.JSONFormatter
}

func InstanceOfFairFaxFormatter() *FairFaxFormatter {
	return &FairFaxFormatter{
		jsonFormatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "_time",
				logrus.FieldKeyLevel: "_level",
			},
			TimestampFormat: time.RFC3339Nano,
		},
	}
}

func (f *FairFaxFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return f.jsonFormatter.Format(entry)
}
