package main

import (
	"fmt"
	raRegion "github.com/Toorop/goabove/region"
	osStorage "github.com/Toorop/gopenstack/objectStorage/V1"
	"github.com/codegangsta/cli"
	"os"
	"path"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

// Commands

// getComputeCmds return commands for compute section
func getObjectStorageCmds() (objectStorageCmds []cli.Command) {
	// object storage commands
	objectStorageCmds = []cli.Command{
		{
			Name:        "list",
			ShortName:   "ls",
			Usage:       "ra storage list STORAGE_PATH",
			Description: "Lists a path",
			Action: func(c *cli.Context) {
				lines2display := []string{}
				fullPath := path.Clean(c.Args().First()) // fullpath format is REGION/OSPATH
				// add fisrt if needed /
				if fullPath[0] != 47 {
					fullPath = "/" + fullPath
				}
				rawRegion := strings.Split(fullPath, "/")[1]
				rawOsPath := fullPath[len(rawRegion)+1:]

				// List regions
				if rawRegion == "." {
					region, err := raRegion.New(raClient)
					dieOnError(err)
					lines2display, err = region.GetAll()
					dieOnError(err)
				} else {
					// Get a new openstack object storage client
					osClient, err := osStorage.NewClient(osKeyring, rawRegion)
					dieOnError(err)

					osPath := osStorage.NewOsPath(osClient, rawOsPath)

					children, err := osPath.ListChildren()
					dieOnError(err)

					originOfTheWorld := *new(time.Time)
					for _, child := range children {
						line := ""
						line += strconv.FormatInt(int64(child.Count), 10) + "\t"

						// Content-Type
						if child.ContentType != "" {
							line += child.ContentType + "\t"
						} else if child.Ptype != "" {
							line += child.Ptype + "\t"
						}
						// bytes
						line = line + strconv.FormatInt(int64(child.Bytes/1024), 10) + "\t"
						//Date
						if child.LastModified.Time != originOfTheWorld {
							line += child.LastModified.Format(time.RFC3339) + "\t"
						}
						// Etag
						line += child.Etag + "\t"
						// Name
						line += child.Name
						lines2display = append(lines2display, line)
					}

				}
				// output
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 5, 0, 3, ' ', 0)
				for _, l := range lines2display {
					fmt.Fprintln(w, l)
				}
				w.Flush()
				dieOk()
			},
		}, {
			Name:        "copy",
			ShortName:   "cp",
			Usage:       "ra storage cp SRC DEST",
			Description: "Copies SRC to DEST. SRC & DEST can be a local path or a remote path. Remote Path (on Runabove) must start with the storage region (ex /SBG-1/images/oles-nude.jpg).",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 2)
				src := c.Args().Get(0)
				dest := c.Args().Get(1)
				region := ""
				// Get Available régions
				apiRegion, err := raRegion.New(raClient)
				dieOnError(err)
				regions, err := apiRegion.GetAll()
				dieOnError(err)
				paths := []string{src, dest}
				for k, path := range paths {
					for _, r := range regions {
						if strings.HasPrefix(path, r) || strings.HasPrefix(path, "/"+r) {
							// User just give a region without container (should we raise error ?)
							if len(path) < len(r)+1 {
								paths[k] = ""
							} else {
								paths[k] = path[len(r)+1:]
							}
							if region == "" {
								region = r
							} else if region != r {
								dieError("Copy between 2 Runabove regions are not possible (yet).")
							}
						}
					}
				}
				//Get OpenStack object storage client
				osClient, err := osStorage.NewClient(osKeyring, region)
				dieOnError(err)

				// Get swift client
				swift := osStorage.NewSwift(osClient)

				// Do copy
				dieOnError(swift.Copy(paths[0], paths[1]))
				dieOk()
			},
		}, {
			Name:        "remove",
			ShortName:   "rm",
			Usage:       "ra storage rm PATH",
			Description: "Remove PATH (container,object or folder). PATH must start with the storage region (ex /SBG-1/images/oles-naked.jpg).",
			Action: func(c *cli.Context) {
				dieIfArgsMiss(len(c.Args()), 1)
				path := c.Args().First()
				if !strings.HasPrefix(path, "/") {
					path = "/" + path
				}
				// We must have a region
				if strings.Count(path, "/") < 2 {
					dieError("No region specified")
				}
				region := strings.Split(path, "/")[1]
				path2remove := path[len(region)+1:]

				//Get OpenStack object storage client
				osClient, err := osStorage.NewClient(osKeyring, region)
				dieOnError(err)
				dieOnError(osStorage.NewSwift(osClient).DeletePath(path2remove))
				dieOk()
			},
		},
	}
	return
}
