// Based on https://github.com/TheDiveO/gons/blob/master/gonamespaces.c

#define _GNU_SOURCE
#include <sched.h>
#include <sys/syscall.h>
#include <unistd.h>
#include <sys/stat.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <signal.h>
#include <stdarg.h>
#include <limits.h>
#include <errno.h>
#include <fcntl.h>

char *nsMsg;
static unsigned int maxMsgSize;

// Set error message nsMsg for go to read when checking status
static void setErrMsg(const char *format, ...){
    va_list args;
    if (!nsMsg) {
        maxMsgSize = 256 + PATH_MAX;
        nsMsg = (char *) malloc(maxMsgSize);
        if (!nsMsg){
            nsMsg = "malloc error";
            maxMsgSize = 0;
            return;
        }
    }
    va_start(args, format);
    vsnprintf(nsMsg, maxMsgSize, format, args);
    va_end(args);
}

int nsenter(int pid, char *type, long clonetype){
    char nsPath[PATH_MAX];
    snprintf(nsPath, PATH_MAX, "/proc/%d/ns/%s", pid, type);
    int fd = open(nsPath, O_RDONLY);
    if (fd < 0) {
        setErrMsg("invalid user ns reference %s: %s",
                nsPath, strerror(errno));
        return -1;
    }
    long res = syscall(SYS_setns, fd, clonetype);
    close(fd);
    if (res < 0){
        setErrMsg("cannot join %s namespace using reference %s: %s",
                type, nsPath, strerror(errno));
        return -1;
    }
    return 0;
}

void switchNamespace(void){
    // Get pid from environment
    const char* umlPidEnv = getenv("UML_NS_PID");
    if (umlPidEnv == NULL) {
        setErrMsg("environment variable UML_NS_PID is not set");
        return;
    }
    int umlPid = atoi(umlPidEnv);
    if (umlPid == 0){
        setErrMsg("environment variable UML_NS_PID (%s) is not an integer", umlPidEnv);
        return;
    }
    // check pid is running
    struct stat sts;
    if (kill(umlPid, 0) != 0){
        setErrMsg("process %d cannot be found", umlPid);
        return;
    }
    // enter uml user namespace
    int err = nsenter(umlPid, "user", CLONE_NEWUSER);
    if (err < 0){
        return;
    }
    // enter uml mount namespace
    nsenter(umlPid, "mnt", CLONE_NEWNS);
    // change dir to wd of parent
    const char* origWd = getenv("UML_ORIG_WD");
    if (origWd == NULL) {
        setErrMsg("environment variable UML_ORIG_WD is not set");
        return;
    }
    err = chdir(origWd);
    if (err < 0) {
        setErrMsg("could not change dir to UML_ORIG_WD (%s) : %s",
                origWd, strerror(errno));
        return;
    }
}
