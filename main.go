package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

var (
	dsvPort         = flag.Uint("dsv-port", 9090, "DSV Server listen tcp-port")
	dsvAppDir       = flag.String("dsv-app-dir", "dsv-server", "DSV Server application directory")
	dsvLogFile      = flag.String("dsv-log", "dsv.log", "DSV Server log file path")
	dsvLogPropsFile = flag.String("dsv-log-props", "", "DSV Server logging properties file path")
	vpnAppDir       = flag.String("vpn-app-dir", "vpn-client", "VPN Client application directory")
	vpnLogFile      = flag.String("vpn-log", "vpn.log", "VPN Client log file path")
	vpnCongFile     = flag.String("vpn-config", "client-eimzo.conf", "VPN Client config file path")
	vpnLogPropsFile = flag.String("vpn-log-props", "", "VPN Client logging properties file path")
)

func init() {
	flag.Parse()
}

func main() {
	procs := make([]*os.Process, 0)
	mtx := new(sync.Mutex)

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		lf, err := os.OpenFile(*dsvLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			logrus.Errorf("failed to open log file [%v]: %v", *dsvLogFile, err)
			return
		}
		if os.Getenv("TSA_URL") == "" {
			os.Setenv("TSA_URL", "http://e-imzo.uz/cams/tst")
		}
		if os.Getenv("TRUSTSTORE_FILE") == "" {
			os.Setenv("TRUSTSTORE_FILE", "keys/truststore.jks")
		}
		if os.Getenv("TRUSTSTORE_PASSWORD") == "" {
			os.Setenv("TRUSTSTORE_PASSWORD", "12345678")
		}
		args := []string{"-Dfile.encoding=UTF-8", "-jar", "dsv-server.jar", "-l", "0.0.0.0", "-p", strconv.Itoa(int(*dsvPort))}
		if *dsvLogPropsFile != "" {
			args = append([]string{fmt.Sprintf("-Djava.util.logging.config.file=%v", *dsvLogPropsFile)}, args...)
		}
		cmd := exec.Command("java", args...)
		cmd.Dir = *dsvAppDir
		cmd.Env = os.Environ()
		cmd.Stdout = lf
		cmd.Stderr = lf
		err = cmd.Start()
		if err != nil {
			logrus.Errorf("failed to start command [%v]: %v", cmd.String(), err)
			return
		}
		logrus.Infof("started [%v] with PID: %v", cmd.String(), cmd.Process.Pid)

		mtx.Lock()
		procs = append(procs, cmd.Process)
		mtx.Unlock()

		err = cmd.Wait()
		if err != nil {
			logrus.Errorf("failed to wait command [%v]: %v", cmd.String(), err)
			return
		}
		logrus.Warnf("stopped [%v]", cmd.String())
		logrus.Warnf("see [%v]", *dsvLogFile)
	}()

	go func() {
		defer cancel()
		lf, err := os.OpenFile(*vpnLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			logrus.Errorf("failed to open log file [%v]: %v", *vpnLogFile, err)
			return
		}

		args := []string{"-Dfile.encoding=UTF-8", "-jar", "vpn-client.jar", *vpnCongFile}
		if *vpnLogPropsFile != "" {
			args = append([]string{fmt.Sprintf("-Djava.util.logging.config.file=%v", *vpnLogPropsFile)}, args...)
		}
		cmd := exec.Command("java", args...)
		cmd.Dir = *vpnAppDir
		cmd.Env = os.Environ()
		cmd.Stdout = lf
		cmd.Stderr = lf
		err = cmd.Start()
		if err != nil {
			logrus.Errorf("failed to start command [%v]: %v", cmd.String(), err)
			return
		}
		logrus.Infof("started [%v] with PID: %v", cmd.String(), cmd.Process.Pid)

		mtx.Lock()
		procs = append(procs, cmd.Process)
		mtx.Unlock()

		err = cmd.Wait()
		if err != nil {
			logrus.Errorf("failed to wait command [%v]: %v", cmd.String(), err)
			return
		}
		logrus.Warnf("stopped [%v]", cmd.String())
		logrus.Warnf("see [%v]", *vpnLogFile)
	}()

	select {
	case <-ctx.Done():
	case sig := <-cancelChan:
		logrus.Infof("caught SIGTERM %v", sig)
	}
	for _, p := range procs {
		p.Kill()
	}
	os.Exit(1)
}
