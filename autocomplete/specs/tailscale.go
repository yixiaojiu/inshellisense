// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["tailscale"] = model.Subcommand{
		Name: []string{"tailscale"},
		Options: []model.Option{{
			Name:        []string{"--socket"},
			Description: `Path to tailscaled's unix socket`,
			Args: []model.Arg{{
				Name: "socket",
			}},
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"up"},
			Description: `Connect to Tailscale, logging in if needed`,
			Options: []model.Option{{
				Name:        []string{"--accept-dns"},
				Description: `Accept DNS configuration from the admin panel`,
			}, {
				Name:        []string{"--accept-routes"},
				Description: `Accept routes advertised by other Tailscale nodes`,
			}, {
				Name:        []string{"--advertise-exit-node"},
				Description: `Offer to be an exit node for internet traffic for the tailnet`,
			}, {
				Name:        []string{"--advertise-routes"},
				Description: `Routes to advertise to other nodes (comma-separated, e.g. "10.0.0.0/8,192.168.0.0/24") or empty string to not advertise routes`,
				Args: []model.Arg{{
					Name: "routes",
				}},
			}, {
				Name:        []string{"--advertise-tags"},
				Description: `Comma-separated ACL tags to request; each must start with "tag:" (e.g. "tag:eng,tag:montreal,tag:ssh")`,
				Args: []model.Arg{{
					Name: "tags",
				}},
			}, {
				Name:        []string{"--auth-key"},
				Description: `Node authorization key; if it begins with "file:", then it's a path to a file containing the authkey`,
				Args: []model.Arg{{
					Name: "authkey",
				}},
			}, {
				Name:        []string{"--exit-node"},
				Description: `Tailscale exit node (IP or base name) for internet traffic, or empty string to not use an exit node`,
				Args: []model.Arg{{
					Name: "exitnode",
				}},
			}, {
				Name:        []string{"--exit-node-allow-lan-access"},
				Description: `Allow direct access to the local network when routing traffic via an exit node`,
			}, {
				Name:        []string{"--force-reauth"},
				Description: `Force reauthentication`,
			}, {
				Name:        []string{"--host-routes"},
				Description: `Install host routes to other Tailscale nodes`,
			}, {
				Name:        []string{"--hostname"},
				Description: `Hostname to use instead of the one provided by the OS`,
			}, {
				Name:        []string{"--json"},
				Description: `Output in JSON format`,
			}, {
				Name:        []string{"--login-server"},
				Description: `Login server to use`,
			}, {
				Name:        []string{"--operator"},
				Description: `Unix username to allow to operate on tailscaled without sudo`,
				Args: []model.Arg{{
					Name:        "username",
					Suggestions: []model.Suggestion{{Name: []string{`$USER`}}},
				}},
			}, {
				Name:        []string{"--qr"},
				Description: `Show QR code for login URLs`,
			}, {
				Name:        []string{"--reset"},
				Description: `Reset unspecified settings to their default values`,
			}, {
				Name:        []string{"--shields-up"},
				Description: `Don't allow incoming connections`,
			}, {
				Name:        []string{"--ssh"},
				Description: `Run an SSH server, permitting access per tailnet admin's declared policy`,
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"down"},
			Description: `Disconnect from Tailscale`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"logout"},
			Description: `Disconnect from Tailscale and expire current node key`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"netcheck"},
			Description: `Print an analysis of local network conditions`,
			Options: []model.Option{{
				Name:        []string{"--every"},
				Description: `If non-zero, do an incremental report with the given frequency`,
			}, {
				Name:        []string{"--format"},
				Description: `Output format`,
				Args: []model.Arg{{
					Name: "format",
					Suggestions: []model.Suggestion{{
						Name: []string{`json`},
					}, {
						Name: []string{`json-line`},
					}, {
						Name: []string{`human-readable`},
					}},
				}},
			}, {
				Name:        []string{"--verbose"},
				Description: `Verbose logs`,
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"ip"},
			Description: `Show Tailscale IP addresses`,
			Args: []model.Arg{{
				Generator:  nil, // TODO: port over generator
				IsOptional: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--1"},
				Description: `Only print one IP address`,
			}, {
				Name:        []string{"--4"},
				Description: `Only print IPv4 address`,
			}, {
				Name:        []string{"--6"},
				Description: `Only print IPv6 address`,
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"status"},
			Description: `Show state of tailscaled and its connections`,
			Options: []model.Option{{
				Name:        []string{"--active"},
				Description: `Filter output to only peers with active sessions`,
			}, {
				Name:        []string{"--browser"},
				Description: `Open a browser in web mode`,
			}, {
				Name:        []string{"--json"},
				Description: `Output in JSON format`,
			}, {
				Name:        []string{"--listen"},
				Description: `Listen address for web mode; use port 0 for automatic (default 127.0.0.1:8384)`,
				Args: []model.Arg{{
					Name: "address",
				}},
			}, {
				Name:        []string{"--peers"},
				Description: `Show status of peers`,
			}, {
				Name:        []string{"--self"},
				Description: `Show status of local machine`,
			}, {
				Name:        []string{"--web"},
				Description: `Run webserver with HTML showing status`,
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"ping"},
			Description: `Ping a host at the Tailscale layer, see how it routed`,
			Options: []model.Option{{
				Name:        []string{"--c"},
				Description: `Max number of pings to send`,
				Args: []model.Arg{{
					Name: "count",
				}},
			}, {
				Name:        []string{"--timeout"},
				Description: `Timeout before giving up on a ping`,
				Args: []model.Arg{{
					Name: "timeout",
				}},
			}, {
				Name:        []string{"--tsmp"},
				Description: `Do a TSMP-level ping (through IP + wireguard, but not involving host OS stack)`,
			}, {
				Name:        []string{"--until-direct"},
				Description: `Stop once a direct path is established`,
			}, {
				Name:        []string{"--verbose"},
				Description: `Verbose logs`,
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"nc"},
			Description: `Connect to a port on a host, connected to stdin/stdout`,
			Args: []model.Arg{{
				Name: "hosname-or-ip",
			}, {
				Name: "port",
			}},
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"ssh"},
			Description: `SSH to a Tailscale machine`,
			Args: []model.Arg{{
				Name: "[user@]<host>",
			}},
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"version"},
			Description: `Print Tailscale version`,
			Options: []model.Option{{
				Name:        []string{"--daemon"},
				Description: `Also print local node's daemon version`,
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"web"},
			Description: `Run a web server for controlling Tailscale`,
			Options: []model.Option{{
				Name:        []string{"--cgi"},
				Description: `Run as CGI script`,
			}, {
				Name:        []string{"--listen"},
				Description: `Listen address; use port 0 for automatic (default localhost:8088)`,
				Args: []model.Arg{{
					Name: "address",
				}},
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"file"},
			Description: `Send or receive files`,
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"cp"},
				Description: `Copy file(s) to a host`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}, {
					Name:       "file or host",
					Generator:  nil, // TODO: port over generator
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"get"},
				Description: `Move files out of the Tailscale file inbox`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "target-directory",
				}},
			}},
		}, {
			Name:        []string{"bugreport"},
			Description: `Print a shareable identifier to help diagnose issues`,
			Args: []model.Arg{{
				Name:       "note",
				IsOptional: true,
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"cert"},
			Description: `Get TLS certs`,
			Options: []model.Option{{
				Name:        []string{"--cert-file"},
				Description: `Output cert file or "-" for stdout`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Stdout`,
					}},
				}},
			}, {
				Name:        []string{"--key-file"},
				Description: `Output cert file or "-" for stdout`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
					Suggestions: []model.Suggestion{{
						Name:        []string{`-`},
						Description: `Stdout`,
					}},
				}},
			}, {
				Name:        []string{"--serve-demo"},
				Description: `Serve on port :443 using the cert as a demo, instead of writing out the files to disk`,
			}, {
				Name:        []string{"--help"},
				Description: `Show help message`,
			}},
		}, {
			Name:        []string{"help"},
			Description: `Print help message`,
		}},
	}
}