package main

/*
  Copyright (c) IBM Corporation 2016

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific

   Contributors:
     Mark Taylor - Initial Contribution
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	cf "github.com/ibm-messaging/mq-metric-samples/pkg/config"

	log "github.com/sirupsen/logrus"
)

type mqInfluxConfig struct {
	cf cf.Config

	databaseName    string
	databaseAddress string
	userid          string
	password        string
	passwordFile    string

	interval  string
	maxErrors int
}

var config mqInfluxConfig

/*
initConfig parses the command line parameters.
*/
func initConfig() error {
	var err error

	cf.InitConfig(&config.cf)

	flag.StringVar(&config.databaseName, "ibmmq.databaseName", "", "Name of database")
	flag.StringVar(&config.databaseAddress, "ibmmq.databaseAddress", "", "Address of database eg http://example.com:8086")
	flag.StringVar(&config.userid, "ibmmq.databaseUserID", "", "UserID to access the database")
	flag.StringVar(&config.interval, "ibmmq.interval", "10s", "How long between each collection")
	flag.StringVar(&config.passwordFile, "ibmmq.pwFile", "", "Where is password help temporarily")
	flag.IntVar(&config.maxErrors, "ibmmq.maxErrors", 100, "Maximum number of errors communicating with server before considered fatal")

	flag.Parse()

	if len(flag.Args()) > 0 {
		err = fmt.Errorf("Extra command line parameters given")
		flag.PrintDefaults()
	}

	if err == nil {
		err = cf.VerifyConfig(&config.cf)
	}

	// Read password from a file if there is a userid on the command line
	// Delete the file after reading it.
	if err == nil {
		if config.userid != "" {
			config.userid = strings.TrimSpace(config.userid)

			f, err := os.Open(config.passwordFile)
			if err != nil {
				log.Fatalf("Opening file %s: %s", f, err)
			}

			defer os.Remove(config.passwordFile)
			defer f.Close()

			scanner := bufio.NewScanner(f)
			scanner.Scan()
			p := scanner.Text()
			err = scanner.Err()
			if err != nil {
				log.Fatalf("Reading file %s: %s", f, err)
			}
			config.password = strings.TrimSpace(p)
		}
	}

	if err == nil {
		cf.InitLog(config.cf)
	}
	return err
}
