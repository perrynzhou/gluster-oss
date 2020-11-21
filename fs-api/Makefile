all:
	rm -rf test_fapi test
	gcc  -g -std=gnu99 -DFS_API_TEST -O0 fs_api.h fs_api.c -o test_fapi -lgfapi
	gcc  -g -std=gnu99 -O0 main.c -o test -lgfapi
