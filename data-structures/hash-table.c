// https://www.youtube.com/watch?v=2Ti5yvumFTU&list=PLLchAlP_W0GfYWjv6Off6lfk4xNe_l-QB&index=75

#define MAX_NAME 256
#define TABLE_SIZE 10

typedef struct {
char name[MAX_NAME];
int age;

} person




//特性:足够随机
unsigned int hash(char *name) {
    int length = strnlen(name,MAX_NAME)'
    unsigned int hash_val=0;
    for (int i=0;i<length;i++) {
    hash_val+=name[i];
    hash_val = (hash_val * name[i]) %TABLE_SIZE;

    }
    return hash_val;

}

int main() {
printf("Jacob")
}