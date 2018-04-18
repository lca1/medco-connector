package main

import (
	"errors"
	"github.com/dedis/onet/app"
	"github.com/dedis/onet/log"
	"github.com/lca1/medco/app/loader"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"time"
)

func computePerfFromApp(c *cli.Context) error {

	// cli arguments
	groupFilePath := c.String(optionGroupFile)
	entryPointIdx := c.Int(optionEntryPointIdx)
	perfType := c.String(optionPerfType)

	// generate el with group file
	f, err := os.Open(groupFilePath)
	if err != nil {
		log.Error("Error while opening group file", err)
		return cli.NewExitError(err, 1)
	}
	el, err := app.ReadGroupDescToml(f)
	if err != nil {
		log.Error("Error while reading group file", err)
		return cli.NewExitError(err, 1)
	}
	if len(el.Roster.List) <= 0 {
		log.Error("Empty or invalid group file", err)
		return cli.NewExitError(err, 1)
	}

	// perform test
	switch perfType {
	case "encryptAndTag":

		// parameter
		if c.NArg() != 1 {
			err := errors.New("1 argument needed (number of elements to test)")
			log.Error(err)
			return cli.NewExitError(err, 3)
		}

		nbElements, err := strconv.ParseInt(c.Args().Get(0), 10, 64)
		if err != nil {
			log.Error(err)
			return cli.NewExitError(err, 2)
		}

		// allocate array and run test
		testValues := make([]int64, 0, nbElements)

		for i := int64(0); i < nbElements; i++ {
			testValues = append(testValues, i)
		}

		start := time.Now()
		listEncryptedElements := loader.EncryptElements(testValues, el.Roster)
		loader.TagElements(listEncryptedElements, el.Roster, entryPointIdx)
		log.LLvl1("Encrypt and tag for ", nbElements, "... (", time.Since(start), ")")

	}

	return nil
}
