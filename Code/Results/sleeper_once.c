#include<stdio.h>
#include<unistd.h>

int main() {
	while (1) {
		printf("Sleeping for 5 seconds\n");
		sleep(5);
		|\colorbox{SkyBlue}{return 0;}|
	}
}
