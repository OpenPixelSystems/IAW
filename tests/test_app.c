#include <stdio.h>
#include <stdlib.h>

int main()
{
	for (int i = 0; i < 10; i++) {
		printf("test\r\n");
		fflush(stdout);
		sleep(1);
	}
	printf("error\n");
	fflush(stdout);
	for (int i = 0; i < 10; i++) {
		printf("test\r\n");
		fflush(stdout);
		sleep(1);
	}
	printf("error\n");

	return 0;
}
