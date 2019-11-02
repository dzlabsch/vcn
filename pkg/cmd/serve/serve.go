/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package serve

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/internal/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

// NewCommand returns the cobra command for `vcn serve`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start a local API server",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return runServe(cmd)
		},
		Args: cobra.NoArgs,
	}
	cmd.Flags().String("host", "", "host address")
	cmd.Flags().String("port", "8080", "port")
	cmd.Flags().String("tls-cert-file", "", "TLS certificate file")
	cmd.Flags().String("tls-key-file", "", "TLS key file")
	return cmd
}

func runServe(cmd *cobra.Command) error {
	host, err := cmd.Flags().GetString("host")
	if err != nil {
		return nil
	}
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		return nil
	}
	addr := host + ":" + port

	certFile, _ := cmd.Flags().GetString("tls-cert-file")
	keyFile, _ := cmd.Flags().GetString("tls-key-file")
	if certFile != "" && keyFile == "" {
		return fmt.Errorf("--tls-key-file is missing")
	}
	if certFile == "" && keyFile != "" {
		return fmt.Errorf("--tls-cert-file is missing")
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/notarize", signHander(meta.StatusTrusted)).Methods("POST")
	router.HandleFunc("/untrust", signHander(meta.StatusUntrusted)).Methods("POST")
	router.HandleFunc("/unsupport", signHander(meta.StatusUnsupported)).Methods("POST")
	router.HandleFunc("/authenticate/{hash}", verify).Methods("GET")

	logs.LOG.Infof("Log level: %s", logs.LOG.GetLevel().String())
	logs.LOG.Infof("Stage: %s", meta.StageEnvironment().String())

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"content-type", "authorization", "x-notarization-password", "x-notarization-password-empty"}),
	)(router)

	if certFile != "" && keyFile != "" {
		logs.LOG.Infof("Listening on %s (TLS)", addr)
		return http.ListenAndServeTLS(addr, certFile, keyFile, handler)
	}

	logs.LOG.Infof("Listening on %s", addr)
	return http.ListenAndServe(addr, handler)
}

func index(w http.ResponseWriter, r *http.Request) {
	// can be used for healthcheck
	writeResponse(w, http.StatusOK, []byte("OK"))
}
