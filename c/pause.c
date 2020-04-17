#include <stdio.h>
#include <string.h>
#include <sched.h>
#include <unistd.h>
#include <signal.h>
#include <sys/types.h>
#include <stdlib.h>


#define STRINGIFY(x) #x
#define VERSION_STRING(X) STRINGIFY(x)

#ifndef VERSION
#define VERSION HEAD
#endif

static void sigdown(int signo) {
    // Shuting down, got signal: Interrupt
    psignal(signo, "Shuting down, got signal");
    exit(0);
}

static void sigreap(int signo) {
    // wait所有子进程，非阻塞
    // >0 说明wait到了子进程，继续;否则离开while循环
    // NULL代表我们不关心导致子进程状态变化的原因
    while (waitpid(-1, NULL, WNOHANG) > 0)
        ;
}


int main(int argc, char **argv) {
    int i;
    for (i = 1; i < argc; ++i) {        
        if (!strcasecmp(argv[i], "-v")) {
            printf("pause.c %s\n", VERSION_STRING(VERSION));
            return 0;
        }
    }

    if (getpid() != 1) {
        fprintf(stderr, "Warning: pause should be the first process\n");
    }
    // 注册信号处理函数
    // 第三个参数是old disposition,我们不关心，置空.
    // SIGINT,终结前台进程组
    if (sigaction(SIGINT, &(struct sigaction){.sa_handler = sigdown}, NULL) < 0)  {
        return 1;
    }
    // SIGTERM, 终结指定进程
    if (sigaction(SIGTERM, &(struct sigaction){.sa_handler = sigdown}, NULL) < 0) {
        return 2;
    }
    // SA_NOCLDSTOP,不要将子进程变成僵尸进程。
    if (sigaction(SIGCHLD, &(struct sigaction){.sa_handler = sigreap, .sa_flags = SA_NOCLDSTOP}, NULL) < 0) {
        return 3;
    }
    
    for (;;)
        // 等待信号
        // 调用pause挂起当前进程，直到被信号处理函数打断睡眠.
        // The only time pause returns is if a signal handler is executed and that handler returns. 
        // In that case, pause returns −1 with errno set to EINTR.
        pause();

    fprintf(stderr, "Error: infiinite loop terminated\n");
    return 42;    
    
}

