#include <sys/mman.h>
#include <fcntl.h>
#include "tlpi_hdr.h"
#define MEM_SIZE 10

/*
 mmap根据是file-based还是匿名，是PRIVATE还是SHRED，可以出现4中组合
 
 1. file-based, PRIVATE
    用途: Initializing memory from contents of file
 2. file-based, SHARED, 
    用途:Memory-mapped I/O; sharing memory between processes (IPC)
 3. Anonymous, PRIVATE
    用途: Memory allocation
 3. Anonymous, SHARED
    用途: Sharing memory between processes (IPC)
*/

// 本例我们看第二种用途: Memory-Mapped I/O
// Copy from TLPI 49.4.2
// 从这个例子可以看出，利用mmap系统调用的Memory-Mapped I/O操作文件十分方便简洁
// 如果不想用tlpi_hdr.h中的函数，可以自己写这几个函数
int main(int argc, char *argv[]) {
    char *addr;
    int fd;
    if (argc < 2 || strcmp(argv[1], "--help") == 0) 
        usageErr("%s file [new-value]\n", argv[0]);
    
    fd = open(argv[1], O_RDWR);
    if (fd == -1)
        errExit("open");
    // Memory-Map I/O
    // 拿到映射的虚拟地址，可以在用户空间直接读写
    // 用户空间与内核空间共享同一份文件物理内存拷贝
    addr = mmap(NULL, MEM_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
     if (addr == MAP_FAILED)
        errExit("mmap");
    
    if (close(fd) == -1) /* No longer need 'fd' */ 
        errExit("close");
    // %.*s  动态指定精度的宽度
    // 参见 [What does “%.*s” mean in printf?]https://stackoverflow.com/questions/7899119/what-does-s-mean-in-printf
    printf("Current string=%.*s\n", MEM_SIZE, addr);
    
    if (argc > 2) { 
        if (strlen(argv[2]) >= MEM_SIZE)
            cmdLineErr("'new-value' too large\n");
        
        // 直接读写addr
        memset(addr, 0, MEM_SIZE);
        strncpy(addr, argv[2], MEM_SIZE - 1);
        // 不用write,直接调用msync讲所写内容刷入磁盘
        if (msync(addr, MEM_SIZE, MS_SYNC) == -1)
            errExit("msync");
        
        printf("Copied \"%s\" to shared memory\n", argv[2]); 
    }

    exit(EXIT_SUCCESS);
}